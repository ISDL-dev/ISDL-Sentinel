package repositories

import (
	"database/sql"
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func SignUpRepository(user schema.PostUserInformationRequest) error {
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

	selectUserGradeIDQuery := `SELECT id FROM grade WHERE grade_name = ?;`
	row = infrastructures.DB.QueryRow(selectUserGradeIDQuery, user.GradeName)

	var gradeID string
	err = row.Scan(&gradeID)
	if err != nil {
		return fmt.Errorf("grade_name not found")
	}

	insertUserInformationQuery := `
		INSERT INTO user (name, auth_user_name, mail_address, password, number_of_coin, current_entered_at, status_id, place_id, grade_id, avatar_id)
		VALUES (?, ?, ?, ?, 0, NULL, 2, NULL, ?, 1);
	`

	result, err := infrastructures.DB.Exec(insertUserInformationQuery, user.Name, user.AuthUserName, user.MailAddress, user.Password, gradeID)
	if err != nil {
		return fmt.Errorf("failed to register user information: %w", err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve user_id: %w", err)
	}

	insertAvatarQuery := `INSERT INTO user_possession_avatar (user_id, avatar_id) VALUES (?, 1);`
	_, err = infrastructures.DB.Exec(insertAvatarQuery, userID)
	if err != nil {
		return fmt.Errorf("failed to register default avatar: %w", err)
	}

	return nil
}
