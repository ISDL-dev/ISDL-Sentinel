package services

import (
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func PutAvatarService(avatarRequest schema.Avatar) (avatarResponse schema.Avatar, err error) {
	avatarResponse, err = repositories.PutAvatarRepository(avatarRequest)
	if err != nil {
		return schema.Avatar{}, err
	}

	return avatarResponse, nil
}
