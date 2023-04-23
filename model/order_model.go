package model

import (
	"time"
)

type Order struct {
	Id        uint   `gorm:"primarykey"`
	UserId    string ` db:"user_id, omitempty" gorm:"size:255;"`
	CartId    uint   ` db:"cart_id, omitempty"`
	Cart      Cart   `json:"Cart,omitempty" gorm:"foreignKey:CartId" `
	Status    int    `json:"status" db:"status, omitempty"`
	User      User   `json:"user,omitempty" gorm:"foreignKey:UserId" `
	CreatedAt time.Time
	UpdatedAt time.Time
}
