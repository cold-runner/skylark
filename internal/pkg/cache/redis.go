package cache

import (
	"context"
	"fmt"
	"github.com/cold-runner/skylark/internal/pkg/config"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisClient struct {
	client *redis.Client
}

func newRedis(opt *config.Redis) Cache {
	var err error
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", opt.Host, opt.Port),
		Username: opt.Username,
		Password: opt.Password,
		DB:       opt.Database,
	})
	_, err = client.Ping(context.Background()).Result()

	if err != nil {
		panic(fmt.Errorf("redis初始化失败，err: %v", err))
	}

	return &redisClient{client: client}
}

func (r *redisClient) Set(c context.Context, key string, value interface{}) error {
	if err := r.client.Set(c, key, value, redis.KeepTTL).Err(); err != nil {
		return err
	}
	return nil
}

func (r *redisClient) SetExpiration(c context.Context, key string, value interface{}, expiration time.Duration) error {
	result, err := r.client.SetNX(c, key, value, expiration).Result()
	if err != nil {
		return err
	}
	if !result {
		return errors.WithMessagef(err, "设置失败，key: %v, val: %v", key, value)
	}
	return nil
}

func (r *redisClient) Get(c context.Context, key string) (interface{}, error) {
	result, err := r.client.Get(c, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r *redisClient) GetDel(c context.Context, key string) (interface{}, error) {
	result, err := r.client.GetDel(c, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r *redisClient) Del(c context.Context, key string) error {
	result, err := r.client.Del(c, key).Result()
	if err != nil {
		return err
	}
	if result == int64(1) {
		return nil
	}
	return errors.Errorf("redis删除键错误 err: %v", err)
}

func (r *redisClient) IsExpired(c context.Context, key string) (bool, error) {
	result, err := r.client.TTL(c, key).Result()
	if err != nil {
		return false, err
	}
	if result <= 0 {
		return true, nil
	}
	return false, nil
}
