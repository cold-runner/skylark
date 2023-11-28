package share

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/internal/domain/entity"
	"github.com/cold-runner/skylark/internal/domain/model_gen"
	"github.com/cold-runner/skylark/internal/infrastructure/cache"
	"github.com/cold-runner/skylark/internal/infrastructure/errCode"
	"github.com/cold-runner/skylark/internal/infrastructure/store"
	"github.com/cold-runner/skylark/internal/infrastructure/store/mysql"
	"github.com/cold-runner/skylark/internal/infrastructure/util"
	pkgErr "github.com/pkg/errors"
	"gorm.io/gorm"
)

func PasswordLogin(c context.Context, ctx *app.RequestContext, loginUser *entity.LoginUser, store store.Store) (*entity.LoggedInUser, error) {
	// 根据学号拿到密码
	storedUser, err := store.GetLark(c, mysql.LarkByStuNum(loginUser.StuNum))
	switch {
	// 用户不存在
	case err != nil && pkgErr.Is(err, gorm.ErrRecordNotFound):
		return nil, DebugfAndResp(ctx, err, errCode.ErrUserNotExist, "用户不存在！\n登陆信息: %v", *loginUser)

	// 内部错误
	case err != nil:
		return nil, ErrorfAndResp(ctx, err, errCode.ErrUnknown, "密码登陆错误！\n登陆信息： %v", *loginUser)
	}

	// 密码错误
	if err = util.VerifyPassword(loginUser.Password, storedUser.Password); err != nil {
		return nil, DebugfAndResp(ctx, err, errCode.ErrPasswordIncorrect, "密码登陆时密码错误！\n登陆信息: %v", *loginUser)
	}

	// 密码正确返回信息
	return buildLoggedUser(storedUser), nil
}

func QqUnionIdLogin(c context.Context, ctx *app.RequestContext, loginUser *entity.LoginUser, store store.Store) (*entity.LoggedInUser, error) {
	// 根据qqUnionId获取用户
	storedUser, err := store.GetLark(c, mysql.LarkByQqUnionId(loginUser.QqUnionId))
	switch {
	// 用户不存在
	case pkgErr.Is(err, gorm.ErrRecordNotFound):
		return nil, DebugfAndResp(ctx, err, errCode.ErrUserNotExist, "qqUnionId登陆时用户不存在！\n登陆信息: %v", *loginUser)

	// 内部错误
	case err != nil:
		return nil, ErrorfAndResp(ctx, err, errCode.ErrUnknown, "qqUnionId登陆失败！\n")
	}

	// 存在用户，返回信息
	return buildLoggedUser(storedUser), nil
}

func PhoneLogin(c context.Context, ctx *app.RequestContext, loginUser *entity.LoginUser, store store.Store, cache cache.Cache) (*entity.LoggedInUser, error) {
	// 数据库中是否存在该手机号
	storedUser, err := store.GetLark(c, mysql.LarkByPhone(loginUser.Phone))
	switch {
	// 用户不存在
	case pkgErr.Is(err, gorm.ErrRecordNotFound):
		return nil, DebugfAndResp(ctx, err, errCode.ErrUserNotExist, "收集登陆时用户不存在！\n登陆信息: %v", *loginUser)

	// 内部错误
	case err != nil:
		return nil, ErrorfAndResp(ctx, err, errCode.ErrUnknown, "手机登陆失败！\n")
	}

	// 从缓存获取验证码
	inCacheSmsCode, err := GetSmsCode(c, ctx, loginUser.Phone, cache)
	switch {
	// 缓存中不存在验证码
	case pkgErr.Is(err, nil) && inCacheSmsCode == "":
		return nil, DebugfAndResp(ctx, err, errCode.ErrSmsCode, "验证码已过期或不存在！\n登陆信息: %v", *loginUser)

	// 内部错误
	case err != nil:
		return nil, ErrorfAndResp(ctx, err, errCode.ErrUnknown, "手机登录时错误！\n登陆信息; %v", *loginUser)
	}

	// 校验验证码
	if inCacheSmsCode != loginUser.SmsCode {
		// 验证码不对
		return nil, DebugfAndResp(ctx, err, errCode.ErrSmsCode, "验证码错误！\n登陆信息: %v", *loginUser)
	}

	// 验证码正确则删除缓存中的验证码
	if err = DeleteSmsCode(c, ctx, loginUser.Phone, cache); err != nil {
		ctx.Error(err)
		hlog.Errorf("手机登陆时删除验证码失败！err: %v", ctx.Errors.Errors())
	}

	// 登陆成功返回信息
	return buildLoggedUser(storedUser), nil
}

func buildLoggedUser(storedUser *model_gen.Lark) *entity.LoggedInUser {
	return &entity.LoggedInUser{
		StuNum: storedUser.StuNum,
	}
}
