package model

import "time"

type Book struct {
	Id           int            `json:"id" db:"id, omitempty"`
	BookName     string         `json:"title" db:"title, omitempty"`
	Author       string         `json:"author" db:"author, omitempty"`
	Content      string         `json:"content" db:"content, omitempty"`
	Img          string         `json:"img" db:"img, omitempty"`
	Audio        string         `json:"audio" db:"audio, omitempty"`
	Price        int            `json:"price" db:"price, omitempty"`
	IsHot        bool           `json:"is_hot" db:"is_hot, omitempty"`
	IsNew        bool           `json:"is_new" db:"is_new, omitempty"`
	IsBestSeller bool           `json:"is_best_seller" db:"is_best_seller, omitempty"`
	IsSale       bool           `json:"is_sale" db:"is_sale, omitempty"`
	IsFree       bool           `json:"is_free" db:"is_free, omitempty"`
	Status       int            `json:"status" gorm:"default:0"`
	CreatedAt    time.Time      `json:"created_at" db:"created_at, omitempty"`
	UpdatedAt    time.Time      `json:"updated_at" db:"updated_at, omitempty"`
	BookCategory []BookCategory `json:"book_category,omitempty" gorm:"many2many:meta_book_category;"`
}

type AudioBookChapter struct {
	Id        int    `json:"id" db:"id, omitempty"`
	Name      string `json:"name" db:"name, omitempty"`
	BookId    int    `json:"book_id" db:"book_id, omitempty"`
	Audio     string `json:"audio" db:"audio, omitempty"`
	IsDeleted bool   `json:"is_deleted" db:"is_deleted, omitempty"`
	CreatedAt string `json:"created_at" db:"created_at, omitempty"`
	UpdatedAt string `json:"updated_at" db:"updated_at, omitempty"`
}
