package errCode

import (
	stdErr "errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var (
	ErrSuccess            = errors.New(stdErr.New("OK"), errSuccess, nil)
	ErrUnknown            = errors.New(stdErr.New("服务器内部错误"), errUnknown, nil)
	ErrValidation         = errors.New(stdErr.New("校验不通过"), errValidation, nil)
	ErrSmsCode            = errors.New(stdErr.New("验证码错误"), errSmsCode, nil)
	ErrSmsCodeExpired     = errors.New(stdErr.New("验证码过期"), errSmsCodeExpired, nil)
	ErrAlreadySendSmsCode = errors.New(stdErr.New("验证码已发送且未过期"), errAlreadySendSmsCode, nil)
	ErrResourceNotFound   = errors.New(stdErr.New("资源未找到"), errResourceNotFound, nil)
	ErrOss                = errors.New(stdErr.New("对象存储错误"), errOss, nil)
	ErrTokenInvalid       = errors.New(stdErr.New("无效的Token"), errTokenInvalid, nil)
	ErrInvalidAuthHeader  = errors.New(stdErr.New("无效的认证头"), errInvalidAuthHeader, nil)
	ErrDatabase           = errors.New(stdErr.New("数据库错误"), errDatabase, nil)

	ErrCreateDraft = errors.New(stdErr.New("创建草稿错误"), errCreateDraft, nil)
	ErrSavePost    = errors.New(stdErr.New("保存文章错误"), errSavePost, nil)

	ErrUserAlreadyExist  = errors.New(stdErr.New("用户已存在"), errUserAlreadyExist, nil)
	ErrUserNotExist      = errors.New(stdErr.New("不存在的用户"), errUserNotExist, nil)
	ErrLoginType         = errors.New(stdErr.New("不支持的登陆类型"), errLoginType, nil)
	ErrPasswordIncorrect = errors.New(stdErr.New("密码错误"), errPasswordIncorrect, nil)
	ErrPermissionDenied  = errors.New(stdErr.New("权限不足"), errPermissionDenied, nil)
	ErrRepeatBindQq      = errors.New(stdErr.New("重复绑定qq错误"), errRepeatBindQq, nil)
	ErrDeleteCache       = errors.New(stdErr.New("删除缓存错误"), errDeleteCache, nil)
)

func ResponseOk(ctx *app.RequestContext, data interface{}) {
	ctx.JSON(consts.StatusOK, utils.H{
		"code": ErrSuccess.Type,
		"msg":  ErrSuccess.Error(),
		"data": data,
	})
}

func ResponseWithError(ctx *app.RequestContext, err *errors.Error, data interface{}) {
	ctx.JSON(consts.StatusOK, utils.H{
		"code": err.Type,
		"msg":  err.Error(),
		"data": data,
	})
}

func ResponseError(ctx *app.RequestContext, data interface{}) {
	bizErr := ctx.Errors.Last()
	ctx.JSON(consts.StatusOK, utils.H{
		"code": bizErr.Type,
		"msg":  bizErr.Error(),
		"data": data,
	})
}
