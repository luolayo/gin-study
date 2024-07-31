package ip

import (
	"github.com/xiaoqidun/qqwry"
	"os"
)

func (d *V4) InitFile() (err error) {
	// check file status
	_, err = os.Stat(d.FilePath)
	if err != nil {
		return err
	}
	if err := qqwry.LoadFile(d.FilePath); err != nil {
		return err
	}
	return nil
}

func (d *V4) Find(ip string) *qqwry.Location {
	location, _ := qqwry.QueryIP(ip)
	return location
}
