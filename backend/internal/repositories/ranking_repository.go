package repositories

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetRankingRepository() (rankingList []schema.GetRanking200ResponseInner, err error) {
	var user schema.GetRanking200ResponseInner
	getRankingListQuery := `
	SELECT 
		u.id AS user_id,
		u.name AS user_name,
		g.grade_name,
		u.avatar_id,
		a.img_path,
		COALESCE(t.total_stay_time, '00:00:00') AS total_stay_time,
		COALESCE(e.unique_enter_days, 0) AS unique_enter_days
	FROM 
		user u
	LEFT JOIN 
		grade g ON u.grade_id = g.id
	LEFT JOIN 
		avatar a ON u.avatar_id = a.id
	LEFT JOIN 
		(SELECT 
			user_id, 
			SEC_TO_TIME(SUM(TIME_TO_SEC(stay_time))) AS total_stay_time
		FROM 
			leaving_history
		GROUP BY 
			user_id) t ON u.id = t.user_id
	LEFT JOIN 
		(SELECT 
			user_id, 
			COUNT(DISTINCT DATE(entered_at)) AS unique_enter_days
		FROM 
			entering_history
		GROUP BY 
			user_id) e ON u.id = e.user_id
	ORDER BY 
		u.id;`

	getRows, err := infrastructures.DB.Query(getRankingListQuery)
	if err != nil {
		return []schema.GetRanking200ResponseInner{}, fmt.Errorf("getRows getRankingListQuery Query error err:%w", err)
	}
	for getRows.Next() {
		err := getRows.Scan(
			&user.UserId,
			&user.UserName,
			&user.Grade,
			&user.AvatarId,
			&user.AvatarImgPath,
			&user.StayTime,
			&user.AttendanceDays)
		if err != nil {
			return []schema.GetRanking200ResponseInner{}, fmt.Errorf("failed to find user: %v", err)
		}
		rankingList = append(rankingList, user)
	}
	return rankingList, nil
}
