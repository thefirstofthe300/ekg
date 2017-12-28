package dns

import (
	"reflect"
	"testing"
)

func init() {
	isTesting = true
}

func TestNewConfig(t *testing.T) {
	expectedConf, err := NewConfig(&ResolvConf{}, false)

	if err != nil {
		t.Fatalf("creating new expected DNS config structs failed: %s", err)
	}

	dnsConf, err := NewConfig(&ResolvConf{}, false)

	if err != nil {
		t.Fatalf("creating new DNS config structs failed: %s", err)
	}

	if reflect.DeepEqual(expectedConf, dnsConf) != true {
		t.Fatalf("DNS config structs did not match: %s", err)
	}
}

func TestNewResolvConf(t *testing.T) {
	expectedConf := &ResolvConf{
		Nameservers: []string{"192.168.0.100", "8.8.4.4"},
		Domains:     []string{"domain.com"},
		Search:      []string{"example.com", "company.net"},
		Ndots:       "5",
		Attempts:    "4",
		Timeout:     "3",
	}

	actualConf, err := NewResolvConf("testdata/resolve.conf")
	if err.Error() != "open testdata/resolve.conf: no such file or directory" {
		t.Fatalf("creating new ResolvConf from non-existant file didn't error as expected: %s", err)
	}

	actualConf, err = NewResolvConf("testdata/resolv.conf")

	if reflect.DeepEqual(actualConf, expectedConf) != true {
		t.Fatalf("resolv config structs did not match: %s", err)
	}
}
