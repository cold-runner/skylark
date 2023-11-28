package v1

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/domain/entity"
	"gorm.io/gorm"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/internal/application/share"
	"github.com/cold-runner/skylark/internal/infrastructure/config"
	"github.com/cold-runner/skylark/internal/infrastructure/errCode"
	"github.com/cold-runner/skylark/internal/infrastructure/util"
	pkgErr "github.com/pkg/errors"
)

func (s *serviceV1) SendSmsCode(c context.Context, ctx *app.RequestContext, phone string) error {
	// 查看是否已经发送过验证码
	code, err := share.GetSmsCode(c, ctx, phone, s.cacheIns)
	switch {
	// 已发送过且未过期
	case err == nil && code != "":
		return share.DebugfAndResp(ctx, err, errCode.ErrAlreadySendSmsCode, "已发送过验证码，且未过期！\n手机号：%v", phone)

	// 内部错误
	case err != nil:
		return share.ErrorfAndResp(ctx, err, errCode.ErrUnknown, "发送验证码错误！\n手机号：%v", phone)
	}

	serverConf := config.GetConfig().ServerConfig()
	expirationConfig := serverConf.SmsExpiration          // 从全局配置对象获取验证码过期时间
	randCode, _ := util.RandSmsCode(serverConf.SmsNumber) // 从全局配置对象获取验证码长度，根据验证码长度生成验证码

	// 缓存验证码
	if err = s.cacheIns.SetExpiration(c, phone, randCode, time.Duration(expirationConfig)*time.Minute); err != nil {
		return share.ErrorfAndResp(ctx, err, errCode.ErrUnknown, "缓存验证码错误！\n手机号：%v\t验证码:%v", phone, randCode)
	}

	// 发送验证码
	if err = s.smsClient.SendToSingle(c, phone, []string{randCode, strconv.Itoa(expirationConfig)}); err != nil {
		return share.ErrorfAndResp(ctx, err, errCode.ErrUnknown, "发送验证码错误！\n手机号：%v\t验证码：%v", phone, randCode)
	}

	hlog.Infof("向%v发送验证码：%v\t过期时间：%v 分钟", phone, randCode, expirationConfig)
	return nil
}

func (s *serviceV1) Register(c context.Context, ctx *app.RequestContext, register *entity.Register) error {
	// 从缓存中获取验证码
	storeSmsCode, err := share.GetSmsCode(c, ctx, register.Phone, s.cacheIns)
	switch {
	// 内部错误
	case err != nil && pkgErr.Is(err, errCode.ErrUnknown):
		return share.ErrorfAndResp(ctx, err, errCode.ErrUnknown, "从缓存中获取注册验证码错误！手机号：%v", register.Phone)

	// 缓存中不存在验证码
	case err != nil:
		return share.DebugfAndResp(ctx, err, errCode.ErrSmsCodeExpired, "缓存中不存在验证码！手机号：%v", register.Phone)
	}

	// 验证码校验
	if register.SmsCode != storeSmsCode {
		// 用户传入的验证码不正确
		return share.DebugfAndResp(ctx, err, errCode.ErrSmsCode, "用户输入的验证码错误！\n用户手机号：%v\t用户输入：%v", register.Phone, register.SmsCode)
	}

	// 验证码正确，删除缓存中的验证码
	if err = share.DeleteSmsCode(c, ctx, register.Phone, s.cacheIns); err != nil {
		ctx.Error(err)
		hlog.Errorf("注册时删除验证码错误！\n用户手机号：%v\nerr: %v", register.Phone, ctx.Errors.Errors())
	}

	// 实体对象执行业务
	registerOV, err := register.PreCreat(c, ctx, s.ossIns)
	if err != nil {
		return share.ErrorfAndResp(ctx, err, errCode.ErrUnknown, "处理用户注册数据失败！\n注册信息：%v", *register)
	}

	// 存储进数据库
	err = register.Store(c, s.storeIns, registerOV)
	switch {
	// 用户已注册
	case err != nil && errors.Is(err, gorm.ErrDuplicatedKey):
		return share.DebugfAndResp(ctx, err, errCode.ErrUserAlreadyExist, "用户已注册！\n注册信息：%v", register)

	// 内部错误
	case err != nil:
		return share.ErrorfAndResp(ctx, err, errCode.ErrUnknown, "存储用户至数据库错误！\n注册信息：%v", register)
	}

	return nil
}
