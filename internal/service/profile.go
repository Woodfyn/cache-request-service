package service

import (
	"context"

	"github.com/Woodfyn/chat-api-cache-service/pkg/core"
	cache_service "github.com/Woodfyn/chat-api-cache-service/pkg/proto"
)

func (c *Cache) CreateProfile(ctx context.Context, req *cache_service.Request_Create_Profile) (*cache_service.Response_Empty, error) {
	err := c.repoRedis.CreateProfile(ctx, req.GetKey(), core.NewProfileFromCreateReq(req), req.GetTtl().AsDuration())

	return &cache_service.Response_Empty{}, err
}

func (c *Cache) GetProfile(ctx context.Context, req *cache_service.Request_Key) (*cache_service.Response_Get_Profile, error) {
	value, err := c.repoRedis.GetProfile(ctx, req.GetKey())
	if err != nil {
		return nil, err
	}

	return core.NewGetRespFromProfile(value), nil
}

func (c *Cache) UpdateProfile(ctx context.Context, req *cache_service.Request_Update_Profile) (*cache_service.Response_Empty, error) {
	err := c.repoRedis.UpdateProfile(ctx, req.GetKey(), core.NewProfileFromUpdateReq(req), req.GetTtl().AsDuration())

	return &cache_service.Response_Empty{}, err
}
