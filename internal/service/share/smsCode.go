package share

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cold-runner/skylark/internal/pkg/cache"
	"github.com/cold-runner/skylark/internal/pkg/errCode"
	pkgErr "github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func GetSmsCode(c context.Context, ctx *app.RequestContext, phone string, cache cache.Cache) (string, *errors.Error) {
	value, err := cache.Get(c, phone)
	switch {
	case err == nil:
		return value.(string), nil

	// 处理业务错误，缓存中没有验证码
	case pkgErr.Is(err, redis.Nil):
		return "", nil

	// 服务内部错误
	default:
		ctx.Error(err)
		return "", ctx.Error(errCode.ErrUnknown)
	}
}

func DeleteSmsCode(c context.Context, ctx *app.RequestContext, phone string, cache cache.Cache) *errors.Error {
	if err := cache.Del(c, phone); err != nil {
		ctx.Error(err)
		return ctx.Error(errCode.ErrUnknown)
	}
	return nil
}
