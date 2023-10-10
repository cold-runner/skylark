package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/cold-runner/skylark/internal/pkg/sms"
	"github.com/cold-runner/skylark/internal/pkg/util"
	"github.com/marmotedu/errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
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
	storeSmsCode, err := c.serviceIns.GetDelete(newer.Phone)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			e := errors.WithCode(code.ErrSmsCodeExpired, "验证码过期, err: %v", err)
			code.WriteResponse(context, e, nil)
			return
		}
		e := errors.WithCode(code.ErrUnknown, "缓存获取验证码失败, err: %v", err)
		// TODO 记录日志
		code.WriteResponse(context, e, nil)
		return
	}

	// 验证码校验
	if storeSmsCode != newer.SmsCode {
		code.WriteResponse(context, errors.WithCode(code.ErrSmsCode, "验证码错误"), nil)
		return
	}

	// 签发token

	// 移交服务层
	err = c.serviceIns.Register(ctx, newer)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			e := errors.WithCode(code.ErrUserAlreadyExist, "用户已注册, err: %v", err)
			code.WriteResponse(context, e, nil)
			return
		}
		e := errors.WithCode(code.ErrRegister, "注册失败, err: %v", err)
		// TODO 记录日志
		code.WriteResponse(context, e, nil)
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
	if err := c.serviceIns.SetExpiration(tmp.Phone, randCode, expiration); err != nil {
		// TODO 日志记录
		e := errors.WithCode(code.ErrUnknown, "缓存设置失败！")
		code.WriteResponse(context, e, "缓存设置失败")
		return
	}

	// 发送验证码
	if err := c.serviceIns.SendRegisterSms(tmp.Phone, []string{randCode, strconv.Itoa(expiration)}); err != nil {
		// TODO 日志记录
		e := errors.WithCode(code.ErrUnknown, "短信发送失败")
		code.WriteResponse(context, e, "短信发送失败")
		return
	}

	code.WriteResponse(context, errors.WithCode(code.ErrSuccess, ""), nil)
}
