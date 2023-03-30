package controller

import (
	"backend/biedeptrai"
	_ "backend/docs"
	"backend/model"
	"backend/model/request"
	"backend/repository"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// GetByID godoc
// @Summary Get BusinessGroup By ID
// @Description Get BusinessGroup By ID
// @Tags BusinessGroup
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Security ApiKeyAuth
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} model.Response{data=model.User} "BusinessGroup Info"
// @Failure 400,401,404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /business_groups/{id} [get]
type BookController struct {
	BookRepo repository.BookRepo
}

func (b *BookController) SaveBook(c echo.Context) error {
	request := request.BookSaveRequest{}

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
	bookModel := model.Book{
		BookName:  request.BookName,
		Author:    request.Author,
		Price:     request.Price,
		Content:   request.Content,
		Img:       request.Img,
		Audio:     request.Audio,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		Status:    1,
	}

	book, err := b.BookRepo.SaveBook(bookModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Thành công",
		Data:    book,
	})

}

func (b *BookController) SelectAllBook(c echo.Context) error {
	books, err := b.BookRepo.SelectAllBook()
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Thành công",
		Data:    books,
	})
}
