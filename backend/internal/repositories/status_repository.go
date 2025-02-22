package repositories

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func JudgeNoMemberInRoom(tx *sql.Tx, kc104PlaceId int32) (isFirstEntering bool, retrunTx *sql.Tx, err error) {
	getRows, err := tx.Query("SELECT id FROM user WHERE place_id = ?;", kc104PlaceId)
	if err != nil {
		return false, tx, fmt.Errorf("getRows JudgeNoMemberInRoom Query error err:%w", err)
	}
	defer getRows.Close()

	return !getRows.Next(), tx, nil
}

func GetStatusId(status string) (statusId int32, err error) {
	getRows, err := infrastructures.DB.Query("SELECT id FROM status WHERE status_name = ?;", status)
	if err != nil {
		return 0, fmt.Errorf("getRows db.Query error err:%w", err)
	}
	for getRows.Next() {
		err := getRows.Scan(&statusId)
		if err != nil {
			return 0, fmt.Errorf("failed to find target status id: %v", err)
		}
	}
	return statusId, nil
}

func GetPlaceId() (placeId int32, err error) {
	getRows, err := infrastructures.DB.Query("SELECT id FROM place WHERE place_name = ?;", model.KC104)
	if err != nil {
		return 0, fmt.Errorf("getRows GetInRoomStatusId Query error err:%w", err)
	}
	for getRows.Next() {
		err := getRows.Scan(&placeId)
		if err != nil {
			return 0, fmt.Errorf("failed to find target status id: %v", err)
		}
	}
	return placeId, nil
}

