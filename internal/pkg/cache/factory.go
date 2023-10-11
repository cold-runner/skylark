package cache

import (
	"github.com/cold-runner/skylark/internal/pkg/config"
)

const (
	REDIS = "redis"
)

// 简单工厂模式

func NewInstance(conf config.Config) Cache {
	switch conf.ServerConfig().Cache {
	case REDIS:
		return NewRedis(conf.RedisConfig())
	}
	panic("无效的缓存实例")
}
