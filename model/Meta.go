package model

type Meta struct {
	Mid         uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"size:150"`
	Slug        string `gorm:"size:150;unique"`
	Type        string `gorm:"size:32"`
	Description string `gorm:"size:150"`
	Count       uint
	Order       uint
	Parent      uint

	Relationships []Relationship `gorm:"foreignKey:Mid"`
}
