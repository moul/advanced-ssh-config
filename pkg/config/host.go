package config

import (
	"fmt"
	"io"
	"strings"
)

// Host defines the configuration flags of a host
type Host struct {
	// ssh-config fields
	AddressFamily                    string `yaml:"addressfamily,omitempty,flow" json:"AddressFamily,omitempty"`
	AskPassGUI                       string `yaml:"askpassgui,omitempty,flow" json:"AskPassGUI,omitempty"`
	BatchMode                        string `yaml:"batchmode,omitempty,flow" json:"BatchMode,omitempty"`
	BindAddress                      string `yaml:"bindaddress,omitempty,flow" json:"BindAddress,omitempty"`
	CanonicalDomains                 string `yaml:"canonicaldomains,omitempty,flow" json:"CanonicalDomains,omitempty"`
	CanonicalizeFallbackLocal        string `yaml:"canonicalizefallbacklocal,omitempty,flow" json:"CanonicalizeFallbackLocal,omitempty"`
	CanonicalizeHostname             string `yaml:"canonicalizehostname,omitempty,flow" json:"CanonicalizeHostname,omitempty"`
	CanonicalizeMaxDots              string `yaml:"canonicalizemaxDots,omitempty,flow" json:"CanonicalizeMaxDots,omitempty"`
	CanonicalizePermittedCNAMEs      string `yaml:"canonicalizepermittedcnames,omitempty,flow" json:"CanonicalizePermittedCNAMEs,omitempty"`
	ChallengeResponseAuthentication  string `yaml:"challengeresponseauthentication,omitempty,flow" json:"ChallengeResponseAuthentication,omitempty"`
	CheckHostIP                      string `yaml:"checkhostip,omitempty,flow" json:"CheckHostIP,omitempty"`
	Cipher                           string `yaml:"cipher,omitempty,flow" json:"Cipher,omitempty"`
	Ciphers                          string `yaml:"ciphers,omitempty,flow" json:"Ciphers,omitempty"`
	ClearAllForwardings              string `yaml:"clearallforwardings,omitempty,flow" json:"ClearAllForwardings,omitempty"`
	Compression                      string `yaml:"compression,omitempty,flow" json:"Compression,omitempty"`
	CompressionLevel                 int    `yaml:"compressionlevel,omitempty,flow" json:"CompressionLevel,omitempty"`
	ConnectionAttempts               string `yaml:"connectionattempts,omitempty,flow" json:"ConnectionAttempts,omitempty"`
	ConnectTimeout                   int    `yaml:"connecttimeout,omitempty,flow" json:"ConnectTimeout,omitempty"`
	ControlMaster                    string `yaml:"controlmaster,omitempty,flow" json:"ControlMaster,omitempty"`
	ControlPath                      string `yaml:"controlpath,omitempty,flow" json:"ControlPath,omitempty"`
	ControlPersist                   string `yaml:"controlpersist,omitempty,flow" json:"ControlPersist,omitempty"`
	DynamicForward                   string `yaml:"dynamicforward,omitempty,flow" json:"DynamicForward,omitempty"`
	EnableSSHKeysign                 string `yaml:"enablesshkeysign,omitempty,flow" json:"EnableSSHKeysign,omitempty"`
	EscapeChar                       string `yaml:"escapechar,omitempty,flow" json:"EscapeChar,omitempty"`
	ExitOnForwardFailure             string `yaml:"exitonforwardfailure,omitempty,flow" json:"ExitOnForwardFailure,omitempty"`
	FingerprintHash                  string `yaml:"fingerprinthash,omitempty,flow" json:"FingerprintHash,omitempty"`
	ForwardAgent                     string `yaml:"forwardagent,omitempty,flow" json:"ForwardAgent,omitempty"`
	ForwardX11                       string `yaml:"forwardx11,omitempty,flow" json:"ForwardX11,omitempty"`
	ForwardX11Timeout                int    `yaml:"forwardx11timeout,omitempty,flow" json:"ForwardX11Timeout,omitempty"`
	ForwardX11Trusted                string `yaml:"forwardx11trusted,omitempty,flow" json:"ForwardX11Trusted,omitempty"`
	GatewayPorts                     string `yaml:"gatewayports,omitempty,flow" json:"GatewayPorts,omitempty"`
	GlobalKnownHostsFile             string `yaml:"globalknownhostsfile,omitempty,flow" json:"GlobalKnownHostsFile,omitempty"`
	GSSAPIAuthentication             string `yaml:"gssapiauthentication,omitempty,flow" json:"GSSAPIAuthentication,omitempty"`
	GSSAPIClientIdentity             string `yaml:"gssapiclientidentity,omitempty,flow" json:"GSSAPIClientIdentity,omitempty"`
	GSSAPIDelegateCredentials        string `yaml:"gssapidelegatecredentials,omitempty,flow" json:"GSSAPIDelegateCredentials,omitempty"`
	GSSAPIKeyExchange                string `yaml:"gssapikeyexchange,omitempty,flow" json:"GSSAPIKeyExchange,omitempty"`
	GSSAPIRenewalForcesRekey         string `yaml:"gssapirenewalforcesrekey,omitempty,flow" json:"GSSAPIRenewalForcesRekey,omitempty"`
	GSSAPIServerIdentity             string `yaml:"gssapiserveridentity,omitempty,flow" json:"GSSAPIServerIdentity,omitempty"`
	GSSAPITrustDns                   string `yaml:"gssapitrustdns,omitempty,flow" json:"GSSAPITrustDns,omitempty"`
	HashKnownHosts                   string `yaml:"hashknownhosts,omitempty,flow" json:"HashKnownHosts,omitempty"`
	HostbasedAuthentication          string `yaml:"hostbasedauthentication,omitempty,flow" json:"HostbasedAuthentication,omitempty"`
	HostbasedKeyTypes                string `yaml:"hostbasedkeytypes,omitempty,flow" json:"HostbasedKeyTypes,omitempty"`
	HostKeyAlgorithms                string `yaml:"hostkeyalgorithms,omitempty,flow" json:"HostKeyAlgorithms,omitempty"`
	HostKeyAlias                     string `yaml:"hostkeyalias,omitempty,flow" json:"HostKeyAlias,omitempty"`
	IdentitiesOnly                   string `yaml:"identitiesonly,omitempty,flow" json:"IdentitiesOnly,omitempty"`
	IdentityFile                     string `yaml:"identityfile,omitempty,flow" json:"IdentityFile,omitempty"`
	IgnoreUnknown                    string `yaml:"ignoreunknown,omitempty,flow" json:"IgnoreUnknown,omitempty"`
	IPQoS                            string `yaml:"ipqos,omitempty,flow" json:"IPQoS,omitempty"`
	KbdInteractiveAuthentication     string `yaml:"kbdinteractiveauthentication,omitempty,flow" json:"KbdInteractiveAuthentication,omitempty"`
	KbdInteractiveDevices            string `yaml:"kbdinteractivedevices,omitempty,flow" json:"KbdInteractiveDevices,omitempty"`
	KexAlgorithms                    string `yaml:"kexalgorithms,omitempty,flow" json:"KexAlgorithms,omitempty"`
	KeychainIntegration              string `yaml:"keychainintegration,omitempty,flow" json:"KeychainIntegration,omitempty"`
	LocalCommand                     string `yaml:"localcommand,omitempty,flow" json:"LocalCommand,omitempty"`
	LocalForward                     string `yaml:"localforward,omitempty,flow" json:"LocalForward,omitempty"`
	LogLevel                         string `yaml:"loglevel,omitempty,flow" json:"LogLevel,omitempty"`
	MACs                             string `yaml:"macs,omitempty,flow" json:"MACs,omitempty"`
	Match                            string `yaml:"match,omitempty,flow" json:"Match,omitempty"`
	NoHostAuthenticationForLocalhost string `yaml:"nohostauthenticationforlocalhost,omitempty,flow" json:"NoHostAuthenticationForLocalhost,omitempty"`
	NumberOfPasswordPrompts          string `yaml:"numberofpasswordprompts,omitempty,flow" json:"NumberOfPasswordPrompts,omitempty"`
	PasswordAuthentication           string `yaml:"passwordauthentication,omitempty,flow" json:"PasswordAuthentication,omitempty"`
	PermitLocalCommand               string `yaml:"permitlocalcommand,omitempty,flow" json:"PermitLocalCommand,omitempty"`
	PKCS11Provider                   string `yaml:"pkcs11provider,omitempty,flow" json:"PKCS11Provider,omitempty"`
	Port                             string `yaml:"port,omitempty,flow" json:"Port,omitempty"`
	PreferredAuthentications         string `yaml:"preferredauthentications,omitempty,flow" json:"PreferredAuthentications,omitempty"`
	Protocol                         string `yaml:"protocol,omitempty,flow" json:"Protocol,omitempty"`
	ProxyUseFdpass                   string `yaml:"proxyusefdpass,omitempty,flow" json:"ProxyUseFdpass,omitempty"`
	PubkeyAuthentication             string `yaml:"pubkeyauthentication,omitempty,flow" json:"PubkeyAuthentication,omitempty"`
	RekeyLimit                       string `yaml:"rekeylimit,omitempty,flow" json:"RekeyLimit,omitempty"`
	RemoteForward                    string `yaml:"remoteforward,omitempty,flow" json:"RemoteForward,omitempty"`
	RequestTTY                       string `yaml:"requesttty,omitempty,flow" json:"RequestTTY,omitempty"`
	RevokedHostKeys                  string `yaml:"revokedhostkeys,omitempty,flow" json:"RevokedHostKeys,omitempty"`
	RhostsRSAAuthentication          string `yaml:"rhostsrsaauthentication,omitempty,flow" json:"RhostsRSAAuthentication,omitempty"`
	RSAAuthentication                string `yaml:"rsaauthentication,omitempty,flow" json:"RSAAuthentication,omitempty"`
	SendEnv                          string `yaml:"sendenv,omitempty,flow" json:"SendEnv,omitempty"`
	ServerAliveCountMax              int    `yaml:"serveralivecountmax,omitempty,flow" json:"ServerAliveCountMax,omitempty"`
	ServerAliveInterval              int    `yaml:"serveraliveinterval,omitempty,flow" json:"ServerAliveInterval,omitempty"`
	StreamLocalBindMask              string `yaml:"streamlocalbindmask,omitempty,flow" json:"StreamLocalBindMask,omitempty"`
	StreamLocalBindUnlink            string `yaml:"streamlocalbindunlink,omitempty,flow" json:"StreamLocalBindUnlink,omitempty"`
	StrictHostKeyChecking            string `yaml:"stricthostkeychecking,omitempty,flow" json:"StrictHostKeyChecking,omitempty"`
	TCPKeepAlive                     string `yaml:"tcpkeepalive,omitempty,flow" json:"TCPKeepAlive,omitempty"`
	Tunnel                           string `yaml:"tunnel,omitempty,flow" json:"Tunnel,omitempty"`
	TunnelDevice                     string `yaml:"tunneldevice,omitempty,flow" json:"TunnelDevice,omitempty"`
	UpdateHostKeys                   string `yaml:"updatehostkeys,omitempty,flow" json:"UpdateHostKeys,omitempty"`
	UsePrivilegedPort                string `yaml:"useprivilegedport,omitempty,flow" json:"UsePrivilegedPort,omitempty"`
	User                             string `yaml:"user,omitempty,flow" json:"User,omitempty"`
	UserKnownHostsFile               string `yaml:"userknownhostsfile,omitempty,flow" json:"UserKnownHostsFile,omitempty"`
	VerifyHostKeyDNS                 string `yaml:"verifyhostkeydns,omitempty,flow" json:"VerifyHostKeyDNS,omitempty"`
	VisualHostKey                    string `yaml:"visualhostkey,omitempty,flow" json:"VisualHostKey,omitempty"`
	XAuthLocation                    string `yaml:"xauthlocation,omitempty,flow" json:"XAuthLocation,omitempty"`

	// ssh-config fields with a different behavior
	HostName     string `yaml:"hostname,omitempty,flow" json:"HostName,omitempty"`
	ProxyCommand string `yaml:"proxycommand,omitempty,flow" json:"ProxyCommand,omitempty"`

	// exposed assh fields
	Inherits             []string `yaml:"inherits,omitempty,flow" json:"Inherits,omitempty"`
	Gateways             []string `yaml:"gateways,omitempty,flow" json:"Gateways,omitempty"`
	ResolveNameservers   []string `yaml:"resolvenameservers,omitempty,flow" json:"ResolveNameservers,omitempty"`
	ResolveCommand       string   `yaml:"resolvecommand,omitempty,flow" json:"ResolveCommand,omitempty"`
	NoControlMasterMkdir string   `yaml:"nocontrolmastermkdir,omitempty,flow" json:"NoControlMasterMkdir,omitempty"`
	Aliases              []string `yaml:"aliases,omitempty,flow" json:"Aliases,omitempty"`

	// private assh fields
	knownHosts []string        `yaml:"-" json:"-"`
	pattern    string          `yaml:"-" json:"-"`
	name       string          `yaml:"-" json:"-"`
	inputName  string          `yaml:"-" json:"-"`
	isDefault  bool            `yaml:"-" json:"-"`
	isTemplate bool            `yaml:"-" json:"-"`
	inherited  map[string]bool `yaml:"-" json:"-"`
}

