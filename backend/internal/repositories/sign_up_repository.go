package repositories

import (
	"database/sql"
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func SignUpRepository(user schema.PostUserSignUpRequest) error {
    getUserCredentialQuery := `
        SELECT name
        FROM user
        WHERE name = ? OR mail_address = ? OR auth_user_name = ?;
    `

    row := infrastructures.DB.QueryRow(getUserCredentialQuery, user.Name, user.MailAddress, user.AuthUserName)
	
	var existingName string
	err := row.Scan(&existingName)

    if err == nil {
        return fmt.Errorf("user already exists")
    } else if err != sql.ErrNoRows {
		return fmt.Errorf("error querying user: %v", err)
	} 

	InsertUserInformationQuery := `
		INSERT INTO user (name, auth_user_name, mail_address, password, number_of_coin, current_entered_at, status_id, place_id, grade_id, avatar_id)
		VALUES (?, ?, ?, ?, 0, NULL, 2, NULL, ?, 1);
	`

	_ , err = infrastructures.DB.Exec(InsertUserInformationQuery, user.Name, user.AuthUserName, user.MailAddress, user.Password, user.GradeID)

	if err != nil {
		return fmt.Errorf("failed to register user information: %w", err)
	}

	return nil
}