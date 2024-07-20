package services

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetAttendeesListService() (userList schema.GetAttendeesList200ResponseInner, err error) {
	// in room のStatusIdを取ってくる
	inRoomUserList, err := repositories.GetInRoomUserListRepository(1)
	if err != nil {
		return schema.GetAttendeesList200ResponseInner{}, fmt.Errorf("failed to execute query to get in room user list: %v", err)
	}
	

	PURPOSE := "purpose"
	return userInformation, nil
}