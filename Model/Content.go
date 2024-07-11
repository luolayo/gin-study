package Model

import (
	"gorm.io/gorm"
	"time"
)

// Content represents a blog post or page
type Content struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Slug      string `gorm:"unique;not null"`
	Text      string `gorm:"type:text;not null"`
	Type      string `gorm:"not null"` // post, page, etc.
	Status    string `gorm:"not null"` // publish, draft, etc.
	AuthorID  uint   `gorm:"not null"`
	Author    User   `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Comments  []Comment      `gorm:"foreignKey:ContentID"`
	Metas     []Meta         `gorm:"many2many:relationships"`
}
