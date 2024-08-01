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

func DeleteAvatarService(avatarRequest schema.Avatar) (err error) {
	err = repositories.DeleteAvatarRepository(avatarRequest)
	if err != nil {
		return err
	}

	return nil
}
