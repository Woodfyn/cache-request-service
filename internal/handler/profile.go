package handler

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Woodfyn/nexus-chat.cache-service/pkg/core"
	cache_service "github.com/Woodfyn/nexus-chat.cache-service/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (h *Handler) CreateProfile(ctx context.Context, req *cache_service.Request_Create_Profile) (*cache_service.Response_Empty, error) {
	slog.Info("Take request", "request", req)
	empty, err := h.cacheService.CreateProfile(ctx, req)
	if err != nil {
		return nil, NewErrorResponse(codes.Internal, err)
	}

	return empty, nil
}

func (h *Handler) GetProfile(ctx context.Context, req *cache_service.Request_Key) (*cache_service.Response_Get_Profile, error) {
	slog.Info("Take request", "request", req)
	resp, err := h.cacheService.GetProfile(ctx, req)
	if errors.Is(err, core.ErrCacheIsExpiredOrNotFound) {
		return nil, NewErrorResponse(codes.NotFound, err)
	} else if err != nil {
		return nil, NewErrorResponse(codes.Internal, err)
	}

	return resp, nil
}

func (h *Handler) UpdateProfile(ctx context.Context, req *cache_service.Request_Update_Profile) (*cache_service.Response_Empty, error) {
	slog.Info("Take request", "request", req)
	empty, err := h.cacheService.UpdateProfile(ctx, req)
	if errors.Is(err, core.ErrCacheIsExpiredOrNotFound) {
		return nil, NewErrorResponse(codes.NotFound, err)
	} else if err != nil {
		return nil, NewErrorResponse(codes.Internal, err)
	}

	return empty, nil
}
