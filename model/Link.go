package model

type Link struct {
	ID    int    `gorm:"primary_key;auto_increment"`
	Name  string `gorm:"type:varchar(100);not null"`
	URL   string `gorm:"type:varchar(100);not null"`
	Sort  int    `gorm:"type:int;not null;sort:asc"`
	Image string `gorm:"type:varchar(100);not null"`
}
