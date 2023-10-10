package cache

import (
	"context"
	"fmt"
	"github.com/cold-runner/skylark/internal/pkg/option"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisClient struct {
	context context.Context
	client  *redis.Client
}

func (r *redisClient) Set(key, value string) error {
	if err := r.client.Set(r.context, key, value, redis.KeepTTL).Err(); err != nil {
		return err
	}
	return nil
}

func (r *redisClient) SetExpiration(key string, value string, expiration time.Duration) error {
	result, err := r.client.SetNX(r.context, key, value, expiration).Result()
	if err != nil {
		return err
	}
	if !result {
		return errors.WithMessagef(err, "设置失败，key: %v, val: %v", key, value)
	}
	return nil
}

func (r *redisClient) Get(key string) (string, error) {
	result, err := r.client.Get(r.context, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r *redisClient) GetDel(key string) (string, error) {
	result, err := r.client.GetDel(r.context, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func NewRedis(opt *option.Redis) Cache {
	var err error
	c := context.Background()
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

	return &redisClient{
		context: c,
		client:  client,
	}
}
