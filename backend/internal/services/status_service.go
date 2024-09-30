package services

import (
	"fmt"

	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func PutStatusService(status schema.Status) (user schema.Status, err error) {
	var placeId int32

	if status.Status == model.OUT_ROOM {
		placeId = 0
	} else {
		placeId, err = repositories.GetPlaceId()
		if err != nil {
			return schema.Status{}, fmt.Errorf("failed to get place id: %v", err)
		}
	}

	err = repositories.PutStatusRepository(status, placeId)
	if err != nil {
		return schema.Status{}, fmt.Errorf("failed to change user status: %v", err)
	}

	return status, nil
}
