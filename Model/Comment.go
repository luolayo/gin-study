package Model

import (
	"gorm.io/gorm"
	"time"
)

// Comment represents a comment on a content
type Comment struct {
	ID        uint    `gorm:"primaryKey"`
	ContentID uint    `gorm:"not null"`
	Content   Content `gorm:"foreignKey:ContentID"`
	Author    string  `gorm:"not null"`
	Email     string  `gorm:"not null"`
	URL       string
	Text      string   `gorm:"type:text;not null"`
	ParentID  uint     `gorm:"default:0"`
	Parent    *Comment `gorm:"foreignKey:ParentID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
