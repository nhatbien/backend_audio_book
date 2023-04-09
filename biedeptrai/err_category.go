package biedeptrai

import "errors"

var (
	ErrorCategoryConflict   = errors.New("chuyên mục đã tồn tại")
	ErrorCategoryDeleteBook = errors.New("vui lòng xóa hết sách trong chuyên mục này trước khi xóa chuyên mục")
	ErrorCategoryNotFound   = errors.New("chuyên mục không tồn tại")
	ErrorCategoryUpdate     = errors.New("lỗi khi cập nhật chuyên mục")
)
