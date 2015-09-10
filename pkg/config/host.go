package config

import (
	"fmt"
	"io"
	"strings"
)

// Host defines the configuration flags of a host
type Host struct {
	// ssh-config fields
	AddressFamily                    string `toml:"AddressFamily" json:"AddressFamily,omitempty"`
	BatchMode                        string `toml:"BatchMode" json:"BatchMode,omitempty"`
	BindAddress                      string `toml:"BindAddress" json:"BindAddress,omitempty"`
	CanonicalDomains                 string `toml:"CanonicalDomains" json:"CanonicalDomains,omitempty"`
	CanonicalizeFallbackLocal        string `toml:"CanonicalizeFallbackLocal" json:"CanonicalizeFallbackLocal,omitempty"`
	CanonicalizeHostname             string `toml:"CanonicalizeHostname" json:"CanonicalizeHostname,omitempty"`
	CanonicalizeMaxDots              string `toml:"CanonicalizeMaxDots" json:"CanonicalizeMaxDots,omitempty"`
	CanonicalizePermittedCNAMEs      string `toml:"CanonicalizePermittedCNAMEs" json:"CanonicalizePermittedCNAMEs,omitempty"`
	ChallengeResponseAuthentication  string `toml:"ChallengeResponseAuthentication" json:"ChallengeResponseAuthentication,omitempty"`
	CheckHostIP                      string `toml:"CheckHostIP" json:"CheckHostIP,omitempty"`
	Cipher                           string `toml:"Cipher" json:"Cipher,omitempty"`
	Ciphers                          string `toml:"Ciphers" json:"Ciphers,omitempty"`
	ClearAllForwardings              string `toml:"ClearAllForwardings" json:"ClearAllForwardings,omitempty"`
	Compression                      string `toml:"Compression" json:"Compression,omitempty"`
	CompressionLevel                 int    `toml:"CompressionLevel" json:"CompressionLevel,omitempty"`
	ConnectionAttempts               string `toml:"ConnectionAttempts" json:"ConnectionAttempts,omitempty"`
	ConnectTimeout                   int    `toml:"ConnectTimeout" json:"ConnectTimeout,omitempty"`
	ControlMaster                    string `toml:"ControlMaster" json:"ControlMaster,omitempty"`
	ControlPath                      string `toml:"ControlPath" json:"ControlPath,omitempty"`
	ControlPersist                   string `toml:"ControlPersist" json:"ControlPersist,omitempty"`
	DynamicForward                   string `toml:"DynamicForward" json:"DynamicForward,omitempty"`
	EnableSSHKeysign                 string `toml:"EnableSSHKeysign" json:"EnableSSHKeysign,omitempty"`
	EscapeChar                       string `toml:"EscapeChar" json:"EscapeChar,omitempty"`
	ExitOnForwardFailure             string `toml:"ExitOnForwardFailure" json:"ExitOnForwardFailure,omitempty"`
	FingerprintHash                  string `toml:"FingerprintHash" json:"FingerprintHash,omitempty"`
	ForwardAgent                     string `toml:"ForwardAgent" json:"ForwardAgent,omitempty"`
	ForwardX11                       string `toml:"ForwardX11" json:"ForwardX11,omitempty"`
	ForwardX11Timeout                int    `toml:"ForwardX11Timeout" json:"ForwardX11Timeout,omitempty"`
	ForwardX11Trusted                string `toml:"ForwardX11Trusted" json:"ForwardX11Trusted,omitempty"`
	GatewayPorts                     string `toml:"GatewayPorts" json:"GatewayPorts,omitempty"`
	GlobalKnownHostsFile             string `toml:"GlobalKnownHostsFile" json:"GlobalKnownHostsFile,omitempty"`
	GSSAPIAuthentication             string `toml:"GSSAPIAuthentication" json:"GSSAPIAuthentication,omitempty"`
	GSSAPIDelegateCredentials        string `toml:"GSSAPIDelegateCredentials" json:"GSSAPIDelegateCredentials,omitempty"`
	HashKnownHosts                   string `toml:"HashKnownHosts" json:"HashKnownHosts,omitempty"`
	HostbasedAuthentication          string `toml:"HostbasedAuthentication" json:"HostbasedAuthentication,omitempty"`
	HostbasedKeyTypes                string `toml:"HostbasedKeyTypes" json:"HostbasedKeyTypes,omitempty"`
	HostKeyAlgorithms                string `toml:"HostKeyAlgorithms" json:"HostKeyAlgorithms,omitempty"`
	HostKeyAlias                     string `toml:"HostKeyAlias" json:"HostKeyAlias,omitempty"`
	HostName                         string `toml:"HostName" json:"HostName,omitempty"`
	IdentitiesOnly                   string `toml:"IdentitiesOnly" json:"IdentitiesOnly,omitempty"`
	IdentityFile                     string `toml:"IdentityFile" json:"IdentityFile,omitempty"`
	IgnoreUnknown                    string `toml:"IgnoreUnknown" json:"IgnoreUnknown,omitempty"`
	IPQoS                            string `toml:"IPQoS" json:"IPQoS,omitempty"`
	KbdInteractiveAuthentication     string `toml:"KbdInteractiveAuthentication" json:"KbdInteractiveAuthentication,omitempty"`
	KbdInteractiveDevices            string `toml:"KbdInteractiveDevices" json:"KbdInteractiveDevices,omitempty"`
	KexAlgorithms                    string `toml:"KexAlgorithms" json:"KexAlgorithms,omitempty"`
	LocalCommand                     string `toml:"LocalCommand" json:"LocalCommand,omitempty"`
	LocalForward                     string `toml:"LocalForward" json:"LocalForward,omitempty"`
	LogLevel                         string `toml:"LogLevel" json:"LogLevel,omitempty"`
	MACs                             string `toml:"MACs" json:"MACs,omitempty"`
	Match                            string `toml:"Match" json:"Match,omitempty"`
	NoHostAuthenticationForLocalhost string `toml:"NoHostAuthenticationForLocalhost" json:"NoHostAuthenticationForLocalhost,omitempty"`
	NumberOfPasswordPrompts          string `toml:"NumberOfPasswordPrompts" json:"NumberOfPasswordPrompts,omitempty"`
	PasswordAuthentication           string `toml:"PasswordAuthentication" json:"PasswordAuthentication,omitempty"`
	PermitLocalCommand               string `toml:"PermitLocalCommand" json:"PermitLocalCommand,omitempty"`
	PKCS11Provider                   string `toml:"PKCS11Provider" json:"PKCS11Provider,omitempty"`
	Port                             uint   `toml:"Port" json:"Port,omitempty"`
	PreferredAuthentications         string `toml:"PreferredAuthentications" json:"PreferredAuthentications,omitempty"`
	Protocol                         string `toml:"Protocol" json:"Protocol,omitempty"`
	ProxyUseFdpass                   string `toml:"ProxyUseFdpass" json:"ProxyUseFdpass,omitempty"`
	PubkeyAuthentication             string `toml:"PubkeyAuthentication" json:"PubkeyAuthentication,omitempty"`
	RekeyLimit                       string `toml:"RekeyLimit" json:"RekeyLimit,omitempty"`
	RemoteForward                    string `toml:"RemoteForward" json:"RemoteForward,omitempty"`
	RequestTTY                       string `toml:"RequestTTY" json:"RequestTTY,omitempty"`
	RevokedHostKeys                  string `toml:"RevokedHostKeys" json:"RevokedHostKeys,omitempty"`
	RhostsRSAAuthentication          string `toml:"RhostsRSAAuthentication" json:"RhostsRSAAuthentication,omitempty"`
	RSAAuthentication                string `toml:"RSAAuthentication" json:"RSAAuthentication,omitempty"`
	SendEnv                          string `toml:"SendEnv" json:"SendEnv,omitempty"`
	ServerAliveCountMax              int    `toml:"ServerAliveCountMax" json:"ServerAliveCountMax,omitempty"`
	ServerAliveInterval              int    `toml:"ServerAliveInterval" json:"ServerAliveInterval,omitempty"`
	StreamLocalBindMask              string `toml:"StreamLocalBindMask" json:"StreamLocalBindMask,omitempty"`
	StreamLocalBindUnlink            string `toml:"StreamLocalBindUnlink" json:"StreamLocalBindUnlink,omitempty"`
	StrictHostKeyChecking            string `toml:"StrictHostKeyChecking" json:"StrictHostKeyChecking,omitempty"`
	TCPKeepAlive                     string `toml:"TCPKeepAlive" json:"TCPKeepAlive,omitempty"`
	Tunnel                           string `toml:"Tunnel" json:"Tunnel,omitempty"`
	TunnelDevice                     string `toml:"TunnelDevice" json:"TunnelDevice,omitempty"`
	UpdateHostKeys                   string `toml:"UpdateHostKeys" json:"UpdateHostKeys,omitempty"`
	UsePrivilegedPort                string `toml:"UsePrivilegedPort" json:"UsePrivilegedPort,omitempty"`
	User                             string `toml:"User" json:"User,omitempty"`
	UserKnownHostsFile               string `toml:"UserKnownHostsFile" json:"UserKnownHostsFile,omitempty"`
	VerifyHostKeyDNS                 string `toml:"VerifyHostKeyDNS" json:"VerifyHostKeyDNS,omitempty"`
	VisualHostKey                    string `toml:"VisualHostKey" json:"VisualHostKey,omitempty"`
	XAuthLocation                    string `toml:"XAuthLocation" json:"XAuthLocation,omitempty"`

	// ssh-config fields with a different behavior
	ProxyCommand string `toml:"ProxyCommand" json:"ProxyCommand,omitempty"`

	// exposed assh fields
	Inherits           []string `toml:"Inherits" json:"Inherits,omitempty"`
	Gateways           []string `toml:"Gateways" json:"Gateways,omitempty"`
	ResolveNameservers []string `toml:"ResolveNameservers" json:"ResolveNameservers,omitempty"`
	ResolveCommand     string   `toml:"ResolveCommand" json:"ResolveCommand,omitempty"`

	// private assh fields
	name      string          `toml:- json:-`
	isDefault bool            `toml:- json:-`
	inherited map[string]bool `toml:- json:-`
}

