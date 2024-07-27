package repositories

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"github.com/go-webauthn/webauthn/webauthn"
)

// GetSession retrieves the session data from the database using the session ID.
func GetSession(sessionID string) (*webauthn.SessionData, error) {
	getSessionQuery := `
		SELECT session_data FROM session WHERE session_id = ?;
	`

	var sessionDataJSON string
	err := infrastructures.DB.QueryRow(getSessionQuery, sessionID).Scan(&sessionDataJSON)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error getting session '%s': does not exist", sessionID)
		}
		return nil, fmt.Errorf("failed to execute query to get session: %v", err)
	}

	var sessionData webauthn.SessionData
	err = json.Unmarshal([]byte(sessionDataJSON), &sessionData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal session data: %v", err)
	}

	return &sessionData, nil
}

// DeleteSession deletes the session data from the database using the session ID.
func DeleteSession(sessionID string) error {
	deleteSessionQuery := `
		DELETE FROM session WHERE session_id = ?;
	`

	_, err := infrastructures.DB.Exec(deleteSessionQuery, sessionID)
	if err != nil {
		return fmt.Errorf("failed to execute query to delete session: %v", err)
	}

	return nil
}

// StartSession inserts the session data into the database and returns the session ID.
func StartSession(data *webauthn.SessionData) (string, error) {
	sessionID, err := random(32)
	if err != nil {
		return "", fmt.Errorf("failed to generate random session ID: %v", err)
	}

	sessionDataJSON, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal session data: %v", err)
	}

	insertSessionQuery := `
		INSERT INTO session (session_id, session_data) VALUES (?, ?);
	`

	_, err = infrastructures.DB.Exec(insertSessionQuery, sessionID, sessionDataJSON)
	if err != nil {
		return "", fmt.Errorf("failed to execute query to insert session: %v", err)
	}

	return sessionID, nil
}

// random generates a random string of the given length.
func random(length int) (string, error) {
	randomData := make([]byte, length)
	_, err := rand.Read(randomData)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(randomData), nil
}