// NewHost returns a host with name
func NewHost(name string) *Host {
	return &Host{
		name: name,
	}
}

// Name returns the name of a host
func (h *Host) Name() string {
	return h.name
}

// Clone returns a copy of an existing Host
func (h *Host) Clone() *Host {
	newHost := *h
	return &newHost
}

// ApplyDefaults ensures a Host is valid by filling the missing fields with defaults
func (h *Host) ApplyDefaults(defaults *Host) {
	// ssh-config fields
	if h.AddressFamily == "" {
		h.AddressFamily = defaults.AddressFamily
	}
	h.AddressFamily = expandField(h.AddressFamily)

	if h.AskPassGUI == "" {
		h.AskPassGUI = defaults.AskPassGUI
	}
	h.AskPassGUI = expandField(h.AskPassGUI)

	if h.BatchMode == "" {
		h.BatchMode = defaults.BatchMode
	}
	h.BatchMode = expandField(h.BatchMode)

	if h.BindAddress == "" {
		h.BindAddress = defaults.BindAddress
	}
	h.BindAddress = expandField(h.BindAddress)

	if h.CanonicalDomains == "" {
		h.CanonicalDomains = defaults.CanonicalDomains
	}
	h.CanonicalDomains = expandField(h.CanonicalDomains)

	if h.CanonicalizeFallbackLocal == "" {
		h.CanonicalizeFallbackLocal = defaults.CanonicalizeFallbackLocal
	}
	h.CanonicalizeFallbackLocal = expandField(h.CanonicalizeFallbackLocal)

	if h.CanonicalizeHostname == "" {
		h.CanonicalizeHostname = defaults.CanonicalizeHostname
	}
	h.CanonicalizeHostname = expandField(h.CanonicalizeHostname)

	if h.CanonicalizeMaxDots == "" {
		h.CanonicalizeMaxDots = defaults.CanonicalizeMaxDots
	}
	h.CanonicalizeMaxDots = expandField(h.CanonicalizeMaxDots)

	if h.CanonicalizePermittedCNAMEs == "" {
		h.CanonicalizePermittedCNAMEs = defaults.CanonicalizePermittedCNAMEs
	}
	h.CanonicalizePermittedCNAMEs = expandField(h.CanonicalizePermittedCNAMEs)

	if h.ChallengeResponseAuthentication == "" {
		h.ChallengeResponseAuthentication = defaults.ChallengeResponseAuthentication
	}
	h.ChallengeResponseAuthentication = expandField(h.ChallengeResponseAuthentication)

	if h.CheckHostIP == "" {
		h.CheckHostIP = defaults.CheckHostIP
	}
	h.CheckHostIP = expandField(h.CheckHostIP)

	if h.Cipher == "" {
		h.Cipher = defaults.Cipher
	}
	h.Cipher = expandField(h.Cipher)

	if h.Ciphers == "" {
		h.Ciphers = defaults.Ciphers
	}
	h.Ciphers = expandField(h.Ciphers)

	if h.ClearAllForwardings == "" {
		h.ClearAllForwardings = defaults.ClearAllForwardings
	}
	h.ClearAllForwardings = expandField(h.ClearAllForwardings)

	if h.Compression == "" {
		h.Compression = defaults.Compression
	}
	h.Compression = expandField(h.Compression)

	if h.CompressionLevel == 0 {
		h.CompressionLevel = defaults.CompressionLevel
	}
	// h.CompressionLevel = expandField(h.CompressionLevel)

	if h.ConnectionAttempts == "" {
		h.ConnectionAttempts = defaults.ConnectionAttempts
	}
	h.ConnectionAttempts = expandField(h.ConnectionAttempts)

	if h.ConnectTimeout == 0 {
		h.ConnectTimeout = defaults.ConnectTimeout
	}
	// h.ConnectTimeout = expandField(h.ConnectTimeout)

	if h.ControlMaster == "" {
		h.ControlMaster = defaults.ControlMaster
	}
	h.ControlMaster = expandField(h.ControlMaster)

	if h.ControlPath == "" {
		h.ControlPath = defaults.ControlPath
	}
	h.ControlPath = expandField(h.ControlPath)

	if h.ControlPersist == "" {
		h.ControlPersist = defaults.ControlPersist
	}
	h.ControlPersist = expandField(h.ControlPersist)

	if h.DynamicForward == "" {
		h.DynamicForward = defaults.DynamicForward
	}
	h.DynamicForward = expandField(h.DynamicForward)

	if h.EnableSSHKeysign == "" {
		h.EnableSSHKeysign = defaults.EnableSSHKeysign
	}
	h.EnableSSHKeysign = expandField(h.EnableSSHKeysign)

	if h.EscapeChar == "" {
		h.EscapeChar = defaults.EscapeChar
	}
	h.EscapeChar = expandField(h.EscapeChar)

	if h.ExitOnForwardFailure == "" {
		h.ExitOnForwardFailure = defaults.ExitOnForwardFailure
	}
	h.ExitOnForwardFailure = expandField(h.ExitOnForwardFailure)

	if h.FingerprintHash == "" {
		h.FingerprintHash = defaults.FingerprintHash
	}
	h.FingerprintHash = expandField(h.FingerprintHash)

	if h.ForwardAgent == "" {
		h.ForwardAgent = defaults.ForwardAgent
	}
	h.ForwardAgent = expandField(h.ForwardAgent)

	if h.ForwardX11 == "" {
		h.ForwardX11 = defaults.ForwardX11
	}
	h.ForwardX11 = expandField(h.ForwardX11)

	if h.ForwardX11Timeout == 0 {
		h.ForwardX11Timeout = defaults.ForwardX11Timeout
	}
	// h.ForwardX11Timeout = expandField(h.ForwardX11Timeout)

	if h.ForwardX11Trusted == "" {
		h.ForwardX11Trusted = defaults.ForwardX11Trusted
	}
	h.ForwardX11Trusted = expandField(h.ForwardX11Trusted)

	if h.GatewayPorts == "" {
		h.GatewayPorts = defaults.GatewayPorts
	}
	h.GatewayPorts = expandField(h.GatewayPorts)

	if h.GlobalKnownHostsFile == "" {
		h.GlobalKnownHostsFile = defaults.GlobalKnownHostsFile
	}
	h.GlobalKnownHostsFile = expandField(h.GlobalKnownHostsFile)

	if h.GSSAPIAuthentication == "" {
		h.GSSAPIAuthentication = defaults.GSSAPIAuthentication
	}
	h.GSSAPIAuthentication = expandField(h.GSSAPIAuthentication)

	if h.GSSAPIClientIdentity == "" {
		h.GSSAPIClientIdentity = defaults.GSSAPIClientIdentity
	}
	h.GSSAPIClientIdentity = expandField(h.GSSAPIClientIdentity)

	if h.GSSAPIDelegateCredentials == "" {
		h.GSSAPIDelegateCredentials = defaults.GSSAPIDelegateCredentials
	}
	h.GSSAPIDelegateCredentials = expandField(h.GSSAPIDelegateCredentials)

	if h.GSSAPIKeyExchange == "" {
		h.GSSAPIKeyExchange = defaults.GSSAPIKeyExchange
	}
	h.GSSAPIKeyExchange = expandField(h.GSSAPIKeyExchange)

	if h.GSSAPIRenewalForcesRekey == "" {
		h.GSSAPIRenewalForcesRekey = defaults.GSSAPIRenewalForcesRekey
	}
	h.GSSAPIRenewalForcesRekey = expandField(h.GSSAPIRenewalForcesRekey)

	if h.GSSAPIServerIdentity == "" {
		h.GSSAPIServerIdentity = defaults.GSSAPIServerIdentity
	}
	h.GSSAPIServerIdentity = expandField(h.GSSAPIServerIdentity)

	if h.GSSAPITrustDns == "" {
		h.GSSAPITrustDns = defaults.GSSAPITrustDns
	}
	h.GSSAPITrustDns = expandField(h.GSSAPITrustDns)

	if h.HashKnownHosts == "" {
		h.HashKnownHosts = defaults.HashKnownHosts
	}
	h.HashKnownHosts = expandField(h.HashKnownHosts)

	if h.HostbasedAuthentication == "" {
		h.HostbasedAuthentication = defaults.HostbasedAuthentication
	}
	h.HostbasedAuthentication = expandField(h.HostbasedAuthentication)

	if h.HostbasedKeyTypes == "" {
		h.HostbasedKeyTypes = defaults.HostbasedKeyTypes
	}
	h.HostbasedKeyTypes = expandField(h.HostbasedKeyTypes)

	if h.HostKeyAlgorithms == "" {
		h.HostKeyAlgorithms = defaults.HostKeyAlgorithms
	}
	h.HostKeyAlgorithms = expandField(h.HostKeyAlgorithms)

	if h.HostKeyAlias == "" {
		h.HostKeyAlias = defaults.HostKeyAlias
	}
	h.HostKeyAlias = expandField(h.HostKeyAlias)

	if h.HostName == "" {
		h.HostName = defaults.HostName
	}
	h.HostName = expandField(h.HostName)

	if h.IdentitiesOnly == "" {
		h.IdentitiesOnly = defaults.IdentitiesOnly
	}
	h.IdentitiesOnly = expandField(h.IdentitiesOnly)

	if h.IdentityFile == "" {
		h.IdentityFile = defaults.IdentityFile
	}
	h.IdentityFile = expandField(h.IdentityFile)

	if h.IgnoreUnknown == "" {
		h.IgnoreUnknown = defaults.IgnoreUnknown
	}
	h.IgnoreUnknown = expandField(h.IgnoreUnknown)

	if h.IPQoS == "" {
		h.IPQoS = defaults.IPQoS
	}
	h.IPQoS = expandField(h.IPQoS)

	if h.KbdInteractiveAuthentication == "" {
		h.KbdInteractiveAuthentication = defaults.KbdInteractiveAuthentication
	}
	h.KbdInteractiveAuthentication = expandField(h.KbdInteractiveAuthentication)

	if h.KbdInteractiveDevices == "" {
		h.KbdInteractiveDevices = defaults.KbdInteractiveDevices
	}
	h.KbdInteractiveDevices = expandField(h.KbdInteractiveDevices)

	if h.KexAlgorithms == "" {
		h.KexAlgorithms = defaults.KexAlgorithms
	}
	h.KexAlgorithms = expandField(h.KexAlgorithms)

	if h.KeychainIntegration == "" {
		h.KeychainIntegration = defaults.KeychainIntegration
	}
	h.KeychainIntegration = expandField(h.KeychainIntegration)

	if h.LocalCommand == "" {
		h.LocalCommand = defaults.LocalCommand
	}
	h.LocalCommand = expandField(h.LocalCommand)

	if h.LocalForward == "" {
		h.LocalForward = defaults.LocalForward
	}
	h.LocalForward = expandField(h.LocalForward)

	if h.LogLevel == "" {
		h.LogLevel = defaults.LogLevel
	}
	h.LogLevel = expandField(h.LogLevel)

	if h.MACs == "" {
		h.MACs = defaults.MACs
	}
	h.MACs = expandField(h.MACs)

	if h.Match == "" {
		h.Match = defaults.Match
	}
	h.Match = expandField(h.Match)

	if h.NoHostAuthenticationForLocalhost == "" {
		h.NoHostAuthenticationForLocalhost = defaults.NoHostAuthenticationForLocalhost
	}
	h.NoHostAuthenticationForLocalhost = expandField(h.NoHostAuthenticationForLocalhost)

	if h.NumberOfPasswordPrompts == "" {
		h.NumberOfPasswordPrompts = defaults.NumberOfPasswordPrompts
	}
	h.NumberOfPasswordPrompts = expandField(h.NumberOfPasswordPrompts)

	if h.PasswordAuthentication == "" {
		h.PasswordAuthentication = defaults.PasswordAuthentication
	}
	h.PasswordAuthentication = expandField(h.PasswordAuthentication)

	if h.PermitLocalCommand == "" {
		h.PermitLocalCommand = defaults.PermitLocalCommand
	}
	h.PermitLocalCommand = expandField(h.PermitLocalCommand)

	if h.PKCS11Provider == "" {
		h.PKCS11Provider = defaults.PKCS11Provider
	}
	h.PKCS11Provider = expandField(h.PKCS11Provider)

	if h.Port == "" {
		h.Port = defaults.Port
	}
	h.Port = expandField(h.Port)

	if h.PreferredAuthentications == "" {
		h.PreferredAuthentications = defaults.PreferredAuthentications
	}
	h.PreferredAuthentications = expandField(h.PreferredAuthentications)

	if h.Protocol == "" {
		h.Protocol = defaults.Protocol
	}
	h.Protocol = expandField(h.Protocol)

	if h.ProxyCommand == "" {
		h.ProxyCommand = defaults.ProxyCommand
	}
	h.ProxyCommand = expandField(h.ProxyCommand)

	if h.ProxyUseFdpass == "" {
		h.ProxyUseFdpass = defaults.ProxyUseFdpass
	}
	h.ProxyUseFdpass = expandField(h.ProxyUseFdpass)

	if h.PubkeyAuthentication == "" {
		h.PubkeyAuthentication = defaults.PubkeyAuthentication
	}
	h.PubkeyAuthentication = expandField(h.PubkeyAuthentication)

	if h.RekeyLimit == "" {
		h.RekeyLimit = defaults.RekeyLimit
	}
	h.RekeyLimit = expandField(h.RekeyLimit)

	if h.RemoteForward == "" {
		h.RemoteForward = defaults.RemoteForward
	}
	h.RemoteForward = expandField(h.RemoteForward)

	if h.RequestTTY == "" {
		h.RequestTTY = defaults.RequestTTY
	}
	h.RequestTTY = expandField(h.RequestTTY)

	if h.RevokedHostKeys == "" {
		h.RevokedHostKeys = defaults.RevokedHostKeys
	}
	h.RevokedHostKeys = expandField(h.RevokedHostKeys)

	if h.RhostsRSAAuthentication == "" {
		h.RhostsRSAAuthentication = defaults.RhostsRSAAuthentication
	}
	h.RhostsRSAAuthentication = expandField(h.RhostsRSAAuthentication)

	if h.RSAAuthentication == "" {
		h.RSAAuthentication = defaults.RSAAuthentication
	}
	h.RSAAuthentication = expandField(h.RSAAuthentication)

	if h.SendEnv == "" {
		h.SendEnv = defaults.SendEnv
	}
	h.SendEnv = expandField(h.SendEnv)

	if h.ServerAliveCountMax == 0 {
		h.ServerAliveCountMax = defaults.ServerAliveCountMax
	}
	// h.ServerAliveCountMax = expandField(h.ServerAliveCountMax)

	if h.ServerAliveInterval == 0 {
		h.ServerAliveInterval = defaults.ServerAliveInterval
	}
	// h.ServerAliveInterval = expandField(h.ServerAliveInterval)

	if h.StreamLocalBindMask == "" {
		h.StreamLocalBindMask = defaults.StreamLocalBindMask
	}
	h.StreamLocalBindMask = expandField(h.StreamLocalBindMask)

	if h.StreamLocalBindUnlink == "" {
		h.StreamLocalBindUnlink = defaults.StreamLocalBindUnlink
	}
	h.StreamLocalBindUnlink = expandField(h.StreamLocalBindUnlink)

	if h.StrictHostKeyChecking == "" {
		h.StrictHostKeyChecking = defaults.StrictHostKeyChecking
	}
	h.StrictHostKeyChecking = expandField(h.StrictHostKeyChecking)

	if h.TCPKeepAlive == "" {
		h.TCPKeepAlive = defaults.TCPKeepAlive
	}
	h.TCPKeepAlive = expandField(h.TCPKeepAlive)

	if h.Tunnel == "" {
		h.Tunnel = defaults.Tunnel
	}
	h.Tunnel = expandField(h.Tunnel)

	if h.TunnelDevice == "" {
		h.TunnelDevice = defaults.TunnelDevice
	}
	h.TunnelDevice = expandField(h.TunnelDevice)

	if h.UpdateHostKeys == "" {
		h.UpdateHostKeys = defaults.UpdateHostKeys
	}
	h.UpdateHostKeys = expandField(h.UpdateHostKeys)

	if h.UsePrivilegedPort == "" {
		h.UsePrivilegedPort = defaults.UsePrivilegedPort
	}
	h.UsePrivilegedPort = expandField(h.UsePrivilegedPort)

	if h.User == "" {
		h.User = defaults.User
	}
	h.User = expandField(h.User)

	if h.UserKnownHostsFile == "" {
		h.UserKnownHostsFile = defaults.UserKnownHostsFile
	}
	h.UserKnownHostsFile = expandField(h.UserKnownHostsFile)

	if h.VerifyHostKeyDNS == "" {
		h.VerifyHostKeyDNS = defaults.VerifyHostKeyDNS
	}
	h.VerifyHostKeyDNS = expandField(h.VerifyHostKeyDNS)

	if h.VisualHostKey == "" {
		h.VisualHostKey = defaults.VisualHostKey
	}
	h.VisualHostKey = expandField(h.VisualHostKey)

	if h.XAuthLocation == "" {
		h.XAuthLocation = defaults.XAuthLocation
	}
	h.XAuthLocation = expandField(h.XAuthLocation)

	// ssh-config fields with a different behavior
	if h.ProxyCommand == "" {
		h.ProxyCommand = defaults.ProxyCommand
	}
	h.ProxyCommand = expandField(h.ProxyCommand)

	// exposed assh fields
	if len(h.ResolveNameservers) == 0 {
		h.ResolveNameservers = defaults.ResolveNameservers
	}
	// h.ResolveNameservers = expandField(h.ResolveNameservers)

	if h.ResolveCommand == "" {
		h.ResolveCommand = defaults.ResolveCommand
	}
	h.ResolveCommand = expandField(h.ResolveCommand)

	if h.NoControlMasterMkdir == "" {
		h.NoControlMasterMkdir = defaults.NoControlMasterMkdir
	}
	h.NoControlMasterMkdir = expandField(h.NoControlMasterMkdir)

	if len(h.Gateways) == 0 {
		h.Gateways = defaults.Gateways
	}
	// h.Gateways = expandField(h.Gateways)

	if len(h.Aliases) == 0 {
		h.Aliases = defaults.Aliases
	}

	if len(h.Inherits) == 0 {
		h.Inherits = defaults.Inherits
	}
	// h.Inherits = expandField(h.Inherits)

	// private assh fields
	// h.inherited = make(map[string]bool, 0)
	if h.inputName == "" {
		h.inputName = h.name
	}
	h.inputName = expandField(h.inputName)

	// Extra defaults
	if h.Port == "" {
		h.Port = "22"
	}
}

