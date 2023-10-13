package code

//go:generate ./codegen -type=int
const (
	// ErrUserAlreadyExist - 400: 用户已存在
	ErrUserAlreadyExist = iota + 110001

	// ErrUserNotExist - 400: 不存在的用户
	ErrUserNotExist

	// ErrRepeatBindQq - 400: 重复绑定qq错误
	ErrRepeatBindQq
)
