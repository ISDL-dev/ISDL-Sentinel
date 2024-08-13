package services

import (
	"fmt"
	"time"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetUsersByIdService(userId int) (userInformation schema.GetUserById200Response, err error) {
	now := time.Now()
	date := now.Format("2006-01")

	userInformation, err = repositories.GetUsersRepository(userId, date)
	if err != nil {
		return schema.GetUserById200Response{}, fmt.Errorf("failed to execute query to get staytime and attendance days: %v", err)
	}

	return userInformation, nil
}