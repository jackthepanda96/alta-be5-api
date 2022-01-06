package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama  string
	HP    string
	Books []Book `gorm:"foreignKey:Author_ID"`
}
