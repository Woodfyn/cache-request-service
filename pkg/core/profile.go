package core

import cache_service "github.com/Woodfyn/nexus-chat.cache-service/pkg/proto"

type Profile struct {
	ID        int64
	Phone     string
	Username  string
	AvatarUrl string
}

func NewProfileFromCreateReq(req *cache_service.Request_Create_Profile) *Profile {
	return &Profile{
		ID:        req.GetProfileCache().GetId(),
		Phone:     req.GetProfileCache().GetPhone(),
		Username:  req.GetProfileCache().GetUsername(),
		AvatarUrl: req.GetProfileCache().GetAvatarUrl(),
	}
}

func NewGetRespFromProfile(profile *Profile) *cache_service.Response_Get_Profile {
	return &cache_service.Response_Get_Profile{
		ProfileCache: &cache_service.Cache_Profile{
			Id:        profile.ID,
			Phone:     profile.Phone,
			Username:  profile.Username,
			AvatarUrl: profile.AvatarUrl,
		},
	}
}

func NewProfileFromUpdateReq(req *cache_service.Request_Update_Profile) *Profile {
	return &Profile{
		ID:        req.GetProfileCache().GetId(),
		Phone:     req.GetProfileCache().GetPhone(),
		Username:  req.GetProfileCache().GetUsername(),
		AvatarUrl: req.GetProfileCache().GetAvatarUrl(),
	}
}
