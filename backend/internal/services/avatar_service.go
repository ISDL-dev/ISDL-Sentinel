package services

import (
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func PutAvatarService(avatarRequest schema.PutAvatarRequest) (err error) {
	err = repositories.PutAvatarRepository(avatarRequest)
	if err != nil {
		return err
	}

	return nil
}
