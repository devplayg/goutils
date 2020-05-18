package network

import (
	"encoding/binary"
	"net"
)

// IPv4ToInt converts IPv4 to int
func IPv4ToInt(ip net.IP) uint32 {
	if ip == nil {
		return 0
	}
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

// IntToIPv4 converts int to IPv4
func IntToIPv4(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
