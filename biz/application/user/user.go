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

	// 获取token

	// 返回响应
	return &user.RegisterResp{
		Status: errCode.SuccessStatus,
		Token:  "",
	}, nil
}
