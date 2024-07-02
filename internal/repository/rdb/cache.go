package rdb

import (
	"context"
	"errors"

	"github.com/Woodfyn/chat-api-cache-service/pkg/core"
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

func (c *Cache) Delete(ctx context.Context, key string) error {
	err := c.rdb.Del(ctx, key).Err()
	if errors.Is(err, redis.Nil) {
		return core.ErrCacheIsExpiredOrNotFound
	}

	return err
}
