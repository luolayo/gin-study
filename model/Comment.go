package model

type Comment struct {
	Coid       uint `gorm:"primaryKey;autoIncrement"`
	Cid        uint `gorm:"index"`
	Created    uint
	AuthorName string `gorm:"size:150"`
	AuthorId   uint
	OwnerId    uint
	Mail       string `gorm:"size:150"`
	Url        string `gorm:"size:255"`
	Ip         string `gorm:"size:64"`
	Agent      string `gorm:"size:511"`
	Text       string `gorm:"type:text"`
	Type       string `gorm:"size:16;default:comment"`
	Status     string `gorm:"size:16;default:approved"`
	Parent     uint

	Content Content `gorm:"foreignKey:Cid"`
	Author  User    `gorm:"foreignKey:AuthorId"`
}
