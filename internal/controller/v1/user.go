package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/marmotedu/errors"
)

func (c *controllerV1) BindQq(ctx context.Context, context *app.RequestContext) {
	var tmp struct {
		StuNum  string `vd:"stuNum($)" json:"stuNum"`
		UnionId string `vd:"qqUnionId($)" json:"unionId"`
	}
	// 参数校验
	if err := context.BindAndValidate(&tmp); err != nil {
		code.WriteResponse(context, errors.WithCode(code.ErrValidation, "校验失败"), nil)
		return
	}

	// 移交服务层
	if err := c.serviceIns.BindQq(ctx, tmp.StuNum, tmp.UnionId); err != nil {
		code.WriteResponse(context, err, nil)
		return
	}

	code.WriteResponse(context, errors.WithCode(code.ErrSuccess, "", nil), nil)
}

func (c *controllerV1) ChangePassword(ctx context.Context, context *app.RequestContext) {

}
