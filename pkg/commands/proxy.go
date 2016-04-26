package commands

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
	shlex "github.com/flynn/go-shlex"

	"github.com/moul/advanced-ssh-config/pkg/config"
	. "github.com/moul/advanced-ssh-config/pkg/logger"
)

func cmdProxy(c *cli.Context) {
	Logger.Debugf("assh args: %s", c.Args())

	if len(c.Args()) < 1 {
		Logger.Fatalf("assh: \"proxy\" requires 1 argument. See 'assh proxy --help'.")
	}

	// dry-run option
	// Setting the 'ASSH_DRYRUN=1' environment variable,
	// so 'assh' can use gateways using sub-SSH commands.
	if c.Bool("dry-run") == true {
		os.Setenv("ASSH_DRYRUN", "1")
	}
	dryRun := os.Getenv("ASSH_DRYRUN") == "1"

	conf, err := config.Open()
	if err != nil {
		Logger.Fatalf("Cannot open configuration file: %v", err)
	}

	if err = conf.LoadKnownHosts(); err != nil {
		Logger.Debugf("Failed to load assh known_hosts: %v", err)
	}

	target := c.Args()[0]

	isOutdated, err := conf.IsConfigOutdated(target)
	if err != nil {
		Logger.Warnf("Cannot check if ~/.ssh/config is outdated.")
	}
	if isOutdated {
		Logger.Debugf("The configuration file is outdated, rebuilding it before calling ssh")
		Logger.Warnf("'~/.ssh/config' has been rewritten.  SSH needs to be restarted.  See https://github.com/moul/advanced-ssh-config/issues/122 for more information.")
		Logger.Debugf("Saving SSH config")
		err = conf.SaveSshConfig()
		if err != nil {
			Logger.Fatalf("Cannot save SSH config file: %v", err)
		}
	}

	// FIXME: handle complete host with json

	host, err := computeHost(target, c.Int("port"), conf)
	if err != nil {
		Logger.Fatalf("Cannot get host '%s': %v", target, err)
	}
	w := Logger.Writer()
	defer w.Close()
	host.WriteSshConfigTo(w)

	Logger.Debugf("Proxying")
	err = proxy(host, conf, dryRun)
	if err != nil {
		Logger.Fatalf("Proxy error: %v", err)
	}
}

func computeHost(dest string, portOverride int, conf *config.Config) (*config.Host, error) {
	host := conf.GetHostSafe(dest)

	if portOverride > 0 {
		host.Port = strconv.Itoa(portOverride)
	}

	return host, nil
}

func prepareHostControlPath(host, gateway *config.Host) error {
	controlPathDir := path.Dir(os.ExpandEnv(strings.Replace(host.ControlPath, "~", "$HOME", -1)))
	gatewayControlPath := path.Join(controlPathDir, gateway.Name())
	if config.BoolVal(host.NoControlMasterMkdir) {
		return nil
	}
	return os.MkdirAll(gatewayControlPath, 0700)
}

func proxy(host *config.Host, conf *config.Config, dryRun bool) error {
	if len(host.Gateways) > 0 {
		Logger.Debugf("Trying gateways: %s", host.Gateways)
		for _, gateway := range host.Gateways {
			if gateway == "direct" {
				err := proxyDirect(host, dryRun)
				if err != nil {
					Logger.Errorf("Failed to use 'direct' connection: %v", err)
				}
			} else {
				hostCopy := host.Clone()
				gatewayHost := conf.GetGatewaySafe(gateway)

				err := prepareHostControlPath(hostCopy, gatewayHost)
				if err != nil {
					return err
				}

				if hostCopy.ProxyCommand == "" {
					hostCopy.ProxyCommand = "nc %h %p"
				}
				// FIXME: dynamically add "-v" flags

				if err = hostPrepare(hostCopy); err != nil {
					return err
				}

				command := "ssh %name -- " + hostCopy.ExpandString(hostCopy.ProxyCommand)

				Logger.Debugf("Using gateway '%s': %s", gateway, command)
				err = proxyCommand(gatewayHost, command, dryRun)
				if err == nil {
					return nil
				}
				Logger.Errorf("Cannot use gateway '%s': %v", gateway, err)
			}
		}
		return fmt.Errorf("No such available gateway")
	}

	Logger.Debugf("Connecting without gateway")
	return proxyDirect(host, dryRun)
}

