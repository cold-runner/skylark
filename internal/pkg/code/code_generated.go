// Code generated by "codegen -type=int"; DO NOT EDIT.

package code

// init register error codes defines in this source code to `github.com/marmotedu/errors`
func init() {
	register(ErrSuccess, 200, "OK")
	register(ErrUnknown, 500, "服务器内部错误")
	register(ErrBind, 400, "将请求正文绑定到结构时出错")
	register(ErrValidation, 400, "校验失败")
	register(ErrSmsCode, 400, "验证码错误")
	register(ErrSmsCodeExpired, 400, "验证码过期")
	register(ErrTokenInvalid, 401, "无效的Token")
	register(ErrCasdoorIssue, 401, "Casdoor签发token失败")
	register(ErrCasdoorRefresh, 401, "Casdoor刷新token失败")
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
}
