package code

// Common: 基本错误
// Code must start with 1xxxxx.
const (
	// ErrSuccess - 200: OK
	ErrSuccess int = iota + 100001

	// ErrUnknown - 500: 服务器内部错误
	ErrUnknown

	// ErrBind - 500: 将请求正文绑定到结构时出错
	ErrBind

	// ErrValidation - 400: 校验失败
	ErrValidation

	// ErrSmsCode - 400: 验证码错误
	ErrSmsCode

	// ErrSmsCodeExpired - 400: 验证码过期
	ErrSmsCodeExpired

	// ErrAlreadySendSmsCode - 400: 验证码已发送且未过期
	ErrAlreadySendSmsCode

	// ErrTokenInvalid - 400: 无效的Token
	ErrTokenInvalid

	// ErrPageNotFound - 404: 资源未找到
	ErrPageNotFound

	// ErrOss - 500: 对象存储错误
	ErrOss
)

// common: 数据库错误
const (
	// ErrDatabase - 500: 数据库错误
	ErrDatabase int = iota + 100101
)

// common: 校验相关错误
const (
	// ErrExpired - 400: Token已过期
	ErrExpired = iota + 100201

	// ErrSignatureInvalid - 400: 无效的签名
	ErrSignatureInvalid

	// ErrInvalidAuthHeader - 400: 无效的认证头
	ErrInvalidAuthHeader

	// ErrMissingHeader - 400: 认证头为空
	ErrMissingHeader

	// ErrLoginType - 400: 不支持的登陆类型
	ErrLoginType

	// ErrPasswordIncorrect - 400: 密码错误
	ErrPasswordIncorrect

	// ErrSocialNotExist - 400: 请先注册后再进行第三方登陆
	ErrSocialNotExist

	// ErrPermissionDenied - 403: 权限不足，拒绝请求
	ErrPermissionDenied
)

// common: 编解码错误
const (
	// ErrEncodingFailed - 500: 由于数据有误，编码失败
	ErrEncodingFailed int = iota + 100301

	// ErrDecodingFailed - 500: 由于数据错误，解码失败
	ErrDecodingFailed

	// ErrInvalidJSON - 500: 数据不是有效的json格式
	ErrInvalidJSON

	// ErrEncodingJSON - 500: 无法对JSON数据进行编码
	ErrEncodingJSON

	// ErrDecodingJSON - 500: 无法解码JSON数据
	ErrDecodingJSON

	// ErrInvalidYaml - 500: 无效的yaml数据
	ErrInvalidYaml

	// ErrEncodingYaml - 500: 无法对yaml数据进行编码
	ErrEncodingYaml

	// ErrDecodingYaml - 500: 无法解码yaml数据
	ErrDecodingYaml
)
