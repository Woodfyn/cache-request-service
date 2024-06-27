package service

import (
	"context"
	"time"

	cache_service "github.com/Woodfyn/chat-api-cache-service/pkg/proto"
)

type CacheRepositoryRedis interface {
	Create(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Get(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
	Update(ctx context.Context, key string, value []byte, ttl time.Duration) error
}

type Cache struct {
	repo CacheRepositoryRedis
}

func NewCache(repoRedis CacheRepositoryRedis) *Cache {
	return &Cache{
		repo: repoRedis,
	}
}

func (c *Cache) Create(ctx context.Context, req *cache_service.CreateCacheRequest) (*cache_service.Empty, error) {
	if err := c.repo.Create(ctx, req.GetKey(), req.GetValue(), req.GetTtl().AsDuration()); err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Cache) Get(ctx context.Context, req *cache_service.KeyCacheRequest) (*cache_service.CacheResponse, error) {
	value, err := c.repo.Get(ctx, req.GetKey())
	if err != nil {
		return nil, err
	}

	return &cache_service.CacheResponse{Value: value}, nil
}

func (c *Cache) Delete(ctx context.Context, req *cache_service.KeyCacheRequest) (*cache_service.Empty, error) {
	err := c.repo.Delete(ctx, req.GetKey())

	return nil, err
}

func (c *Cache) Update(ctx context.Context, req *cache_service.UpdateCacheRequest) (*cache_service.Empty, error) {
	err := c.repo.Update(ctx, req.GetKey(), req.GetValue(), req.GetTtl().AsDuration())

	return nil, err
}
