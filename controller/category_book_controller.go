package controller

import (
	"backend/biedeptrai"
	"backend/model"
	"backend/model/request"
	"backend/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type CategoryBookController struct {
	CategoryBookRepo repository.CategoryBookRepo
}

func (b *CategoryBookController) SaveCategoryBook(c echo.Context) error {
	request := request.CategoryBookSave{}
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

	if claims.Role.RoleName != "admin" {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: biedeptrai.ErrorRoleUser.Error(),
			Data:    nil,
		})
	}

	category := model.BookCategory{
		Name:        request.Name,
		Description: request.Description,
		Images:      request.Images,
	}

	response, err := b.CategoryBookRepo.SaveCategory(category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lưu thành công",
		Data:    response,
	})

}

func (b *CategoryBookController) UpdateCategoryBook(c echo.Context) error {
	request := request.CategoryBookSave{}

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

	if claims.Role.RoleName != "admin" {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: biedeptrai.ErrorRoleUser.Error(),
			Data:    nil,
		})
	}
	idOrder, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: "Vui lòng điền đúng ID chuyên mục",
			Data:    nil,
		})
	}

	category := model.BookCategory{
		Name:        request.Name,
		Description: request.Description,
		Images:      request.Images,
	}

	response, err := b.CategoryBookRepo.UpdateCategory(category, uint(idOrder))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lưu thành công",
		Data:    response,
	})

}

func (b *CategoryBookController) GetAllCategoryBook(c echo.Context) error {

	response, err := b.CategoryBookRepo.SelectAllCategory()
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lấy thành công",
		Data:    response,
	})

}

func (b *CategoryBookController) GetCategoryBookById(c echo.Context) error {
	idOrder, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: "Vui lòng điền đúng ID chuyên mục",
			Data:    nil,
		})
	}
	response, err := b.CategoryBookRepo.SelectCategoryById(idOrder)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lấy thành công",
		Data:    response,
	})
}
