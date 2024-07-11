package Model

// Relationship is the join table for the many-to-many relationship between Content and Meta
type Relationship struct {
	ContentID uint `gorm:"primaryKey"`
	MetaID    uint `gorm:"primaryKey"`
}
