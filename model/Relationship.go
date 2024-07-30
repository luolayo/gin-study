package model

type Relationship struct {
	Cid uint `gorm:"primaryKey"`
	Mid uint `gorm:"primaryKey"`

	Content Content `gorm:"foreignKey:Cid"`
	Meta    Meta    `gorm:"foreignKey:Mid"`
}
