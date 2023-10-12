package code

import "net/http"

// init register error codes defines in this source code to `github.com/marmotedu/errors`
func init() {
	register(ErrSuccess, http.StatusOK, "OK")
	register(ErrUnknown, http.StatusInternalServerError, "服务器内部错误")
	register(ErrBind, 400, "将请求正文绑定到结构时出错")
	register(ErrValidation, 400, "校验失败")
	register(ErrSmsCode, 400, "验证码错误")
	register(ErrSmsCodeExpired, 400, "验证码过期")
	register(ErrTokenInvalid, 401, "无效的Token")
	register(ErrPageNotFound, 404, "资源未找到")
	register(ErrOss, 500, "对象存储错误")
	register(ErrDatabase, 500, "数据库错误")
	register(ErrEncodingFailed, 500, "由于数据有误，编码失败")
	register(ErrDecodingFailed, 500, "由于数据错误，解码失败")
	register(ErrInvalidJSON, 500, "数据不是有效的json格式")
	register(ErrEncodingJSON, 500, "无法对JSON数据进行编码")
	register(ErrDecodingJSON, 500, "无法解码JSON数据")
	register(ErrInvalidYaml, 500, "无效的yaml数据")
	register(ErrEncodingYaml, 500, "无法对yaml数据进行编码")
	register(ErrDecodingYaml, 500, "无法解码yaml数据")
	register(ErrExpired, 400, "token已过期")
	register(ErrSignatureInvalid, 400, "无效的签名")
	register(ErrInvalidAuthHeader, 400, "无效的认证头")
	register(ErrMissingHeader, 400, "认证头为空")
	register(ErrPasswordIncorrect, 400, "密码错误")
	register(ErrPermissionDenied, 401, "权限不足")
	register(ErrUserAlreadyExist, 400, "用户已注册")
	register(ErrRegister, 400, "用户注册失败")
	register(ErrLoginType, 401, "不支持的登陆类型")
	register(ErrSocialNotExist, 401, "请先注册后再进行第三方登陆")
	register(ErrUserNotExist, 400, "不存在的用户")
	register(ErrAlreadySendSmsCode, 400, "发送验证码频率过高")
}
