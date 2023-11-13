package share

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/model"
	"github.com/cold-runner/skylark/internal/pkg/errCode"
	"strconv"

	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/cache"
	"github.com/cold-runner/skylark/internal/pkg/util"
	"github.com/cold-runner/skylark/internal/store"
	"github.com/cold-runner/skylark/internal/store/mysql"
	pkgErr "github.com/pkg/errors"
	"gorm.io/gorm"
)

func PasswordLogin(c context.Context, ctx *app.RequestContext, loginUser *user.LoginUser, store store.Store) (*user.LoggedUser, *errors.Error) {
	// 根据学号拿到密码
	storedUser, err := store.GetLark(c, mysql.LarkByStuNum(loginUser.StuNum))
	switch {
	// 用户不存在
	case err != nil && pkgErr.Is(err, gorm.ErrRecordNotFound):
		ctx.Error(err)
		return nil, ctx.Error(errCode.ErrUserNotExist)

	// 内部错误
	case err != nil:
		hlog.Errorf("%v", err)
		ctx.Error(err)
		return nil, ctx.Error(errCode.ErrUnknown)
	}

	// 密码错误
	if err = util.VerifyPassword(loginUser.Password, storedUser.Password); err != nil {
		ctx.Error(err)
		return nil, ctx.Error(errCode.ErrPasswordIncorrect)
	}

	// 密码正确返回信息
	return buildLoggedUser(storedUser), nil
}

func QqUnionIdLogin(c context.Context, ctx *app.RequestContext, loginUser *user.LoginUser, store store.Store) (*user.LoggedUser, *errors.Error) {
	// 根据qqUnionId获取用户
	storedUser, err := store.GetLark(c, mysql.LarkByQqUnionId(loginUser.QqUnionId))
	switch {
	// 用户不存在
	case pkgErr.Is(err, gorm.ErrRecordNotFound):
		ctx.Error(err)
		return nil, ctx.Error(errCode.ErrUserNotExist)

	// 内部错误
	case err != nil:
		hlog.Errorf("qqUnionId登陆失败！err: %v", err)
		ctx.Error(err)
		return nil, ctx.Error(errCode.ErrUnknown)
	}

	// 存在用户，返回信息
	return buildLoggedUser(storedUser), nil
}

func PhoneLogin(c context.Context, ctx *app.RequestContext, loginUser *user.LoginUser, store store.Store, cache cache.Cache) (*user.LoggedUser, *errors.Error) {
	// 数据库中是否存在该手机号
	storedUser, err := store.GetLark(c, mysql.LarkByPhone(loginUser.Phone))
	switch {
	// 用户不存在
	case pkgErr.Is(err, gorm.ErrRecordNotFound):
		ctx.Error(err)
		return nil, ctx.Error(errCode.ErrUserNotExist)

	// 内部错误
	case err != nil:
		hlog.Errorf("手机登陆失败！err: %v", err)
		ctx.Error(err)
		return nil, ctx.Error(errCode.ErrUnknown)
	}

	// 从缓存获取验证码
	inCacheSmsCode, err := GetSmsCode(c, ctx, loginUser.Phone, cache)
	switch {
	// 缓存中不存在验证码 TODO 验证码
	case pkgErr.Is(err, nil):
		ctx.Error(err)
		return nil, ctx.Error(errCode.ErrSmsCode)

	// 内部错误
	case err != nil:
		hlog.Errorf("获取验证码失败！err: %v", err)
		ctx.Error(err)
		return nil, ctx.Error(errCode.ErrUnknown)
	}

	// 校验验证码
	if inCacheSmsCode != loginUser.SmsCode {
		// 验证码不对
		ctx.Error(err)
		return nil, ctx.Error(errCode.ErrSmsCode)
	}

	// 验证码正确则删除缓存中的验证码
	if err = DeleteSmsCode(c, ctx, loginUser.Phone, cache); err != nil {
		hlog.Errorf("手机登陆时删除验证码失败！err: %v", err)
		ctx.Error(err)
	}

	// 登陆成功返回信息
	return buildLoggedUser(storedUser), nil
}

func buildLoggedUser(storedUser *model.Lark) *user.LoggedUser {
	return &user.LoggedUser{
		UserId:  strconv.Itoa(int(storedUser.ID)),
		Avatar:  storedUser.Avatar,
		StuNum:  storedUser.StuNum,
		Name:    storedUser.Name,
		College: storedUser.College,
		Major:   storedUser.Major,
		Grade:   strconv.FormatInt(storedUser.Grade, 10),
	}
}
