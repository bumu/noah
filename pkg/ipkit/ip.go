package ipkit

import (
	"encoding/binary"
	"net"
)

func Ip2Int(ipstr string) uint32 {
	// Parse the IP address.
	ip := net.ParseIP(ipstr)

	// Convert the IP address to a byte slice.
	ipBytes := ip.To4()

	// Convert the byte slice to an integer.
	ipInt := binary.BigEndian.Uint32(ipBytes[:])

	// Print the integer representation of the IP address.
	return ipInt
}

func IsIPv4(ipstr string) bool {
	ip := net.ParseIP(ipstr)
	return ip != nil && ip.To4() != nil
}

func IsIPv6(ipstr string) bool {
	ip := net.ParseIP(ipstr)
	return ip != nil && ip.To4() == nil && ip.To16() != nil
}
