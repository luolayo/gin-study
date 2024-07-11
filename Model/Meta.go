package Model

import (
	"gorm.io/gorm"
	"time"
)

// Meta represents a category or tag
type Meta struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Slug      string    `gorm:"unique;not null"`
	Type      string    `gorm:"not null"` // category, tag, etc.
	ParentID  uint      `gorm:"default:0"`
	Parent    *Meta     `gorm:"foreignKey:ParentID"`
	Contents  []Content `gorm:"many2many:relationships"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
