package services

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetAttendeesListService() (attendeeList []schema.GetAttendeesList200ResponseInner, err error) {
	inRoomStatusId, err := repositories.GetInRoomStatusIdRepository()
	if err != nil {
		return []schema.GetAttendeesList200ResponseInner{}, fmt.Errorf("failed to execute query to get in room status id: %v", err)
	}
	infrastructures.GetCalendarList()
	attendeeList, err = repositories.GetInRoomUserListRepository(inRoomStatusId)
	if err != nil {
		return []schema.GetAttendeesList200ResponseInner{}, fmt.Errorf("failed to execute query to get in room user list: %v", err)
	}

	return attendeeList, nil
}
