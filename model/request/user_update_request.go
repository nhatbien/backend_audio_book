package request

type UserUpdateRequest struct {
	Username string `json:"username,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" `
	Phone    string `json:"phone,omitempty" `
	Photo    string `json:"photo,omitempty" `
	Age      int    `json:"age,omitempty" `
	Address  string `json:"address,omitempty" db:"address, omitempty"`

	FullName string `json:"full_name,omitempty"`
	Password string `json:"password,omitempty"`
	Status   int    `json:"status,omitempty" `
	Role     string `json:"role,omitempty" `
}
