package mygobase

import (
	"fmt"
	"net"
	"net/url"
)

func ValidURL(input string) bool {
	myo, error := url.Parse(input)
	return error == nil && myo.Scheme != "" && myo.Host != ""
}

func ValidIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	if parsedIP.To4() != nil {
		return true
	}
	return true
}

func ValidIP4(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	if parsedIP.To4() != nil {
		return true
	}
	return false
}

func ValidIP6(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	if parsedIP.To4() != nil {
		return false
	}
	return true
}

func lockPort(port int) bool {
	_, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return false
	}
	return true
}
