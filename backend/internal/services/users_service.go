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

func GetAuthUserInfoService() (map[string]schema.PostSignInRequest, error) {
	authUsers := make(map[string]schema.PostSignInRequest)
	
	rows, err := repositories.GetAllAuthInfoRepository()
	if err != nil {
		return authUsers, fmt.Errorf("failed to execute query to get auth user information: %v", err)
	}

	for rows.Next() {
        var mailAddress, password string
        err := rows.Scan(&mailAddress, &password)
        if err != nil {
            return nil, err
        }

        authUsers[mailAddress] = schema.PostSignInRequest{
            MailAddress: mailAddress,
            Password:    password,
        }
    }

	return authUsers, nil
}
