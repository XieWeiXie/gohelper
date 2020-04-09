package gohelper

import (
	"fmt"
	"net"
	"testing"
)

func TestIP(t *testing.T) {
	ip := NewIP(byte(10), byte(109), byte(48), byte(95))
	fmt.Println(ip.Mask().String())
	IpNet := net.IPNet{
		IP:   ip.ip,
		Mask: ip.Mask(),
	}
	fmt.Println(IpNet.String())
	i, n, _ := net.ParseCIDR("16.158.165.91/22")
	fmt.Println(i.String(), i.DefaultMask().String(), n.String())
}
