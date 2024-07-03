package handler

import (
	"context"
	"errors"

	"github.com/Woodfyn/nexus-chat.cache-service/pkg/core"
	cache_service "github.com/Woodfyn/nexus-chat.cache-service/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (h *Handler) CreateAvatars(ctx context.Context, req *cache_service.Request_Create_Avatars) (*cache_service.Response_Empty, error) {
	empty, err := h.cacheService.CreateAvatars(ctx, req)
	if err != nil {
		return nil, NewErrorResponse(codes.Internal, err)
	}

	return empty, nil
}

func (h *Handler) GetAvatars(ctx context.Context, req *cache_service.Request_Key) (*cache_service.Response_Get_Avatars, error) {
	resp, err := h.cacheService.GetAvatars(ctx, req)
	if errors.Is(err, core.ErrCacheIsExpiredOrNotFound) {
		return nil, NewErrorResponse(codes.NotFound, err)
	} else if err != nil {
		return nil, NewErrorResponse(codes.Internal, err)
	}

	return resp, nil
}

func (h *Handler) UpdateAvatars(ctx context.Context, req *cache_service.Request_Update_Avatars) (*cache_service.Response_Empty, error) {
	empty, err := h.cacheService.UpdateAvatars(ctx, req)
	if errors.Is(err, core.ErrCacheIsExpiredOrNotFound) {
		return nil, NewErrorResponse(codes.NotFound, err)
	} else if err != nil {
		return nil, NewErrorResponse(codes.Internal, err)
	}

	return empty, nil
}
