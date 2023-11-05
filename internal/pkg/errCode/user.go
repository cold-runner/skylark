package code

import "github.com/cloudwego/hertz/pkg/common/errors"

// user: 用户服务类错误
// code must start with 40xxxx.

const (
	ErrUserAlreadyExist  errors.ErrorType = iota + 400001 // 400: 用户已存在
	ErrLoginType                                          // 400: 不支持的登陆类型
	ErrPasswordIncorrect                                  // 400: 密码错误
	ErrUserNotExist                                       // 400: 不存在的用户
	ErrPermissionDenied                                   // 403: 权限不足
	ErrRepeatBindQq                                       // 400: 重复绑定qq错误
)
