package services

import (
	"fmt"
	"time"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetAllUsersService() (userInformationList []schema.GetUsersInfo200ResponseInner, err error) {
	userInformationList, err = repositories.GetAllUsersRepository()
	if err != nil {
		return []schema.GetUsersInfo200ResponseInner{}, fmt.Errorf("failed to execute query to get all users info: %v", err)
	}

	return userInformationList, nil
}

func GetUsersByIdService(userId int) (userInformation schema.GetUserById200Response, err error) {
	now := time.Now()
	date := now.Format("2006-01")

	userInformation, err = repositories.GetUsersRepository(userId, date)
	if err != nil {
		return schema.GetUserById200Response{}, fmt.Errorf("failed to execute query to get staytime and attendance days: %v", err)
	}

	return userInformation, nil
}

func PutUsersByIdService(userId int, userInformation schema.PutUserByIdRequest) (err error) {
	err = repositories.PutUsersRepository(userId, userInformation)
	if err != nil {
		return fmt.Errorf("failed to execute query to put user info: %v", err)
	}

	return nil
}
