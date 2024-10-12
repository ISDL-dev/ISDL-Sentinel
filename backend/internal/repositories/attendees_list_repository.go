package repositories

import (
	"fmt"
	"strings"
	"time"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetInRoomStatusIdRepository() (getStatusId int32, err error) {
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
		CASE
			WHEN p.place_name = ? THEN IFNULL(eh.latest_entered_at, '2024-07-16 09:00:00')
			ELSE IFNULL(u.current_entered_at, '2024-07-16 09:00:00')
		END AS EnteredAt
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
		s.status_name IN (?, ?)
		AND g.grade_name != 'OB';`
	getRows, err := infrastructures.DB.Query(getInRoomUserListQuery, model.KC104, model.IN_ROOM, model.OVERNIGHT)
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

func UpdateInRoomUserFromCalendarRepository(eventList []model.Calendar) (err error) {
	now := time.Now().UTC()
	layout := time.RFC3339
	for _, room := range eventList {
		startTime, err := time.Parse(layout, room.StartDate)
		if err != nil {
			return fmt.Errorf("Error parsing date: %v", err)
		}
		endTime, err := time.Parse(layout, room.EndDate)
		if err != nil {
			return fmt.Errorf("Error parsing date: %v", err)
		}
		if startTime.Before(now) && endTime.After(now) {
			fmt.Printf("calendar %s.\n", room.AttendeeMail[0])
			fmt.Printf("room %s.\n", room.RoomName)

			emailPlaceholders := strings.Repeat("?,", len(room.AttendeeMail))
			emailPlaceholders = strings.TrimSuffix(emailPlaceholders, ",")

			UpdateInRoomUserFromCalendarQuery := fmt.Sprintf(`
				UPDATE user
				SET 
					place_id = (
						SELECT id
						FROM place
						WHERE place_name = ?
					),
					current_entered_at = ?
				WHERE mail_address IN (%s)`, emailPlaceholders)

			args := append([]interface{}{room.RoomName, startTime.Add(9 * time.Hour)}, model.ToInterfaceSlice(room.AttendeeMail)...)
			_, err := infrastructures.DB.Exec(UpdateInRoomUserFromCalendarQuery, args...)
			if err != nil {
				return fmt.Errorf("failed to execute query: %v", err)
			}
			return nil
		}
	}
	return nil
}

func DeleteRoomFromCalendarRepository() (err error) {
	DeleteRoomFromCalendarQuery := `
		UPDATE user
		SET place_id = (SELECT id FROM place WHERE place_name = ?)
		WHERE place_id IS NOT NULL 
		AND place_id != (SELECT id FROM place WHERE place_name = ?);`

	_, err = infrastructures.DB.Exec(DeleteRoomFromCalendarQuery, model.KC104, model.KC104)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}
	return nil
}
