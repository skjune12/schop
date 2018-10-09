package schop

import (
	"encoding/json"
	"net"
	"strings"
)

type Result struct {
	Fqdn string `json:"fqdn"`
	IPv4 string `json:"ipv4"`
	IPv6 string `json:"ipv6"`
}

func Search(ip string) Result {
	var r Result
	fqdn, err := GetFQDN(ip)

	if err != nil {
		if IsIPv4(ip) {
			r.IPv4 = ip
		}
		if IsIPv6(ip) {
			r.IPv6 = ip
		}
	}

	r.Fqdn = fqdn

	addrs, err := GetAddrs(fqdn)

	for _, addr := range addrs {
		if IsIPv6(addr) == true {
			r.IPv6 = addr
		}

		if IsIPv4(addr) {
			r.IPv4 = addr
		}
	}

	return r
}

func GetFQDN(ip string) (string, error) {
	fqdn, err := net.LookupAddr(ip)
	if err != nil {
		return "", err
	}
	return fqdn[0], nil
}

func GetAddrs(fqdn string) ([]string, error) {
	addrs, err := net.LookupHost(fqdn)
	if err != nil {
		return nil, err
	}

	return addrs, nil
}

func (r Result) ToJson() (string, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func IsIPv6(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ":")
}

func IsIPv4(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ".")
}
