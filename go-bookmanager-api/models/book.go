package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Isbn    string    `json:"isbn"`
	Title   string    `json:"title"`
	Authors []*Author `json:"authors" gorm:"many2many:book_authors;"`
}
