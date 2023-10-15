package share

import (
	"context"
	"github.com/cold-runner/skylark/internal/pkg/cache"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/marmotedu/errors"
	"github.com/redis/go-redis/v9"
)

func GetSmsCode(c context.Context, phone string, cache cache.Cache) (smsCode string, err error) {
	value, err := cache.Get(c, phone)
	switch {
	case err == nil:
		return value.(string), nil

	// 处理业务错误
	case errors.Is(err, redis.Nil):
		return "", errors.WithCode(code.ErrSmsCodeExpired, "缓存中不存在验证码, err: %v", err)

	// 服务内部错误
	default:
		e := errors.WithCode(code.ErrUnknown, "公共服务GetSmsCode方法失败！err: %v", err)
		return "", e
	}
}

func DeleteSmsCode(c context.Context, phone string, cache cache.Cache) error {
	if err := cache.Del(c, phone); err != nil {
		return errors.WrapC(err, code.ErrUnknown, "公共服务validateSmsCode方法失败！err: %v", nil)
	}
	return nil
}
