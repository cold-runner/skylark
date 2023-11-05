package errCode

import "github.com/cloudwego/hertz/pkg/common/errors"

// user: 用户服务类错误
// errCode must start with 40xxxx.

const (
	errUserAlreadyExist  errors.ErrorType = iota + 400001 // 400: 用户已存在
	errUserNotExist                                       // 400: 不存在的用户
	errLoginType                                          // 400: 不支持的登陆类型
	errPasswordIncorrect                                  // 400: 密码错误
	errPermissionDenied                                   // 403: 权限不足
	errRepeatBindQq                                       // 400: 重复绑定qq错误
)
