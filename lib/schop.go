/*
Copyright (c) 2012-2018 Scott Chacon and others

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

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
