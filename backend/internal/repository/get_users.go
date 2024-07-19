package repository

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetUsers(userId int, date string) (userInformation schema.GetUserById200Response, err error) {
	var avatar schema.GetUserById200ResponseAvatarListInner
	var avatarList []schema.GetUserById200ResponseAvatarListInner

	getUserInfoQuery := `
		SELECT 
			user.id AS UserId,
			user.name AS UserName,
			user.mail_address AS MailAddress,
			user.number_of_coin AS NumberOfCoin,
			status.status_name AS Status,
			IFNULL(place.place_name, '') AS Place,
			grade.grade_name AS Grade,
			user.avatar_id AS AvatarId,
			avatar.img_path AS AvatarImgPath
		FROM 
			user
		JOIN 
			status ON user.status_id = status.id
		LEFT JOIN 
			place ON user.place_id = place.id
		JOIN 
			grade ON user.grade_id = grade.id
		LEFT JOIN 
			avatar ON user.avatar_id = avatar.id
		WHERE 
			user.id = ?;`
	if err := db.QueryRow(
		getUserInfoQuery,
		userId,
	).Scan(
		&userInformation.UserId,
		&userInformation.UserName,
		&userInformation.MailAddress,
		&userInformation.NumberOfCoin,
		&userInformation.Status,
		&userInformation.Place,
		&userInformation.Grade,
		&userInformation.AvatarId,
		&userInformation.AvatarImgPath); err != nil {
		return schema.GetUserById200Response{}, fmt.Errorf("failed to execute query to get user information: %v", err)
	}

	getMonthStaytimeAndDaysQuery := `
		SELECT 
			COUNT(DISTINCT DATE(left_at)) AS AttendanceDays,
			IFNULL(SEC_TO_TIME(SUM(TIME_TO_SEC(stay_time))), '0:00') AS StayTime
		FROM 
			leaving_history
		WHERE 
			user_id = ?
			AND DATE_FORMAT(left_at, '%Y-%m') = ?;`
	if err := db.QueryRow(getMonthStaytimeAndDaysQuery, userId, date).Scan(&userInformation.AttendanceDays, &userInformation.StayTime); err != nil {
		return schema.GetUserById200Response{}, fmt.Errorf("failed to execute query to get staytime and attendance days: %v", err)
	}

	getAvatarListQuery := `
		SELECT 
			avatar.id AS AvatarId,
			avatar.avatar_name AS AvatarName,
			avatar.rarity AS Rarity,
			avatar.img_path AS ImgPath 
		FROM 
			user_possession_avatar 
		JOIN 
			avatar ON user_possession_avatar.avatar_id = avatar.id 
		WHERE 
			user_possession_avatar.user_id = ?;`
	rows, err := db.Query(getAvatarListQuery, userId)
	if err != nil {
		return schema.GetUserById200Response{}, fmt.Errorf("getRows db.Query error err:%w", err)
	}
	for rows.Next() {
		err := rows.Scan(&avatar.AvatarId, &avatar.AvatarName, &avatar.Rarity, &avatar.ImgPath)
		if err != nil {
			return schema.GetUserById200Response{}, fmt.Errorf("failed to execute query to get avatar list:%w", err)
		}
		avatarList = append(avatarList, avatar)
	}
	userInformation.AvatarList = avatarList

	return userInformation, nil
}
