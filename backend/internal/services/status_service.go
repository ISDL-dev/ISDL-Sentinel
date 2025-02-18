package services

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func PutStatusService(status schema.Status) (user schema.Status, err error) {
	var placeId int32

	placeId, err = repositories.GetPlaceId()
	if err != nil {
		return schema.Status{}, fmt.Errorf("failed to get place id: %v", err)
	}

	err = repositories.PutStatusRepository(status, placeId)
	if err != nil {
		return schema.Status{}, fmt.Errorf("failed to change user status: %v", err)
	}

	return status, nil
}
