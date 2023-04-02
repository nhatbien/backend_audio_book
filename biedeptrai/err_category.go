package biedeptrai

import "errors"

var (
	ErrorCategoryConflict = errors.New("chuyên mục đã tồn tại")
	ErrorCategoryNotFound = errors.New("chuyên mục không tồn tại")
	ErrorCategoryUpdate   = errors.New("lỗi khi cập nhật chuyên mục")
)
