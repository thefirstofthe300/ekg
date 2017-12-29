package route

import (
	"bufio"
	"encoding/hex"
	"io"
	"net"
	"os"
	"strings"
)

// NewTable generates a routing table from the specified route file. By default,
// this is /proc/net/route
func NewTable(routeFile string) (*Table, error) {
	table := &Table{}
	fileHandle, err := os.Open(routeFile)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewReader(fileHandle)

	for {
		line, err := scanner.ReadString('\n')
		if err == io.EOF {
			break
		}
		if strings.HasPrefix(line, "Iface") {
			continue
		}
		fields := strings.Fields(line)

		iface := fields[0]
		destIP, err := hexIPtoIPAddr(fields[1])
		if err != nil {
			return nil, err
		}
		gatewayIP, err := hexIPtoIPAddr(fields[2])
		if err != nil {
			return nil, err
		}
		metric := fields[6]
		mask, err := hexIPtoIPAddr(fields[7])
		if err != nil {
			return nil, err
		}
		table.Add(&Route{
			Interface:   iface,
			Gateway:     gatewayIP,
			Metric:      metric,
			Destination: destIP,
			Mask:        mask,
		})
	}
	return table, nil
}

// Table acts as a structure to hold the current routing table information
// for the machine. Currently, that is limited to an array of routes.
type Table struct {
	Routes []*Route
}

// Add a route to the routing table
func (rt *Table) Add(r *Route) {
	rt.Routes = append(rt.Routes, r)
}

// Route is a single route from the machines routing table and includes the
// interface, destination IP, the net mask, the associated gateway, and weight
// metric.
type Route struct {
	Interface   string
	Destination *net.IP
	Gateway     *net.IP
	Mask        *net.IP
	Metric      string
}

func hexIPtoIPAddr(hexIP string) (*net.IP, error) {
	decoded, err := hex.DecodeString(hexIP)

	if err != nil {
		return nil, err
	}

	ip := net.IPv4(decoded[3], decoded[2], decoded[1], decoded[0])
	return &ip, nil
}
