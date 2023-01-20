package request

type UserLoginRequest struct {
	Username string `json:"username,omitempty" db:"username, omitempty" validate:"required"`
	Password string `json:"password,omitempty" db:"password, omitempty" validate:"required"`
}
