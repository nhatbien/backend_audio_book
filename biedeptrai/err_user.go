package biedeptrai

import "errors"

var (
	ErrorUserConflict   = errors.New("người dùng đã tồn tại")
	ErrorUserNotFound   = errors.New("người dùng không tồn tại")
	ErrorUserNotUpdated = errors.New("cập nhật thông tin người dùng thất bại")
	ErrorSignUpFail     = errors.New("đăng ký thất bại")
	ErrorAddPostFail    = errors.New("đăng bài thất bại")
	///////
	ErrorUserDupEmail    = errors.New("email này đã có người sử dụng")
	ErrorUserDupPhone    = errors.New("số điện thoại này đã có người sử dụng")
	ErrorUserDupUsername = errors.New("username này đã có người sử dụng")

	ErrorRoleUser = errors.New("bạn không phải là admin")
)
