package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/go-webauthn/webauthn/protocol"
)

func GetBeginLoginService(userName string, w http.ResponseWriter, r *http.Request) (*protocol.CredentialAssertion, error) {
	userCredential, err := repositories.GetUserCredential(userName)
	if err != nil {
		return nil, fmt.Errorf("failed to get user credential for %s: %w", userName, err)
	}

	options, sessionData, err := Wc.BeginLogin(userCredential)
	if err != nil {
		return nil, fmt.Errorf("failed to begin login: %w", err)
	}

	sessionID, err := repositories.StartSession(sessionData)
	if err != nil {
		return nil, fmt.Errorf("failed to start session: %w", err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "authentication",
		Value: sessionID,
		Path:  "/",
	})

	return options, nil
}

func GetFinishLoginService(userName string, w http.ResponseWriter, r *http.Request) error {
	userCredential, err := repositories.GetUserCredential(userName)
	if err != nil {
		return fmt.Errorf("failed to get user credential for %s: %w", userName, err)
	}

	cookie, err := r.Cookie("authentication")
	if err != nil {
		return fmt.Errorf("failed to get authentication cookie: %w", err)
	}

	sessionData, err := repositories.GetSession(cookie.Value)
	if err != nil {
		return fmt.Errorf("failed to get session for cookie %s: %w", cookie.Value, err)
	}

	credential, err := Wc.FinishLogin(userCredential, *sessionData, r)
	if err != nil {
		return fmt.Errorf("failed to finish login: %w", err)
	}

	if credential.Authenticator.CloneWarning {
		return fmt.Errorf("cloned key detected")
	}

	err = repositories.DeleteSession(cookie.Value)
	if err != nil {
		return fmt.Errorf("failed to delete session for cookie %s: %w", cookie.Value, err)
	}
	log.Printf("User %s finished login successfully", userName)

	return nil
}
