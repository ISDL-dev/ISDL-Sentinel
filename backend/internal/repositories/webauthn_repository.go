package repositories

import (
	"fmt"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/go-webauthn/webauthn/webauthn"
)

func GetUserCredential(userName string) (userCredential model.UserCredential, err error) {
	getUserCredentialQuery := `
		SELECT id, auth_user_name, COALESCE(display_name, auth_user_name)
		FROM user
		WHERE auth_user_name = ?
	`

	if err := infrastructures.DB.QueryRow(getUserCredentialQuery, userName).Scan(&userCredential.Id, &userCredential.Name, &userCredential.DisplayName); err != nil {
		return model.UserCredential{}, fmt.Errorf("failed to execute a query to get user credential: %v", err)
	}

	getUserCredentialsQuery := `
		SELECT
			c.credential_id,
			c.public_key,
			c.attestation_type,
			cf.user_present,
			cf.user_verified,
			cf.backup_eligible,
			cf.backup_state,
			a.aaguid,
			a.sign_count,
			a.cloneWarning,
			a.Attachment
		FROM credential c
		LEFT JOIN credential_flags cf ON c.flags_id = cf.id
		LEFT JOIN credential_authenticator a ON c.authenticator_id = a.id
		WHERE c.user_id = ?
	`
	rows, err := infrastructures.DB.Query(getUserCredentialsQuery, userCredential.Id)
	if err != nil {
		return model.UserCredential{}, fmt.Errorf("failed to execute a query to get credentials: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var credential webauthn.Credential
		err := rows.Scan(
			&credential.ID,
			&credential.PublicKey,
			&credential.AttestationType,
			&credential.Flags.UserPresent,
			&credential.Flags.UserVerified,
			&credential.Flags.BackupEligible,
			&credential.Flags.BackupState,
			&credential.Authenticator.AAGUID,
			&credential.Authenticator.SignCount,
			&credential.Authenticator.CloneWarning,
			&credential.Authenticator.Attachment,
		)
		if err != nil {
			return model.UserCredential{}, fmt.Errorf("failed to scan credentials: %v", err)
		}

		userCredential.Credentials = append(userCredential.Credentials, credential)
	}

	if err = rows.Err(); err != nil {
		return model.UserCredential{}, fmt.Errorf("failed to iterate over credentials: %v", err)
	}

	return userCredential, nil
}

func UpdateUserCredential(userCredential model.UserCredential) error {
	insertCredentialQuery := `
		INSERT INTO credential (user_id, credential_id, public_key, attestation_type, flags_id, authenticator_id)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	for _, credential := range userCredential.Credentials {
		// Insert credential flags data
		flagsID, err := insertCredentialFlags(credential.Flags)
		if err != nil {
			return fmt.Errorf("failed to insert credential flags: %v", err)
		}

		// Insert authenticator data
		authenticatorID, err := insertAuthenticator(credential.Authenticator)
		if err != nil {
			return fmt.Errorf("failed to insert authenticator: %v", err)
		}

		// Insert credential data
		_, err = infrastructures.DB.Exec(insertCredentialQuery,
			userCredential.Id,
			credential.ID,
			credential.PublicKey,
			credential.AttestationType,
			flagsID,
			authenticatorID,
		)
		if err != nil {
			return fmt.Errorf("failed to execute a query to insert credential: %v", err)
		}
	}

	return nil
}

func insertCredentialFlags(flags webauthn.CredentialFlags) (int64, error) {
	insertFlagsQuery := `
		INSERT INTO credential_flags (user_present, user_verified, backup_eligible, backup_state)
		VALUES (?, ?, ?, ?)
	`
	result, err := infrastructures.DB.Exec(
		insertFlagsQuery,
		flags.UserPresent,
		flags.UserVerified,
		flags.BackupEligible,
		flags.BackupState,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to insert credential flags: %v", err)
	}

	flagsID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID for credential flags: %v", err)
	}

	return flagsID, nil
}

func insertAuthenticator(authenticator webauthn.Authenticator) (int64, error) {
	insertAuthenticatorQuery := `
		INSERT INTO credential_authenticator (aaguid, sign_count, cloneWarning, Attachment)
		VALUES (?, ?, ?, ?)
	`
	result, err := infrastructures.DB.Exec(
		insertAuthenticatorQuery,
		authenticator.AAGUID,
		authenticator.SignCount,
		authenticator.CloneWarning,
		authenticator.Attachment,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to insert authenticator: %v", err)
	}

	authenticatorID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID for authenticator: %v", err)
	}

	return authenticatorID, nil
}
