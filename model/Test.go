package model

type Test struct {
	Msg string `form:"msg" json:"msg" binding:"required" gorm:"column:msg varchar(255) not null"`
}
