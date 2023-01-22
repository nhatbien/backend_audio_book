package repository

import (
	"backend/model"
	"backend/model/request"
	"context"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckLogin(context context.Context, loginRequest request.UserLoginRequest) (model.User, error)
	UpdateUser(context context.Context, user model.User) error
	UpdateRole(context context.Context, userRole request.UserUpdateRoleRequest) error
}
