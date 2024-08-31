package repositories

import (
	"fmt"
	"strings"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetLabAssistantMemberRepository() (labAssistantMemberList []schema.GetLabAssistantMember200ResponseInner, err error) {
	var gradeId int32
	var labAssistantMember schema.GetLabAssistantMember200ResponseInner

	getGradeIdQuery := `SELECT id FROM grade WHERE grade_name = ?;`
	if err := infrastructures.DB.QueryRow(getGradeIdQuery, model.U4).Scan(&gradeId); err != nil {
		return []schema.GetLabAssistantMember200ResponseInner{}, fmt.Errorf("failed to execute a query to get grade_id: %v", err)
	}

	getLabAssistantMemberQuery := `
		SELECT 
			u.id, 
			u.name, 
			u.avatar_id, 
			a.img_path, 
			IFNULL((SELECT MAX(las.shift_day) FROM lab_assistant_shift las WHERE las.user_id = u.id), "") AS last_shift_date,
			(SELECT COUNT(*) FROM lab_assistant_shift las WHERE las.user_id = u.id) AS count
		FROM 
			user u
		LEFT JOIN 
			avatar a ON u.avatar_id = a.id
		WHERE 
			u.grade_id = ?;
	`
	getRows, err := infrastructures.DB.Query(getLabAssistantMemberQuery, gradeId)
	if err != nil {
		return []schema.GetLabAssistantMember200ResponseInner{}, fmt.Errorf("getRows getLabAssistantMember Query error err:%w", err)
	}
	defer getRows.Close()

	for getRows.Next() {
		var lastShiftDate string
		err := getRows.Scan(
			&labAssistantMember.UserId,
			&labAssistantMember.UserName,
			&labAssistantMember.AvatarId,
			&labAssistantMember.AvatarImgPath,
			&lastShiftDate,
			&labAssistantMember.Count,
		)
		if err != nil {
			return []schema.GetLabAssistantMember200ResponseInner{}, fmt.Errorf("getRows getLabAssistantMember Query error err: %v", err)
		}

		if lastShiftDate != "" {
			labAssistantMember.LastShiftDate = strings.Split(lastShiftDate, "T")[0]
		}
		labAssistantMemberList = append(labAssistantMemberList, labAssistantMember)
	}

	if err := getRows.Err(); err != nil {
		return []schema.GetLabAssistantMember200ResponseInner{}, fmt.Errorf("error occurred during iteration: %v", err)
	}

	return labAssistantMemberList, nil
}
