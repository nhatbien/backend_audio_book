package request

type UserUpdateRoleRequest struct {
	UserId string `json:"user_id,omitempty" db:"user_id, omitempty" validate:"required"`
	RoleId int    `json:"role_id,omitempty" db:"role_id, omitempty"   validate:"required"`
}
