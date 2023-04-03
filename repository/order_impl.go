package repository

import "backend/model"

type OrderRepo interface {
	SaveOrder(order model.Order) (model.Order, error)
	UpdateOrder(order model.Order) (model.Order, error)
	DeleteOrder(orderId uint) error
	SelectAllOrder() ([]model.Order, error)
	SelectOrderById(orderId uint) (model.Order, error)
}
