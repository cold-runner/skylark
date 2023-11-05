package errCode

import "github.com/cloudwego/hertz/pkg/common/errors"

// Common: 基本错误
// Code must start with 10xxxx.

const (
	errSuccess            errors.ErrorType = iota + 100001 // 200: OK
	errUnknown                                             // 500: 服务器内部错误
	errValidation                                          // 400: 校验不通过
	errSmsCode                                             // 400: 验证码错误
	errSmsCodeExpired                                      // 400: 验证码过期
	errAlreadySendSmsCode                                  // 400: 验证码已发送且未过期
	errResourceNotFound                                    // 404: 资源未找到
	errOss                                                 // 500: 对象存储错误
	errTokenInvalid                                        // 400: 无效的Token
	errInvalidAuthHeader                                   // 400: 无效的认证头
	errDatabase                                            // 500: 数据库错误
	errDeleteCache                                         // 500: 删除缓存错误
)
