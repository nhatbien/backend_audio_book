package model

import (
	"time"
)

type Cart struct {
	Id         uint       `gorm:"primarykey"`
	UserId     string     `json:"-"  gorm:"size:255" `
	TotalPrice float64    ` gorm:"default:0"`
	IsCurrent  bool       ` gorm:"default:true"`
	Items      []CartItem ``
	Status     int        ` gorm:"default:0"`
	User       User       `json:"-" gorm:"foreignKey:UserId"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CartItem struct {
	Id          uint    `gorm:"primarykey"`
	CartId      uint    `json:"-" `
	BookId      uint    `json:"-" `
	Quantity    int     ` gorm:"default:1" `
	TotalAmount float64 ` gorm:"default:0"`
	Cart        Cart    `json:"-" gorm:"foreignKey:CartId"`
	Book        Book    `json:"Book,omitempty" gorm:"foreignKey:BookId" `
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
