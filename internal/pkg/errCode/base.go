package code

import "github.com/cloudwego/hertz/pkg/common/errors"

// Common: 基本错误
// Code must start with 10xxxx.

const (
	ErrSuccess            errors.ErrorType = iota + 100001 // 200: OK
	ErrUnknown                                             // 500: 服务器内部错误
	ErrBind                                                // 500: 将请求正文绑定到结构时出错
	ErrValidation                                          // 400: 校验失败
	ErrSmsCode                                             // 400: 验证码错误
	ErrSmsCodeExpired                                      // 400: 验证码过期
	ErrAlreadySendSmsCode                                  // 400: 验证码已发送且未过期
	ErrPageNotFound                                        // 404: 资源未找到
	ErrOss                                                 // 500: 对象存储错误
	ErrTokenInvalid                                        // 400: 无效的Token
	ErrInvalidAuthHeader                                   // 400: 无效的认证头
	ErrDatabase                                            //500: 数据库错误
)

var (
	ErrorSuccess = errors.New(nil, ErrSuccess, "OK")
)
