package services

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetLabAssistantMemberService() (labAssistantMemberList []schema.GetLabAssistantMember200ResponseInner, err error) {
	labAssistantMemberList, err = repositories.GetLabAssistantMemberRepository()
	if err != nil {
		return []schema.GetLabAssistantMember200ResponseInner{}, fmt.Errorf("failed to execute query to get lab assistant member: %v", err)
	}

	return labAssistantMemberList, nil
}
