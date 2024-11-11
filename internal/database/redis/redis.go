package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedis(addr string) *Redis {
	redisDB := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &Redis{
		client: redisDB,
		ctx:    context.Background(),
	}
}

func (r *Redis) Set(key string, value string) error {
	return r.client.Set(r.ctx, key, value, 5*60*time.Second).Err()
}

func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r *Redis) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}
