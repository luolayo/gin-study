package ip

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
)

func (d *V6) InitFile() (err error) {
	// check file status
	_, err = os.Stat(d.FilePath)
	if err != nil {
		return
	}

	// check file can open
	d.F, err = os.OpenFile(d.FilePath, os.O_RDONLY, 0400)
	if err != nil {
		return
	}

	// check read
	d.Data, err = os.ReadFile(d.FilePath)
	if err != nil {
		return
	}

	// 0-3 string "IPDB"
	if tag := d.Data[:4]; string(tag) != "IPDB" {
		return errors.New("file format error")
	}

	// Version 4-5, now it is 2. Ensure compatibility between version numbers 0x01 and 0xFF
	if d.Version = binary.LittleEndian.Uint16(d.Data[4:6]); d.Version > 2 {
		return errors.New("file version error")
	}

	// 6 bytes offset address length (2-8) 3
	d.Index.Offlen = uint64(d.Data[6])

	// 8~15 int64 records, as it cannot be negative, Uint64 has sufficient capacity
	d.Index.Count = binary.LittleEndian.Uint64(d.Data[8:16])

	// 16~23 int64 Offset of the first record in the index area
	d.Index.Start = binary.LittleEndian.Uint64(d.Data[16:24])

	d.Index.End = d.Index.Start + (d.Index.Count-1)*(d.Index.Offlen+8)
	defer func(F *os.File) {
		err := F.Close()
		if err != nil {

		}
	}(d.F)
	return nil
}

func (d *V6) Find(ip string) (res Result) {
	if ip == "" {
		return
	}

	ipv6 := net.ParseIP(ip)
	if ipv6 == nil {
		return
	}

	res.IP = fmt.Sprintf("%X", ipv6)
	r := binary.BigEndian.Uint64(ipv6[:8])
	res.Number = r
	o := d.SearchIndex(r, 0, d.Index.Count)
	res.Offset = d.GetUint64(o+8, d.Index.Offlen)
	res.Country, res.Area = d.getAddr(res.Offset)
	res.IP = ipv6.String()
	return res
}

func (d *V6) SearchIndex(ip, l, r uint64) (offset uint64) {
	// Using binary search to find index records of IP addresses encoded in network bytes
	// Return index location
	if r-l <= 1 {
		return d.Index.Start + l*(8+d.Index.Offlen)
	}

	m := (l + r) / 2
	o := d.Index.Start + m*(8+d.Index.Offlen)
	newIp := d.GetUint64(o, 8)
	if ip < newIp {
		return d.SearchIndex(ip, l, m)
	} else {
		return d.SearchIndex(ip, m, r)
	}

}

func (d *V6) getAddr(offset uint64) (string, string) {
	flag := d.GetUint64(offset, 1)
	if flag == 1 {
		//# redirection mode 1
		//# [IP] [0x01] [Absolute offset address of country and region information]
		//# Use the next 3 bytes as offset to call bytes for information retrieval
		return d.getAddr(d.GetUint64(offset+1, d.Index.Offlen))
	} else {
		//# redirection mode 2+normal mode
		//# [IP] [0x02] [Absolute offset of information] [...]
		cArea := d.getAreaAddr(offset)
		if flag == 2 {
			offset += 1 + d.Index.Offlen
		} else {
			offset = d.getOffset(offset) + 1
		}
		aArea := d.getAreaAddr(offset)
		return cArea, aArea
	}
}

func (d *V6) getAreaAddr(offset uint64) string {
	flag := d.GetUint64(offset, 1)
	if flag == 1 || flag == 2 {
		p := d.GetUint64(offset+1, d.Index.Offlen)
		return d.getAreaAddr(p)
	} else {
		return d.getString(offset)
	}
}

func (d *V6) getString(offset uint64) string {
	var res []byte
	for {
		if buf := d.Data[offset]; buf == 0 {
			break
		} else {
			res = append(res, buf)
			offset += 1
		}
	}
	return string(res)
}

func (d *V6) getOffset(offset uint64) uint64 {
	for {
		if buf := d.Data[offset]; buf == 0 {
			return offset
		} else {
			offset += 1
		}
	}
}

func (d *V6) GetUint64(offset, size uint64) uint64 {
	return byte2UInt64(d.Data[offset : offset+size])
}

func byte2UInt64(data []byte) uint64 {
	var i = []byte{0, 0, 0, 0, 0, 0, 0, 0}
	for j := 0; j < len(data); j++ {
		i[j] = data[j]
	}

	return binary.LittleEndian.Uint64(i)
}
