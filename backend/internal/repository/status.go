package repository

import (
	"fmt"

	"github.com/ISDL-dev/ISDL_Sentinel/backend/internal/model"
	"github.com/ISDL-dev/ISDL_Sentinel/backend/internal/schema"
)

func SelectApplyStatusId(status string) (id int32, statusName string) {
	var getStatusId int32 
	var getStatusName string
	var targetStatusName string
	getRows, err := db.Query("SELECT id, status_name FROM status")
	if err != nil {
		return 0, targetStatusName
	}
	if (model.IsInRoom(status)){
		targetStatusName = model.OUT_ROOM
	} else {
		targetStatusName = model.IN_ROOM
	}
	for getRows.Next() {
		err := getRows.Scan(&getStatusId, &getStatusName)
		if err != nil {
			return 0, targetStatusName
		}
		if getStatusName == targetStatusName{
			return getStatusId, targetStatusName 
		}
	}
	return 0, targetStatusName
}

func PutStatus(user schema.Status) (schema.Status, error) {
	var applyStatusId int32
	var targetStatusName string
	var returnUser schema.Status
	applyStatusId, targetStatusName = SelectApplyStatusId(user.Status)
	_, err := db.Query("UPDATE user SET status_id = ? WHERE id = ?", applyStatusId, user.UserId)
	if err != nil {
		return user, fmt.Errorf("failed to select change status: %v", err)
	}
	returnUser = schema.Status{UserId: user.UserId, Status: targetStatusName}
	return returnUser, nil
}