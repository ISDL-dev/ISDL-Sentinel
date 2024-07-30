package services

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetLabAsistantScheduleService(month string) (labAsistantSchedule []schema.GetLabAsistantSchedule200ResponseInner, err error) {
	labAsistantSchedule, err = repositories.GetLabAsistantScheduleRepository(month)
	if err != nil {
		return []schema.GetLabAsistantSchedule200ResponseInner{}, fmt.Errorf("failed to execute query to get lab asistant schedule: %v", err)
	}

	return labAsistantSchedule, nil
}

func PostLabAsistantScheduleService(month string, labAsistantScheduleRequest []schema.PostLabAsistantScheduleRequestInner) (labAsistantSchedule []schema.GetLabAsistantSchedule200ResponseInner, err error) {
	err = repositories.PostLabAsistantScheduleRepository(month, labAsistantScheduleRequest)
	if err != nil {
		return []schema.GetLabAsistantSchedule200ResponseInner{}, fmt.Errorf("failed to execute query to post lab asistant schedule: %v", err)
	}

	labAsistantSchedule, err = repositories.GetLabAsistantScheduleRepository(month)
	if err != nil {
		return []schema.GetLabAsistantSchedule200ResponseInner{}, fmt.Errorf("failed to execute query to get lab asistant schedule: %v", err)
	}

	return labAsistantSchedule, nil
}
