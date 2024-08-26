package repositories

import (
	"database/sql"
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

// GetUserCredential retrieves user credentials based on the provided name
func GetDigestCredential(name string) (userInfo schema.PostUserInformationRequest, err error) {

    getUserCredentialQuery := `
        SELECT 
            auth_user_name,
            mail_address,
            password
        FROM user
        WHERE mail_address = ? OR auth_user_name = ?;
    `

    row := infrastructures.DB.QueryRow(getUserCredentialQuery, name, name)

    err = row.Scan(
        &userInfo.Name,
        &userInfo.MailAddress,
        &userInfo.Password,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            return schema.PostUserInformationRequest{}, fmt.Errorf("user not found: %w", err)
        }
        return schema.PostUserInformationRequest{}, fmt.Errorf("failed to get user credential: %w", err)
    }

    return userInfo, nil
}