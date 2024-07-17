package repository

import (
	"fmt"

	"github.com/ISDL-dev/ISDL_Sentinel/backend/internal/model"
	"github.com/ISDL-dev/ISDL_Sentinel/backend/internal/schema"
)

func SelectApplyStatusId(status string) (id int32, err error) {
	var getStatusId int32 
	var getStatusName string
	var targetStatusName string
	getRows, err := db.Query("SELECT id, status_name FROM status")
	if err != nil {
		return 0, fmt.Errorf("getRows db.Query error err:%w", err)
	}
	if (model.IsInRoom(status)){
		targetStatusName = model.OUT_ROOM
	} else {
		targetStatusName = model.IN_ROOM
	}
	for getRows.Next() {
		err := getRows.Scan(&getStatusId, &getStatusName)
		if err != nil {
			return 0, fmt.Errorf("getRows rows_title.Scan error err:%w", err)
		}
		if getStatusName == targetStatusName{
			return getStatusId, nil 
		}
	}
	return 0, nil
}

func PutStatus(user schema.PutStatusRequest) error {
	var applyStatusId int32
	update, err := db.Prepare("UPDATE user SET status_id = ? WHERE id ?")
	if err != nil {
		return fmt.Errorf("failed to prepare for a query to update signal: %v", err)
	}
	applyStatusId, err = SelectApplyStatusId(user.Status)
	if err != nil {
		return fmt.Errorf("failed to select change status: %v", err)
	}
	_, err = update.Exec(applyStatusId, user.UserId)
	if err != nil {
		return fmt.Errorf("db.Query error err:%w", err)
	}

	return nil
}