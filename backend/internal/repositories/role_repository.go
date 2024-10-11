package repositories

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
)

func GetRoleRepository() (roleList []string, err error) {
	var role string

	getRows, err := infrastructures.DB.Query("SELECT role_name FROM role;")
	if err != nil {
		return []string{}, fmt.Errorf("getRows GetRoleName Query error err:%w", err)
	}
	for getRows.Next() {
		err := getRows.Scan(&role)
		if err != nil {
			return []string{}, fmt.Errorf("failed to find target role name: %v", err)
		}
		roleList = append(roleList, role)
	}
	return roleList, nil
}
