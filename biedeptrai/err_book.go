package biedeptrai

import "errors"

var (
	ErrorBookConflict      = errors.New("sách đã tồn tại")
	ErrorBookNotFound      = errors.New("sách không tồn tại")
	ErrorBookAlreadyInCart = errors.New("sách đã có trong giỏ hàng")
)
