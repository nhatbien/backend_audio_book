package repo_impl

import (
	"backend/biedeptrai"
	"backend/db"
	"backend/model"
	"backend/repository"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepoImpl struct {
	sql *db.Sql
}

func NewOrderRepo(sql *db.Sql) repository.OrderRepo {
	return &OrderRepoImpl{sql: sql}
}

func (n *OrderRepoImpl) SaveOrder(order model.Order) (model.Order, error) {
	cartModel := model.Cart{}

	err := n.sql.Db.Transaction(func(tx *gorm.DB) error {
		if tx.First(&cartModel, order.CartId).RowsAffected <= 0 {
			return biedeptrai.ErrCartNotFound
		}
		cartModel.IsCurrent = false
		cartModel.UpdatedAt = time.Now()
		if err := tx.Model(&cartModel).Updates(
			map[string]interface{}{
				"is_current": false,
				"updated_at": time.Now(),
			},
		).Error; err != nil {
			return err
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return order, err
	}
	return order, nil
}

func (n *OrderRepoImpl) PutOrderStatus(order model.Order) (model.Order, error) {
	orderModel := model.Order{}
	if n.sql.Db.Preload("Cart.Items.Book").First(&orderModel, order.Id).RowsAffected <= 0 {

		return orderModel, biedeptrai.ErrOrderNotFound
	}
	orderModel.Status = order.Status
	orderModel.UpdatedAt = order.UpdatedAt

	err := n.sql.Db.Where(&model.Order{Id: order.Id}).Updates(&orderModel).Error
	if err != nil {
		return orderModel, err
	}
	return orderModel, nil
}

func (n *OrderRepoImpl) UpdateOrder(order model.Order) (model.Order, error) {
	err := n.sql.Db.Updates(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (n *OrderRepoImpl) DeleteOrder(orderId uint) error {
	err := n.sql.Db.Delete(&model.Order{}, orderId).Error
	if err != nil {
		return err
	}
	return nil
}

func (n *OrderRepoImpl) SelectAllOrder() ([]model.Order, error) {
	var orders []model.Order
	err := n.sql.Db.Preload(clause.Associations).Find(&orders).Error
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (n *OrderRepoImpl) SelectOrderById(orderId uint) (model.Order, error) {
	if count := n.sql.Db.First(new(model.Order), orderId).RowsAffected; count <= 0 {
		return model.Order{}, biedeptrai.ErrOrderNotFound
	}
	var order model.Order
	err := n.sql.Db.Preload("Cart.Items.Book").First(&order, orderId).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (n *OrderRepoImpl) SelectOrderByStatus(status int) ([]model.Order, error) {
	var orders []model.Order
	err := n.sql.Db.Where(&model.Order{Status: status}).Preload(clause.Associations).Preload("Cart.Items.Book").Find(&orders).Error
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (n *OrderRepoImpl) SelectOrderbyStatusAndUserId(userId string, status int) ([]model.Book, error) {
	var orders []model.Order
	var books []model.Book
	err := n.sql.Db.Where(&model.Order{UserId: userId, Status: status}).Preload("Cart.Items.Book").Find(&orders).Error

	if err != nil {
		return books, err
	}

	for _, order := range orders {
		for _, item := range order.Cart.Items {
			books = append(books, item.Book)
		}
	}

	return books, nil
}
