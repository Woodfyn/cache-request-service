package service

import (
	"context"
	"time"

	"github.com/Woodfyn/chat-api-cache-service/pkg/core"
	cache_service "github.com/Woodfyn/chat-api-cache-service/pkg/proto"
)

type CacheRepositoryRedis interface {
	CreateAvatars(ctx context.Context, key string, input []*core.Avatar, ttl time.Duration) error
	CreateProfile(ctx context.Context, key string, input *core.Profile, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	GetAvatars(ctx context.Context, key string) ([]*core.Avatar, error)
	GetProfile(ctx context.Context, key string) (*core.Profile, error)
	UpdateAvatars(ctx context.Context, key string, input []*core.Avatar, ttl time.Duration) error
	UpdateProfile(ctx context.Context, key string, input *core.Profile, ttl time.Duration) error
}

type Cache struct {
	repoRedis CacheRepositoryRedis
}

func NewCache(repoRedis CacheRepositoryRedis) *Cache {
	return &Cache{
		repoRedis: repoRedis,
	}
}

func (c *Cache) Delete(ctx context.Context, req *cache_service.Request_Key) (*cache_service.Response_Empty, error) {
	err := c.repoRedis.Delete(ctx, req.GetKey())

	return &cache_service.Response_Empty{}, err
}
