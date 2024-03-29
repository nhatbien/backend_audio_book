package repository

import "backend/model"

type CartRepo interface {
	SaveCart(cart model.Cart) (model.Cart, error)
	UpdateCart(cart model.Cart) (model.Cart, error)
	DeleteCart(userId string) error
	SelectCartById(cartId int) (model.Cart, error)
	SelectMyCart(userID string) (model.Cart, error)
	AddItemToCart(userID string, cartItem model.CartItem) (model.Cart, error)
	DeleteCartItem(cartItemId int) error
}
