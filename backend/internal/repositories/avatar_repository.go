package repositories

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func PutAvatarRepository(avatarRequest schema.PutAvatarRequest) (err error) {
	putAvatarQuery := `UPDATE user SET avatar_id = ? WHERE id = ?;`
	_, err = infrastructures.DB.Exec(putAvatarQuery, avatarRequest.AvatarId, avatarRequest.UserId)
	if err != nil {
		return fmt.Errorf("failed to execute query to put avatar: %v", err)
	}

	return nil
}