func PutStatusRepository(status schema.Status, placeId int32) (err error) {
	var enteringHistoryId int32
	var enteredAt time.Time

	statusId, err := GetStatusId(status.Status)
	if err != nil {
		return fmt.Errorf("failed to get status id: %v", err)
	}

	tx, err := infrastructures.DB.Begin()
	if err != nil {
		return fmt.Errorf("fail to begin transaction error err:%w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if status.Status == model.OUT_ROOM {
		_, err = tx.Exec(`
        UPDATE user 
        SET place_id = NULL, status_id = ?
        WHERE id = ?;`, statusId, status.UserId)
		if err != nil {
			return fmt.Errorf("putOutRoomQuery error err:%w", err)
		}

		err = tx.QueryRow(`
        SELECT id, entered_at 
        FROM entering_history 
        WHERE user_id = ? 
        ORDER BY entered_at DESC LIMIT 1;`, status.UserId).Scan(&enteringHistoryId, &enteredAt)
		if err != nil {
			return fmt.Errorf("failed to find entering history: %v", err)
		}

		isLastLeaving, tx, err := JudgeNoMemberInRoom(tx, placeId)
		if err != nil {
			return fmt.Errorf("failed to check if last leaving: %w", err)
		}

		_, err = tx.Exec(`
        INSERT INTO leaving_history (user_id, entering_history_id, left_at, stay_time, is_last_leaving)
        VALUES (?, ?, NOW(), TIMEDIFF(NOW(), ?), ?);`, status.UserId, enteringHistoryId, enteredAt, isLastLeaving)
		if err != nil {
			return fmt.Errorf("insertOutRoomQuery error err:%w", err)
		}
	} else if status.Status == model.IN_ROOM {
		isFirstEntering, tx, err := JudgeNoMemberInRoom(tx, placeId)
		if err != nil {
			return fmt.Errorf("failed to check first entering: %w", err)
		}

		_, err = tx.Exec(`
        UPDATE user 
        SET place_id = ?, status_id = ?, current_entered_at = NOW()
        WHERE id = ?;`, placeId, statusId, status.UserId)
		if err != nil {
			return fmt.Errorf("putInRoomQuery error err:%w", err)
		}

		_, err = tx.Exec(`
        INSERT INTO entering_history (user_id, entered_at, is_first_entering)
        VALUES (?, NOW(), ?);`, status.UserId, isFirstEntering)
		if err != nil {
			return fmt.Errorf("insertInRoomQuery error err:%w", err)
		}
	} else {
		_, err = tx.Exec(`
        UPDATE user SET status_id = ? WHERE id = ?`, statusId, status.UserId)
		if err != nil {
			return fmt.Errorf("putOvernightQuery error err:%w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail to commit transaction: %w", err)
	}

	return nil
}

func UpdateUserStatusToOutRoom() error {
	var inRoomCount int
	var inRoomIDs []int

	query := `
		SELECT 
			s.status_name,
			COUNT(u.id) as count,
			GROUP_CONCAT(u.id) as ids
		FROM user u
		JOIN status s ON u.status_id = s.id
		WHERE s.status_name = ?
		GROUP BY s.status_name
	`
	rows, err := infrastructures.DB.Query(query, model.IN_ROOM)
	if err != nil {
		return fmt.Errorf("failed to query user counts: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var statusName string
		var count int
		var ids string
		if err := rows.Scan(&statusName, &count, &ids); err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}

		if statusName == model.IN_ROOM {
			inRoomCount = count
			inRoomIDs = stringToIntSlice(ids)
		}
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating rows: %w", err)
	}

	var outRoomStatusId int
	err = infrastructures.DB.QueryRow("SELECT id FROM status WHERE status_name = ?", model.OUT_ROOM).Scan(&outRoomStatusId)
	if err != nil {
		return fmt.Errorf("failed to get OUT_ROOM status id: %w", err)
	}

	var overnightStatusId int
	err = infrastructures.DB.QueryRow("SELECT id FROM status WHERE status_name = ?", model.OVERNIGHT).Scan(&overnightStatusId)
	if err != nil {
		return fmt.Errorf("failed to get OVERNIGHT status id: %w", err)
	}

	tx, err := infrastructures.DB.Begin()
	if err != nil {
		return fmt.Errorf("fail to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if inRoomCount > 0 {
		idArgs := make([]interface{}, len(inRoomIDs))
		for i, id := range inRoomIDs {
			idArgs[i] = id
		}

		placeholders := strings.Repeat("?,", len(inRoomIDs)-1) + "?"

		_, err = tx.Exec(fmt.Sprintf(`
			UPDATE user
			SET 
				status_id = ?,
				place_id = NULL
			WHERE id IN (%s)
		`, placeholders), append([]interface{}{outRoomStatusId}, idArgs...)...)
		if err != nil {
			return fmt.Errorf("failed to execute force logout query: %w", err)
		}

		insertQuery := fmt.Sprintf(`
			INSERT INTO leaving_history (user_id, entering_history_id, left_at, stay_time, is_last_leaving)
			WITH LastUser AS (
				SELECT MAX(u.id) as last_user_id
				FROM user u
				WHERE u.id IN (%s)
			)
			SELECT 
				u.id,
				eh.id,
				NOW(),
				'00:00:00',
				CASE 
					WHEN EXISTS (SELECT 1 FROM user WHERE status_id = ?) THEN false
					WHEN u.id = (SELECT last_user_id FROM LastUser) THEN true
					ELSE false
				END
			FROM user u
			JOIN (
				SELECT user_id, MAX(id) as id
				FROM entering_history
				WHERE user_id IN (%s)
				GROUP BY user_id
			) eh ON u.id = eh.user_id
			WHERE u.id IN (%s)
		`, placeholders, placeholders, placeholders)

		_, err = tx.Exec(insertQuery, append(append(append(idArgs, overnightStatusId), idArgs...), idArgs...)...)
		if err != nil {
			return fmt.Errorf("failed to execute force logout query: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail to commit transaction: %w", err)
	}

	return nil
}

func stringToIntSlice(s string) []int {
	var result []int
	for _, idStr := range strings.Split(s, ",") {
		if id, err := strconv.Atoi(idStr); err == nil {
			result = append(result, id)
		}
	}
	return result
}
