package repositories

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
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
