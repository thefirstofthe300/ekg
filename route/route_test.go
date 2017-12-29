package route

import (
	"net"
	"testing"
)

func TestNewTable(t *testing.T) {
	expected := &Table{Routes: []*Route{
		&Route{
			Destination: &net.IP{byte(0), byte(0), byte(0), byte(0)},
			Gateway:     &net.IP{byte(192), byte(168), byte(50), byte(1)},
			Mask:        &net.IP{byte(0), byte(0), byte(0), byte(0)},
			Interface:   "wlp2s0",
			Metric:      "600",
		},
		&Route{
			Destination: &net.IP{byte(169), byte(254), byte(0), byte(0)},
			Gateway:     &net.IP{byte(0), byte(0), byte(0), byte(0)},
			Mask:        &net.IP{byte(255), byte(255), byte(0), byte(0)},
			Interface:   "wlp2s0",
			Metric:      "1000",
		},
		&Route{
			Destination: &net.IP{byte(192), byte(168), byte(50), byte(0)},
			Gateway:     &net.IP{byte(0), byte(0), byte(0), byte(0)},
			Mask:        &net.IP{byte(255), byte(255), byte(255), byte(0)},
			Interface:   "wlp2s0",
			Metric:      "600",
		},
	}}

	actual, err := NewTable("testdata/routefile")

	if err != nil {
		t.Fatalf("failed to generate route table from file: %s", err)
	}

	for index, item := range actual.Routes {
		if !item.Destination.Equal(*expected.Routes[index].Destination) {
			t.Error("generated destination", item.Destination, "for route", item, "not equal to expected", expected.Routes[index].Destination)
		}
		if !item.Gateway.Equal(*expected.Routes[index].Gateway) {
			t.Error("generated gateway", item.Gateway, "for route", item, "not equal to expected", expected.Routes[index].Gateway)
		}
		if !item.Mask.Equal(*expected.Routes[index].Mask) {
			t.Error("generated mask", item.Mask, "for route", item, "not equal to expected", expected.Routes[index].Mask)
		}
		if item.Interface != expected.Routes[index].Interface {
			t.Error("generated interface", item.Interface, "for route", item, "not equal to expected", expected.Routes[index].Interface)
		}
		if item.Metric != expected.Routes[index].Metric {
			t.Error("generated metric", item.Metric, "for route", item, "not equal to expected", expected.Routes[index].Metric)
		}
	}
}

func TestHexIPtoIPAddr(t *testing.T) {
	inputs := []struct {
		hexIP    string
		expected *net.IP
	}{
		{"00000000", &net.IP{byte(0), byte(0), byte(0), byte(0)}},
		{"0000FEA9", &net.IP{byte(169), byte(254), byte(0), byte(0)}},
		{"0000FEA9", &net.IP{byte(169), byte(254), byte(0), byte(0)}},
	}

	errInputs := []struct {
		hexIP    string
		expected string
	}{
		{"0000GEA9", "encoding/hex: invalid byte: U+0047 'G'"},
	}

	for _, item := range inputs {
		convertedIP, err := hexIPtoIPAddr(item.hexIP)
		if err != nil {
			t.Fatalf("unable to decode hex IP address string: %s", err)
		}
		if !convertedIP.Equal(*item.expected) {
			t.Errorf("converting hex IP %s to IP had unexpected result: %s", item.expected, convertedIP)
		}
	}

	for _, item := range errInputs {
		convertedIP, err := hexIPtoIPAddr(item.hexIP)
		if err == nil {
			t.Fatalf("able to decode incorrect hex address: %s", convertedIP)
		}
		message := err.Error()
		if message != item.expected {
			t.Errorf("converting hex IP %s to IP had unexpected error message: %s", item.expected, message)
		}
	}
}
