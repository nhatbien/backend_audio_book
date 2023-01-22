package model

import "time"

type User struct {
	Id        string    `json:"-" db:"id, omitempty"`
	Username  string    `json:"username,omitempty" db:"username, omitempty" `
	Email     string    `json:"email,omitempty" db:"email, omitempty" `
	Phone     string    `json:"phone,omitempty" db:"phone, omitempty" `
	Password  string    `json:"password,omitempty" db:"password, omitempty"`
	FullName  string    `json:"full_name,omitempty" db:"full_name, omitempty"`
	Age       int       `json:"age,omitempty" db:"age, omitempty"`
	Address   string    `json:"address,omitempty" db:"address, omitempty"`
	Photo     string    `json:"photo,omitempty" db:"photo, omitempty"`
	Status    int       `json:"status,omitempty" db:"status, omitempty"`
	RoleId    int       `json:"-" db:"role_id, omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at, omitempty"`

	Role Role `json:"role"`

	Token string `json:"token" db:"-, omitempty" `
}

type Error struct {
	ResponseCode      int    `json:"rc"`
	Message           string `json:"message"`
	Detail            string `json:"detail"`
	ExternalReference string `json:"ext_ref"`
}
