package model

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null;comment:文章标题"`
	Context string `gorm:"type:text;not null;comment:文章内容"`
	User    User   `gorm:"foreignKey:UserID"`
	UserID  uint   `gorm:"comment:用户ID;onDelete:CASCADE;onCreate:CASCADE;not null"`
}

type ArticleInfo struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Context string `json:"context" form:"context" binding:"required"`
	UserID  uint   `json:"user_id" form:"user_id" binding:"required"`
}

type ArticleResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Context   string    `json:"context"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