func (h *Host) Name() string {
	return h.name
}

// ApplyDefaults ensures a Host is valid by filling the missing fields with defaults
func (h *Host) ApplyDefaults(defaults *Host) {
	// ssh-config fields
	if h.AddressFamily == "" {
		h.AddressFamily = defaults.AddressFamily
	}
	if h.BatchMode == "" {
		h.BatchMode = defaults.BatchMode
	}
	if h.BindAddress == "" {
		h.BindAddress = defaults.BindAddress
	}
	if h.CanonicalDomains == "" {
		h.CanonicalDomains = defaults.CanonicalDomains
	}
	if h.CanonicalizeFallbackLocal == "" {
		h.CanonicalizeFallbackLocal = defaults.CanonicalizeFallbackLocal
	}
	if h.CanonicalizeHostname == "" {
		h.CanonicalizeHostname = defaults.CanonicalizeHostname
	}
	if h.CanonicalizeMaxDots == "" {
		h.CanonicalizeMaxDots = defaults.CanonicalizeMaxDots
	}
	if h.CanonicalizePermittedCNAMEs == "" {
		h.CanonicalizePermittedCNAMEs = defaults.CanonicalizePermittedCNAMEs
	}
	if h.ChallengeResponseAuthentication == "" {
		h.ChallengeResponseAuthentication = defaults.ChallengeResponseAuthentication
	}
	if h.CheckHostIP == "" {
		h.CheckHostIP = defaults.CheckHostIP
	}
	if h.Cipher == "" {
		h.Cipher = defaults.Cipher
	}
	if h.Ciphers == "" {
		h.Ciphers = defaults.Ciphers
	}
	if h.ClearAllForwardings == "" {
		h.ClearAllForwardings = defaults.ClearAllForwardings
	}
	if h.Compression == "" {
		h.Compression = defaults.Compression
	}
	if h.CompressionLevel == 0 {
		h.CompressionLevel = defaults.CompressionLevel
	}
	if h.ConnectionAttempts == "" {
		h.ConnectionAttempts = defaults.ConnectionAttempts
	}
	if h.ConnectTimeout == 0 {
		h.ConnectTimeout = defaults.ConnectTimeout
	}
	if h.ControlMaster == "" {
		h.ControlMaster = defaults.ControlMaster
	}
	if h.ControlPath == "" {
		h.ControlPath = defaults.ControlPath
	}
	if h.ControlPersist == "" {
		h.ControlPersist = defaults.ControlPersist
	}
	if h.DynamicForward == "" {
		h.DynamicForward = defaults.DynamicForward
	}
	if h.EnableSSHKeysign == "" {
		h.EnableSSHKeysign = defaults.EnableSSHKeysign
	}
	if h.EscapeChar == "" {
		h.EscapeChar = defaults.EscapeChar
	}
	if h.ExitOnForwardFailure == "" {
		h.ExitOnForwardFailure = defaults.ExitOnForwardFailure
	}
	if h.FingerprintHash == "" {
		h.FingerprintHash = defaults.FingerprintHash
	}
	if h.ForwardAgent == "" {
		h.ForwardAgent = defaults.ForwardAgent
	}
	if h.ForwardX11 == "" {
		h.ForwardX11 = defaults.ForwardX11
	}
	if h.ForwardX11Timeout == 0 {
		h.ForwardX11Timeout = defaults.ForwardX11Timeout
	}
	if h.ForwardX11Trusted == "" {
		h.ForwardX11Trusted = defaults.ForwardX11Trusted
	}
	if h.GatewayPorts == "" {
		h.GatewayPorts = defaults.GatewayPorts
	}
	if h.GlobalKnownHostsFile == "" {
		h.GlobalKnownHostsFile = defaults.GlobalKnownHostsFile
	}
	if h.GSSAPIAuthentication == "" {
		h.GSSAPIAuthentication = defaults.GSSAPIAuthentication
	}
	if h.GSSAPIDelegateCredentials == "" {
		h.GSSAPIDelegateCredentials = defaults.GSSAPIDelegateCredentials
	}
	if h.HashKnownHosts == "" {
		h.HashKnownHosts = defaults.HashKnownHosts
	}
	if h.HostbasedAuthentication == "" {
		h.HostbasedAuthentication = defaults.HostbasedAuthentication
	}
	if h.HostbasedKeyTypes == "" {
		h.HostbasedKeyTypes = defaults.HostbasedKeyTypes
	}
	if h.HostKeyAlgorithms == "" {
		h.HostKeyAlgorithms = defaults.HostKeyAlgorithms
	}
	if h.HostKeyAlias == "" {
		h.HostKeyAlias = defaults.HostKeyAlias
	}
	if h.HostName == "" {
		h.HostName = defaults.HostName
	}
	if h.IdentitiesOnly == "" {
		h.IdentitiesOnly = defaults.IdentitiesOnly
	}
	if h.IdentityFile == "" {
		h.IdentityFile = defaults.IdentityFile
	}
	if h.IgnoreUnknown == "" {
		h.IgnoreUnknown = defaults.IgnoreUnknown
	}
	if h.IPQoS == "" {
		h.IPQoS = defaults.IPQoS
	}
	if h.KbdInteractiveAuthentication == "" {
		h.KbdInteractiveAuthentication = defaults.KbdInteractiveAuthentication
	}
	if h.KbdInteractiveDevices == "" {
		h.KbdInteractiveDevices = defaults.KbdInteractiveDevices
	}
	if h.KexAlgorithms == "" {
		h.KexAlgorithms = defaults.KexAlgorithms
	}
	if h.LocalCommand == "" {
		h.LocalCommand = defaults.LocalCommand
	}
	if h.LocalForward == "" {
		h.LocalForward = defaults.LocalForward
	}
	if h.LogLevel == "" {
		h.LogLevel = defaults.LogLevel
	}
	if h.MACs == "" {
		h.MACs = defaults.MACs
	}
	if h.Match == "" {
		h.Match = defaults.Match
	}
	if h.NoHostAuthenticationForLocalhost == "" {
		h.NoHostAuthenticationForLocalhost = defaults.NoHostAuthenticationForLocalhost
	}
	if h.NumberOfPasswordPrompts == "" {
		h.NumberOfPasswordPrompts = defaults.NumberOfPasswordPrompts
	}
	if h.PasswordAuthentication == "" {
		h.PasswordAuthentication = defaults.PasswordAuthentication
	}
	if h.PermitLocalCommand == "" {
		h.PermitLocalCommand = defaults.PermitLocalCommand
	}
	if h.PKCS11Provider == "" {
		h.PKCS11Provider = defaults.PKCS11Provider
	}
	if h.Port == 0 {
		h.Port = defaults.Port
	}
	if h.PreferredAuthentications == "" {
		h.PreferredAuthentications = defaults.PreferredAuthentications
	}
	if h.Protocol == "" {
		h.Protocol = defaults.Protocol
	}
	if h.ProxyCommand == "" {
		h.ProxyCommand = defaults.ProxyCommand
	}
	if h.ProxyUseFdpass == "" {
		h.ProxyUseFdpass = defaults.ProxyUseFdpass
	}
	if h.PubkeyAuthentication == "" {
		h.PubkeyAuthentication = defaults.PubkeyAuthentication
	}
	if h.RekeyLimit == "" {
		h.RekeyLimit = defaults.RekeyLimit
	}
	if h.RemoteForward == "" {
		h.RemoteForward = defaults.RemoteForward
	}
	if h.RequestTTY == "" {
		h.RequestTTY = defaults.RequestTTY
	}
	if h.RevokedHostKeys == "" {
		h.RevokedHostKeys = defaults.RevokedHostKeys
	}
	if h.RhostsRSAAuthentication == "" {
		h.RhostsRSAAuthentication = defaults.RhostsRSAAuthentication
	}
	if h.RSAAuthentication == "" {
		h.RSAAuthentication = defaults.RSAAuthentication
	}
	if h.SendEnv == "" {
		h.SendEnv = defaults.SendEnv
	}
	if h.ServerAliveCountMax == 0 {
		h.ServerAliveCountMax = defaults.ServerAliveCountMax
	}
	if h.ServerAliveInterval == 6 {
		h.ServerAliveInterval = defaults.ServerAliveInterval
	}
	if h.StreamLocalBindMask == "" {
		h.StreamLocalBindMask = defaults.StreamLocalBindMask
	}
	if h.StreamLocalBindUnlink == "" {
		h.StreamLocalBindUnlink = defaults.StreamLocalBindUnlink
	}
	if h.StrictHostKeyChecking == "" {
		h.StrictHostKeyChecking = defaults.StrictHostKeyChecking
	}
	if h.TCPKeepAlive == "" {
		h.TCPKeepAlive = defaults.TCPKeepAlive
	}
	if h.Tunnel == "" {
		h.Tunnel = defaults.Tunnel
	}
	if h.TunnelDevice == "" {
		h.TunnelDevice = defaults.TunnelDevice
	}
	if h.UpdateHostKeys == "" {
		h.UpdateHostKeys = defaults.UpdateHostKeys
	}
	if h.UsePrivilegedPort == "" {
		h.UsePrivilegedPort = defaults.UsePrivilegedPort
	}
	if h.User == "" {
		h.User = defaults.User
	}
	if h.UserKnownHostsFile == "" {
		h.UserKnownHostsFile = defaults.UserKnownHostsFile
	}
	if h.VerifyHostKeyDNS == "" {
		h.VerifyHostKeyDNS = defaults.VerifyHostKeyDNS
	}
	if h.VisualHostKey == "" {
		h.VisualHostKey = defaults.VisualHostKey
	}
	if h.XAuthLocation == "" {
		h.XAuthLocation = defaults.XAuthLocation
	}

	// ssh-config fields with a different behavior
	if h.ProxyCommand == "" {
		h.ProxyCommand = defaults.ProxyCommand
	}

	// exposed assh fields
	if len(h.ResolveNameservers) == 0 {
		h.ResolveNameservers = defaults.ResolveNameservers
	}
	if h.ResolveCommand == "" {
		h.ResolveCommand = defaults.ResolveCommand
	}
	if len(h.Gateways) == 0 {
		h.Gateways = defaults.Gateways
	}
	if len(h.Inherits) == 0 {
		h.Inherits = defaults.Inherits
	}

	// private assh fields
	// h.inherited = make(map[string]bool, 0)

	// Extra defaults
	if h.Port == 0 {
		h.Port = 22
	}
}