func proxyDirect(host *config.Host, dryRun bool) error {
	if host.ProxyCommand != "" {
		return proxyCommand(host, host.ProxyCommand, dryRun)
	}
	return proxyGo(host, dryRun)
}

func proxyCommand(host *config.Host, command string, dryRun bool) error {
	command = host.ExpandString(command)
	Logger.Debugf("ProxyCommand: %s", command)
	args, err := shlex.Split(command)
	if err != nil {
		return err
	}

	if dryRun {
		return fmt.Errorf("dry-run: Execute %s", args)
	}

	spawn := exec.Command(args[0], args[1:]...)
	spawn.Stdout = os.Stdout
	spawn.Stdin = os.Stdin
	spawn.Stderr = os.Stderr
	return spawn.Run()
}

func hostPrepare(host *config.Host) error {
	if host.HostName == "" {
		host.HostName = host.Name()
	}

	if len(host.ResolveNameservers) > 0 {
		Logger.Debugf("Resolving host: '%s' using nameservers %s", host.HostName, host.ResolveNameservers)
		// FIXME: resolve using custom dns server
		results, err := net.LookupAddr(host.HostName)
		if err != nil {
			return err
		}
		if len(results) > 0 {
			host.HostName = results[0]
		}
		Logger.Debugf("Resolved host is: %s", host.HostName)
	}

	if host.ResolveCommand != "" {
		command := host.ExpandString(host.ResolveCommand)
		Logger.Debugf("Resolving host: %q using command: %q", host.HostName, command)

		args, err := shlex.Split(command)
		if err != nil {
			return err
		}

		cmd := exec.Command(args[0], args[1:]...)
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		if err := cmd.Run(); err != nil {
			Logger.Errorf("ResolveCommand failed: %s", stderr.String())
			return err
		}

		host.HostName = strings.TrimSpace(fmt.Sprintf("%s", stdout.String()))
		Logger.Debugf("Resolved host is: %s", host.HostName)
	}
	return nil
}

type exportReadWrite struct {
	written uint64
	err     error
}

func proxyGo(host *config.Host, dryRun bool) error {
	Logger.Debugf("Preparing host object")
	if err := hostPrepare(host); err != nil {
		return err
	}

	if dryRun {
		return fmt.Errorf("dry-run: Golang native TCP connection to '%s:%s'", host.HostName, host.Port)
	}

	Logger.Debugf("Connecting to %s:%s", host.HostName, host.Port)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host.HostName, host.Port))
	if err != nil {
		return err
	}
	Logger.Debugf("Connected to %s:%s", host.HostName, host.Port)

	// Ignore SIGHUP
	signal.Ignore(syscall.SIGHUP)

	waitGroup := sync.WaitGroup{}
	result := exportReadWrite{}

	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "sync", &waitGroup)

	waitGroup.Add(2)
	c1 := readAndWrite(ctx, conn, os.Stdout)
	c2 := readAndWrite(ctx, os.Stdin, conn)
	select {
	case result = <-c1:
	case result = <-c2:
	}
	if result.err != nil && result.err == io.EOF {
		result.err = nil
	}
	select {
	case res := <-c2:
		Logger.Debugf("Byte written %v", res.written)
	default:
		Logger.Debugf("Byte written %v", result.written)
	}
	conn.Close()
	cancel()
	waitGroup.Wait()
	return result.err
}

func readAndWrite(ctx context.Context, r io.Reader, w io.Writer) <-chan exportReadWrite {
	buff := make([]byte, 1024)
	c := make(chan exportReadWrite, 1)

	go func() {
		defer ctx.Value("sync").(*sync.WaitGroup).Done()

		export := exportReadWrite{}
		for {
			select {
			case <-ctx.Done():
				c <- export
				return
			default:
				nr, err := r.Read(buff)
				if err != nil {
					export.err = err
					c <- export
					return
				}
				if nr > 0 {
					wr, err := w.Write(buff[:nr])
					if err != nil {
						export.err = err
						c <- export
						return
					}
					if wr > 0 {
						export.written += uint64(wr)
					}
				}
			}
		}
	}()
	return c
}
