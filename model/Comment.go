package model

type CommentStatus string

const (
	Approved CommentStatus = "approved"
	Pending  CommentStatus = "pending"
)

type Comment struct {
	Coid       uint          `gorm:"primaryKey;autoIncrement" json:"coid"`
	Cid        uint          `gorm:"index" json:"cid"`
	Created    uint          `gorm:"autoCreateTime" json:"created"`
	AuthorName string        `gorm:"size:150" json:"authorName"`
	AuthorId   uint          `json:"authorId"`
	Url        string        `gorm:"size:255" json:"url"`
	Ip         string        `gorm:"size:64" json:"ip"`
	Agent      string        `gorm:"size:511" json:"agent"`
	Text       string        `gorm:"type:text" json:"text"`
	Status     CommentStatus `gorm:"size:16;default:approved" json:"status"`

	Content Content `gorm:"foreignKey:Cid"`
	Author  User    `gorm:"foreignKey:AuthorId"`
}

type CommentRequest struct {
	Text string `json:"text" binding:"required"`
}
