package handler

import (
	"context"

	cache_service "github.com/Woodfyn/chat-api-cache-service/pkg/proto"
)

type CacheService interface {
	CreateAvatars(ctx context.Context, req *cache_service.Request_Create_Avatars) (*cache_service.Response_Empty, error)
	CreateProfile(ctx context.Context, req *cache_service.Request_Create_Profile) (*cache_service.Response_Empty, error)
	Delete(ctx context.Context, req *cache_service.Request_Key) (*cache_service.Response_Empty, error)
	GetAvatars(ctx context.Context, req *cache_service.Request_Key) (*cache_service.Response_Get_Avatars, error)
	GetProfile(ctx context.Context, req *cache_service.Request_Key) (*cache_service.Response_Get_Profile, error)
	UpdateAvatars(ctx context.Context, req *cache_service.Request_Update_Avatars) (*cache_service.Response_Empty, error)
	UpdateProfile(ctx context.Context, req *cache_service.Request_Update_Profile) (*cache_service.Response_Empty, error)
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
