package repositories

import (
	"database/sql"
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func ChangePasswordRepository(user schema.PutChangePasswordRequest) error {
    getUserCredentialQuery := `
        SELECT auth_user_name, mail_address, password
        FROM user
        WHERE auth_user_name = ? OR mail_address = ?;`

    row := infrastructures.DB.QueryRow(getUserCredentialQuery, user.AuthUserName,user.AuthUserName)
	
	var authUserName string
	var mailAddress string
	var password string
	
    err := row.Scan(
        &authUserName,
		&mailAddress,
		&password,
	)
    if err != nil {
        if err == sql.ErrNoRows {
            return fmt.Errorf("user not found: %w", err)
        }
        return fmt.Errorf("failed to get user credential: %w", err)
    }

	if !((authUserName == user.AuthUserName || mailAddress == user.AuthUserName) && password == user.BeforePassword) {
		return fmt.Errorf("worng username or password")
	}

	UpdateUserPasswordQuery := `update user set password = ? where auth_user_name = ?;`

	_ , err = infrastructures.DB.Exec(UpdateUserPasswordQuery, user.AfterPassword, authUserName)
	if err != nil {
		return fmt.Errorf("failed to change password: %w", err)
	}

	return nil
}