// AddKnownHost append target to the host' known hosts list
func (h *Host) AddKnownHost(target string) {
	h.knownHosts = append(h.knownHosts, target)
}

// WriteSshConfigTo writes an ~/.ssh/config file compatible host definition to a writable stream
func (h *Host) WriteSshConfigTo(w io.Writer) error {
	aliases := append([]string{h.Name()}, h.Aliases...)
	aliases = append(aliases, h.knownHosts...)
	aliasIdx := 0
	for _, alias := range aliases {
		// FIXME: skip complex patterns

		if aliasIdx > 0 {
			fmt.Fprint(w, "\n")
		}

		fmt.Fprintf(w, "Host %s\n", alias)

		// ssh-config fields
		if h.AddressFamily != "" {
			fmt.Fprintf(w, "  AddressFamily %s\n", h.AddressFamily)
		}
		if h.AskPassGUI != "" {
			fmt.Fprintf(w, "  AskPassGUI %s\n", h.AskPassGUI)
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
		if h.GSSAPIClientIdentity != "" {
			fmt.Fprintf(w, "  GSSAPIClientIdentity %s\n", h.GSSAPIClientIdentity)
		}
		if h.GSSAPIDelegateCredentials != "" {
			fmt.Fprintf(w, "  GSSAPIDelegateCredentials %s\n", h.GSSAPIDelegateCredentials)
		}
		if h.GSSAPIKeyExchange != "" {
			fmt.Fprintf(w, "  GSSAPIKeyExchange %s\n", h.GSSAPIKeyExchange)
		}
		if h.GSSAPIRenewalForcesRekey != "" {
			fmt.Fprintf(w, "  GSSAPIRenewalForcesRekey %s\n", h.GSSAPIRenewalForcesRekey)
		}
		if h.GSSAPIServerIdentity != "" {
			fmt.Fprintf(w, "  GSSAPIServerIdentity %s\n", h.GSSAPIServerIdentity)
		}
		if h.GSSAPITrustDns != "" {
			fmt.Fprintf(w, "  GSSAPITrustDns %s\n", h.GSSAPITrustDns)
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
		if h.KeychainIntegration != "" {
			fmt.Fprintf(w, "  KeychainIntegration %s\n", h.KeychainIntegration)
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
		if h.Port != "" {
			fmt.Fprintf(w, "  Port %s\n", h.Port)
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
			fmt.Fprintf(w, "  ProxyCommand %s proxy --port=%%p %%h\n", ASSHBinary)
		} else {
			if h.ProxyCommand != "" {
				fmt.Fprintf(w, "  # ProxyCommand %s\n", h.ProxyCommand)
			}
		}

		// assh fields
		if h.HostName != "" {
			fmt.Fprintf(w, "  # HostName: %s\n", h.HostName)
		}
		if BoolVal(h.NoControlMasterMkdir) {
			fmt.Fprintf(w, "  # NoControlMasterMkdir: true\n")
		}
		if len(h.Inherits) > 0 {
			fmt.Fprintf(w, "  # Inherits: [%s]\n", strings.Join(h.Inherits, ", "))
		}
		if len(h.Gateways) > 0 {
			fmt.Fprintf(w, "  # Gateways: [%s]\n", strings.Join(h.Gateways, ", "))
		}
		if len(h.Aliases) > 0 {
			if aliasIdx == 0 {
				fmt.Fprintf(w, "  # Aliases: [%s]\n", strings.Join(h.Aliases, ", "))
			} else {
				fmt.Fprintf(w, "  # AliasOf: %s\n", h.Name())
			}
		}
		if len(h.knownHosts) > 0 {
			if aliasIdx == 0 {
				fmt.Fprintf(w, "  # KnownHosts: [%s]\n", strings.Join(h.knownHosts, ", "))
			} else {
				fmt.Fprintf(w, "  # KnownHostOf: %s\n", h.Name())
			}
		}

		if len(h.ResolveNameservers) > 0 {
			fmt.Fprintf(w, "  # ResolveNameservers: [%s]\n", strings.Join(h.ResolveNameservers, ", "))
		}
		if h.ResolveCommand != "" {
			fmt.Fprintf(w, "  # ResolveCommand: %s\n", h.ResolveCommand)
		}

		aliasIdx++
	}
	return nil
}

func (h *Host) ExpandString(input string) string {
	output := input

	// name of the host in config
	output = strings.Replace(output, "%name", h.Name(), -1)

	// original target host name specified on the command line
	output = strings.Replace(output, "%n", h.inputName, -1)

	// target host name
	output = strings.Replace(output, "%h", h.HostName, -1)

	// port
	output = strings.Replace(output, "%p", fmt.Sprintf("%s", h.Port), -1)

	// FIXME: add
	//   %L -> first component of the local host name
	//   %l -> local host name
	//   %r -> remote login username
	//   %u -> username of the user running assh
	//   %r -> remote login username

	return output
}
