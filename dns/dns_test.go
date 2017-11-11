package dns

import (
	"reflect"
	"testing"
)

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
