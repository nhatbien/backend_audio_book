package controller

import (
	"backend/model"
	"backend/model/request"
	"backend/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type CartController struct {
	CartRepo repository.CartRepo
}

func (n *CartController) AddItemToCart(c echo.Context) error {
	request := request.CartItemSave{}
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
	cartItem := model.CartItem{
		BookId:    request.BookId,
		Quantity:  request.Quantity,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	cart, err := n.CartRepo.AddItemToCart(claims.Id, cartItem)

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
		Data:    cart,
	})
}

func (n *CartController) SelectMyCart(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	cart, err := n.CartRepo.SelectMyCart(claims.Id)
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
		Data:    cart,
	})

}

func (n *CartController) DeleteCart(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	err := n.CartRepo.DeleteCart(claims.Id)

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

func (n *CartController) DeleteCartItem(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	err = n.CartRepo.DeleteCartItem(id)

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
