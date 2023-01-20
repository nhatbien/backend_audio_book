package request

type UserSignupRequest struct {
	Username string `json:"username,omitempty" db:"username, omitempty" validate:"required"`
	Email    string `json:"email,omitempty" db:"email, omitempty" `
	Phone    string `json:"phone,omitempty" db:"phone, omitempty" validate:"required"`
	FullName string `json:"fullname,omitempty" db:"fullname, omitempty" validate:"required"`
	Password string `json:"password,omitempty" db:"password, omitempty" validate:"required"`
	Photo    string `json:"photo,omitempty" db:"password, omitempty" `
}
