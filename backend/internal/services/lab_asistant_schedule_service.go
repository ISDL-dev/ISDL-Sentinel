package services

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetLabAssistantScheduleService(month string) (labAssistantSchedule []schema.GetLabAssistantSchedule200ResponseInner, err error) {
	labAssistantSchedule, err = repositories.GetLabAssistantScheduleRepository(month)
	if err != nil {
		return []schema.GetLabAssistantSchedule200ResponseInner{}, fmt.Errorf("failed to execute query to get lab assistant schedule: %v", err)
	}

	return labAssistantSchedule, nil
}

func PostLabAssistantScheduleService(month string, labAssistantScheduleRequest []schema.PostLabAssistantScheduleRequestInner) (labAssistantSchedule []schema.GetLabAssistantSchedule200ResponseInner, err error) {
	err = repositories.PostLabAssistantScheduleRepository(month, labAssistantScheduleRequest)
	if err != nil {
		return []schema.GetLabAssistantSchedule200ResponseInner{}, fmt.Errorf("failed to execute query to post lab assistant schedule: %v", err)
	}

	labAssistantSchedule, err = repositories.GetLabAssistantScheduleRepository(month)
	if err != nil {
		return []schema.GetLabAssistantSchedule200ResponseInner{}, fmt.Errorf("failed to execute query to get lab assistant schedule: %v", err)
	}

	return labAssistantSchedule, nil
}
