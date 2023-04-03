package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID           uint            `gorm:"primarykey"`
	BookName     string          ``
	Author       string          ``
	Content      string          ``
	Img          string          ``
	Audio        string          ``
	Price        float64         ``
	IsHot        bool            ``
	IsNew        bool            ``
	IsBestSeller bool            ``
	IsSale       bool            ``
	IsFree       bool            ``
	Status       int             ` gorm:"default:0"`
	CreatedAt    time.Time       ``
	UpdatedAt    time.Time       ``
	BookCategory []*BookCategory `json:"BookCategory,omitempty" gorm:"many2many:meta_book_category;;ForeignKey:id,locale;References:id"`
}

type AudioBookChapter struct {
	gorm.Model
	Name      string `json:"name" db:"name, omitempty"`
	BookId    uint   `json:"-" db:"book_id, omitempty"`
	Audio     string `json:"audio" db:"audio, omitempty"`
	IsDeleted bool   `json:"is_deleted" db:"is_deleted, omitempty"`
	Status    int    `json:"status" gorm:"default:0"`
}
