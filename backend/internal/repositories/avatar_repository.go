package repositories

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func PutAvatarRepository(avatarRequest schema.Avatar) (avatarResponse schema.Avatar, err error) {
	putAvatarQuery := `UPDATE user SET avatar_id = ? WHERE id = ?;`
	_, err = infrastructures.DB.Exec(putAvatarQuery, avatarRequest.AvatarId, avatarRequest.UserId)
	if err != nil {
		return schema.Avatar{}, fmt.Errorf("failed to execute query to put avatar: %v", err)
	}

	getAvatarIdQuery := `SELECT id, avatar_id FROM user WHERE id = ?;`
	if err := infrastructures.DB.QueryRow(getAvatarIdQuery, avatarRequest.UserId).Scan(&avatarResponse.UserId, &avatarResponse.AvatarId); err != nil {
		return schema.Avatar{}, fmt.Errorf("failed to execute a query to get avatar_id: %v", err)
	}

	return avatarResponse, nil
}
