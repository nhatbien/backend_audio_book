package controller

import (
	"backend/biedeptrai"
	_ "backend/docs"
	"backend/model"
	"backend/model/request"
	"backend/repository"
	"net/http"
	"strconv"
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
		BookName:     request.BookName,
		Author:       request.Author,
		Price:        request.Price,
		Content:      request.Content,
		Img:          request.Img,
		Audio:        request.Audio,
		Status:       1,
		BookCategory: request.BookCategory,
		UpdatedAt:    time.Now(),
		CreatedAt:    time.Now(),
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

func (b *BookController) UpdateBook(c echo.Context) error {
	request := request.BookUpdateRequest{}
	bookId, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
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
		ID:           uint(bookId),
		BookName:     request.BookName,
		Author:       request.Author,
		Price:        request.Price,
		Content:      request.Content,
		Img:          request.Img,
		Audio:        request.Audio,
		IsBestSeller: request.IsBestSeller,
		Status:       1,
		UpdatedAt:    time.Now(),
		CreatedAt:    time.Now(),
	}

	books, err := b.BookRepo.UpdateBook(bookModel, request.BookCategory)
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

func (b *BookController) SelectBookById(c echo.Context) error {
	idOrder, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	books, err := b.BookRepo.SelectBookById(idOrder)
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

func (e *BookController) DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
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
	err = e.BookRepo.DeleteBook(id)
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
		Data:    nil,
	})
}

func (b *BookController) SearchBookByName(c echo.Context) error {
	bookName := c.QueryParam("name")
	books, err := b.BookRepo.SearchBookByName(bookName)
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
