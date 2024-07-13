package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string `gorm:"type:varchar(255);not null"`
	Phone  string `gorm:"type:varchar(11);not null;unique"`
	Passwd string `gorm:"type:varchar(255);not null"`
}

type UserRegister struct {
	Name            string `form:"name" json:"name" binding:"required"`
	Phone           string `form:"phone" json:"phone" binding:"required"`
	Passwd          string `form:"passwd" json:"passwd" binding:"required"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" binding:"required"`
	Code            string `form:"code" json:"code" binding:"required"`
}

type UserLogin struct {
	Phone  string `form:"phone" json:"phone" binding:"required"`
	Passwd string `form:"passwd" json:"passwd" binding:"required"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Token string `json:"token"`
}
