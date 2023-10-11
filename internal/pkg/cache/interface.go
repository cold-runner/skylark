package cache

import (
	"context"
	"time"
)

type Cache interface {
	Set(c context.Context, key, value string) error
	SetExpiration(c context.Context, key string, value string, expiration time.Duration) error
	Get(c context.Context, key string) (value string, err error)
	GetDel(c context.Context, key string) (value string, err error)
	Del(c context.Context, key string) error
}
