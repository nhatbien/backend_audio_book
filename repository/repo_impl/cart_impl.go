package repo_impl

import (
	"backend/biedeptrai"
	"backend/db"
	"backend/model"
	"backend/repository"

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
	if n.sql.Db.Where("id = ? AND is_current = ?", userId, true).Preload("Items").First(&cart).RowsAffected <= 0 {

		cart = model.Cart{
			UserId: userId,
		}
		err := n.sql.Db.Create(&cart).Error
		if err != nil {
			return cart, err
		}

	}
	if n.sql.Db.First(&book, cartItem.BookId).RowsAffected <= 0 {
		return cart, biedeptrai.ErrorBookNotFound
	}

	if n.sql.Db.Where("book_id = ? AND cart_id = ?", cartItem.BookId, cart.Id).First(&cartItem).RowsAffected > 0 {
		cartItem.Quantity += 1
		cartItem.TotalAmount = book.Price * float64(cartItem.Quantity)
		if err := n.sql.Db.Updates(&cartItem).Error; err != nil {
			return cart, err
		}
		if n.sql.Db.Preload(clause.Associations).First(&cart, cart.Id).RowsAffected <= 0 {

			return cart, biedeptrai.ErrCartNotFound
		}
		for _, item := range cart.Items {
			cart.TotalPrice += item.TotalAmount
		}
		if err := n.sql.Db.Updates(&cart).Error; err != nil {
			return cart, err
		}

		if n.sql.Db.Preload("Items.Book").Find(&cart, cart.Id).RowsAffected <= 0 {

			return cart, biedeptrai.ErrCartNotFound
		}

		return cart, nil
	}

	cartItem.CartId = cart.Id
	cartItem.TotalAmount = book.Price * float64(cartItem.Quantity)
	err := n.sql.Db.Create(&cartItem).Error
	if err != nil {
		return cart, err
	}

	if n.sql.Db.Preload("Items.Book").First(&cart, cart.Id).RowsAffected <= 0 {

		return cart, biedeptrai.ErrCartNotFound
	}

	for _, item := range cart.Items {
		cart.TotalPrice += item.TotalAmount
	}

	if err := n.sql.Db.Updates(&cart).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

func (n *CartRepoImpl) SelectMyCart(userId string) (model.Cart, error) {
	var cart model.Cart
	if n.sql.Db.Where("user_id = ? AND is_current = ?", userId, true).Preload("Items.Book").Find(&cart).RowsAffected <= 0 {
		return cart, biedeptrai.ErrCartNotFound
	}
	return cart, nil

}
