package mygobase

import (
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
