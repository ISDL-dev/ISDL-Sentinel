package services

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
)

func GetRoleService() (roleList []string, err error) {
	roleList, err = repositories.GetRoleRepository()
	if err != nil {
		return []string{}, fmt.Errorf("failed to execute query to get role list: %v", err)
	}

	return roleList, nil
}
