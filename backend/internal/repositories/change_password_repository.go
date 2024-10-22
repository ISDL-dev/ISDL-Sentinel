package repositories

import (
    "database/sql"
    "fmt"

    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func ChangePasswordRepository(user schema.PutChangePasswordRequest, userID int) error {
    getUserCredentialQuery := `
        SELECT id, password
        FROM user
        WHERE id = ?;`

    row := infrastructures.DB.QueryRow(getUserCredentialQuery, userID)
    
    var userId int
    var password string
    
    err := row.Scan(
        &userId,
        &password,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            return fmt.Errorf("user not found: %w", err)
        }
        return fmt.Errorf("failed to get user credential: %w", err)
    }

    if !(id == userID && password == user.BeforePassword) {
        return fmt.Errorf("worng username or password")
    }

    UpdateUserPasswordQuery := `update user set password = ? where id = ?;`

    _ , err = infrastructures.DB.Exec(UpdateUserPasswordQuery, user.AfterPassword, id)
    if err != nil {
        return fmt.Errorf("failed to change password: %w", err)
    }

    return nil
}