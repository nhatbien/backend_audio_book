package request

type UserUpdateRoleRequest struct {
	UserId string ` validate:"required"`
	RoleId int    `   validate:"required"`
}
