package ip

import (
	"log"
	"os"
	"path"
	"testing"
)

func TestIpV6_Find(t *testing.T) {
	dir, _ := os.Getwd()
	FilePath := path.Join(dir, "/data/ipv6wry.db")
	d := V6{FilePath: FilePath}

	err := d.InitFile()
	if err != nil {
		println(err.Error())
	}
	res := d.Find("2409:8a38:4222:8460:d635:38ff:fe65:2015")
	log.Println(res)
}
