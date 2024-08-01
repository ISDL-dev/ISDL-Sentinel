package repositories

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func PutAvatarRepository(avatarRequest schema.Avatar) (avatarResponse schema.Avatar, err error) {
	tx, err := infrastructures.DB.Begin()
	if err != nil {
		return schema.Avatar{}, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	putAvatarQuery := `UPDATE user SET avatar_id = ? WHERE id = ?;`
	_, err = tx.Exec(putAvatarQuery, avatarRequest.AvatarId, avatarRequest.UserId)
	if err != nil {
		return schema.Avatar{}, fmt.Errorf("failed to execute query to put avatar: %v", err)
	}

	getAvatarIdQuery := `SELECT id, avatar_id FROM user WHERE id = ?;`
	if err := tx.QueryRow(getAvatarIdQuery, avatarRequest.UserId).Scan(&avatarResponse.UserId, &avatarResponse.AvatarId); err != nil {
		return schema.Avatar{}, fmt.Errorf("failed to execute query to get avatar_id: %v", err)
	}

	return avatarResponse, nil
}

func DeleteAvatarRepository(avatarRequest schema.Avatar) (err error) {
	tx, err := infrastructures.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	deleteAvatarQuery := `DELETE FROM avatar WHERE id = ?;`
	_, err = tx.Exec(deleteAvatarQuery, avatarRequest.AvatarId)
	if err != nil {
		return fmt.Errorf("failed to execute query to delete avatar: %v", err)
	}

	getCurrentAvatarIdQuery := `SELECT avatar_id FROM user WHERE id = ?;`
	var currentAvatarId int32
	if err := tx.QueryRow(getCurrentAvatarIdQuery, avatarRequest.UserId).Scan(&currentAvatarId); err != nil {
		return fmt.Errorf("failed to execute query to get current avatar id: %v", err)
	}

	if currentAvatarId == avatarRequest.AvatarId {
		getNewAvatarQuery := `
			SELECT avatar_id
			FROM user_possession_avatar
			WHERE user_id = ?
			LIMIT 1;`
		var newAvatarId int32
		if err := tx.QueryRow(getNewAvatarQuery, avatarRequest.UserId).Scan(&newAvatarId); err != nil {
			return fmt.Errorf("failed to execute query to get new avatar: %v", err)
		}

		updateUserAvatarQuery := `UPDATE user SET avatar_id = ? WHERE id = ?;`
		_, err = tx.Exec(updateUserAvatarQuery, newAvatarId, avatarRequest.UserId)
		if err != nil {
			return fmt.Errorf("failed to execute query to update user avatar: %v", err)
		}
	}

	return nil
}
