package model

import "time"

type User struct {
	Id        string    `json:"-" gorm:"primaryKey" `
	Username  string    `json:"username"  gorm:"size:255;uniqueIndex" `
	Email     string    `json:"email"  gorm:"size:255;uniqueIndex"`
	Phone     string    `json:"phone"  gorm:"size:255;uniqueIndex" `
	Password  string    `json:"password" `
	FullName  string    `json:"full_name" `
	Age       int       `json:"age" `
	Address   string    `json:"address" `
	Photo     string    `json:"photo" `
	Status    int       `json:"status" `
	RoleId    int       `json:"role_id"   `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role      Role      `json:"role" gorm:"foreignKey:RoleId"`
	Token     string    `json:"token"  gorm:"-"`
}

type Error struct {
	ResponseCode      int    `json:"rc"`
	Message           string `json:"message"`
	Detail            string `json:"detail"`
	ExternalReference string `json:"ext_ref"`
}
