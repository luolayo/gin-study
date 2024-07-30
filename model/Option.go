package model

type Option struct {
	Name  string `gorm:"primary_key" json:"name"`
	Value string `json:"value"`
}
