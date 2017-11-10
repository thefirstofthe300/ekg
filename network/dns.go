package dns

import (
	"fmt"
	"net"
)

// DnsConfig is a struct that describes the current DNS configuration for the
// machine. It contains the current configuration of /etc/resolv.conf as well as
// information regarding whether certain important servers are accessible.
type DNSConfig struct {
	ResolvConf     *ResolvConf
	MetadataServer []string
	RemoteServer   []string
}

type ResolvConf struct {
	Nameservers []string
	Ndots       int
	Domains     []string
	Search      []string
}

func NewDNSConfig(resolveConf *ResolvConf, remoteHost string, lookupMetadata bool) (*DNSConfig, error) {
	remoteHostIPs, err := net.LookupHost(remoteHost)

	if err != nil {
		return nil, fmt.Errorf("could not look up host %s: %s", remoteHost, err)
	}

	dnsConf := &DNSConfig{
		ResolvConf:     resolveConf,
		MetadataServer: nil,
		RemoteServer:   remoteHostIPs,
	}

	if lookupMetadata == true {
		dnsConf.MetadataServer, err = net.LookupHost("metadata.google.com")

		if err != nil {
			return nil, fmt.Errorf("unable to resolve metadata server %s:", err)
		}
	}

	return dnsConf, nil
}

func NewResolveConf() (*ResolvConf, error) {
	return &ResolvConf{}, nil
}
