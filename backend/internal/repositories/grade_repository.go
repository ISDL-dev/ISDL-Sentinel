package repositories

import (
	"fmt"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
)

func GetGradeRepository() (gradeList []string, err error) {
	var grade string

	getRows, err := infrastructures.DB.Query("SELECT grade_name FROM grade;")
	if err != nil {
		return []string{}, fmt.Errorf("getRows GetGradeName Query error err:%w", err)
	}
	for getRows.Next() {
		err := getRows.Scan(&grade)
		if err != nil {
			return []string{}, fmt.Errorf("failed to find target grade name: %v", err)
		}
		gradeList = append(gradeList, grade)
	}
	return gradeList, nil
}

func MoveUpGradeRepository() (err error) {
	tx, err := infrastructures.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Update the user grade based on current grade_name using model constants
	updateGradeQuery := `
		UPDATE user u
		JOIN grade g ON u.grade_id = g.id
		SET u.grade_id = (
			SELECT g2.id 
			FROM grade g2 
			WHERE 
				(g.grade_name = ? AND g2.grade_name = ?) OR
				(g.grade_name = ? AND g2.grade_name = ?) OR
				(g.grade_name = ? AND g2.grade_name = ?)
		)
		WHERE g.grade_name IN (?, ?, ?);
	`

	_, err = tx.Exec(updateGradeQuery,
		model.U4, model.M1, // U4 -> M1
		model.M1, model.M2, // M1 -> M2
		model.M2, model.OB, // M2 -> OB
		model.U4, model.M1, model.M2)
	if err != nil {
		return fmt.Errorf("failed to update user grade: %v", err)
	}

	return nil
}
