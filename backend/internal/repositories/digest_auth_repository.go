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
            name,
            mail_address,
            password
        FROM users
        WHERE mail_address = ?;
    `

    row := infrastructures.DB.QueryRow(getUserCredentialQuery, name)

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

// UpdateUserCredential updates user information in the database
func UpdateDigestCredential(userInfo schema.PostUserInformationRequest) error {
	updateUserCredentialQuery := `
		UPDATE user
		SET 
			mail_address = ?,
			password = ?
		WHERE name = ?;
	`

	result, err := infrastructures.DB.Exec(updateUserCredentialQuery,
		userInfo.MailAddress,
		userInfo.Password,
		userInfo.Name,
	)

	if err != nil {
		return fmt.Errorf("failed to update user credential: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with the provided name")
	}

	return nil
}

// CreateUser creates a new user in the database
func CreateUser(userInfo schema.PostUserInformationRequest) error {
	createUserQuery := `
		INSERT INTO user (name, mail_address, password)
		VALUES (?, ?, ?);
	`

	_, err := infrastructures.DB.Exec(createUserQuery,
		userInfo.Name,
		userInfo.MailAddress,
		userInfo.Password,
	)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}