package server

import (
	"context"

	cache_service "github.com/Woodfyn/cache-service/pkg/proto"
)

func (h *Handler) Create(ctx context.Context, req *cache_service.CreateCacheRequest) (*cache_service.Empty, error) {
	return nil, nil
}

func (h *Handler) Get(ctx context.Context, req *cache_service.KeyCacheRequest) (*cache_service.CacheResponse, error) {
	return nil, nil
}

func (h *Handler) Delete(ctx context.Context, req *cache_service.KeyCacheRequest) (*cache_service.Empty, error) {
	return nil, nil
}

func (h *Handler) Update(ctx context.Context, req *cache_service.UpdateCacheRequest) (*cache_service.Empty, error) {
	return nil, nil
}
