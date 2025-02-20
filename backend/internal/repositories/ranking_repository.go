package repositories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func convertToDateRange(termValue string) (startDate, endDate time.Time) {
	if strings.Contains(termValue, "-") {
		t, err := time.Parse("2006-01", termValue)
		if err != nil {
			year, month, _ := time.Now().Date()
			startDate = time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
			endDate = startDate.AddDate(0, 1, -1).Add(23*time.Hour + 59*time.Minute + 59*time.Second)
		} else {
			year, month, _ := t.Date()
			startDate = time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
			endDate = startDate.AddDate(0, 1, -1).Add(23*time.Hour + 59*time.Minute + 59*time.Second)
		}
	} else {
		year, err := strconv.Atoi(termValue)
		if err != nil {
			year = time.Now().Year()
		}
		startDate = time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
		endDate = time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)
	}
	return
}

func GetRankingRepository(term string) (rankingList []schema.GetRanking200ResponseInner, err error) {
	var user schema.GetRanking200ResponseInner
	startDate, endDate := convertToDateRange(term)
	log.Printf("startDate: %v, endDate: %v", startDate, endDate)
	getRankingListQuery := `
		SELECT 
			u.id AS user_id,
			u.name AS user_name,
			g.grade_name,
			u.avatar_id,
			a.img_path,
			SEC_TO_TIME(SUM(TIME_TO_SEC(IFNULL(l.stay_time, '00:00:00')))) AS total_stay_time,
			COALESCE(COUNT(DISTINCT DATE(e.entered_at)), 0) AS unique_enter_days
		FROM 
			user u
		LEFT JOIN 
			grade g ON u.grade_id = g.id
		LEFT JOIN 
			avatar a ON u.avatar_id = a.id
		LEFT JOIN 
			entering_history e ON u.id = e.user_id AND e.entered_at BETWEEN ? AND ?
		LEFT JOIN 
			leaving_history l ON e.id = l.entering_history_id
		GROUP BY 
			u.id, u.name, g.grade_name, u.avatar_id, a.img_path
		ORDER BY 
			u.id;`
	getRows, err := infrastructures.DB.Query(getRankingListQuery, startDate, endDate)
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
