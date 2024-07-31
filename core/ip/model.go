package ip

import (
	"database/sql/driver"
	"fmt"
	"github.com/goccy/go-json"
	"os"
)

type Type int

const (
	V4Type Type = iota
	V6Type
)

type Ip struct {
	v6 *V6
	v4 *V4
}

type V6 struct {
	Data     []byte
	Version  uint16
	FilePath string
	F        *os.File
	Index    Index
	Offset   uint32
}

type V4 struct {
	FilePath string
}

type Index struct {
	Start  uint64
	End    uint64
	Offlen uint64
	Count  uint64
	Data   [][]uint64
}

type Result struct {
	IP      string `json:"ip"`
	Number  uint64 `json:"number"`
	Country string `json:"country"`
	Area    string `json:"area"`
	Offset  uint64 `json:"offset"`
}

type Address struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	IpType  string `json:"ip_type"`
	Isp     string `json:"isp"`
}

func (a *Address) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	// Convert the value from the database to bytes
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Address: %v", value)
	}
	// Unmarshal the JSON-encoded data into the Address struct
	return json.Unmarshal(bytes, a)
}

func (a Address) Value() (driver.Value, error) {
	// Marshal the Address struct into JSON
	bytes, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	// Return the JSON-encoded data
	return string(bytes), nil
}
