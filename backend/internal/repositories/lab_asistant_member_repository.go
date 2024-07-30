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
		return []schema.GetLabAsistantMember200ResponseInner{}, fmt.Errorf("failed to execute a query to get grade_id: %v", err)
	}

	getLabAsistantMemberQuery := `
		SELECT 
			u.id, 
			u.name, 
			u.avatar_id, 
			a.img_path, 
			(SELECT COUNT(*) FROM lab_asistant_shift las WHERE las.user_id = u.id) AS count
		FROM 
			user u
		LEFT JOIN 
			avatar a ON u.avatar_id = a.id
		WHERE 
			u.grade_id = ?;
	`
	getRows, err := infrastructures.DB.Query(getLabAsistantMemberQuery, gradeId)
	if err != nil {
		return []schema.GetLabAsistantMember200ResponseInner{}, fmt.Errorf("getRows getLabAsistantMember Query error err:%w", err)
	}
	defer getRows.Close()

	for getRows.Next() {
		err := getRows.Scan(
			&labAsistantMember.UserId,
			&labAsistantMember.UserName,
			&labAsistantMember.AvatarId,
			&labAsistantMember.AvatarImgPath,
			&labAsistantMember.Count,
		)
		if err != nil {
			return []schema.GetLabAsistantMember200ResponseInner{}, fmt.Errorf("getRows getLabAsistantMember Query error err: %v", err)
		}
		labAsistantMemberList = append(labAsistantMemberList, labAsistantMember)
	}

	if err := getRows.Err(); err != nil {
		return []schema.GetLabAsistantMember200ResponseInner{}, fmt.Errorf("error occurred during iteration: %v", err)
	}

	return labAsistantMemberList, nil
}
