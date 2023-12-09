package errCode

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
)

func WrapBizErr(ctx *app.RequestContext, err error, bizErrCode BizErrCode) *errors.Error {
	return ctx.Error(errors.New(err, errors.ErrorTypePublic, bizErrCode))
}
