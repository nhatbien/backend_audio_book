package request

type UserLoginRequest struct {
	Username string `json:"username" db:"username, omitempty" validate:"required"`
	Password string `json:"password" db:"password, omitempty" validate:"required"`
}
