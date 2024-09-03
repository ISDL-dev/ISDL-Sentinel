package services

import (
	"fmt"
	"mime/multipart"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func PostAvatarService(userId int, avatarFile *multipart.FileHeader) (err error) {
	avatarImgPath, err := infrastructures.UploadAvatarFile(avatarFile)
	if err != nil {
		return fmt.Errorf("failed to upload avatar file: %w", err)
	}

	err = repositories.PostAvatarRepository(userId, avatarImgPath)
	if err != nil {
		return fmt.Errorf("failed to save avatar link: %w", err)
	}

	return nil
}

func PutAvatarService(avatarRequest schema.Avatar) (avatarResponse schema.Avatar, err error) {
	avatarResponse, err = repositories.PutAvatarRepository(avatarRequest)
	if err != nil {
		return schema.Avatar{}, err
	}

	return avatarResponse, nil
}

func DeleteAvatarService(avatarRequest schema.Avatar) (err error) {
	avatarImgPath, err := repositories.DeleteAvatarRepository(avatarRequest)
	if err != nil {
		return err
	}

	err = infrastructures.DeleteAvatarFile(avatarImgPath)
	if err != nil {
		return fmt.Errorf("failed to delete avatar file: %w", err)
	}

	return nil
}
