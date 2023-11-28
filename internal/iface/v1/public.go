package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/internal/domain/entity"
	"github.com/cold-runner/skylark/internal/infrastructure/errCode"
)

func (c *userInterface) Register(ctx context.Context, context *app.RequestContext) {
	hlog.Debug("调用注册方法")

	newer := &entity.Register{}
	if err := context.BindAndValidate(newer); err != nil {
		errCode.ResponseWithError(context, errCode.ErrValidation, nil)
		return
	}

	if err := c.application.Register(ctx, context, newer); err != nil {
		errCode.ResponseError(context, nil)
		return
	}

	errCode.ResponseOk(context, nil)
}

func (c *userInterface) SendSms(ctx context.Context, context *app.RequestContext) {
	hlog.Debug("调用发送验证码方法")

	var tmp struct {
		Phone string `vd:"phone($)" json:"phone,required"`
	}

	if err := context.BindAndValidate(&tmp); err != nil {
		errCode.ResponseWithError(context, errCode.ErrValidation, nil)
		return
	}

	if err := c.application.SendSmsCode(ctx, context, tmp.Phone); err != nil {
		errCode.ResponseError(context, nil)
		return
	}

	errCode.ResponseOk(context, nil)
}
