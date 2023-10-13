package code

import "net/http"

// init register error codes defines in this source code to `github.com/marmotedu/errors`
func init() {
	register(ErrSuccess, http.StatusOK, "OK")
	register(ErrUnknown, http.StatusInternalServerError, "服务器内部错误")
	register(ErrBind, http.StatusInternalServerError, "将请求正文绑定到结构时出错")
	register(ErrValidation, http.StatusBadRequest, "校验失败")
	register(ErrSmsCode, http.StatusBadRequest, "验证码错误")
	register(ErrSmsCodeExpired, http.StatusBadRequest, "验证码过期")
	register(ErrTokenInvalid, http.StatusBadRequest, "无效的Token")
	register(ErrPageNotFound, http.StatusNotFound, "资源未找到")
	register(ErrOss, http.StatusInternalServerError, "对象存储错误")
	register(ErrDatabase, http.StatusInternalServerError, "数据库错误")
	register(ErrEncodingFailed, http.StatusInternalServerError, "由于数据有误，编码失败")
	register(ErrDecodingFailed, http.StatusInternalServerError, "由于数据错误，解码失败")
	register(ErrInvalidJSON, http.StatusBadRequest, "数据不是有效的json格式")
	register(ErrEncodingJSON, http.StatusInternalServerError, "无法对JSON数据进行编码")
	register(ErrDecodingJSON, http.StatusInternalServerError, "无法解码JSON数据")
	register(ErrInvalidYaml, http.StatusBadRequest, "无效的yaml数据")
	register(ErrEncodingYaml, http.StatusInternalServerError, "无法对yaml数据进行编码")
	register(ErrDecodingYaml, http.StatusInternalServerError, "无法解码yaml数据")
	register(ErrExpired, http.StatusBadRequest, "token已过期")
	register(ErrSignatureInvalid, http.StatusBadRequest, "无效的签名")
	register(ErrInvalidAuthHeader, http.StatusBadRequest, "无效的认证头")
	register(ErrMissingHeader, http.StatusBadRequest, "认证头为空")
	register(ErrPasswordIncorrect, http.StatusBadRequest, "密码错误")
	register(ErrPermissionDenied, http.StatusForbidden, "权限不足")
	register(ErrUserAlreadyExist, http.StatusBadRequest, "用户已注册")
	register(ErrLoginType, http.StatusBadRequest, "不支持的登陆类型")
	register(ErrSocialNotExist, http.StatusBadRequest, "请先注册后再进行第三方登陆")
	register(ErrUserNotExist, http.StatusBadRequest, "不存在的用户")
	register(ErrAlreadySendSmsCode, http.StatusBadRequest, "发送验证码频率过高")
}
