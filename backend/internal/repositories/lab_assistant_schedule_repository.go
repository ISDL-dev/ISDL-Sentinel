package repositories

import (
	"fmt"
	"time"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetLabAssistantScheduleRepository(month string) (labAssistantSchedule []schema.GetLabAssistantSchedule200ResponseInner, err error) {
	var labAssistantScheduleInner schema.GetLabAssistantSchedule200ResponseInner

	getLabAssistantScheduleQuery := `
		SELECT 
			u.name AS user_name, 
			las.shift_day AS shift_date
		FROM 
			lab_assistant_shift las
		JOIN 
			user u ON las.user_id = u.id
		WHERE 
			DATE_FORMAT(las.shift_day, '%Y-%m') = ?
	`
	getLabAssistantScheduleRows, err := infrastructures.DB.Query(getLabAssistantScheduleQuery, month)
	if err != nil {
		return []schema.GetLabAssistantSchedule200ResponseInner{}, fmt.Errorf("failed to execute query to get lab assistant schedule: %w", err)
	}
	defer getLabAssistantScheduleRows.Close()

	for getLabAssistantScheduleRows.Next() {
		var shiftDate time.Time
		err := getLabAssistantScheduleRows.Scan(
			&labAssistantScheduleInner.UserName,
			&shiftDate)
		if err != nil {
			return []schema.GetLabAssistantSchedule200ResponseInner{}, fmt.Errorf("failed to scan row for lab assistant schedule: %v", err)
		}
		labAssistantScheduleInner.ShiftDate = shiftDate.Format("2006-01-02")
		labAssistantSchedule = append(labAssistantSchedule, labAssistantScheduleInner)
	}

	if err := getLabAssistantScheduleRows.Err(); err != nil {
		return []schema.GetLabAssistantSchedule200ResponseInner{}, fmt.Errorf("error occurred during iteration: %v", err)
	}

	return labAssistantSchedule, nil
}

func PostLabAssistantScheduleRepository(month string, labAssistantScheduleRequest []schema.PostLabAssistantScheduleRequestInner) (err error) {
	tx, err := infrastructures.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	deleteLabAssistantScheduleQuery := `DELETE FROM lab_assistant_shift WHERE DATE_FORMAT(shift_day, '%Y-%m') = ?;`
	_, err = tx.Exec(deleteLabAssistantScheduleQuery, month)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to execute query to delete lab assistant schedule: %v", err)
	}

	postLabAssistantScheduleQuery := `INSERT INTO lab_assistant_shift (user_id, shift_day) VALUES (?, ?);`
	for _, schedule := range labAssistantScheduleRequest {
		_, err = tx.Exec(postLabAssistantScheduleQuery, schedule.UserId, schedule.ShiftDate)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to execute query to insert lab assistant schedule: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
