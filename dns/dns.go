// This package extracts information about the current state of the operating
// system's DNS resolution configuration through a combination of performing
// actual hostname look ups and looking at the resolv.conf file.
//
// Note that that this package makes the assumption that you are running on a
// glibc-based system without any custom patches made to the DNS resolution logic.

package dns

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
)

var isTesting bool

// Config is a struct that describes the current DNS configuration for the
// machine. It contains the current configuration of /etc/resolv.conf as well as
// the IP addresses to which certain hostnames resolve.
type Config struct {
	ResolvConf        *ResolvConf
	MetadataServerIPs []string
	RemoteServerIPs   []string
}

// ResolvConf holds the values that are parsed from the resolv.conf file. There
// are a couple of values that are not yet supported. For an explanation of
// these values, see https://linux.die.net/man/5/resolv.conf
type ResolvConf struct {
	Nameservers []string
	Ndots       string
	Timeout     string
	Attempts    string
	Domains     []string
	Search      []string
}

// NewConfig is a helper function to generate the current DNS configuration.
func NewConfig(resolvConf *ResolvConf, lookupMetadata bool) (*Config, error) {
	remoteHostIPs, err := net.LookupHost("en.wikipedia.org")

	if err != nil {
		return nil, fmt.Errorf("could not look up host en.wikipedia.org: %s", err)
	}

	dnsConf := &Config{
		ResolvConf:        resolvConf,
		MetadataServerIPs: nil,
		RemoteServerIPs:   remoteHostIPs,
	}

	if lookupMetadata == true {
		dnsConf.MetadataServerIPs, err = net.LookupHost("metadata.google.com")

		if err != nil {
			return nil, fmt.Errorf("unable to resolve metadata server: %s", err)
		}
	}

	return dnsConf, nil
}

// NewResolvConf is a helper function to generate a new ResolvConf struct from
// the current machines /etc/resolv.conf file.
func NewResolvConf(resolvePath string) (*ResolvConf, error) {
	// Per http://sourceware.org/git/?p=glibc.git;a=blob;f=resolv/resolv.h;h=80a523e5e40982adbe2e6ec761615f46b69817c2;hb=HEAD#l94
	// /etc/resolv.conf is hardcoded to be the default configuration file. This path
	// should not change on any stock glibc-based system. If this errors, the
	// person should know what the heck they are doing.
	if resolvePath == "" {
		resolvePath = "/etc/resolv.conf"
	}

	file, err := ioutil.ReadFile(resolvePath)

	if err != nil {
		return nil, err
	}

	// Create array out of the lines in resolv.conf
	fileString := string(file)
	lineArray := strings.Split(fileString, "\n")

	resolv := &ResolvConf{}

	for _, line := range lineArray {
		option := strings.Split(line, " ")
		if option[0] == "nameserver" {
			resolv.Nameservers = append(resolv.Nameservers, option[1])
		}
		if option[0] == "domain" {
			resolv.Domains = append(resolv.Domains, option[1])
		}
		if option[0] == "search" {
			for _, i := range option[1:] {
				resolv.Search = append(resolv.Search, i)
			}
		}
		if option[0] == "options" {
			val := strings.Split(option[1], ":")
			if val[0] == "ndots" {
				resolv.Ndots = val[1]
			}
			if val[0] == "timeout" {
				resolv.Timeout = val[1]
			}
			// TODO: The number of attempts is silently capped at 5, see
			// http://sourceware.org/git/?p=glibc.git;a=blob;f=resolv/resolv.h;h=80a523e5e40982adbe2e6ec761615f46b69817c2;hb=HEAD#l71
			if val[0] == "attempts" {
				resolv.Attempts = val[1]
			}
		}
	}

	return resolv, nil
}
