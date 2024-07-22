package model

type Link struct {
	ID     int    `gorm:"primary_key;auto_increment"`
	Name   string `gorm:"type:varchar(100);not null"`
	URL    string `gorm:"type:varchar(100);not null"`
	Sort   int    `gorm:"type:int;not null;sort:asc"`
	Image  string `gorm:"type:varchar(100);not null"`
	Stutas int    `gorm:"type:int;not null;default:0"`
}

type LinkRequest struct {
	Name  string `json:"name" binding:"required" form:"name"`
	Image string `json:"avatar" binding:"required" form:"avatar"`
	URL   string `json:"url" binding:"required" form:"url"`
}

type LinkUpdate struct {
	Name  string `json:"name" form:"name"`
	Image string `json:"avatar" form:"avatar"`
	URL   string `json:"url" form:"url"`
	Sort  int    `json:"sort" form:"sort"`
}
