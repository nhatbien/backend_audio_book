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

type CartRepoImpl struct {
	sql *db.Sql
}

func NewCartRepo(sql *db.Sql) repository.CartRepo {
	return &CartRepoImpl{sql: sql}
}

func (n *CartRepoImpl) SaveCart(cart model.Cart) (model.Cart, error) {

	err := n.sql.Db.Create(&cart).Error

	if err != nil {
		return cart, err
	}
	return cart, nil

}

func (n *CartRepoImpl) UpdateCart(cart model.Cart) (model.Cart, error) {

	err := n.sql.Db.Updates(&cart).Error

	if err != nil {
		return cart, err
	}
	return cart, nil

}

func (n *CartRepoImpl) DeleteCart(cartId int) error {

	err := n.sql.Db.Delete(&model.Cart{}, cartId).Error

	if err != nil {
		return err
	}
	return nil

}
func (n *CartRepoImpl) SelectCartById(cartId int) (model.Cart, error) {
	var cart model.Cart
	err := n.sql.Db.First(&cart, cartId).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (n *CartRepoImpl) AddItemToCart(userId string, cartItem model.CartItem) (model.Cart, error) {

	var cart model.Cart
	var book model.Book
	if n.sql.Db.Where("id = ?", userId).First(new(model.User)).RowsAffected <= 0 {
		return cart, biedeptrai.ErrorUserNotFound
	}

	n.sql.Db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')

		if tx.Where("user_id = ? AND is_current = ?", userId, 1).Preload("Items").First(&cart).RowsAffected <= 0 {

			cart = model.Cart{
				UserId: userId,
			}
			err := tx.Create(&cart).Error
			if err != nil {
				return err
			}

		}
		if tx.First(&book, cartItem.BookId).RowsAffected <= 0 {
			return biedeptrai.ErrorBookNotFound
		}
		if tx.Where("book_id = ? AND cart_id = ?", cartItem.BookId, cart.Id).First(&cartItem).RowsAffected > 0 {
			cartItem.Quantity += 1
			cartItem.TotalAmount = book.Price * float64(cartItem.Quantity)
			if err := tx.Updates(&cartItem).Error; err != nil {
				return err
			}
			if tx.Preload(clause.Associations).First(&cart, cart.Id).RowsAffected <= 0 {

				return biedeptrai.ErrCartNotFound
			}
			for _, item := range cart.Items {
				cart.TotalPrice += item.TotalAmount
			}
			if err := tx.Updates(&cart).Error; err != nil {
				return err
			}

			if tx.Preload("Items.Book").Find(&cart, cart.Id).RowsAffected <= 0 {

				return biedeptrai.ErrCartNotFound
			}

		}

		cartItem.CartId = cart.Id
		cartItem.TotalAmount = book.Price * float64(cartItem.Quantity)
		err := tx.Create(&cartItem).Error
		if err != nil {
			return err
		}

		if tx.Preload("Items.Book").First(&cart, cart.Id).RowsAffected <= 0 {

			return biedeptrai.ErrCartNotFound
		}

		for _, item := range cart.Items {
			cart.TotalPrice += item.TotalAmount
		}
		cart.UpdatedAt = time.Now()
		if err := tx.Updates(&cart).Error; err != nil {
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

	return cart, nil
}

func (n *CartRepoImpl) SelectMyCart(userId string) (model.Cart, error) {
	var cart model.Cart
	if n.sql.Db.Where("user_id = ? AND is_current = ?", userId, 1).Preload("Items.Book").Find(&cart).RowsAffected <= 0 {
		return cart, biedeptrai.ErrCartNotFound
	}
	return cart, nil

}
