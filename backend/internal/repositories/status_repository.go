package repositories

import (
	"database/sql"
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
)

func GetAllStatusRepository() (statusList *sql.Rows, err error) {
	getRows, err := infrastructures.DB.Query("SELECT id, status_name FROM status;")
	if err != nil {
		return nil, fmt.Errorf("getRows db.Query error err:%w", err)
	}
	return getRows, nil
}

func PutStatusRepository(userId int32, statusId int32) (err error) {
	_, err = infrastructures.DB.Query("UPDATE user SET status_id = ? WHERE id = ?;", statusId, userId)
	if err != nil {
		return fmt.Errorf("getRows db.Query error err:%w", err)
	} else {
		return nil
	}
}