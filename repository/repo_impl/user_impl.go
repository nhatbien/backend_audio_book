package repo_impl

import (
	"backend/biedeptrai"
	"backend/db"
	"backend/model"
	"backend/model/request"
	"backend/repository"
	"context"
	"time"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{sql: sql}
}

func (n *UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {

	if count := n.sql.Db.Where(&model.User{Username: user.Username}).First(new(model.User)).RowsAffected; count > 0 {
		return user, biedeptrai.ErrorUserConflict
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err := n.sql.Db.Create(user).Error
	if err != nil {
		return user, biedeptrai.ErrorSignUpFail
	}
	return user, nil

}

func (n *UserRepoImpl) Login(context context.Context, loginRequest request.UserLoginRequest) (model.User, error) {
	var user = model.User{}

	if res := n.sql.Db.Where(
		&model.User{Username: loginRequest.Username},
	).Preload("Role").First(&user); res.RowsAffected <= 0 {
		return user, biedeptrai.ErrorUserNotFound
	}
	return user, nil

}

func (n *UserRepoImpl) SelectUserId(context context.Context, userId string) (model.User, error) {
	var user model.User
	if res := n.sql.Db.Where(
		&model.User{Id: userId},
	).First(&user); res.RowsAffected <= 0 {
		return user, biedeptrai.ErrorUserNotFound
	}
	return user, nil

}

func (n *UserRepoImpl) UpdateUser(context context.Context, user model.User) (model.User, error) {
	user.UpdatedAt = time.Now()

	if res := n.sql.Db.Where(
		&model.User{Id: user.Id},
	).Save(&user); res.RowsAffected <= 0 {
		return user, biedeptrai.ErrorUserNotFound
	}

	return user, nil

}

func (n *UserRepoImpl) UpdateRole(context context.Context, userRole request.UserUpdateRoleRequest) error {
	user := model.User{
		Id:     userRole.UserId,
		RoleId: userRole.RoleId,
	}
	if res := n.sql.Db.Where(
		&model.User{Id: user.Id},
	).Save(&user); res.RowsAffected <= 0 {
		return biedeptrai.ErrorUserNotFound
	}
	return nil
}

func (n *UserRepoImpl) DeleteUser(context context.Context, userId string) error {
	if res := n.sql.Db.Where(
		&model.User{Id: userId},
	).Delete(&model.User{}); res.RowsAffected <= 0 {
		return biedeptrai.ErrorUserNotFound
	}
	return nil
}

func (n *UserRepoImpl) SelectAllUser(context context.Context) ([]model.User, error) {
	var users []model.User
	if res := n.sql.Db.Find(&users); res.RowsAffected <= 0 {
		return users, biedeptrai.ErrorUserNotFound
	}
	return users, nil
}

func (n *UserRepoImpl) SelectAllUserByRole(context context.Context, roleId int) ([]model.User, error) {
	var users []model.User
	if res := n.sql.Db.Where(
		&model.User{RoleId: roleId},
	).Find(&users); res.RowsAffected <= 0 {
		return users, biedeptrai.ErrorUserNotFound
	}
	return users, nil
}
