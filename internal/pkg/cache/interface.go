package cache

import (
	"context"
	"time"
)

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
