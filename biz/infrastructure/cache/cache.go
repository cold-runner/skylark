package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/cold-runner/skylark/biz/config"
	"github.com/pkg/errors"
)

var (
	cacheIns Cache
)

type CacheType string

const (
	LOCAL CacheType = "local"
	REDIS CacheType = "redis"
)

func (c CacheType) string() string {
	return string(c)
}

func Init(c config.Conf) {
	cache := c.GetServer().Cache
	switch cache {
	case REDIS.string():
		cacheIns = NewRedis(c.GetRedis())
	case LOCAL.string():
		fmt.Println("还未实现本地缓存")
	default:
		panic(errors.Errorf("无效的缓存实例类型，支持的类型：redis、local。当前传入类型为：%s", cache))
	}

}

// NewCache 简单工厂模式
func NewCache(conf config.Conf) Cache {
	switch conf.GetServer().Cache {
	case LOCAL.string():
		panic("还未实现本地缓存")
	case REDIS.string():
		return NewRedis(conf.GetRedis())
	default:
		panic(errors.Errorf("无效的缓存实例类型，支持的类型：redis、local。当前传入类型为：%s", conf.GetServer().Cache))
	}
}

func SetDefault(c Cache) {
	if c == nil {
		panic("cache is nil")
	}
	cacheIns = c
}

func GetCache() Cache {
	if cacheIns == nil {
		panic("cache is nil")
	}
	return cacheIns
}

type Cache interface {
	Set(c context.Context, key string, value interface{}) error
	SetExpiration(c context.Context, key string, value interface{}, expiration time.Duration) error
	SetHash(c context.Context, hashName string, key string, value interface{}) error
	SetHashExpiration(c context.Context, hashName string, key string, value interface{}, expiration time.Duration) error
	SetHashMulti(c context.Context, hashName string, kvPair map[string]interface{}) error
	Get(c context.Context, key string) (value interface{}, err error)
	GetDel(c context.Context, key string) (value interface{}, err error)
	Del(c context.Context, key string) error
	IsExpired(c context.Context, key string) (isExpired bool, err error)
	Expire(c context.Context, key string, expiration time.Duration) error
}
