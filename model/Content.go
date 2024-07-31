package model

import "github.com/luolayo/gin-study/core/ip"

type ViewStatus string
type Type string

const (
	ViewStatusPrivate ViewStatus = "private"
	ViewStatusPublic  ViewStatus = "public"
)

const (
	TypePost Type = "post"
	TypePage Type = "page"
)

type Content struct {
	Cid           uint   `gorm:"primaryKey;autoIncrement"`
	Title         string `gorm:"unique_index"`
	Slug          string `gorm:"unique_index"`
	Created       uint   `gorm:"index"`
	Modified      uint
	Text          string `gorm:"type:longtext"`
	Order         uint
	AuthorId      uint
	Type          Type       `gorm:"size:16;default:post"`
	Status        ViewStatus `gorm:"size:16;default:public"`
	Password      string     `gorm:"size:32"`
	CommentsNum   uint
	Parent        uint
	Views         uint `gorm:"default:0"`
	Ip            string
	Address       *ip.Address
	Author        User           `gorm:"foreignKey:AuthorId"`
	Comments      []Comment      `gorm:"foreignKey:Cid" json:"-"`
	Relationships []Relationship `gorm:"foreignKey:Cid" json:"-"`
}
