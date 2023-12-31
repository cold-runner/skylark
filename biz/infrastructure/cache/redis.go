package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/cold-runner/skylark/biz/config"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	client *redis.Client
}

func NewRedis(opt *config.Redis) Cache {
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

func (r *redisClient) SetHash(c context.Context, hashName string, key string, value interface{}) error {
	result, err := r.client.HSet(c, hashName, key, value).Result()
	if err != nil {
		return err
	}
	if result != 1 {
		return errors.WithMessagef(err, "设置HSET失败，haseName: %v, key: %v, val: %v", hashName, key, value)
	}
	return nil
}

func (r *redisClient) SetHashExpiration(c context.Context, hashName string, key string, value interface{}, expiration time.Duration) error {
	_, err := r.client.HSet(c, hashName, key, value).Result()
	if err != nil {
		return err
	}
	return r.Expire(c, hashName, expiration)
}

func (r *redisClient) SetHashMulti(c context.Context, hashName string, kvPair map[string]interface{}) error {
	result, err := r.client.HSet(c, hashName, kvPair).Result()
	if err != nil {
		return err
	}
	if result != int64(len(kvPair)) {
		return errors.WithMessagef(err, "批量设置HSET失败，haseName: %v, kvPair: %v", hashName, kvPair)
	}
	return nil
}

func (r *redisClient) Get(c context.Context, key string) (interface{}, error) {
	result, err := r.client.Get(c, key).Result()
	if err == redis.Nil {
		return "", nil
	}
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

func (r *redisClient) Expire(c context.Context, key string, expiration time.Duration) error {
	success, err := r.client.Expire(c, key, expiration).Result()
	if err != nil {
		return err
	}
	if !success {
		return errors.New("设置超时时间失败")
	}
	return nil
}
