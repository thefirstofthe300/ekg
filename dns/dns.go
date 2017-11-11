package dns

import (
	"fmt"
	"net"
)

// DnsConfig is a struct that describes the current DNS configuration for the
// machine. It contains the current configuration of /etc/resolv.conf as well as
// information regarding what IP the metadata server resolves to and the
type DNSConfig struct {
	ResolvConf        *ResolvConf
	MetadataServerIPs []string
	RemoteServerIPs   []string
}

type ResolvConf struct {
	Nameservers []string
	Ndots       int
	Domains     []string
	Search      []string
}

func NewDNSConfig(resolveConf *ResolvConf, lookupMetadata bool) (*DNSConfig, error) {
	remoteHostIPs, err := net.LookupHost("en.wikipedia.org")

	if err != nil {
		return nil, fmt.Errorf("could not look up host en.wikipedia.org: %s", err)
	}

	dnsConf := &DNSConfig{
		ResolvConf:        resolveConf,
		MetadataServerIPs: nil,
		RemoteServerIPs:   remoteHostIPs,
	}

	if lookupMetadata == true {
		dnsConf.MetadataServerIPs, err = net.LookupHost("metadata.google.com")

		if err != nil {
			return nil, fmt.Errorf("unable to resolve metadata server %s:", err)
		}
	}

	return dnsConf, nil
}

func NewResolvConf() (*ResolvConf, error) {
	return &ResolvConf{}, nil
}
