package repository_impl

import (
	"backend/biedeptrai"
	"backend/db"
	"backend/helper"
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

	err := helper.WithTransaction(c, n.sql.Db, func(tx helper.Transaction) error {
		statement := `INSERT INTO users(
			id, username, email, phone, password, full_name, age, address, photo, status,role_id, created_at, updated_at)
			VALUES (:id, :username, :email, :phone, :password, :full_name, :age ,:address, :photo , :status, :role_id, :created_at, :updated_at)`

		_, err := tx.NamedExecContext(c, statement, user)

		if err != nil {
			return err
		}
		/* _, err = tx.ExecContext(c, "INSERT INTO user_role(user_id, role_id) VALUES($1, $2)", user.Id, 1)

		if err != nil {
			return err
		} */
		return nil

	})
	///insert user to sql

	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				if err.Constraint == "users_username_key" {
					return user, biedeptrai.ErrorUserDupUsername
				}
				if err.Constraint == "users_phone_key" {
					return user, biedeptrai.ErrorUserDupPhone
				}
				if err.Constraint == "users_email_key" {
					return user, biedeptrai.ErrorUserDupEmail
				}
				return user, biedeptrai.ErrorUserConflict
			}

		}
		return user, biedeptrai.ErrorSignUpFail
	}

	return user, nil

}

func (n *UserRepoImpl) CheckLogin(context context.Context, loginRequest request.UserLoginRequest) (model.User, error) {
	var user = model.User{}
	var role = model.Role{}
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
	statement2 := `SELECT * FROM role WHERE id=$1`

	err = n.sql.Db.GetContext(context, &role, statement2, user.RoleId)
	user.Role = role

	if err != nil {
		return user, err
	}

	return user, nil

}

func (n *UserRepoImpl) UpdateUser(context context.Context, user model.User) error {

	statement := `UPDATE users
	SET  age=:age,  address=:address,  email=:email, phone= :phone , photo = :photo, full_name = :full_name
	WHERE id = :id
	`

	_, err := n.sql.Db.NamedExecContext(context, statement, user)

	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", err.Constraint)
			if err.Code.Name() == "unique_violation" {
				if err.Constraint == "users_username_key" {
					return biedeptrai.ErrorUserDupUsername
				}
				if err.Constraint == "users_phone_key" {
					return biedeptrai.ErrorUserDupPhone
				}
				if err.Constraint == "users_email_key" {
					return biedeptrai.ErrorUserDupEmail
				}

			}
		}
		return biedeptrai.ErrorUserNotUpdated
	}
	return nil

}

func (n *UserRepoImpl) UpdateRole(context context.Context, userRole request.UserUpdateRoleRequest) error {

	statement := `UPDATE user_role
	SET role_id = :role_id
	WHERE user_id = :user_id
	`

	_, err := n.sql.Db.NamedExecContext(context, statement, userRole)

	if err != nil {
		log.Error(err.Error())
		return biedeptrai.ErrorUserNotUpdated
	}
	return nil

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
