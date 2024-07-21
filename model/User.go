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
	// User ID
	Uid uint `gorm:"primaryKey;autoIncrement" json:"uid" example:"1"`
	// User name
	Name string `gorm:"size:32;unique" json:"name" example:"admin"`
	// User password
	Password string `gorm:"size:64" json:"-"`
	// User phone number
	Phone string `gorm:"size:150;unique" json:"phone" example:"18888888888"`
	// User url
	Url string `gorm:"size:150" json:"url" example:"https://www.luola.me"`
	// User nickname
	ScreenName string `gorm:"size:32" json:"screenName" example:"罗拉"`
	// User registration time
	Created time.Time `gorm:"autoCreateTime" json:"-"`
	// User activation time
	Activated *time.Time `gorm:"autoCreateTime" json:"activated" example:"2021-07-01 00:00:00"`
	// User last login time
	Logged *time.Time `gorm:"autoCreateTime" json:"logged" example:"2021-07-01 00:00:00"`
	// User group
	Group Group `gorm:"default:'guest'" json:"group" enum:"admin,user,guest"`
	// User token
	Token string `gorm:"-" json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsIm5hbWUiOiJhZG1pbiIsImV4cCI6MTYyNjMwNzQwMCwiaWF0IjoxNjI2MzA3MjAwfQ"`

	Contents []Content `gorm:"foreignKey:AuthorId" json:"-"`
}

type UserRegister struct {
	// User name
	Name string `json:"name" form:"name" binding:"required" example:"admin"`
	// User password
	Password string `json:"password" form:"password" binding:"required" example:"123456"`
	// Confirm password is the same as password
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" binding:"required"  example:"123456"`
	// User phone number
	Phone string `json:"phone" form:"phone" binding:"required"  example:"18888888888"`
	// User avatar
	Url string `json:"url" form:"url"  example:"https://www.luola.me"`
	// User nickname
	ScreenName string `json:"screenName" form:"screenName"  example:"罗拉"`
	// Verification code
	Code string `json:"code" form:"code" binding:"required"  example:"123456"`
}

type UserLogin struct {
	// User name
	Name string `json:"name" form:"name" binding:"required" example:"admin"`
	// User password
	Password string `json:"password" form:"password" binding:"required" example:"123456"`
}

type UserUpdate struct {
	// User url
	Url string `json:"url" form:"url" example:"https://www.luola.me"`
	// User nickname
	ScreenName string `json:"screenName" example:"罗拉"`
}
