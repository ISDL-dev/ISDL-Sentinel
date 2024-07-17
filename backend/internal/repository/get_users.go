package repository

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetUsers(userId int) (userInformation schema.UserInformation, err error) {
	var avatar schema.UserInformationAvatarListInner
	var avatarList []schema.UserInformationAvatarListInner

	query :=
		`SELECT 
			avatar.id AS AvatarId,
			avatar.avatar_name AS AvatarName,
			avatar.rarity AS Rarity,
			avatar.img_path AS ImgPath
		FROM 
			user_possession_avatar
		JOIN 
			avatar ON user_possession_avatar.avatar_id = avatar.id
		WHERE 
			user_possession_avatar.user_id = ?;
		`
	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("getRows db.Query error err:%w", err)
	}
	for rows.Next() {
		err := rows.Scan(&avatar.AvatarId, &avatar.AvatarName, &avatar.Rarity, &avatar.ImgPath)
		if err != nil {
			return nil, fmt.Errorf("getRows rows_title.Scan error err:%w", err)
		}
		avatarList = append(avatarList, avatar)
	}

	for rows.Next() {
		err := rows.Scan(&userInformation.UserId, &userInformation.UserName, &userInformation.MailAddress)
		if err != nil {
			return nil, fmt.Errorf("getRows rows_title.Scan error err:%w", err)
		}
		imagesList = append(imagesList, image)
	}

	return userInformation, nil
}
