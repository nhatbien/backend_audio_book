package request

type UserSignupRequest struct {
	Username string ` validate:"required"`
	Email    string ` `
	Phone    string ` validate:"required"`
	FullName string ` validate:"required"`
	Age      int    ` validate:"required"`
	Address  string ``
	Password string ` validate:"required"`
	Photo    string ` `
}
