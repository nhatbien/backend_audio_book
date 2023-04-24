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

type OrderController struct {
	OrderRepo repository.OrderRepo
}

func (o *OrderController) SaveOrder(e echo.Context) error {
	request := request.OrderSave{}
	tokenData := e.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	if err := e.Bind(&request); err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	if err := e.Validate(request); err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	orderModel := model.Order{
		UserId:    claims.Id,
		CartId:    request.CartId,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	order, err := o.OrderRepo.SaveOrder(orderModel)
	if err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return e.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Thành công",
		Data:    order,
	})

}
func (o *OrderController) PutOrderStatus(e echo.Context) error {
	request := request.OrderStatusChange{}
	id, err := strconv.Atoi(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := e.Bind(&request); err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})

	}

	if err := e.Validate(request); err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	order := model.Order{
		Id:        uint(id),
		UpdatedAt: time.Now(),
		Status:    request.Status,
	}
	orders, err := o.OrderRepo.PutOrderStatus(order)
	if err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return e.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Thành công",
		Data:    orders,
	})
}

func (o *OrderController) SelectOrderById(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	order, err := o.OrderRepo.SelectOrderById(uint(id))
	if err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return e.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Thành công",
		Data:    order,
	})
}

func (o *OrderController) SelectOrderByStatus(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("status"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	orders, err := o.OrderRepo.SelectOrderByStatus(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return e.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Thành công",
		Data:    orders,
	})
}

func (o *OrderController) SelectAllBookOrderbyStatusAndUserId(e echo.Context) error {
	tokenData := e.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)
	id, err := strconv.Atoi(e.Param("status"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	books, err := o.OrderRepo.SelectAllBookOrderbyStatusAndUserId(claims.Id, id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return e.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Thành công",
		Data:    books,
	})
}

func (o *OrderController) SelectAllOrderbyStatusAndUserId(e echo.Context) error {
	tokenData := e.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)
	id, err := strconv.Atoi(e.Param("status"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	orders, err := o.OrderRepo.SelectOrderbyStatusAndUserId(claims.Id, id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return e.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Thành công",
		Data:    orders,
	})
}
