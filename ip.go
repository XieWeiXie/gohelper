package gohelper

import (
	"net"
)

type IP struct {
	ip net.IP
}

func NewIP(a,b,c,d byte) *IP{
	return &IP{
		ip: net.IPv4(a,b,c,d),
	}
}
func (i IP) Mask() net.IPMask{
	return i.ip.DefaultMask()
}

func (i IP)ToMask(ip []byte) net.IP{
	mask := net.IPMask(ip)
	return i.ip.Mask(mask)
}