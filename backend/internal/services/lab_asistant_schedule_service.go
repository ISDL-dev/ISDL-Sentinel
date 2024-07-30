package services

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetLabAsistantScheduleService(month string) (labAsistantSchedule []schema.GetLabAsistantSchedule200ResponseInner, err error) {
	labAsistantSchedule, err = repositories.GetLabAsistantScheduleRepository(month)
	if err != nil {
		return []schema.GetLabAsistantSchedule200ResponseInner{}, fmt.Errorf("failed to execute query to get access history: %v", err)
	}

	return labAsistantSchedule, nil
}
