package Model

import (
	"gorm.io/gorm"
	"time"
)

// User represents a user in the system
type User struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"unique;not null"`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	ScreenName string `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
