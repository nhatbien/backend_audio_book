package model

import "time"

type User struct {
	Id        string    ` gorm:"primaryKey" `
	Username  string    `  gorm:"size:255;uniqueIndex" `
	Email     string    `  gorm:"size:255;uniqueIndex"`
	Phone     string    `  gorm:"size:255;uniqueIndex" `
	Password  string    ` `
	FullName  string    ` `
	Age       int       ` `
	Address   string    ` `
	Photo     string    ` `
	Status    int       ` `
	RoleId    int       `   `
	CreatedAt time.Time ``
	UpdatedAt time.Time ``
	Role      Role      ` gorm:"foreignKey:RoleId"`
	Token     string    `json:"Token,omitempty" gorm:"-"`
}

type Error struct {
	ResponseCode      int    `json:"rc"`
	Message           string `json:"message"`
	Detail            string `json:"detail"`
	ExternalReference string `json:"ext_ref"`
}
