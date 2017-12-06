package main

import (
	"context"
	"fmt"
	"net"
)

func openDNSDialer(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{}
	// OpenDNS server ip address.
	return d.DialContext(ctx, "udp", "208.67.222.222:53")
}

// https://github.com/golang/go/issues/19268
func main() {
	r := net.Resolver{
		Dial:     openDNSDialer,
		PreferGo: true,
	}

	ctx := context.Background()
	ipaddr, err := r.LookupIPAddr(ctx, "myip.opendns.com")

	if err != nil {
		panic(err)
	}

	fmt.Println(ipaddr)
}
