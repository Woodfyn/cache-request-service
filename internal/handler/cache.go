package handler

import (
	"context"
	"errors"

	"github.com/Woodfyn/chat-api-cache-service/pkg/core"
	cache_service "github.com/Woodfyn/chat-api-cache-service/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (h *Handler) Delete(ctx context.Context, req *cache_service.Request_Key) (*cache_service.Response_Empty, error) {
	empty, err := h.cacheService.Delete(ctx, req)
	if err != nil {
		if errors.Is(err, core.ErrCacheIsExpiredOrNotFound) {
			return nil, NewErrorResponse(codes.NotFound, err)
		}

		return nil, NewErrorResponse(codes.Internal, err)
	}

	return empty, nil
}
