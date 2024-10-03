package services

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetAccessHistoryService(month string) (accessHistrory []schema.GetAccessHistory200ResponseInner, err error) {
	accessHistrory, err = repositories.GetAccessHistoryRepository(month)
	if err != nil {
		return []schema.GetAccessHistory200ResponseInner{}, fmt.Errorf("failed to execute query to get access history: %v", err)
	}

	return accessHistrory, nil
}