func (h *Host) WriteSshConfigTo(w io.Writer) error {
	fmt.Fprintf(w, "Host %s\n", h.Name())

	// ssh-config fields
	if h.AddressFamily != "" {
		fmt.Fprintf(w, "  AddressFamily %s\n", h.AddressFamily)
	}
	if h.BatchMode != "" {
		fmt.Fprintf(w, "  BatchMode %s\n", h.BatchMode)
	}
	if h.BindAddress != "" {
		fmt.Fprintf(w, "  BindAddress %s\n", h.BindAddress)
	}
	if h.CanonicalDomains != "" {
		fmt.Fprintf(w, "  CanonicalDomains %s\n", h.CanonicalDomains)
	}
	if h.CanonicalizeFallbackLocal != "" {
		fmt.Fprintf(w, "  CanonicalizeFallbackLocal %s\n", h.CanonicalizeFallbackLocal)
	}
	if h.CanonicalizeHostname != "" {
		fmt.Fprintf(w, "  CanonicalizeHostname %s\n", h.CanonicalizeHostname)
	}
	if h.CanonicalizeMaxDots != "" {
		fmt.Fprintf(w, "  CanonicalizeMaxDots %s\n", h.CanonicalizeMaxDots)
	}
	if h.CanonicalizePermittedCNAMEs != "" {
		fmt.Fprintf(w, "  CanonicalizePermittedCNAMEs %s\n", h.CanonicalizePermittedCNAMEs)
	}
	if h.ChallengeResponseAuthentication != "" {
		fmt.Fprintf(w, "  ChallengeResponseAuthentication %s\n", h.ChallengeResponseAuthentication)
	}
	if h.CheckHostIP != "" {
		fmt.Fprintf(w, "  CheckHostIP %s\n", h.CheckHostIP)
	}
	if h.Cipher != "" {
		fmt.Fprintf(w, "  Cipher %s\n", h.Cipher)
	}
	if h.Ciphers != "" {
		fmt.Fprintf(w, "  Ciphers %s\n", h.Ciphers)
	}
	if h.ClearAllForwardings != "" {
		fmt.Fprintf(w, "  ClearAllForwardings %s\n", h.ClearAllForwardings)
	}
	if h.Compression != "" {
		fmt.Fprintf(w, "  Compression %s\n", h.Compression)
	}
	if h.CompressionLevel != 0 {
		fmt.Fprintf(w, "  CompressionLevel %d\n", h.CompressionLevel)
	}
	if h.ConnectionAttempts != "" {
		fmt.Fprintf(w, "  ConnectionAttempts %s\n", h.ConnectionAttempts)
	}
	if h.ConnectTimeout != 0 {
		fmt.Fprintf(w, "  ConnectTimeout %d\n", h.ConnectTimeout)
	}
	if h.ControlMaster != "" {
		fmt.Fprintf(w, "  ControlMaster %s\n", h.ControlMaster)
	}
	if h.ControlPath != "" {
		fmt.Fprintf(w, "  ControlPath %s\n", h.ControlPath)
	}
	if h.ControlPersist != "" {
		fmt.Fprintf(w, "  ControlPersist %s\n", h.ControlPersist)
	}
	if h.DynamicForward != "" {
		fmt.Fprintf(w, "  DynamicForward %s\n", h.DynamicForward)
	}
	if h.EnableSSHKeysign != "" {
		fmt.Fprintf(w, "  EnableSSHKeysign %s\n", h.EnableSSHKeysign)
	}
	if h.EscapeChar != "" {
		fmt.Fprintf(w, "  EscapeChar %s\n", h.EscapeChar)
	}
	if h.ExitOnForwardFailure != "" {
		fmt.Fprintf(w, "  ExitOnForwardFailure %s\n", h.ExitOnForwardFailure)
	}
	if h.FingerprintHash != "" {
		fmt.Fprintf(w, "  FingerprintHash %s\n", h.FingerprintHash)
	}
	if h.ForwardAgent != "" {
		fmt.Fprintf(w, "  ForwardAgent %s\n", h.ForwardAgent)
	}
	if h.ForwardX11 != "" {
		fmt.Fprintf(w, "  ForwardX11 %s\n", h.ForwardX11)
	}
	if h.ForwardX11Timeout != 0 {
		fmt.Fprintf(w, "  ForwardX11Timeout %d\n", h.ForwardX11Timeout)
	}
	if h.ForwardX11Trusted != "" {
		fmt.Fprintf(w, "  ForwardX11Trusted %s\n", h.ForwardX11Trusted)
	}
	if h.GatewayPorts != "" {
		fmt.Fprintf(w, "  GatewayPorts %s\n", h.GatewayPorts)
	}
	if h.GlobalKnownHostsFile != "" {
		fmt.Fprintf(w, "  GlobalKnownHostsFile %s\n", h.GlobalKnownHostsFile)
	}
	if h.GSSAPIAuthentication != "" {
		fmt.Fprintf(w, "  GSSAPIAuthentication %s\n", h.GSSAPIAuthentication)
	}
	if h.GSSAPIDelegateCredentials != "" {
		fmt.Fprintf(w, "  GSSAPIDelegateCredentials %s\n", h.GSSAPIDelegateCredentials)
	}
	if h.HashKnownHosts != "" {
		fmt.Fprintf(w, "  HashKnownHosts %s\n", h.HashKnownHosts)
	}
	if h.HostbasedAuthentication != "" {
		fmt.Fprintf(w, "  HostbasedAuthentication %s\n", h.HostbasedAuthentication)
	}
	if h.HostbasedKeyTypes != "" {
		fmt.Fprintf(w, "  HostbasedKeyTypes %s\n", h.HostbasedKeyTypes)
	}
	if h.HostKeyAlgorithms != "" {
		fmt.Fprintf(w, "  HostKeyAlgorithms %s\n", h.HostKeyAlgorithms)
	}
	if h.HostKeyAlias != "" {
		fmt.Fprintf(w, "  HostKeyAlias %s\n", h.HostKeyAlias)
	}
	if h.HostName != "" {
		fmt.Fprintf(w, "  HostName %s\n", h.HostName)
	}
	if h.IdentitiesOnly != "" {
		fmt.Fprintf(w, "  IdentitiesOnly %s\n", h.IdentitiesOnly)
	}
	if h.IdentityFile != "" {
		fmt.Fprintf(w, "  IdentityFile %s\n", h.IdentityFile)
	}
	if h.IgnoreUnknown != "" {
		fmt.Fprintf(w, "  IgnoreUnknown %s\n", h.IgnoreUnknown)
	}
	if h.IPQoS != "" {
		fmt.Fprintf(w, "  IPQoS %s\n", h.IPQoS)
	}
	if h.KbdInteractiveAuthentication != "" {
		fmt.Fprintf(w, "  KbdInteractiveAuthentication %s\n", h.KbdInteractiveAuthentication)
	}
	if h.KbdInteractiveDevices != "" {
		fmt.Fprintf(w, "  KbdInteractiveDevices %s\n", h.KbdInteractiveDevices)
	}
	if h.KexAlgorithms != "" {
		fmt.Fprintf(w, "  KexAlgorithms %s\n", h.KexAlgorithms)
	}
	if h.LocalCommand != "" {
		fmt.Fprintf(w, "  LocalCommand %s\n", h.LocalCommand)
	}
	if h.LocalForward != "" {
		fmt.Fprintf(w, "  LocalForward %s\n", h.LocalForward)
	}
	if h.LogLevel != "" {
		fmt.Fprintf(w, "  LogLevel %s\n", h.LogLevel)
	}
	if h.MACs != "" {
		fmt.Fprintf(w, "  MACs %s\n", h.MACs)
	}
	if h.Match != "" {
		fmt.Fprintf(w, "  Match %s\n", h.Match)
	}
	if h.NoHostAuthenticationForLocalhost != "" {
		fmt.Fprintf(w, "  NoHostAuthenticationForLocalhost %s\n", h.NoHostAuthenticationForLocalhost)
	}
	if h.NumberOfPasswordPrompts != "" {
		fmt.Fprintf(w, "  NumberOfPasswordPrompts %s\n", h.NumberOfPasswordPrompts)
	}
	if h.PasswordAuthentication != "" {
		fmt.Fprintf(w, "  PasswordAuthentication %s\n", h.PasswordAuthentication)
	}
	if h.PermitLocalCommand != "" {
		fmt.Fprintf(w, "  PermitLocalCommand %s\n", h.PermitLocalCommand)
	}
	if h.PKCS11Provider != "" {
		fmt.Fprintf(w, "  PKCS11Provider %s\n", h.PKCS11Provider)
	}
	if h.Port != 0 {
		fmt.Fprintf(w, "  Port %d\n", h.Port)
	}
	if h.PreferredAuthentications != "" {
		fmt.Fprintf(w, "  PreferredAuthentications %s\n", h.PreferredAuthentications)
	}
	if h.Protocol != "" {
		fmt.Fprintf(w, "  Protocol %s\n", h.Protocol)
	}
	if h.ProxyUseFdpass != "" {
		fmt.Fprintf(w, "  ProxyUseFdpass %s\n", h.ProxyUseFdpass)
	}
	if h.PubkeyAuthentication != "" {
		fmt.Fprintf(w, "  PubkeyAuthentication %s\n", h.PubkeyAuthentication)
	}
	if h.RekeyLimit != "" {
		fmt.Fprintf(w, "  RekeyLimit %s\n", h.RekeyLimit)
	}
	if h.RemoteForward != "" {
		fmt.Fprintf(w, "  RemoteForward %s\n", h.RemoteForward)
	}
	if h.RequestTTY != "" {
		fmt.Fprintf(w, "  RequestTTY %s\n", h.RequestTTY)
	}
	if h.RevokedHostKeys != "" {
		fmt.Fprintf(w, "  RevokedHostKeys %s\n", h.RevokedHostKeys)
	}
	if h.RhostsRSAAuthentication != "" {
		fmt.Fprintf(w, "  RhostsRSAAuthentication %s\n", h.RhostsRSAAuthentication)
	}
	if h.RSAAuthentication != "" {
		fmt.Fprintf(w, "  RSAAuthentication %s\n", h.RSAAuthentication)
	}
	if h.SendEnv != "" {
		fmt.Fprintf(w, "  SendEnv %s\n", h.SendEnv)
	}
	if h.ServerAliveCountMax != 0 {
		fmt.Fprintf(w, "  ServerAliveCountMax %d\n", h.ServerAliveCountMax)
	}
	if h.ServerAliveInterval != 0 {
		fmt.Fprintf(w, "  ServerAliveInterval %d\n", h.ServerAliveInterval)
	}
	if h.StreamLocalBindMask != "" {
		fmt.Fprintf(w, "  StreamLocalBindMask %s\n", h.StreamLocalBindMask)
	}
	if h.StreamLocalBindUnlink != "" {
		fmt.Fprintf(w, "  StreamLocalBindUnlink %s\n", h.StreamLocalBindUnlink)
	}
	if h.StrictHostKeyChecking != "" {
		fmt.Fprintf(w, "  StrictHostKeyChecking %s\n", h.StrictHostKeyChecking)
	}
	if h.TCPKeepAlive != "" {
		fmt.Fprintf(w, "  TCPKeepAlive %s\n", h.TCPKeepAlive)
	}
	if h.Tunnel != "" {
		fmt.Fprintf(w, "  Tunnel %s\n", h.Tunnel)
	}
	if h.TunnelDevice != "" {
		fmt.Fprintf(w, "  TunnelDevice %s\n", h.TunnelDevice)
	}
	if h.UpdateHostKeys != "" {
		fmt.Fprintf(w, "  UpdateHostKeys %s\n", h.UpdateHostKeys)
	}
	if h.UsePrivilegedPort != "" {
		fmt.Fprintf(w, "  UsePrivilegedPort %s\n", h.UsePrivilegedPort)
	}
	if h.User != "" {
		fmt.Fprintf(w, "  User %s\n", h.User)
	}
	if h.UserKnownHostsFile != "" {
		fmt.Fprintf(w, "  UserKnownHostsFile %s\n", h.UserKnownHostsFile)
	}
	if h.VerifyHostKeyDNS != "" {
		fmt.Fprintf(w, "  VerifyHostKeyDNS %s\n", h.VerifyHostKeyDNS)
	}
	if h.VisualHostKey != "" {
		fmt.Fprintf(w, "  VisualHostKey %s\n", h.VisualHostKey)
	}
	if h.XAuthLocation != "" {
		fmt.Fprintf(w, "  XAuthLocation %s\n", h.XAuthLocation)
	}

	// ssh-config fields with a different behavior
	if h.isDefault {
		asshBinary := "assh"
		fmt.Fprintf(w, "  ProxyCommand %s proxy --port=%%p %%h\n", asshBinary)
	} else {
		if h.ProxyCommand != "" {
			fmt.Fprintf(w, "  # ProxyCommand %s\n", h.ProxyCommand)
		}
	}

	// assh fields
	if len(h.Inherits) > 0 {
		fmt.Fprintf(w, "  # Inherits: [%s]\n", strings.Join(h.Inherits, ", "))
	}
	if len(h.Gateways) > 0 {
		fmt.Fprintf(w, "  # Gateways: [%s]\n", strings.Join(h.Gateways, ", "))
	}
	if len(h.ResolveNameservers) > 0 {
		fmt.Fprintf(w, "  # ResolveNameservers: [%s]\n", strings.Join(h.ResolveNameservers, ", "))
	}
	if h.ResolveCommand != "" {
		fmt.Fprintf(w, "  # ResolveCommand: %s\n", h.ResolveCommand)
	}

	return nil
}
