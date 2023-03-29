package request

type UserSignupRequest struct {
	Username string `json:"username" db:"username, omitempty" validate:"required"`
	Email    string `json:"email" db:"email, omitempty" `
	Phone    string `json:"phone" db:"phone, omitempty" validate:"required"`
	FullName string `json:"fullname" db:"fullname, omitempty" validate:"required"`
	Age      int    `json:"age" db:"age, omitempty" validate:"required"`
	Address  string `json:"address" db:"address, omitempty"`
	Password string `json:"password" db:"password, omitempty" validate:"required"`
	Photo    string `json:"photo" db:"password, omitempty" `
}
