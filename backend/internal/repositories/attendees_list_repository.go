package repositories

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetInRoomStatusIdRepository() (getStatusId int32, err error){
	getRows, err := infrastructures.DB.Query("SELECT id FROM status WHERE status_name = ?;", model.IN_ROOM)
	if err != nil {
		return 0, fmt.Errorf("getRows GetInRoomStatusId Query error err:%w", err)
	}
	for getRows.Next() {
		err := getRows.Scan(&getStatusId)
		if err != nil {
			return 0, fmt.Errorf("failed to find target status id: %v", err)
		}
	}
	return getStatusId, nil
}

func GetInRoomUserListRepository(inRoomStatusId int32) (userList []schema.GetAttendeesList200ResponseInner, err error) {
	var attendee schema.GetAttendeesList200ResponseInner
	getInRoomUserListQuery := `
		SELECT 
			u.id AS UserId,
			u.name AS UserName,
			IFNULL(p.place_name, '') AS Place,
			s.status_name AS Status,
			g.grade_name AS Grade,
			u.avatar_id AS AvatarId,
			a.img_path AS AvatarImgPath,
			IFNULL(eh.latest_entered_at, '2024-07-16 09:00:00') AS EnteredAt
		FROM 
			user u
		JOIN 
			place p ON u.place_id = p.id
		JOIN 
			status s ON u.status_id = s.id
		JOIN 
			grade g ON u.grade_id = g.id
		JOIN 
			avatar a ON u.avatar_id = a.id
		LEFT JOIN 
			(
				SELECT 
					user_id, 
					MAX(entered_at) AS latest_entered_at
				FROM 
					entering_history
				GROUP BY 
					user_id
			) eh ON u.id = eh.user_id
		WHERE 
			u.status_id = ?;`
	getRows, err := infrastructures.DB.Query(getInRoomUserListQuery,inRoomStatusId)
	if err != nil {
		return []schema.GetAttendeesList200ResponseInner{}, fmt.Errorf("getRows getInRoomUserList Query error err:%w", err)
	}
	for getRows.Next() {
		err := getRows.Scan(
			&attendee.UserId,
			&attendee.UserName,
			&attendee.Place,
			&attendee.Status,
			&attendee.Grade,
			&attendee.AvatarId,
			&attendee.AvatarImgPath,
			&attendee.EnteredAt)
		if err != nil {
			return []schema.GetAttendeesList200ResponseInner{}, fmt.Errorf("failed to find target status id: %v", err)
		}
		userList = append(userList, attendee)
	}
	return userList, nil
}
