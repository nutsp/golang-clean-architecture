package datasource

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nutsp/golang-clean-architecture/config"
)

type IRedisClient interface {
	GetKeyName(prefix string, key string) string
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
}

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(cfg config.Redis) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})
	return &RedisClient{
		client: rdb,
	}
}

func (r *RedisClient) GetKeyName(prefix string, key string) string {
	return fmt.Sprintf("%s:%s", prefix, key)
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisClient) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}
