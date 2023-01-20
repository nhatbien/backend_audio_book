package repository_impl

import (
	"backend/biedeptrai"
	"backend/db"
	"backend/model"
	"backend/model/request"
	"backend/repository"
	"context"
	"database/sql"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{sql: sql}
}

func (n *UserRepoImpl) SaveUser(c context.Context, user model.User) (model.User, error) {

	statement := `INSERT INTO users(
		id, username, email, phone, photo, full_name, password, status, role, created_at, updated_at)
		VALUES (:id, :username, :email, :phone, :photo, :full_name, :password, :status, :role, :created_at, :updated_at)`

	_, err := n.sql.Db.NamedExecContext(c, statement, user)

	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, biedeptrai.ErrorUserConflict
			}

		}
		return user, biedeptrai.ErrorSignUpFail
	}

	return user, nil

}

func (n *UserRepoImpl) CheckLogin(context context.Context, loginRequest request.UserLoginRequest) (model.User, error) {
	var user = model.User{}
	statement := `SELECT * FROM users WHERE username=$1`

	err := n.sql.Db.GetContext(context, &user, statement, loginRequest.Username)

	if err != nil {
		log.Error(err)
		if errors.Cause(err) == sql.ErrNoRows {
			return user, biedeptrai.ErrorUserNotFound
		}
		if err, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", err.Code.Name())
		}
		return user, err
	}
	return user, nil

}

func (n *UserRepoImpl) UpdateUser(context context.Context, user model.User) (model.User, error) {

	statement := `UPDATE users
	SET username = :username,  email=:email, phone= :phone , photo = :photo, full_name = :full_name, role=:role
	WHERE id = :id
	`

	_, err := n.sql.Db.NamedExecContext(context, statement, user)

	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", err.Constraint)
			if err.Code.Name() == "unique_violation" {
				if err.Constraint == "idx_users_username" {
					return user, biedeptrai.ErrorUserDupUsername
				}
				if err.Constraint == "idx_users_phone" {
					return user, biedeptrai.ErrorUserDupPhone
				}
				if err.Constraint == "idx_users_email" {
					return user, biedeptrai.ErrorUserDupEmail
				}

			}
		}
		return user, biedeptrai.ErrorUserNotUpdated
	}
	return user, nil

}

/*
func (n *UserRepoImpl) SelectUserId(context context.Context, userId string) (model.User, error) {
	var user model.User
	if res := n.sql.Db.Where(
		&model.User{Id: userId},
	).First(&user); res.RowsAffected <= 0 {
		return user, banana.ErrorUserNotFound
	}
	return user, nil

} */
/*
func (n *UserRepoImpl) UpdateUser(context context.Context, user model.User) (model.User, error) {
	user.UpdatedAt = time.Now()

	if res := n.sql.Db.Where(
		&model.User{Id: user.Id},
	).Save(&user); res.RowsAffected <= 0 {
		return user, banana.ErrorUserNotFound
	}

	return user, nil

} */
