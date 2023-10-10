package cache

import "time"

const (
	REDIS = "redis"
)

type Cache interface {
	Set(key, value string) error
	SetExpiration(key, value string, expiration time.Duration) error
	Get(key string) (value string, err error)
	GetDel(key string) (value string, err error)
}
