package services

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetRankingService() (rankingList []schema.GetRanking200ResponseInner, err error) {
	rankingList, err = repositories.GetRankingRepository()
	if err != nil {
		return []schema.GetRanking200ResponseInner{}, fmt.Errorf("failed to execute query to get ranking list: %v", err)
	}

	return rankingList, nil
}
