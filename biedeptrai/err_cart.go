package biedeptrai

import "errors"

var (
	ErrCartNotFound        = errors.New("cart not found")
	ErrorItemAlreadyInCart = errors.New("item already in cart")
	ErrCartItemNotFound    = errors.New("cart item not found")
)
