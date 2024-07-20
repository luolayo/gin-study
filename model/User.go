package model

import "time"

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
	Uid        uint      `gorm:"primaryKey;autoIncrement" json:"uid" description:"User ID" example:"1"`
	Name       string    `gorm:"size:32;unique" json:"name" description:"User name" example:"admin"`
	Password   string    `gorm:"size:64" json:"-"`
	Phone      string    `gorm:"size:150;unique" json:"phone" description:"User phone number" example:"18888888888"`
	Url        string    `gorm:"size:150" json:"url" description:"User avatar" example:"https://www.luola.me"`
	ScreenName string    `gorm:"size:32" json:"screenName" description:"User nickname" example:"罗拉"`
	Created    time.Time `gorm:"autoCreateTime" json:"-"`
	// Activated 给一个默认的当前时间
	Activated *time.Time `gorm:"autoCreateTime" json:"activated" description:"User activation time" example:"2021-07-01 00:00:00"`
	Logged    *time.Time `gorm:"autoCreateTime" json:"logged" description:"Last login time" example:"2021-07-01 00:00:00"`
	Group     Group      `gorm:"default:'guest'" json:"group" description:"User group" example:"guest" enum:"admin,user,guest"`
	Token     string     `gorm:"-" json:"token" description:"User token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsIm5hbWUiOiJhZG1pbiIsImV4cCI6MTYyNjMwNzQwMCwiaWF0IjoxNjI2MzA3MjAwfQ"`

	Contents []Content `gorm:"foreignKey:AuthorId" json:"-"`
}

type UserRegister struct {
	Name            string `json:"name" form:"name" binding:"required" description:"User name" example:"admin"`
	Password        string `json:"password" form:"password" binding:"required" description:"User password" example:"123456"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" binding:"required" description:"Confirm password" example:"123456"`
	Phone           string `json:"phone" form:"phone" binding:"required" description:"User phone number" example:"18888888888"`
	Url             string `json:"url" form:"url" description:"User avatar" example:"https://www.luola.me"`
	ScreenName      string `json:"screenName" form:"screenName" description:"User nickname" example:"罗拉"`
	Code            string `json:"code" form:"code" binding:"required" description:"Verification code" example:"123456"`
}

type UserLogin struct {
	Name     string `json:"name" form:"name" binding:"required" example:"admin" description:"User name"`
	Password string `json:"password" form:"password" binding:"required" example:"123456" description:"User password"`
}
