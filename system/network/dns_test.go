package dns

import (
	"reflect"
	"testing"
)

func TestNewDNSConfig(t *testing.T) {
	expectedConf, err := &NewDNSConfig(&ResolvConf{}, "www.google.com", false)

	dnsConf, err := NewDNSConfig(&ResolvConf, "www.google.com", false)

	if err != nil {
		t.Fatalf("creating new DNS config structs failed: %s", err)
	}

	if reflect.DeepEqual(expectedConf, NewDNSConfig(&ResolvConf, "www.google.com", false)) != true {
		return t.Fatalf("DNS config structs did not match: %s", err)
	}
}
