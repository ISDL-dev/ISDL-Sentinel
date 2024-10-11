package repositories

import (
	"fmt"
	"log"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetUsersRepository(userId int, date string) (userInformation schema.GetUserById200Response, err error) {
	var avatar schema.GetUserById200ResponseAvatarListInner
	var avatarList []schema.GetUserById200ResponseAvatarListInner
	var roleName string
	var roleList []string

	// Query to get user information
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
	if err := infrastructures.DB.QueryRow(
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

	// Query to get attendance days and stay time
	getMonthStaytimeAndDaysQuery := `
		SELECT 
			COUNT(DISTINCT DATE(left_at)) AS AttendanceDays,
			IFNULL(SEC_TO_TIME(SUM(TIME_TO_SEC(stay_time))), '0:00') AS StayTime
		FROM 
			leaving_history
		WHERE 
			user_id = ?
			AND DATE_FORMAT(left_at, '%Y-%m') = ?;`
	if err := infrastructures.DB.QueryRow(getMonthStaytimeAndDaysQuery, userId, date).Scan(&userInformation.AttendanceDays, &userInformation.StayTime); err != nil {
		return schema.GetUserById200Response{}, fmt.Errorf("failed to execute query to get staytime and attendance days: %v", err)
	}

	// Query to get the list of avatars
	getAvatarListQuery := `
		SELECT 
			avatar.id AS AvatarId,
			avatar.img_path AS ImgPath 
		FROM 
			user_possession_avatar 
		JOIN 
			avatar ON user_possession_avatar.avatar_id = avatar.id 
		WHERE 
			user_possession_avatar.user_id = ?;`
	rows, err := infrastructures.DB.Query(getAvatarListQuery, userId)
	if err != nil {
		return schema.GetUserById200Response{}, fmt.Errorf("getRows db.Query error err:%w", err)
	}
	for rows.Next() {
		err := rows.Scan(&avatar.AvatarId, &avatar.ImgPath)
		if err != nil {
			return schema.GetUserById200Response{}, fmt.Errorf("failed to execute query to get avatar list:%w", err)
		}
		avatarList = append(avatarList, avatar)
	}
	userInformation.AvatarList = avatarList

	// Query to get the list of roles for the user
	getRoleListQuery := `
		SELECT 
			role.role_name 
		FROM 
			user_possession_role
		JOIN 
			role ON user_possession_role.role_id = role.id
		WHERE 
			user_possession_role.user_id = ?;`
	roleRows, err := infrastructures.DB.Query(getRoleListQuery, userId)
	if err != nil {
		return schema.GetUserById200Response{}, fmt.Errorf("getRoles db.Query error err:%w", err)
	}
	defer roleRows.Close()

	// Append the role names to the RoleList
	for roleRows.Next() {
		err := roleRows.Scan(&roleName)
		if err != nil {
			return schema.GetUserById200Response{}, fmt.Errorf("failed to execute query to get role list:%w", err)
		}
		roleList = append(roleList, roleName)
	}
	userInformation.RoleList = roleList

	return userInformation, nil
}

func GetTeacherMailAddress() (teacherMailAddress []string, err error) {
	var mailAddress string

	getUsersMailAddressListQuery := `
		SELECT 
			u.mail_address 
		FROM 
			user u
		INNER JOIN 
			grade g ON u.grade_id = g.id
		WHERE 
			g.grade_name = ?;`
	rows, err := infrastructures.DB.Query(getUsersMailAddressListQuery, model.Teacher)
	if err != nil {
		return nil, fmt.Errorf("getRows db.Query error err:%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&mailAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to scan mail address: %w", err)
		}
		teacherMailAddress = append(teacherMailAddress, mailAddress)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return teacherMailAddress, nil
}

func PutUsersRepository(userId int, userInformation schema.PutUserByIdRequest) (err error) {
	log.Printf("Grade from JSON: '%s'", userInformation.UserName)
	tx, err := infrastructures.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Single query to update user's name, mail_address, and grade_id
	updateUserQuery := `
		UPDATE user
		SET name = ?, mail_address = ?, grade_id = (SELECT id FROM grade WHERE grade_name = ?)
		WHERE id = ?;
	`
	_, err = tx.Exec(updateUserQuery, userInformation.UserName, userInformation.MailAddress, userInformation.Grade, userId)
	if err != nil {
		return fmt.Errorf("failed to update user information and grade: %v", err)
	}

	// Delete existing roles for the user in user_possession_role
	deleteRolesQuery := `DELETE FROM user_possession_role WHERE user_id = ?;`
	_, err = tx.Exec(deleteRolesQuery, userId)
	if err != nil {
		return fmt.Errorf("failed to delete user's existing roles: %v", err)
	}

	// Insert new roles by retrieving role IDs from the role table
	insertRoleQuery := `INSERT INTO user_possession_role (user_id, role_id) VALUES (?, (SELECT id FROM role WHERE role_name = ?));`
	for _, roleName := range userInformation.RoleList {
		_, err = tx.Exec(insertRoleQuery, userId, roleName)
		if err != nil {
			return fmt.Errorf("failed to insert role for user: %v", err)
		}
	}

	return nil
}
