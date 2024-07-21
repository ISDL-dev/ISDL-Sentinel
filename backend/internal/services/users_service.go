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

func GetAuthUserInfoService() (authUsers map[string]schema.PostUserRequest, error) {
	authUsers := make(map[string]schema.PostUserRequest)
	
	rows, err = repositories.GetAllUserRepository()
	if err != nil {
		return schema.PostUserInformationRequest{}, fmt.Errorf("failed to execute query to get auth user information: %v", err)
	}

	for rows.Next() {
        var mailAddress, password, name string
        err := rows.Scan(&mailAddress, &password, &name)
        if err != nil {
            return nil, err
        }

        authUsers[mailAddress] = schema.PostUserRequest{
            MailAddress: mailAddress,
            Password:    password,
            Name:        name,
        }
    }

	return authUsers, nil
}
