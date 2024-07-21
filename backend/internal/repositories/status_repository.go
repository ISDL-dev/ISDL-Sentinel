package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
)

func GetAllStatusRepository() (statusList *sql.Rows, err error) {
	getRows, err := infrastructures.DB.Query("SELECT id, status_name FROM status;")
	if err != nil {
		return nil, fmt.Errorf("getRows db.Query error err:%w", err)
	}
	return getRows, nil
}

func GetInRoomPlaceIdRepository() (getStatusId int32, err error){
	getRows, err := infrastructures.DB.Query("SELECT id FROM place WHERE place_name = ?;", model.KC104)
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

func PutStatusRepository(userId int32, statusId int32, placeId int32) (err error) {
	var enteringHistoryId int32
	var enteredAt time.Time

	tx, err := infrastructures.DB.Begin()
	if err != nil {
		return fmt.Errorf("fail to begin transaction error err:%w", err)
	}
	if placeId == 0 {
		putOutRoomQuery := `
		UPDATE user 
		SET 
			place_id = ?, 
			status_id = ?
		WHERE 
			id = ?;`
		_, err = tx.Exec(putOutRoomQuery, sql.NullInt32{}, statusId, userId)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("putInRoomQuery error err:%w", err)
		}
		getLatestEnteringHistoryQuery :=`
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
		getRows, err := infrastructures.DB.Query(getLatestEnteringHistoryQuery, userId)
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
		insertOutRoomQuery := `
		INSERT INTO leaving_history (user_id, entering_history_id, left_at, stay_time)
		VALUES (
			?, 
			?, 
			NOW(), 
			TIMEDIFF(NOW(), ?)
		);`
		_, err = tx.Exec(insertOutRoomQuery, userId, enteringHistoryId, enteredAt)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insertOutRoomQuery error err:%w", err)
		}
	} else {
		putInRoomQuery := `
		UPDATE user 
		SET 
			place_id = ?, 
			status_id = ?
		WHERE 
			id = ?;`
		_, err = tx.Exec(putInRoomQuery, placeId, statusId, userId)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("putInRoomQuery error err:%w", err)
		}
		insertInRoomQuery := `
		INSERT INTO entering_history (user_id, entered_at)
		VALUES (?, NOW());`
		_, err = tx.Exec(insertInRoomQuery, userId)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insertInRoomQuery error err:%w", err)
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail to commit transaction error err:%w", err)
	}
	return nil
}