package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string
	Author_ID uint
}
