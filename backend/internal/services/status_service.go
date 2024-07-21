package services

import (
	"fmt"

	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func SelectApplyStatusId(status string) (id int32, statusName string, err error) {
	var getStatusId int32 
	var getStatusName string
	var targetStatusName string
	getRows, err := repositories.GetAllStatusRepository()
	if err != nil {
		return 0, targetStatusName, fmt.Errorf("failed to get status list: %v", err)
	}
	if (model.IsInRoom(status)){
		targetStatusName = model.OUT_ROOM
	} else {
		targetStatusName = model.IN_ROOM
	}
	for getRows.Next() {
		err := getRows.Scan(&getStatusId, &getStatusName)
		if err != nil {
			return 0, targetStatusName, fmt.Errorf("failed to find target status id: %v", err)
		}
		if getStatusName == targetStatusName{
			return getStatusId, targetStatusName, nil
		}
	}
	return 0, targetStatusName, nil
}

func SelectApplyPlaceId(targetStatusName string) (applyPlaceId int32, err error){
	if targetStatusName == model.IN_ROOM {
		applyPlaceId, err = repositories.GetInRoomPlaceIdRepository() 
		if err != nil {
			return 0, fmt.Errorf("failed to execute query to get in room place id: %v", err)
		}
	} else {
		applyPlaceId = 0
	}
	return applyPlaceId, nil
}

func PutStatusService(status schema.Status) (user schema.Status, err error) {
	var applyStatusId int32
	var applyPlaceId int32
	var targetStatusName string
	applyStatusId, targetStatusName, err = SelectApplyStatusId(status.Status)
	if err != nil {
		return schema.Status{}, fmt.Errorf("failed to find target status id: %v", err)
	}
	applyPlaceId, err = SelectApplyPlaceId(targetStatusName)
	if err != nil {
		return schema.Status{}, fmt.Errorf("failed to find target place id: %v", err)
	}
	err = repositories.PutStatusRepository(status.UserId, applyStatusId, applyPlaceId)
	if err != nil {
		return schema.Status{}, fmt.Errorf("failed to change user status: %v", err)
	}

	return schema.Status{UserId: status.UserId, Status: targetStatusName}, nil
}