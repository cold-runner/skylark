package errCode

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cold-runner/skylark/biz/model/response"
)

var SuccessStatus = &response.Status{
	Code: int32(ErrSuccess),
	Msg:  ErrMap[ErrSuccess].Msg,
}

func ResponseFailed(c *app.RequestContext) {
	bizErrCode := c.Errors.Last().Meta.(BizErrCode)
	bizErr := ErrMap[bizErrCode]

	c.JSON(bizErr.HttpStatus, utils.H{
		"code": bizErrCode,
		"msg":  bizErr.Msg,
	})
}

func ResponseValidationFailed(c *app.RequestContext, err error) {
	bizErr := ErrMap[ErrValidation]

	c.JSON(bizErr.HttpStatus, utils.H{
		"code": ErrValidation,
		"msg":  bizErr.Msg + ". err: " + err.Error(),
	})
}

func ResponseResourceNotFound(c *app.RequestContext) {
	bizErr := ErrMap[ErrNotFound]

	c.JSON(bizErr.HttpStatus, utils.H{
		"code": ErrNotFound,
		"msg":  bizErr.Msg,
	})
}
