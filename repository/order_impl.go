package repository

import "backend/model"

type OrderRepo interface {
	SaveOrder(order model.Order) (model.Order, error)
	PutOrderStatus(order model.Order) (model.Order, error)
	UpdateOrder(order model.Order) (model.Order, error)
	DeleteOrder(orderId uint) error
	SelectAllOrder() ([]model.Order, error)
	SelectOrderById(orderId uint) (model.Order, error)
	SelectOrderByStatus(status int) ([]model.Order, error)
	SelectAllBookOrderbyStatusAndUserId(userId string, status int) ([]model.Book, error)
	SelectOrderbyStatusAndUserId(userId string, status int) ([]model.Order, error)
}
