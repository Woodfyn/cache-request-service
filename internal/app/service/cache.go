package service

import (
	"context"
	"fmt"

	cache_service "github.com/Woodfyn/cache-service/pkg/proto"
)

type CacheRepositoryMencache interface {
}

type Cache struct {
	memcacheRepo CacheRepositoryMencache
}

func NewCache(memcacheRepo CacheRepositoryMencache) *Cache {
	return &Cache{
		memcacheRepo: memcacheRepo,
	}
}

func (c *Cache) Create(ctx context.Context, req *cache_service.CreateCacheRequest) (*cache_service.Empty, error) {
	return nil, nil
}

func (c *Cache) Get(ctx context.Context, req *cache_service.KeyCacheRequest) (*cache_service.CacheResponse, error) {
	return nil, nil
}

func (c *Cache) Delete(ctx context.Context, req *cache_service.KeyCacheRequest) (*cache_service.Empty, error) {
	fmt.Println(req)

	return nil, nil
}

func (c *Cache) Update(ctx context.Context, req *cache_service.UpdateCacheRequest) (*cache_service.Empty, error) {
	return nil, nil
}
