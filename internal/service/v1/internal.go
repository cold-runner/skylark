package v1

import (
	"context"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/marmotedu/errors"
	"github.com/redis/go-redis/v9"
)

func (s *serviceV1) getSmsCode(c context.Context, phone string) (string, error) {
	value, err := s.cacheIns.Get(c, phone)
	switch {
	case err == nil:
		return value, nil

	// 处理业务错误
	case errors.Is(err, redis.Nil):
		return "", errors.WithCode(code.ErrSmsCodeExpired, "验证码已过期, err: %v", err)

	// 服务内部错误
	default:
		e := errors.WithCode(code.ErrUnknown, "公共服务GetSmsCode方法失败！err: %v", err)
		return "", e
	}
}

func (s *serviceV1) validateAndDelSmsCode(c context.Context, phone, inCache, userPass string) (bool, error) {
	if inCache != userPass {
		return false, nil
	}
	if err := s.cacheIns.Del(c, phone); err != nil {
		return true, errors.WrapC(err, code.ErrUnknown, "公共服务validateSmsCode方法失败！err: %v", nil)
	}
	return true, nil
}
