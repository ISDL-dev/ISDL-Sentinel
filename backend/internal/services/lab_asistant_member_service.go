package services

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetLabAsistantMemberService() (labAsistantMemberList []schema.GetLabAsistantMember200ResponseInner, err error) {
	labAsistantMemberList, err = repositories.GetLabAsistantMemberRepository()
	if err != nil {
		return []schema.GetLabAsistantMember200ResponseInner{}, fmt.Errorf("failed to execute query to get lab asistant member: %v", err)
	}

	return labAsistantMemberList, nil
}
