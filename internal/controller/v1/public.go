package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"
)

func (c *controllerV1) Register(ctx context.Context, context *app.RequestContext) {
	log.V(log.DebugLevel).Info("调用注册方法")

	newer := &user.Register{}
	// 参数校验
	if err := context.BindAndValidate(newer); err != nil {
		code.WriteResponse(context, errors.WithCode(code.ErrValidation, "校验失败"), nil)
		return
	}

	// 移交服务层
	if err := c.serviceIns.Register(ctx, newer); err != nil {
		code.WriteResponse(context, err, nil)
		return
	}

	code.WriteResponse(context, errors.WithCode(code.ErrSuccess, "", nil), nil)
}

func (c *controllerV1) SendSms(ctx context.Context, context *app.RequestContext) {
	log.V(log.DebugLevel).Info("调用发送验证码方法")

	var tmp struct {
		Phone string `vd:"phone($)" json:"phone,required"`
	}
	// 参数校验
	if err := context.BindAndValidate(&tmp); err != nil {
		code.WriteResponse(context, errors.WithCode(code.ErrValidation, ""), nil)
		return
	}

	// 移交服务层
	if err := c.serviceIns.SendSmsCode(ctx, tmp.Phone); err != nil {
		code.WriteResponse(context, err, nil)
		return
	}

	code.WriteResponse(context, errors.WithCode(code.ErrSuccess, ""), nil)
}
