package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/cold-runner/skylark/internal/pkg/config"
	"github.com/cold-runner/skylark/internal/pkg/util"
	"github.com/marmotedu/errors"
	"strconv"
	"time"
)

func (c *controllerV1) Register(ctx context.Context, context *app.RequestContext) {
	newer := &user.Register{}
	// 基础参数校验
	if err := context.BindAndValidate(newer); err != nil {
		code.WriteResponse(context, errors.WithCode(code.ErrValidation, "校验失败"), nil)
		return
	}

	// 从缓存中获取验证码
	storeSmsCode, err := c.serviceIns.GetSmsCode(ctx, newer.Phone)
	if err != nil {
		code.WriteResponse(context, err, nil)
		return
	}

	// 验证码校验
	if err = c.serviceIns.ValidateSmsCode(ctx, newer.Phone, storeSmsCode, newer.SmsCode); err != nil {
		code.WriteResponse(context, err, nil)
		return
	}

	// 移交服务层
	if err = c.serviceIns.Register(ctx, newer); err != nil {
		code.WriteResponse(context, err, nil)
		return
	}

	code.WriteResponse(context, errors.WithCode(code.ErrSuccess, "", nil), nil)
}

func (c *controllerV1) SendSms(ctx context.Context, context *app.RequestContext) {
	var tmp struct {
		Phone string `vd:"regexp(^(?:(?:\+|00)86)?1\d{10}$)" json:"phone,required"`
	}
	if err := context.BindAndValidate(&tmp); err != nil {
		code.WriteResponse(context, errors.WithCode(code.ErrValidation, "手机号校验失败"), nil)
		return
	}

	serverConf := config.GetConfig().ServerConfig()
	expiration := serverConf.SmsExpiration          // 从全局配置对象获取验证码过期时间
	randCode := util.RandCode(serverConf.SmsNumber) // 从全局配置对象获取验证码长度，根据验证码长度生成验证码
	// 缓存验证码
	if err := c.serviceIns.SetExpiration(ctx, tmp.Phone, randCode, time.Duration(expiration)*time.Minute); err != nil {
		code.WriteResponse(context, err, nil)
		return
	}

	// 发送验证码
	if err := c.serviceIns.SendRegisterSms(ctx, tmp.Phone, []string{randCode, strconv.Itoa(expiration)}); err != nil {
		code.WriteResponse(context, err, nil)
		return
	}

	code.WriteResponse(context, errors.WithCode(code.ErrSuccess, ""), nil)
}
