package request

type UserUpdateRequest struct {
	Username string `json:"username,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" `
	Phone    string `json:"phone,omitempty" `
	Photo    string `json:"photo,omitempty" `
	FullName string `json:"full_name,omitempty"`
	Password string `json:"password,omitempty"`
	Status   string `json:"status,omitempty" `
	Role     string `json:"role,omitempty" `
}
