package services

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
)

func GetGradeService() (gradeList []string, err error) {
	gradeList, err = repositories.GetGradeRepository()
	if err != nil {
		return []string{}, fmt.Errorf("failed to execute query to get Grade list: %v", err)
	}

	return gradeList, nil
}
