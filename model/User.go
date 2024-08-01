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
	Uid        uint       `gorm:"primaryKey;autoIncrement" json:"uid"`
	Name       string     `gorm:"size:32;unique" json:"name"`
	Password   string     `gorm:"size:64" json:"-"`
	Phone      string     `gorm:"size:150;unique" json:"phone"`
	Url        string     `gorm:"size:150" json:"url"`
	ScreenName string     `gorm:"size:32" json:"screenName"`
	Created    *time.Time `gorm:"autoCreateTime" json:"created"`
	IP         string     `gorm:"size:32" json:"ip"`
	Logged     *time.Time `gorm:"autoCreateTime" json:"logged"`
	Group      Group      `json:"group" gorm:"default:'guest'" form:"group" json:"group"`
	Token      string     `json:"token" gorm:"-" json:"token"`

	Contents []Content `gorm:"foreignKey:AuthorId" json:"-"`
}
