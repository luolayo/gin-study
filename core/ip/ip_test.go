package ip

import (
	"log"
	"testing"
)

func Test_Find(t *testing.T) {
	ip := Ip{}
	ip.Init()
	res, err := ip.Find("2409:8900:103f:14f:d7e:cd36:11af:be83")
	if err != nil {
		t.Error("ip type error")
	}
	log.Println(res.IP, res.Country, res.IpType, res.Isp)
}
