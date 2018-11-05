package myip

import (
	"context"
	"net"
)

const (
	// OpenDNS1 Main resolver address of OpenDNS.
	OpenDNS1 = "208.67.222.222:53"

	// OpenDNS2 Secondary resolver address of OpenDNS.
	OpenDNS2 = "208.67.220.220:53"

	// MyIP Route to resolve in order to get the public ip address.
	MyIP = "myip.opendns.com"
)

var (
	// PreferIPv6  If given the option, it should return the ipv6 over ipv4.
	PreferIPv6 = false
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

// Public Retrieves the public IP address for the host by dialing to the
// OpenDNS DNS server. Upon resolving the ip address, it will return the sender's
// ip address.
func Public(ctx context.Context) ([]string, error) {
	publics := make([]string, 0)
	resolver := &net.Resolver{
		Dial:     openDNS1Dialer,
		PreferGo: true,
	}

	ips, err := resolver.LookupIPAddr(ctx, "myip.opendns.com")
	if err != nil {
		resolver.Dial = openDNS2Dialer
		ips, err = resolver.LookupIPAddr(ctx, "myip.opendns.com")

		if err != nil {
			return publics, err
		}
	}

	for _, ip := range ips {
		publics = append(publics, ip.String())
	}

	return publics, nil
}

// Local Retrieves the local IP address within the network by iterating
// through the interfaces, filtering the loopback ip addresses. The format of
// the returned addresses is IPv4, optionally IPv6 upon the value of PreferIPv6.
func Local() ([]string, error) {
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

			if PreferIPv6 && ip6 != nil {
				locals = append(locals, ip6.String())
			} else if ip4 != nil {
				locals = append(locals, ip4.String())
			}
		}
	}

	return locals, nil
}
