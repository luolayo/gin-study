package model

import (
	"time"
)

type Group string

const (
	// GroupAdmin is the group for admin users
	GroupAdmin Group = "admin"
	// GroupUser is the group for normal users
	GroupUser Group = "user"
	// GroupGuest is the group for guest users
	GroupGuest Group = "guest"
)

type User struct {
	Uid        uint       `gorm:"primaryKey;autoIncrement"`
	Name       string     `gorm:"size:32;unique"`
	Password   string     `gorm:"size:64"`
	Phone      string     `gorm:"size:150;unique"`
	Url        string     `gorm:"size:150"`
	ScreenName string     `gorm:"size:32"`
	Created    *time.Time `gorm:"autoCreateTime"`
	IP         string     `gorm:"size:32"`
	Logged     *time.Time `gorm:"autoCreateTime"`
	Group      Group      `json:"group" gorm:"default:'guest'" form:"group"`
	Token      string     `json:"token" gorm:"-"`

	Contents []Content `gorm:"foreignKey:AuthorId" json:"-"`
}
