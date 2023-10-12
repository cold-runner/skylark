package v1

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/cold-runner/skylark/internal/pkg/util"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func (p *passwordLogin) verifyPassword(loginUser *user.LoginUser, storedUser *user.Lark) (*user.LoggedUser, error) {
	// 密码错误
	if err := util.VerifyPassword(loginUser.Password, storedUser.Password); err != nil {
		return nil, errors.WithCode(code.ErrPasswordIncorrect, "", nil)
	}

	// 登陆成功返回信息
	return &user.LoggedUser{
		Avatar:  storedUser.Avatar,
		StuNum:  storedUser.StuNum,
		Name:    storedUser.Name,
		College: storedUser.College,
		Major:   storedUser.Major,
		Grade:   storedUser.Grade,
	}, errors.WithCode(code.ErrSuccess, "", nil)
}

func (p *passwordLogin) login(c context.Context, loginUser *user.LoginUser, s *serviceV1) (*user.LoggedUser, error) {
	// 根据学号拿到密码
	storedUser, err := s.storeIns.QueryByStuNum(c, loginUser.StuNum)
	switch {

	// 校验密码
	case err == nil:
		return p.verifyPassword(loginUser, storedUser)

	// 用户不存在
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.WrapC(err, code.ErrUserNotExist, "", nil)

	// 未知错误
	default:
		e := errors.WrapC(err, code.ErrUnknown, "通过密码登陆失败！err: %v", err)
		log.Errorf("%#v", e)
		return nil, e
	}
}

func (q *qqUnionIdLogin) login(ctx context.Context, loginUser *user.LoginUser, v1 *serviceV1) (*user.LoggedUser, error) {
	//TODO implement me
	panic("implement me")
}

func (p *phoneLogin) login(c context.Context, loginUser *user.LoginUser, s *serviceV1) (*user.LoggedUser, error) {
	// TODO 并发查询数据库数据与缓存数据
	// 根据手机号获取用户信息
	storedUser, err := s.storeIns.QueryByPhone(c, loginUser.Phone)
	switch {
	// 校验验证码
	case err == nil:
		// 从缓存获取验证码
		inCacheSmsCode, err := s.getSmsCode(c, loginUser.Phone)

		// 缓存中不存在验证码
		if errors.Is(err, redis.Nil) {
			return nil, errors.WrapC(err, code.ErrSmsCode, "", err)
		}

		// 内部错误
		if err != nil && errors.ParseCoder(err).Code() == code.ErrUnknown {
			e := errors.WrapC(err, code.ErrUnknown, "查询注册验证码失败！err: %v", err)
			log.Errorf("%#v", e)
			return nil, e
		}

		// 校验验证码
		correct, err := s.validateAndDelSmsCode(c, loginUser.Phone, inCacheSmsCode, loginUser.SmsCode)
		if err != nil {
			e := errors.Wrap(err, "删除验证码失败！")
			log.Errorf("%#v", e)
		}
		if !correct {
			return nil, errors.WithCode(code.ErrSmsCode, "", nil)
		}
		// 登陆成功返回信息
		return &user.LoggedUser{
			Avatar:  storedUser.Avatar,
			StuNum:  storedUser.StuNum,
			Name:    storedUser.Name,
			College: storedUser.College,
			Major:   storedUser.Major,
			Grade:   storedUser.Grade,
		}, errors.WithCode(code.ErrSuccess, "", nil)

	// 用户不存在
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.WrapC(err, code.ErrUserNotExist, "", nil)

	default:
		e := errors.WrapC(err, code.ErrUnknown, "", nil)
		log.Errorf("手机登陆失败！err: %v", e)
		return nil, e
	}
}
