package model

import "time"

type BookCategory struct {
	ID          uint    `gorm:"primarykey"`
	Name        string  ` gorm:"not null"`
	Description string  ` `
	Images      string  ` `
	Book        []*Book `json:"Book,omitempty" gorm:"many2many:meta_book_category;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
