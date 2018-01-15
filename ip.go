package main

import (
	"context"
	"fmt"
	"net"
)

const (
	// OpenDNS1 Main resolver address of OpenDNS.
	OpenDNS1 = "208.67.222.222:53"

	// OpenDNS2 Secondary resolver address of OpenDNS.
	OpenDNS2 = "208.67.220.220:53"

	// MyIP Route to resolve in order to get the external ip address.
	MyIP = "myip.opendns.com"
)

var (
	// PreferIP6  If given the option, it should return the ipv6 over ipv4.
	PreferIP6 = false
)

// openDNS1Dialer Dialer pointing to the DNS1 from OpenDNS.
func openDNS1Dialer(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{}
	return d.DialContext(ctx, "udp", OpenDNS1)
}

// openDNS2Dialer Dialer pointing to the DNS2 from OpenDNS.
func openDNS2Dialer(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{}
	return d.DialContext(ctx, "udp", OpenDNS2)
}

// External Retrieves the external IP address for the host by dialing to the
// OpenDNS DNS server. Upon resolving the ip address, it will return the sender's
// ip address.
func External() ([]string, error) {
	externals := make([]string, 0)
	ctxbg := context.Background()
	resolver := &net.Resolver{
		Dial:     openDNS1Dialer,
		PreferGo: true,
	}

	ips, err := resolver.LookupIPAddr(ctxbg, "myip.opendns.com")
	if err != nil {
		resolver.Dial = openDNS2Dialer
		ips, err = resolver.LookupIPAddr(ctxbg, "myip.opendns.com")

		if err != nil {
			return externals, err
		}
	}

	for _, ip := range ips {
		externals = append(externals, ip.String())
	}

	return externals, nil
}

// Internal TODO
func Internal() ([]string, error) {
	ifaces, err := net.InterfaceAddrs()

	if err != nil {
		return nil, err
	}

	locals := make([]string, 0)
	for _, address := range ifaces {
		ip, ok := address.(*net.IPNet)
		if ok && (!ip.IP.IsLoopback()) {
			ip4 := ip.IP.To4()
			ip6 := ip.IP.To16()

			if PreferIP6 && ip6 != nil {
				locals = append(locals, ip6.String())
			} else if ip4 != nil {
				locals = append(locals, ip4.String())
			}
		}
	}

	return locals, nil
}

func main() {
	PreferIP6 = true
	ext, _ := External()
	inte, _ := Internal()
	fmt.Println("External", ext)
	fmt.Println("Internal", inte)
}
