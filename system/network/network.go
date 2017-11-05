package dns

import (
  "net"
)

type DnsConfig struct {
  ResolveConfConfig *ResolvConf
  MetadataServer int

}

type ResolvConf struct {
  Nameservers []string
  Ndots int
  Domains []string
  Search []string
}

func
