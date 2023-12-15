package user

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/biz/config"
	"github.com/cold-runner/skylark/biz/entity"
	userEntity "github.com/cold-runner/skylark/biz/entity/user"
	"github.com/cold-runner/skylark/biz/infrastructure/cache"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/infrastructure/oss"
	"github.com/cold-runner/skylark/biz/infrastructure/sms"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
	"github.com/cold-runner/skylark/biz/model/user"
)

func Register(ctx *app.RequestContext, req *user.RegisterReq) (*user.RegisterResp, error) {
	c := context.Background()
	cacheIns := cache.GetIns()
	storeIns := store.GetIns()
	ossIns := oss.GetIns()

	// 校验验证码
	smsCodeEntity := &entity.SmsCode{}
	err := smsCodeEntity.Validate(c, ctx, cacheIns, req.Phone, req.SmsCode)
	if err != nil {
		return nil, err
	}
	if err := smsCodeEntity.DeleteSmsCode(c, ctx, cacheIns, req.Phone); err != nil {
		hlog.Warnf("注册时校验验证码通过，但删除验证码失败！err: %v", err)
	}

	// 存储到数据库
	ety := &userEntity.RegisterDto{}
	if err := ety.IsRegistered(c, ctx, storeIns, req); err != nil {
		return nil, err
	}
	ov, err := ety.Convert(c, ctx, ossIns, req)
	if err != nil {
		return nil, err
	}
	err = ety.Store(c, ctx, storeIns, ov)
	if err != nil {
		return nil, err
	}

	// 返回响应
	return &user.RegisterResp{
		Status: errCode.SuccessStatus,
	}, nil
}

func PasswordLogin(c context.Context, ctx *app.RequestContext, req *user.PasswordLoginReq) error {
	storeIns := store.GetIns()

	loginEntity := &userEntity.LoginUser{}
	if err := loginEntity.PasswordLogin(c, ctx, storeIns, req); err != nil {
		return err
	}

	return nil
}

func PhoneLogin(c context.Context, ctx *app.RequestContext, req *user.PhoneLoginReq) error {
	storeIns := store.GetIns()
	cacheIns := cache.GetIns()

	loginEntity := &userEntity.LoginUser{}
	if err := loginEntity.PhoneLogin(c, ctx, storeIns, cacheIns, req); err != nil {
		return err
	}

	return nil
}

func SendSmsCode(c context.Context, ctx *app.RequestContext, req *user.SendSmsCodeReq) (*user.SendSmsCodeRes, error) {
	smsClient := sms.GetClient()
	cacheIns := cache.GetIns()
	conf := config.GetConfig().GetServer()

	smsEntity := &entity.SmsCode{}

	// 是否已经发送过还没过期
	if err := smsEntity.AlreadySend(c, ctx, cacheIns, req.Phone); err != nil {
		return nil, err
	}

	// 根据配置生成指定位数的验证码
	smsCode, err := smsEntity.GenSmsCodeRandom(ctx, conf.SmsNumber)
	if err != nil {
		return nil, err
	}

	// 在缓存中设置验证码和过期时间
	expireTime := time.Duration(conf.SmsExpiration) * time.Minute
	if err = smsEntity.SetCache(c, ctx, cacheIns, req.Phone, smsCode, expireTime); err != nil {
		return nil, err
	}

	// 发送验证码
	if err = smsEntity.SendSmsCode(c, ctx, smsClient, req.Phone, smsCode, conf.SmsExpiration); err != nil {
		return nil, err
	}

	return &user.SendSmsCodeRes{
		Status: errCode.SuccessStatus,
	}, nil
}

func GetUserInfo(c context.Context, ctx *app.RequestContext, req *user.GetUserInfoByIdReq) (*user.GetUserInfoRes, error) {
	storeIns := store.GetIns()

	userUuid, _ := ctx.Get("identity")
	lark := &userEntity.Lark{}
	storedLark, err := lark.GetById(c, ctx, storeIns, userUuid.(string))
	if err != nil {
		return nil, err
	}

	basicInfo := lark.Format(storedLark)

	return &user.GetUserInfoRes{
		Status:    errCode.SuccessStatus,
		BasicInfo: basicInfo,
	}, nil
}
