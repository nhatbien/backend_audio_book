package model

import "time"

type User struct {
	Id        string    `json:"-" gorm:"primaryKey" `
	Username  string    `json:"username,omitempty"  gorm:"size:255;uniqueIndex" `
	Email     string    `json:"email,omitempty"  gorm:"size:255;uniqueIndex"`
	Phone     string    `json:"phone,omitempty"  gorm:"size:255;uniqueIndex" `
	Password  string    `json:"password,omitempty" `
	FullName  string    `json:"full_name,omitempty" `
	Age       int       `json:"age,omitempty" `
	Address   string    `json:"address,omitempty" `
	Photo     string    `json:"photo,omitempty" `
	Status    int       `json:"status,omitempty" `
	RoleId    int       `json:"role_id,omitempty"   `
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Role      Role      `json:"role,omitempty" gorm:"foreignKey:RoleId"`
	Token     string    `json:"token,omitempty"  gorm:"-"`
}

type Error struct {
	ResponseCode      int    `json:"rc"`
	Message           string `json:"message"`
	Detail            string `json:"detail"`
	ExternalReference string `json:"ext_ref"`
}
