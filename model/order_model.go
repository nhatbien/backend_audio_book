package model

import (
	"time"
)

type Order struct {
	Id        uint   `gorm:"primarykey"`
	UserId    string ` db:"user_id, omitempty"`
	CartId    uint   ` db:"cart_id, omitempty"`
	Status    int    `json:"status" db:"status, omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
