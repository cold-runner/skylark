package share

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/cache"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/cold-runner/skylark/internal/pkg/util"
	"github.com/cold-runner/skylark/internal/store"
	"github.com/cold-runner/skylark/internal/store/mysql"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"
	"gorm.io/gorm"
	"strconv"
)

func PasswordLogin(c context.Context, loginUser *user.LoginUser, store store.Store) (*user.LoggedUser, error) {
	// 根据学号拿到密码
	storedUser, err := store.GetLarkInfo(c, mysql.LarkByStuNum(loginUser.StuNum))
	switch {
	// 用户不存在
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.WrapC(err, code.ErrUserNotExist, "", nil)

	// 内部错误
	case err != nil && errors.ParseCoder(err).Code() == code.ErrUnknown:
		e := errors.WrapC(err, code.ErrUnknown, "通过密码登陆失败！err: %v", err)
		log.Errorf("%#v", e)
		return nil, e
	}

	// 密码错误
	if err = util.VerifyPassword(loginUser.Password, storedUser.Password); err != nil {
		return nil, errors.WithCode(code.ErrPasswordIncorrect, "", nil)
	}

	// 密码正确返回信息
	return buildLoggedUser(storedUser), nil
}

func QqUnionIdLogin(c context.Context, loginUser *user.LoginUser, store store.Store) (*user.LoggedUser, error) {
	// 根据qqUnionId获取用户
	storedUser, err := store.GetLarkInfo(c, mysql.LarkByQqUnionId(loginUser.QqUnionId))
	switch {
	// 用户不存在
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.WrapC(err, code.ErrUserNotExist, "", nil)

	// 内部错误
	case err != nil && errors.ParseCoder(err).Code() == code.ErrUnknown:
		e := errors.WrapC(err, code.ErrUnknown, "通过密码登陆失败！err: %v", err)
		log.Errorf("%#v", e)
		return nil, e
	}

	// 存在用户，返回信息
	return buildLoggedUser(storedUser), nil
}

func PhoneLogin(c context.Context, loginUser *user.LoginUser, store store.Store, cache cache.Cache) (*user.LoggedUser, error) {
	// 数据库中是否存在该手机号
	storedUser, err := store.GetLarkInfo(c, mysql.LarkByPhone(loginUser.Phone))
	switch {
	// 用户不存在
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.WrapC(err, code.ErrUserNotExist, "", nil)

	// 内部错误
	case err != nil && errors.ParseCoder(err).Code() == code.ErrUnknown:
		e := errors.WrapC(err, code.ErrUnknown, "", nil)
		log.Errorf("手机登陆失败！err: %v", e)
		return nil, e
	}

	// 从缓存获取验证码
	inCacheSmsCode, err := GetSmsCode(c, loginUser.Phone, cache)
	switch {
	// 内部错误
	case err != nil && errors.ParseCoder(err).Code() == code.ErrUnknown:
		e := errors.WrapC(err, code.ErrUnknown, "查询注册验证码失败！err: %v", err)
		log.Errorf("%#v", e)
		return nil, e

	// 缓存中不存在验证码
	case err != nil:
		return nil, errors.WrapC(err, code.ErrSmsCode, "", err)
	}

	// 校验验证码
	if inCacheSmsCode != loginUser.SmsCode {
		// 验证码不对
		return nil, errors.WithCode(code.ErrSmsCode, "", nil)
	}

	// 验证码正确则删除缓存中的验证码
	if err = DeleteSmsCode(c, loginUser.Phone, cache); err != nil {
		e := errors.Wrap(err, "删除验证码失败！")
		log.Errorf("%#v", e)
	}

	// 登陆成功返回信息
	return buildLoggedUser(storedUser), nil
}

func buildLoggedUser(storedUser *user.Lark) *user.LoggedUser {
	return &user.LoggedUser{
		UserId:  strconv.Itoa(int(storedUser.ID)),
		Avatar:  storedUser.Avatar,
		StuNum:  storedUser.StuNum,
		Name:    storedUser.Name,
		College: storedUser.College,
		Major:   storedUser.Major,
		Grade:   storedUser.Grade,
	}
}
