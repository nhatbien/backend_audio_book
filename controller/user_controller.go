package controller

import (
	"backend/biedeptrai"
	"backend/log"
	"backend/model"
	"backend/model/request"
	"backend/repository"
	"backend/security"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserRepo repository.UserRepo
}

func (u *UserController) Signup(c echo.Context) error {

	request := request.UserSignupRequest{}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	hash := security.HashAndSalt([]byte(request.Password))
	//role := model.MEMBER.String()
	//role := model.ADMIN.String()
	userId, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusForbidden, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	userModel := model.User{
		Id:        userId.String(),
		FullName:  request.FullName,
		Phone:     request.Phone,
		Photo:     request.Photo,
		Email:     request.Email,
		Username:  request.Username,
		Age:       request.Age,
		Address:   request.Address,
		Password:  hash,
		RoleId:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user, err := u.UserRepo.SaveUser(c.Request().Context(), userModel)

	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  false,
		Message: "Đăng ký thành công",
		Data:    user,
	})
}

func (u *UserController) Login(c echo.Context) error {
	request := request.UserLoginRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	user, err := u.UserRepo.Login(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	isTheSame := security.ComparePasswords(user.Password, []byte(request.Password))

	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Status:  false,
			Message: "Đăng nhậP thất bại",
			Data:    nil,
		})
	}

	//gentoken is require
	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	user.Token = token

	return c.JSON(http.StatusOK, user)
}

func (u *UserController) Update(c echo.Context) error {
	request := request.UserUpdateRequest{}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	user := model.User{
		Id:        claims.Id,
		Username:  request.Username,
		Email:     request.Email,
		Phone:     request.Phone,
		Photo:     request.Photo,
		FullName:  request.FullName,
		Status:    request.Status,
		Age:       request.Age,
		Address:   request.Address,
		UpdatedAt: time.Now(),
	}

	_, err := u.UserRepo.UpdateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lưu thành công",
	})
}

func (u *UserController) UpdateRole(c echo.Context) error {
	request := request.UserUpdateRoleRequest{}
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	if claims.Role.RoleName != "ADMIN" {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: biedeptrai.ErrorRoleUser.Error(),
			Data:    nil,
		})
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	/* err := u.UserRepo.UpdateRole(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
		})
	} */
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lưu thành công",
	})
}
