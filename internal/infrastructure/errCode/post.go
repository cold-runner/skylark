package errCode

import "github.com/cloudwego/hertz/pkg/common/errors"

// post: 文章功能错误
// errCode must start with 30xxxx.

const (
	errCreateDraft errors.ErrorType = iota + 300001 // 500: 保存文章错误
	errSavePost
)
