# 错误码

skylark 系统错误码列表，由 `codegen -type=int -doc` 命令生成，不要对此文件做任何更改。

## 功能说明

返回结果中会是如下格式，存在 code 和 message 字段，表示调用 API 接口后的返回结果。例如：

```json
{
  "code": 100101,
  "message": "Database error"
}
```

上述返回中 `code` 表示错误码，`message` 表示该错误的具体信息。每个错误同时也对应一个 HTTP 状态码，比如上述错误码对应了 HTTP 状态码 500(Internal Server Error)。

## 错误码列表

skylark 支持的错误码列表如下：

| Identifier | Code | HTTP Code | Description |
| ---------- | ---- | --------- | ----------- |
| ErrSuccess | 100001 | 200 | OK |
| ErrUnknown | 100002 | 500 | 服务器内部错误 |
| ErrBind | 100003 | 400 | 将请求正文绑定到结构时出错 |
| ErrValidation | 100004 | 400 | 校验失败 |
| ErrSmsCode | 100005 | 400 | 验证码错误 |
| ErrSmsCodeExpired | 100006 | 400 | 验证码过期 |
| ErrTokenInvalid | 100007 | 401 | 无效的Token |
| ErrPageNotFound | 100008 | 404 | 资源未找到 |
| ErrOss | 100009 | 500 | 对象存储错误 |
| ErrDatabase | 100101 | 500 | 数据库错误 |
| ErrEncodingFailed | 100301 | 500 | 由于数据有误，编码失败 |
| ErrDecodingFailed | 100302 | 500 | 由于数据错误，解码失败 |
| ErrInvalidJSON | 100303 | 500 | 数据不是有效的json格式 |
| ErrEncodingJSON | 100304 | 500 | 无法对JSON数据进行编码 |
| ErrDecodingJSON | 100305 | 500 | 无法解码JSON数据 |
| ErrInvalidYaml | 100306 | 500 | 无效的yaml数据 |
| ErrEncodingYaml | 100307 | 500 | 无法对yaml数据进行编码 |
| ErrDecodingYaml | 100308 | 500 | 无法解码yaml数据 |

