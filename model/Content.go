package model

import "time"

type ViewStatus string
type Type string

const (
	ViewStatusPrivate ViewStatus = "private"
	ViewStatusPublic  ViewStatus = "public"
)

const (
	TypePost       Type = "post"
	TypePage       Type = "page"
	TypeAttachment Type = "attachment"
)

type Content struct {
	Cid         uint       `gorm:"primaryKey;autoIncrement" json:"cid"`
	Title       string     `gorm:"unique_index" json:"title"`
	Slug        string     `gorm:"unique_index" json:"slug"`
	Created     *time.Time `gorm:"autoCreateTime" json:"created"`
	Modified    *time.Time `gorm:"autoCreateTime" json:"modified"`
	Text        string     `gorm:"type:longtext" json:"text"`
	Order       uint       `gorm:"not null;default:0" json:"order"`
	AuthorId    uint       `gorm:"not null" json:"authorId"`
	Type        Type       `gorm:"size:16;default:post" json:"type"`
	Status      ViewStatus `gorm:"size:16;default:private" json:"status"`
	CommentsNum uint       `gorm:"default:0" json:"commentsNum"`
	Parent      uint       `gorm:"default:0" json:"parent"`
	Views       uint       `gorm:"default:0" json:"views"`

	Author        User           `gorm:"foreignKey:AuthorId" json:"-"`
	Comments      []Comment      `gorm:"foreignKey:Cid" json:"-"`
	Relationships []Relationship `gorm:"foreignKey:Cid" json:"-"`
}

type ContentRequest struct {
	// Content Title
	Title string `json:"title" binding:"required" form:"title" example:"Hello World"`
	// Content Text
	Text string `json:"text" binding:"required" form:"text" example:"Hello World"`
	// Content Type
	Type Type `json:"type" binding:"required" form:"type" example:"post" enum:"post,page,attachment"`
	// Content Parent
	Parent uint `json:"parent" form:"parent" example:"0"`
	// Content Order
	Order uint `json:"order" form:"order" example:"0"`
	// Content Slug
	Slug string `json:"slug" form:"slug" example:"hello-world"`
}

type ContentUpdate struct {
	// Content Title
	Title string `json:"title" form:"title" example:"Hello World"`
	// Content Text
	Text string `json:"text" form:"text" example:"Hello World"`
	// Order
	Order uint `json:"order" form:"order" example:"0"`
}
