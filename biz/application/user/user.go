package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cold-runner/skylark/biz/entity"
	userEntity "github.com/cold-runner/skylark/biz/entity/user"
	"github.com/cold-runner/skylark/biz/infrastructure/cache"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
	"github.com/cold-runner/skylark/biz/model/user"
)

func Register(ctx *app.RequestContext, req *user.RegisterReq) (*user.RegisterResp, *errors.Error) {
	c := context.Background()
	cacheIns := cache.GetCache()
	storeIns := store.GetStoreIns()

	// 校验验证码
	smsCodeEntity := &entity.SmsCode{}
	err := smsCodeEntity.Validate(c, ctx, cacheIns, req.Phone, req.SmsCode)
	if err != nil {
		return nil, err
	}

	// 存储到数据库
	ety := &userEntity.RegisterDto{}
	ov, err := ety.Convert(c, ctx, req)
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
	storeIns := store.GetStoreIns()

	loginEntity := &userEntity.LoginUser{}
	if err := loginEntity.PasswordLogin(c, ctx, storeIns, req); err != nil {
		return err
	}

	return nil
}

func PhoneLogin(c context.Context, ctx *app.RequestContext, req *user.PhoneLoginReq) error {
	storeIns := store.GetStoreIns()
	cacheIns := cache.GetCache()

	loginEntity := &userEntity.LoginUser{}
	if err := loginEntity.PhoneLogin(c, storeIns, cacheIns, req); err != nil {
		return err
	}

	return nil
}
