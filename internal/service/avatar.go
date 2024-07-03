package service

import (
	"context"

	"github.com/Woodfyn/nexus-chat.cache-service/pkg/core"
	cache_service "github.com/Woodfyn/nexus-chat.cache-service/pkg/proto"
)

func (c *Cache) CreateAvatars(ctx context.Context, req *cache_service.Request_Create_Avatars) (*cache_service.Response_Empty, error) {
	err := c.repoRedis.CreateAvatars(ctx, req.GetKey(), core.NewAvatarFromCreateReq(req), req.GetTtl().AsDuration())

	return &cache_service.Response_Empty{}, err
}

func (c *Cache) GetAvatars(ctx context.Context, req *cache_service.Request_Key) (*cache_service.Response_Get_Avatars, error) {
	avatars, err := c.repoRedis.GetAvatars(ctx, req.GetKey())
	if err != nil {
		return nil, err
	}

	return core.NewGetRespFromAvatar(avatars), nil
}

func (c *Cache) UpdateAvatars(ctx context.Context, req *cache_service.Request_Update_Avatars) (*cache_service.Response_Empty, error) {
	err := c.repoRedis.UpdateAvatars(ctx, req.GetKey(), core.NewAvatarsFromUpdateReq(req), req.GetTtl().AsDuration())

	return &cache_service.Response_Empty{}, err
}
