package handler

import (
	"context"

	cache_service "github.com/Woodfyn/cache-service/pkg/proto"
)

type CacheService interface {
	Create(context.Context, *cache_service.CreateCacheRequest) (*cache_service.Empty, error)
	Get(context.Context, *cache_service.KeyCacheRequest) (*cache_service.CacheResponse, error)
	Delete(context.Context, *cache_service.KeyCacheRequest) (*cache_service.Empty, error)
	Update(context.Context, *cache_service.UpdateCacheRequest) (*cache_service.Empty, error)
}

type Handler struct {
	cacheService CacheService
	cache_service.UnimplementedCacheServiceServer
}

func NewHandler(cacheService CacheService) *Handler {
	return &Handler{
		cacheService: cacheService,
	}
}
