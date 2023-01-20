package model

import "time"

type User struct {
	Id        string    `json:"-" db:"id, omitempty"`
	Username  string    `json:"username,omitempty" db:"username, omitempty" gorm:"uniqueIndex"`
	Email     string    `json:"email,omitempty" db:"email, omitempty" gorm:"uniqueIndex"`
	Phone     string    `json:"phone,omitempty" db:"phone, omitempty" gorm:"uniqueIndex"`
	Photo     string    `json:"photo,omitempty" db:"photo, omitempty"`
	FullName  string    `json:"full_name,omitempty" db:"full_name, omitempty"`
	Password  string    `json:"password,omitempty" db:"password, omitempty"`
	Status    string    `json:"status,omitempty" db:"status, omitempty"`
	Role      string    `json:"role,omitempty" db:"role, omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at, omitempty"`
	Token     string    `json:"token" db:"-, omitempty" gorm:"-:migration"`
}

type Error struct {
	ResponseCode      int    `json:"rc"`
	Message           string `json:"message"`
	Detail            string `json:"detail"`
	ExternalReference string `json:"ext_ref"`
}
