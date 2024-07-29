package repositories

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetLabAsistantMemberRepository() (labAsistantMemberList []schema.GetLabAsistantMember200ResponseInner, err error) {
	var gradeId int32
	var labAsistantMember schema.GetLabAsistantMember200ResponseInner

	getGradeIdQuery := `SELECT id FROM grade WHERE grade_name = ?;`
	if err := infrastructures.DB.QueryRow(getGradeIdQuery, model.U4).Scan(&gradeId); err != nil {
		return []schema.GetLabAsistantMember200ResponseInner{}, fmt.Errorf("failed to execute a query to get avatar_id: %v", err)
	}

	getLabAsistantMemberQuery := `SELECT id, name FROM user WHERE grade_id = ?;`
	getRows, err := infrastructures.DB.Query(getLabAsistantMemberQuery, gradeId)
	if err != nil {
		return []schema.GetLabAsistantMember200ResponseInner{}, fmt.Errorf("getRows getInRoomUserList Query error err:%w", err)
	}
	for getRows.Next() {
		err := getRows.Scan(
			&labAsistantMember.UserId,
			&labAsistantMember.UserName,
		)
		if err != nil {
			return []schema.GetLabAsistantMember200ResponseInner{}, fmt.Errorf("failed to find target status id: %v", err)
		}
		labAsistantMemberList = append(labAsistantMemberList, labAsistantMember)
	}

	return labAsistantMemberList, nil
}
