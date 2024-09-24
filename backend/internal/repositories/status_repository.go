package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func JudgeNoMemberInRoom(kc104PlaceId int32) (isFirstEntering bool, err error) {
	getRows, err := infrastructures.DB.Query("SELECT id FROM user WHERE place_id = ?;", kc104PlaceId)
	if err != nil {
		return false, fmt.Errorf("getRows JudgeNoMemberInRoom Query error err:%w", err)
	}
	return !getRows.Next(), nil
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
	if status.Status == model.OUT_ROOM {
		putOutRoomQuery := `
		UPDATE user 
		SET 
			place_id = ?, 
			status_id = ?
		WHERE 
			id = ?;`
		_, err = tx.Exec(putOutRoomQuery, sql.NullInt32{}, statusId, status.UserId)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("putInRoomQuery error err:%w", err)
		}
		getLatestEnteringHistoryQuery := `
		SELECT 
			id AS enteringHistoryId, 
			entered_at AS enteredAt 
		FROM 
			entering_history 
		WHERE 
			user_id = ?
		ORDER BY 
			entered_at DESC 
		LIMIT 
			1;`
		getRows, err := infrastructures.DB.Query(getLatestEnteringHistoryQuery, status.UserId)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("getLatestEnteringHistoryQuery error err:%w", err)
		}
		for getRows.Next() {
			err := getRows.Scan(&enteringHistoryId, &enteredAt)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to find entering history: %v", err)
			}
		}
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("putInRoomQuery error err:%w", err)
		}
		isLastLeaving, err := JudgeNoMemberInRoom(placeId)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to find user id from leaving history:%w", err)
		}
		insertOutRoomQuery := `
		INSERT INTO leaving_history (user_id, entering_history_id, left_at, stay_time, is_last_leaving)
		VALUES (
			?, 
			?, 
			NOW(), 
			TIMEDIFF(NOW(), ?),
			?
		);`
		_, err = tx.Exec(insertOutRoomQuery, status.UserId, enteringHistoryId, enteredAt, isLastLeaving)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insertOutRoomQuery error err:%w", err)
		}
	} else if status.Status == model.IN_ROOM {
		isFirstEntering, err := JudgeNoMemberInRoom(placeId)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to find user id from entering history table:%w", err)
		}
		putInRoomQuery := `
		UPDATE user 
		SET 
			place_id = ?, 
			status_id = ?,
			current_entered_at = NOW()
		WHERE 
			id = ?;`
		_, err = tx.Exec(putInRoomQuery, placeId, statusId, status.UserId)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("putInRoomQuery error err:%w", err)
		}
		insertInRoomQuery := `
		INSERT INTO entering_history (user_id, entered_at, is_first_entering)
		VALUES (?, NOW(), ?);`
		_, err = tx.Exec(insertInRoomQuery, status.UserId, isFirstEntering)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insertInRoomQuery error err:%w", err)
		}
	} else {
		putOvernightQuery := `
		UPDATE user
		SET status_id = ?
		WHERE id = ?`
		_, err = tx.Exec(putOvernightQuery, statusId, status.UserId)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("putOvernightQuery error err:%w", err)
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail to commit transaction error err:%w", err)
	}
	return nil
}

func UpdateUserStatusToOutRoom() error {
	if err := infrastructures.DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	tx, err := infrastructures.DB.Begin()
	if err != nil {
		return fmt.Errorf("fail to begin transaction: %w", err)
	}

	var overnightCount int
	checkOvernightQuery := `
		SELECT COUNT(*) 
		FROM user u
		JOIN status s ON u.status_id = s.id
		WHERE s.status_name = ?`
	err = tx.QueryRow(checkOvernightQuery, model.OVERNIGHT).Scan(&overnightCount)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to check overnight users: %w", err)
	}

	if overnightCount == 0 {
		// 最新のidを一時テーブルに保存してから更新
		updateLastLeavingQuery := `
			UPDATE leaving_history
			JOIN (
				SELECT id 
				FROM leaving_history 
				ORDER BY left_at DESC 
				LIMIT 1
			) AS latest ON leaving_history.id = latest.id
			SET leaving_history.is_last_leaving = true`

		_, err = tx.Exec(updateLastLeavingQuery)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update is_last_leaving: %w", err)
		}
	}

	outRoomsStatusId, err := GetStatusId(model.OUT_ROOM)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get OutRoom status id: %w", err)
	}

	getUserIdQuery := `
		SELECT u.id 
		FROM user u
		JOIN status s ON u.status_id = s.id
		WHERE s.status_name = ?`
	rows, err := tx.Query(getUserIdQuery, model.IN_ROOM)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close() // リソース解放のために defer で rows をクローズ

	for rows.Next() {
		var userId int32
		if err := rows.Scan(&userId); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to scan user: %w", err)
		}

		updateUserQuery := `
		UPDATE user
		SET status_id = ?
		WHERE id = ?`
		_, err = tx.Exec(updateUserQuery, outRoomsStatusId, userId)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update user status: %w", err)
		}

		getLatestEnteringHistoryQuery := `
		SELECT 
			id AS enteringHistoryId, 
			entered_at AS enteredAt 
		FROM 
			entering_history 
		WHERE 
			user_id = ?
		ORDER BY 
			entered_at DESC 
		LIMIT 
			1;`
		var enteringHistoryId int
		var enteredAt time.Time
		err = tx.QueryRow(getLatestEnteringHistoryQuery, userId).Scan(&enteringHistoryId, &enteredAt)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to find entering history: %w", err)
		}

		insertOutRoomQuery := `
		INSERT INTO leaving_history (user_id, entering_history_id, left_at, stay_time, is_last_leaving)
		VALUES (
			?, 
			?, 
			?, 
			?, 
			?
		);`
		_, err = tx.Exec(insertOutRoomQuery, userId, enteringHistoryId, sql.NullTime{}, "00:00:00", false)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert into leaving_history: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail to commit transaction: %w", err)
	}

	return nil
}
