package repositories

import (
	"database/sql"
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
)

func GetInRoomUserListRepository(inRoomStatusId int32) (userList *sql.Rows, err error) {
	getInRoomUserListQuery := `
		SELECT 
			u.id AS user_id,
			u.name AS user_name,
			p.place_name,
			s.status_name,
			g.grade_name,
			a.avatar_name,
			a.img_path,
			eh.latest_entered_at
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
	getRows, err := infrastructures.DB.Query(getInRoomUserListQuery, inRoomStatusId)
	if err != nil {
		return nil, fmt.Errorf("getRows getInRoomUserListQuery error err:%w", err)
	}
	return getRows, nil
}
