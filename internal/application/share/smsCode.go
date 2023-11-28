package share

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/internal/infrastructure/cache"
	"github.com/cold-runner/skylark/internal/infrastructure/errCode"
	pkgErr "github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func GetSmsCode(c context.Context, ctx *app.RequestContext, phone string, cache cache.Cache) (string, error) {
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

func DeleteSmsCode(c context.Context, ctx *app.RequestContext, phone string, cache cache.Cache) error {
	if err := cache.Del(c, phone); err != nil {
		hlog.Error(err)
		ctx.Error(err)
		return ctx.Error(errCode.ErrUnknown)
	}
	return nil
}
