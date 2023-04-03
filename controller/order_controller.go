package controller

import (
	"backend/model"
	"backend/model/request"
	"backend/repository"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
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
		UserId: claims.Id,
		CartId: request.CartId,
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
