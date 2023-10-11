package cache

import (
	"github.com/cold-runner/skylark/internal/pkg/option"
)

const (
	REDIS = "redis"
)

// 简单工厂模式

func NewCache(conf *option.Conf) Cache {
	switch conf.Server.Cache {
	case REDIS:
		return NewRedis(conf.Redis)
	}
	panic("无效的缓存实例")
}
