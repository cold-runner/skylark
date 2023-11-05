package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/errCode"
)

func (c *controllerV1) Register(ctx context.Context, context *app.RequestContext) {
	hlog.Info("调用注册方法")

	newer := &user.Register{}
	// 参数校验
	if err := context.BindAndValidate(newer); err != nil {
		errCode.ResponseWithError(context, errCode.ErrValidation, nil)
		return
	}

	// 移交服务层
	if err := c.serviceIns.Register(ctx, context, newer); err != nil {
		errCode.ResponseError(context, nil)
		return
	}

	errCode.ResponseOk(context, nil)
}

func (c *controllerV1) SendSms(ctx context.Context, context *app.RequestContext) {
	hlog.Debug("调用发送验证码方法")
	var tmp struct {
		Phone string `vd:"phone($)" json:"phone,required"`
	}
	// 参数校验
	if err := context.BindAndValidate(&tmp); err != nil {
		errCode.ResponseWithError(context, errCode.ErrValidation, nil)
		return
	}

	// 移交服务层
	if err := c.serviceIns.SendSmsCode(ctx, context, tmp.Phone); err != nil {
		errCode.ResponseError(context, nil)
		return
	}

	errCode.ResponseOk(context, nil)
}
