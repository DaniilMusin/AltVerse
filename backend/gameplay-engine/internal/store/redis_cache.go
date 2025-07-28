package store

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	rdb *redis.Client
}

func MustRedis(url string) *RedisCache {
	opt, _ := redis.ParseURL(url)
	return &RedisCache{redis.NewClient(opt)}
}

func (c *RedisCache) GetPlayer(ctx context.Context, id string, dst any) (bool, error) {
	raw, err := c.rdb.Get(ctx, "player:"+id).Bytes()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, json.Unmarshal(raw, dst)
}

func (c *RedisCache) SetPlayer(ctx context.Context, id string, src any) error {
	b, _ := json.Marshal(src)
	return c.rdb.Set(ctx, "player:"+id, b, time.Hour).Err()
}
