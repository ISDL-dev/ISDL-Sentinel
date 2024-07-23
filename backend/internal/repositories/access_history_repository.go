package repositories

import (
	"fmt"
	"strings"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetAccessHistoryRepository(date string) (accessHistory []schema.GetAccessHistory200ResponseInner, err error) {
	var accessHistoryInner schema.GetAccessHistory200ResponseInner
	var firstEntering schema.GetAccessHistory200ResponseInnerEntering
	var firstEnteringList []schema.GetAccessHistory200ResponseInnerEntering
	var lastLeaving schema.GetAccessHistory200ResponseInnerLeaving
	var lastLeavingList []schema.GetAccessHistory200ResponseInnerLeaving

	getFirstEnteringHistoryQuery := `
		SELECT 
			eh.user_id,
			u.name AS user_name,
			u.avatar_id,
			a.img_path AS avatar_img_path,
			eh.entered_at
		FROM 
			entering_history eh
		JOIN 
			user u ON eh.user_id = u.id
		JOIN 
			avatar a ON u.avatar_id = a.id
		WHERE 
			eh.is_first_entering = true
			AND DATE_FORMAT(eh.entered_at, '%Y-%m') = ?;`
	getFirstEnteringRows, err := infrastructures.DB.Query(getFirstEnteringHistoryQuery, date)
	if err != nil {
		return []schema.GetAccessHistory200ResponseInner{}, fmt.Errorf("failed to execute query to get first entering history:%w", err)
	}
	for getFirstEnteringRows.Next() {
		err := getFirstEnteringRows.Scan(
			&firstEntering.UserId,
			&firstEntering.UserName,
			&firstEntering.AvatarId,
			&firstEntering.AvatarImgPath,
			&firstEntering.EnteredAt)
		if err != nil {
			return []schema.GetAccessHistory200ResponseInner{}, fmt.Errorf("failed to execute query to get first entering history: %v", err)
		}
		firstEnteringList = append(firstEnteringList, firstEntering)
	}

	getLastLeavingHistoryQuery := `
		SELECT 
			lh.user_id,
			u.name AS user_name,
			u.avatar_id,
			a.img_path AS avatar_img_path,
			lh.left_at
		FROM 
			leaving_history lh
		JOIN 
			user u ON lh.user_id = u.id
		JOIN 
			avatar a ON u.avatar_id = a.id
		WHERE 
			lh.is_last_leaving = true
			AND DATE_FORMAT(lh.left_at, '%Y-%m') = ?;`
	getLastLeavingRows, err := infrastructures.DB.Query(getLastLeavingHistoryQuery, date)
	if err != nil {
		return []schema.GetAccessHistory200ResponseInner{}, fmt.Errorf("failed to execute query to get last leaving history:%w", err)
	}
	for getLastLeavingRows.Next() {
		err := getLastLeavingRows.Scan(
			&lastLeaving.UserId,
			&lastLeaving.UserName,
			&lastLeaving.AvatarId,
			&lastLeaving.AvatarImgPath,
			&lastLeaving.LeftAt)
		if err != nil {
			return []schema.GetAccessHistory200ResponseInner{}, fmt.Errorf("failed to execute query to get last leaving history: %v", err)
		}
		lastLeavingList = append(lastLeavingList, lastLeaving)
	}

	for i := 0; i < len(firstEnteringList); i++ {
		dateParts := strings.Split(firstEnteringList[i].EnteredAt, "T")
		if len(dateParts) > 0 {
			accessHistoryInner.Date = dateParts[0]
		} else {
			accessHistoryInner.Date = ""
		}
		accessHistoryInner.Entering = firstEnteringList[i]
		if i < len(lastLeavingList) {
			accessHistoryInner.Leaving = lastLeavingList[i]
		} else {
			accessHistoryInner.Leaving = schema.GetAccessHistory200ResponseInnerLeaving{}
		}
		accessHistory = append(accessHistory, accessHistoryInner)
	}

	return accessHistory, nil
}
