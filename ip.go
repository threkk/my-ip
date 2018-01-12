package main

import (
	"context"
	"fmt"
	"net"
)

// OpenDNS1 - Resolver 1 of OpenDNS.
const OpenDNS1 = "208.67.222.222:53"

// OpenDNS2 - Resolver 2 of OpenDNS.
const OpenDNS2 = "208.67.220.220:53"

// GoogleDNS1 - Resolver 1 of Google.
const GoogleDNS1 = "8.8.8.8:53"

// GoogleDNS2 - Resolver 2 of Google.
const GoogleDNS2 = "8.8.4.4:53"

// MyIP TODO
type MyIP struct {
	DNS       string
	PreferIP6 bool
	Ctx       context.Context
}

// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go#23558495

// External TODO
func (myIP *MyIP) External() (string, error) {
	dialer := net.Dialer{}
	d, err := dialer.DialContext(myIP.Ctx, "udp", myIP.DNS)
	defer d.Close()
	if err != nil {
		return "", err
	}

	address := d.LocalAddr().String()
	return address, nil
}

// Internal TODO
func (myIP *MyIP) Internal() ([]string, error) {
	ifaces, err := net.InterfaceAddrs()

	if err != nil {
		return nil, err
	}

	locals := make([]string, 0)
	for _, address := range ifaces {
		if ip, ok := address.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip4 := ip.IP.To4(); ip4 != nil {
				locals = append(locals, ip4.String())
			}

			if ip6 := ip.IP.To16(); ip6 != nil {
				locals = append(locals, ip6.String())
			}
		}
	}

	return locals, nil
}

func openDNSDialer(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{}
	// OpenDNS server ip address.
	return d.DialContext(ctx, "udp", "")
}

func main() {
	myIP := &MyIP{DNS: OpenDNS1, Ctx: context.Background(), PreferIP6: false}

	ext, _ := myIP.External()
	inte, _ := myIP.Internal()

	fmt.Println("External", ext)
	fmt.Println("Internal", inte)
}
