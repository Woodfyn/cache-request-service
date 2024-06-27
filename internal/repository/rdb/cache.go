package rdb

import (
	"context"
	"time"

	"github.com/Woodfyn/cache-service/internal/core"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	rdb *redis.Client
}

func NewCache(rdb *redis.Client) *Cache {
	return &Cache{
		rdb: rdb,
	}
}

func (c *Cache) Create(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return c.rdb.Set(ctx, key, string(value), ttl).Err()
}

func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	result, err := c.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, core.ErrCacheIsExpired
	}

	return []byte(result), err
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.rdb.Del(ctx, key).Err()
}

func (c *Cache) Update(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	if err := c.rdb.Get(ctx, key).Err(); err == redis.Nil {
		return core.ErrCacheIsExpiredOrNotFound
	}

	return c.rdb.Set(ctx, key, string(value), ttl).Err()
}