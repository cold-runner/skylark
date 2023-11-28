package v1

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/internal/infrastructure/errCode"
)

func (c *userInterface) GetLarkInfo(ctx context.Context, context *app.RequestContext) {
	hlog.Debug("调用查询用户信息方法")

	var tmp struct {
		StuNum string `vd:"stuNum($)" json:"stuNum"`
	}

	if err := context.BindAndValidate(&tmp); err != nil {
		errCode.ResponseWithError(context, errCode.ErrValidation, nil)
		return
	}

	larkInfo, err := c.application.GetLarkInfo(ctx, context, tmp.StuNum)
	if err != nil {
		errCode.ResponseError(context, nil)
		return
	}

	errCode.ResponseOk(context, larkInfo)
}

func (c *userInterface) BindQq(ctx context.Context, context *app.RequestContext) {
	hlog.Debug("调用绑定QQ唯一Id")

	var tmp struct {
		StuNum  string `vd:"stuNum($)" json:"stuNum"`
		UnionId string `vd:"qqUnionId($)" json:"unionId"`
	}

	if err := context.BindAndValidate(&tmp); err != nil {
		errCode.ResponseWithError(context, errCode.ErrValidation, nil)
		return
	}

	if err := c.application.BindQq(ctx, context, tmp.StuNum, tmp.UnionId); err != nil {
		errCode.ResponseError(context, nil)
		return
	}

	errCode.ResponseOk(context, nil)
}
