package ipkit

import (
	"encoding/binary"
	"net"
	"net/http"
	"strings"
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

// GetRealIP 通用版：兼容代理/CDN，获取真实客户端IP
func GetRealIP(r *http.Request) string {
	// 1. 优先读取 X-Forwarded-For（多个代理时，第一个是客户端真实IP）
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		// X-Forwarded-For 格式：clientIP, proxy1IP, proxy2IP
		ips := strings.Split(xff, ",")
		for _, ip := range ips {
			ip = strings.TrimSpace(ip)
			// 过滤空值和内网IP（可选，根据需求调整）
			if ip != "" && !isPrivateIP(ip) {
				return ip
			}
		}
		// 若所有IP都过滤了，取第一个非空IP
		for _, ip := range ips {
			ip = strings.TrimSpace(ip)
			if ip != "" {
				return ip
			}
		}
	}

	// 2. 读取 X-Real-IP（Nginx 常用）
	xri := r.Header.Get("X-Real-IP")
	if xri != "" {
		return strings.TrimSpace(xri)
	}

	// 3. 降级到 RemoteAddr
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		idx := strings.LastIndex(r.RemoteAddr, ":")
		if idx == -1 {
			return r.RemoteAddr
		}
		return r.RemoteAddr[:idx]
	}
	return host
}

// isPrivateIP 检测是否为内网IP（可选，用于过滤代理内网IP）
func isPrivateIP(ip string) bool {
	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		return true
	}
	// 内网IP段：10.0.0.0/8、172.16.0.0/12、192.168.0.0/16、fc00::/7（ipv6）
	privateCIDRs := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"fc00::/7",
	}
	for _, cidr := range privateCIDRs {
		_, ipNet, err := net.ParseCIDR(cidr)
		if err != nil {
			continue
		}
		if ipNet.Contains(ipAddr) {
			return true
		}
	}
	return false
}
