package core

import cache_service "github.com/Woodfyn/chat-api-cache-service/pkg/proto"

type Avatar struct {
	ID        int64
	AvatarUrl string
}

func NewAvatarFromCreateReq(req *cache_service.Request_Create_Avatars) []*Avatar {
	var avatars []*Avatar

	for _, value := range req.GetAvatarCache() {
		avatars = append(avatars, &Avatar{
			ID:        value.Id,
			AvatarUrl: value.AvatarUrl,
		})
	}

	return avatars
}

func NewGetRespFromAvatar(avatars []*Avatar) *cache_service.Response_Get_Avatars {
	var resp []*cache_service.Cache_Avatar
	for _, v := range avatars {
		resp = append(resp, &cache_service.Cache_Avatar{
			Id:        v.ID,
			AvatarUrl: v.AvatarUrl,
		})
	}

	return &cache_service.Response_Get_Avatars{
		AvatarCache: resp,
	}
}

func NewAvatarsFromUpdateReq(req *cache_service.Request_Update_Avatars) []*Avatar {
	var avatars []*Avatar

	for _, value := range req.GetAvatarCache() {
		avatars = append(avatars, &Avatar{
			ID:        value.Id,
			AvatarUrl: value.AvatarUrl,
		})
	}

	return avatars
}
