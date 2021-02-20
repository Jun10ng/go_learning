// 判断ip是否属于某个网段
package main

import (
	"fmt"
	"net"
	"strings"
)

type CIDR struct {
	ip         net.IP
	ipnet      *net.IPNet
	onlyIP     bool // 不是网段
	onlyIPAddr string
}

func (c CIDR) Contains(ip string) bool {
	if c.onlyIP {
		return ip == c.GetOnlyIp()
	}
	return c.ipnet.Contains(net.ParseIP(ip))
}
func (c CIDR) GetOnlyIp() string {
	return c.onlyIPAddr
}
func ParseCIDR(s string) (*CIDR, error) {

	if strings.Contains(s, "/32") {
		return &CIDR{onlyIPAddr: s[0 : len(s)-3], onlyIP: true}, nil
	}
	i, n, err := net.ParseCIDR(s)
	if err != nil {
		return nil, err
	}
	return &CIDR{ip: i, ipnet: n}, nil
}

func main() {
	ip := "192.168.1.168"
	c, _ := ParseCIDR("192.168.1.168/32")
	//fmt.Println(c.ipnet.IP.String())
	allowIPMap := make([]*CIDR, 0)
	allowIPMap = append(allowIPMap, c)
	for _, ipnet := range allowIPMap {
		if exist := ipnet.Contains(ip); exist {
			fmt.Println("true")
			return
		}
	}
	fmt.Println("false")
}
