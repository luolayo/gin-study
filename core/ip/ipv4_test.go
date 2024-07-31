package ip

import (
	"log"
	"testing"
)

func TestIpV4_Find(t *testing.T) {
	d := V4{FilePath: "qqwry.dat"}

	err := d.InitFile()
	if err != nil {
		println(err.Error())
	}
	log.Println(d.Find("119.29.29.29"))
}
