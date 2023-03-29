package request

type UserUpdateRoleRequest struct {
	UserId string `json:"user_id" db:"user_id, omitempty" validate:"required"`
	RoleId int    `json:"role_id" db:"role_id, omitempty"   validate:"required"`
}
