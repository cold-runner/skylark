package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/model"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/config"
	"github.com/cold-runner/skylark/internal/pkg/errCode"
	"github.com/cold-runner/skylark/internal/pkg/util"
	"github.com/cold-runner/skylark/internal/service/share"
	pkgErr "github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *serviceV1) SendSmsCode(c context.Context, ctx *app.RequestContext, phone string) *errors.Error {
	// 查看是否已经发送过验证码
	code, err := share.GetSmsCode(c, ctx, phone, s.cacheIns)

	switch {
	// 已发送过且未过期
	case err == nil && code != "":
		ctx.Error(err)
		return ctx.Error(errCode.ErrAlreadySendSmsCode)

	// 内部错误
	case err != nil:
		hlog.Errorf("发送验证码错误！\n\n%v\n", ctx.Errors.String())
		return err
	}

	serverConf := config.GetConfig().ServerConfig()
	expirationConfig := serverConf.SmsExpiration          // 从全局配置对象获取验证码过期时间
	randCode, _ := util.RandSmsCode(serverConf.SmsNumber) // 从全局配置对象获取验证码长度，根据验证码长度生成验证码

	// 缓存验证码
	if err := s.cacheIns.SetExpiration(c, phone, randCode, time.Duration(expirationConfig)*time.Minute); err != nil {
		hlog.Errorf("缓存验证码错误！err: %v", err)
		ctx.Error(err)
		return ctx.Error(errCode.ErrUnknown)
	}

	// 发送验证码
	if err := s.smsClient.SendToSingle(c, phone, []string{randCode, strconv.Itoa(expirationConfig)}); err != nil {
		hlog.Errorf("发送验证码错误！err: %v", err)
		ctx.Error(err)
		return ctx.Error(errCode.ErrUnknown)
	}

	hlog.Infof("向%s发送验证码：%s 过期时间：%d 分钟", phone, randCode, expirationConfig)
	return nil
}

func (s *serviceV1) Register(c context.Context, ctx *app.RequestContext, register *user.Register) *errors.Error {
	// 从缓存中获取验证码
	storeSmsCode, err := share.GetSmsCode(c, ctx, register.Phone, s.cacheIns)
	switch {
	// 内部错误
	case err != nil && pkgErr.Is(err, errCode.ErrUnknown):
		hlog.Errorf("从缓存中获取注册验证码错误！err: %v", err)
		ctx.Error(err)
		return ctx.Error(errCode.ErrUnknown)

	// 缓存中不存在验证码
	case err != nil:
		ctx.Error(err)
		return ctx.Error(errCode.ErrSmsCodeExpired)
	}

	// 验证码校验
	if register.SmsCode != storeSmsCode {
		// 用户传入的验证码不正确
		return ctx.Error(errCode.ErrSmsCode)
	}

	// 验证码正确，删除缓存中的验证码
	if err := share.DeleteSmsCode(c, ctx, register.SmsCode, s.cacheIns); err != nil {
		hlog.Errorf("注册时删除验证码错误！err: %v", err)
		ctx.Error(err)
	}

	// 密码加盐加密
	cryptPassword, err2 := util.Crypt(register.Password)
	if err != nil {
		hlog.Errorf("注册时加盐加密密码错误！err: %v", err2)
		return errCode.ErrUnknown
	}

	// 学生证照片文件上传，文件名根据系统时间戳生成，保证唯一
	fileUrl, err2 := s.ossIns.UploadFormFile(c, register.StuCardPhoto, strconv.FormatInt(util.GenUnixTimeNano(), 10))
	if err != nil {
		hlog.Errorf("注册时上传学生证照片文件错误！err: %v", err2)
		return errCode.ErrUnknown
	}

	// 存储进数据库
	err2 = s.storeIns.CreateLark(c, &model.Lark{
		StuNum:   register.StuNum,
		Name:     register.Name,
		Password: cryptPassword,
		Gender:   int64(util.GenderTransform(register.Gender)),
		College:  register.College,
		Major:    register.Major,
		Grade:    register.Grade,
		Phone:    register.Phone,
		PhotoURL: fileUrl,
		Legal:    int64(0),
	})
	switch {
	// 用户已注册
	case pkgErr.Is(err2, gorm.ErrDuplicatedKey):
		return errCode.ErrUserAlreadyExist

	// 服务内部错误
	case err2 != nil && pkgErr.Is(err2, errCode.ErrUnknown):
		hlog.Errorf("注册时存储用户至数据库错误！%v", err2)
		return errCode.ErrUnknown
	}

	return nil
}
