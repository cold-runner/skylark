package user

import (
	"context"
	"time"

	"github.com/cold-runner/skylark/biz/config"
	"github.com/cold-runner/skylark/biz/entity"
	userEntity "github.com/cold-runner/skylark/biz/entity/user"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/model/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func Register(c context.Context, ctx *app.RequestContext, req *user.RegisterReq) (*user.RegisterResp, error) {
	smsCodeEntity := &entity.SmsCode{}
	// 校验验证码
	err := smsCodeEntity.Validate(c, ctx, req.Phone, req.SmsCode)
	if err != nil {
		return nil, err
	}
	if err := smsCodeEntity.DeleteSmsCode(c, ctx, req.Phone); err != nil {
		hlog.Warnf("注册时校验验证码通过，但删除验证码失败！err: %v", err)
	}

	// 存储到数据库
	ety := &userEntity.RegisterDto{}
	if err = ety.IsRegistered(c, ctx, req); err != nil {
		return nil, err
	}
	ov, err := ety.Convert(c, ctx, req)
	if err != nil {
		return nil, err
	}
	err = ety.Store(c, ctx, ov)
	if err != nil {
		return nil, err
	}

	// 返回响应
	return &user.RegisterResp{
		Code: errCode.SuccessStatus.Code,
		Msg:  errCode.SuccessStatus.Msg,
	}, nil
}

func PasswordLogin(c context.Context, ctx *app.RequestContext, req *user.PasswordLoginReq) error {
	loginEntity := &userEntity.LoginUser{}
	if err := loginEntity.PasswordLogin(c, ctx, req); err != nil {
		return err
	}

	return nil
}

func PhoneLogin(c context.Context, ctx *app.RequestContext, req *user.PhoneLoginReq) error {
	loginEntity := &userEntity.LoginUser{}
	if err := loginEntity.PhoneLogin(c, ctx, req); err != nil {
		return err
	}

	return nil
}

func SendSmsCode(c context.Context, ctx *app.RequestContext, req *user.SendSmsCodeReq) (*user.SendSmsCodeResp, error) {
	conf := config.GetConfig().GetServer()
	smsEntity := &entity.SmsCode{}

	// 是否已经发送过还没过期
	if err := smsEntity.AlreadySend(c, ctx, req.Phone); err != nil {
		return nil, err
	}

	// 根据配置生成指定位数的验证码
	smsCode, err := smsEntity.GenSmsCodeRandom(ctx, conf.SmsNumber)
	if err != nil {
		return nil, err
	}

	// 在缓存中设置验证码和过期时间
	expireTime := time.Duration(conf.SmsExpiration) * time.Minute
	if err = smsEntity.SetCache(c, ctx, req.Phone, smsCode, expireTime); err != nil {
		return nil, err
	}

	// 发送验证码
	if err = smsEntity.SendSmsCode(c, ctx, req.Phone, smsCode, conf.SmsExpiration); err != nil {
		return nil, err
	}

	return &user.SendSmsCodeResp{
		Code: errCode.SuccessStatus.Code,
		Msg:  errCode.SuccessStatus.Msg,
	}, nil
}

func GetUserInfo(c context.Context, ctx *app.RequestContext) (*user.GetUserInfoResp, error) {
	userUuid, _ := ctx.Get("identity")
	lark := &userEntity.Lark{}
	err := lark.GetById(c, ctx, userUuid.(string))
	if err != nil {
		return nil, err
	}

	basicInfo := lark.Format()
	return &user.GetUserInfoResp{
		Code: errCode.SuccessStatus.Code,
		Msg:  errCode.SuccessStatus.Msg,
		Data: &user.GetUserInfoResp_Data{BasicInfo: basicInfo},
	}, nil
}

func GetUserInfoByStuNum(c context.Context, ctx *app.RequestContext, req *user.GetUserInfoByStuNumReq) (*user.GetUserInfoResp, error) {
	lark := &userEntity.Lark{}
	err := lark.GetByStuNum(c, ctx, req.StuNum)
	if err != nil {
		return nil, err
	}

	basicInfo := lark.Format()
	return &user.GetUserInfoResp{
		Code: errCode.SuccessStatus.Code,
		Msg:  errCode.SuccessStatus.Msg,
		Data: &user.GetUserInfoResp_Data{BasicInfo: basicInfo},
	}, nil
}
