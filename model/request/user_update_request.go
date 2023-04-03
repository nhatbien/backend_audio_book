package request

type UserUpdateRequest struct {
	Email    string
	Phone    string
	Photo    string
	Age      int
	Address  string
	FullName string
	Password string
}
