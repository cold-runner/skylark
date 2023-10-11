package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/cold-runner/skylark/internal/pkg/sms"
	"github.com/cold-runner/skylark/internal/pkg/util"
	"github.com/marmotedu/errors"
	"strconv"
)

func (c *controllerV1) Register(ctx context.Context, context *app.RequestContext) {
	newer := &user.Register{}
	// 基础参数校验
	if err := context.BindAndValidate(newer); err != nil {
		code.WriteResponse(context, errors.WithCode(code.ErrValidation, "校验失败"), nil)
		return
	}

	// 从缓存中获取验证码
	storeSmsCode, err := c.serviceIns.GetSmsCode(c.context, newer.Phone)
	if err != nil {
		code.WriteResponse(context, err, nil)
		return
	}

	// 验证码校验
	if err := c.serviceIns.ValidateSmsCode(c.context, newer.Phone, storeSmsCode, newer.SmsCode); err != nil {
		code.WriteResponse(context, err, nil)
		return
	}

	// 移交服务层
	if err = c.serviceIns.Register(c.context, newer); err != nil {
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
	// TODO 更优雅的方式传入配置
	paramMap, _ := c.context.Value("paramMap").(map[string]interface{})
	smsNumber := paramMap[sms.SMS_NUMBER].(int)
	expiration := paramMap[sms.SMS_EXPIRATION].(int)
	randCode := util.RandCode(smsNumber)

	// 缓存验证码
	if err := c.serviceIns.SetExpiration(c.context, tmp.Phone, randCode, expiration); err != nil {
		// TODO 日志记录
		e := errors.WithCode(code.ErrUnknown, "缓存设置失败！")
		code.WriteResponse(context, e, "缓存设置失败")
		return
	}

	// 发送验证码
	if err := c.serviceIns.SendRegisterSms(c.context, tmp.Phone, []string{randCode, strconv.Itoa(expiration)}); err != nil {
		// TODO 日志记录
		e := errors.WithCode(code.ErrUnknown, "短信发送失败")
		code.WriteResponse(context, e, "短信发送失败")
		return
	}

	code.WriteResponse(context, errors.WithCode(code.ErrSuccess, ""), nil)
}
