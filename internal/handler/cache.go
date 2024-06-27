package handler

import (
	"context"

	cache_service "github.com/Woodfyn/cache-service/pkg/proto"
)

func (h *Handler) Create(ctx context.Context, req *cache_service.CreateCacheRequest) (*cache_service.Empty, error) {
	return h.cacheService.Create(ctx, req)
}

func (h *Handler) Get(ctx context.Context, req *cache_service.KeyCacheRequest) (*cache_service.CacheResponse, error) {
	return h.cacheService.Get(ctx, req)
}

func (h *Handler) Delete(ctx context.Context, req *cache_service.KeyCacheRequest) (*cache_service.Empty, error) {
	return h.cacheService.Delete(ctx, req)
}

func (h *Handler) Update(ctx context.Context, req *cache_service.UpdateCacheRequest) (*cache_service.Empty, error) {
	return h.cacheService.Update(ctx, req)
}
