package model

import "gorm.io/gorm"

type Test struct {
	gorm.Model `gorm:"default:CURRENT_TIMESTAMP"`
	Msg        string `form:"msg" json:"msg" binding:"required" gorm:"column:msg varchar(255) not null"`
}
