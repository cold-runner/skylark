package code

//go:generate ./codegen -type=int
const (
	// ErrUserAlreadyExist - 400: 用户已存在
	ErrUserAlreadyExist = iota + 110001
	// ErrRegister - 400: 用户注册失败
	ErrRegister
)
