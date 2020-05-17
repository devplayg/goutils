package network

import (
	"encoding/binary"
	"net"
)

func IpToInt(ip net.IP) uint32 {
	if ip == nil {
		return 0
	}
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func IntToIp(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func CidrHostCount(cidr int) int {
	if cidr == 32 {
		return 1
	} else if cidr == 31 {
		return 2
	}
	return 2 << (31 - uint(cidr))
}
