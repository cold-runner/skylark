package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/internal/pkg/errCode"
)

func (c *controllerV1) BindQq(ctx context.Context, context *app.RequestContext) {
	hlog.Debug("调用绑定QQ唯一Id")
	var tmp struct {
		StuNum  string `vd:"stuNum($)" json:"stuNum"`
		UnionId string `vd:"qqUnionId($)" json:"unionId"`
	}
	// 参数校验
	if err := context.BindAndValidate(&tmp); err != nil {
		errCode.ResponseWithError(context, errCode.ErrValidation, nil)
		return
	}

	// 移交服务层
	if err := c.serviceIns.BindQq(ctx, context, tmp.StuNum, tmp.UnionId); err != nil {
		errCode.ResponseError(context, nil)
		return
	}

	errCode.ResponseOk(context, nil)
}

func (c *controllerV1) ChangePassword(ctx context.Context, context *app.RequestContext) {

}
