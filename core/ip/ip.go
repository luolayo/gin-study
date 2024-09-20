package ip

import (
	"errors"
	"github.com/luolayo/gin-study/enum"
	"github.com/xiaoqidun/qqwry"
	"net"
	"os"
	"path"
	"strings"
)

/**
 * How to use this snippet:
 * 1. Create a new IP object
 * 2. Initialize the IP library
 * 3. Find the IP address
 * @Param ip string
 * @Return Address, error
 */

func NewIp() *Ip {
	return &Ip{}
}

func (i *Ip) Init() {
	dir, _ := os.Getwd()
	FilePath := path.Join(dir, enum.V6Path)
	d := V6{FilePath: FilePath}
	err := d.InitFile()
	if err != nil {
		panic(err)
	}
	i.v6 = &d
	FilePath = path.Join(dir, enum.V4Path)
	d4 := V4{FilePath: FilePath}
	err = d4.InitFile()
	if err != nil {
		panic(err)
	}
	i.v4 = &d4
}

func (i *Ip) Find(ip string) (Address, error) {
	switch checkIpType(ip) {
	case V4Type:
		return ParseIpv4(i.v4.Find(ip)), nil
	case V6Type:
		return ParseIpv6(i.v6.Find(ip)), nil
	}
	return Address{}, errors.New("ip type error")
}

func checkIpType(address string) Type {
	ip := net.ParseIP(address)
	if ip == nil {
		return -1
	}
	if ip.To4() != nil {
		return V4Type
	} else {
		return V6Type
	}
}

func ParseIpv4(res *qqwry.Location) Address {
	return Address{
		IP:      res.IP,
		Country: res.Country + res.Country + res.District,
		Isp:     res.ISP,
		IpType:  "ipv4",
	}
}

func ParseIpv6(res Result) Address {
	return Address{
		IP:      res.IP,
		Country: res.Country,
		IpType:  "ipv6",
		Isp:     res.Area,
	}
}

func FinDIsp(isp string) string {
	var isps = []string{
		"电信",
		"联通",
		"移动",
		"铁通",
		"教育网",
	}
	for _, v := range isps {
		if strings.Contains(v, isp) {
			return v
		}
	}
	return isp
}
