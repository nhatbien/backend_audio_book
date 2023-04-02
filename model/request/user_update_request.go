package request

type UserUpdateRequest struct {
	Email    string `json:"email" `
	Phone    string `json:"phone" `
	Photo    string `json:"photo" `
	Age      int    `json:"age" `
	Address  string `json:"address" db:"address, omitempty"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}
