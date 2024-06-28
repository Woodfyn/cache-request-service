package handler

import (
	"context"
	"errors"

	"github.com/Woodfyn/chat-api-cache-service/internal/core"
	cache_service "github.com/Woodfyn/chat-api-cache-service/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (h *Handler) Create(ctx context.Context, req *cache_service.CreateCacheRequest) (*cache_service.Empty, error) {
	if _, err := h.cacheService.Create(ctx, req); err != nil {
		return nil, NewErrorResponse(codes.Internal, err)
	}

	return nil, nil
}

func (h *Handler) Get(ctx context.Context, req *cache_service.KeyCacheRequest) (*cache_service.CacheResponse, error) {
	resp, err := h.cacheService.Get(ctx, req)
	if errors.Is(err, core.ErrCacheIsExpiredOrNotFound) {
		return nil, NewErrorResponse(codes.NotFound, err)
	} else if err != nil {
		return nil, NewErrorResponse(codes.Internal, err)
	}

	return resp, nil
}

func (h *Handler) Delete(ctx context.Context, req *cache_service.KeyCacheRequest) (*cache_service.Empty, error) {
	if _, err := h.cacheService.Delete(ctx, req); err != nil {
		if errors.Is(err, core.ErrCacheIsExpiredOrNotFound) {
			return nil, NewErrorResponse(codes.NotFound, err)
		}

		return nil, NewErrorResponse(codes.Internal, err)
	}

	return nil, nil
}

func (h *Handler) Update(ctx context.Context, req *cache_service.UpdateCacheRequest) (*cache_service.Empty, error) {
	if _, err := h.cacheService.Update(ctx, req); err != nil {
		if errors.Is(err, core.ErrCacheIsExpiredOrNotFound) {
			return nil, NewErrorResponse(codes.NotFound, err)
		}

		return nil, NewErrorResponse(codes.Internal, err)
	}

	return nil, nil
}
