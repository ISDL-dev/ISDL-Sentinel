package repositories

import (
	"fmt"
	"time"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetLabAsistantScheduleRepository(month string) (labAsistantSchedule []schema.GetLabAsistantSchedule200ResponseInner, err error) {
	var labAsistantScheduleInner schema.GetLabAsistantSchedule200ResponseInner

	getLabAsistantScheduleQuery := `
		SELECT 
			u.name AS user_name, 
			las.shift_day AS shift_date
		FROM 
			lab_asistant_shift las
		JOIN 
			user u ON las.user_id = u.id
		WHERE 
			DATE_FORMAT(las.shift_day, '%Y-%m') = ?
	`
	getLabAsistantScheduleRows, err := infrastructures.DB.Query(getLabAsistantScheduleQuery, month)
	if err != nil {
		return []schema.GetLabAsistantSchedule200ResponseInner{}, fmt.Errorf("failed to execute query to get lab assistant schedule: %w", err)
	}
	defer getLabAsistantScheduleRows.Close()

	for getLabAsistantScheduleRows.Next() {
		var shiftDate time.Time
		err := getLabAsistantScheduleRows.Scan(
			&labAsistantScheduleInner.UserName,
			&shiftDate)
		if err != nil {
			return []schema.GetLabAsistantSchedule200ResponseInner{}, fmt.Errorf("failed to scan row for lab assistant schedule: %v", err)
		}
		labAsistantScheduleInner.ShiftDate = shiftDate.Format("2006-01-02")
		labAsistantSchedule = append(labAsistantSchedule, labAsistantScheduleInner)
	}

	if err := getLabAsistantScheduleRows.Err(); err != nil {
		return []schema.GetLabAsistantSchedule200ResponseInner{}, fmt.Errorf("error occurred during iteration: %v", err)
	}

	return labAsistantSchedule, nil
}
