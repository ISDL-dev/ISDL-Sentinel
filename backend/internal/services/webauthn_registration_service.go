package services

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

var (
	Wc  *webauthn.WebAuthn
	err error
)

func GetBeginRegistrationService(userName string, w http.ResponseWriter, r *http.Request) (*protocol.CredentialCreation, error) {
	rpID := getDynamicRPID(r)
	Wc, err = webauthn.New(&webauthn.Config{
		RPDisplayName: "ISDL-Sentinel",
		RPID:          rpID,
		RPOrigin:      getOrigin(r),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create WebAuthn from config: %w", err)
	}

	userCredential, err := repositories.GetUserCredential(userName)
	if err != nil {
		return nil, fmt.Errorf("failed to get user credential for %s: %w", userName, err)
	}

	registerOptions := func(credCreationOpts *protocol.PublicKeyCredentialCreationOptions) {
		credCreationOpts.CredentialExcludeList = credentialExcludeList(userCredential)
	}

	options, sessionData, err := Wc.BeginRegistration(
		userCredential,
		registerOptions,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to begin registration: %w", err)
	}

	sessionID, err := repositories.StartSession(sessionData)
	if err != nil {
		return nil, fmt.Errorf("failed to start session: %w", err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "registration",
		Value: sessionID,
		Path:  "/",
	})

	return options, nil
}

func GetFinishRegistrationService(userName string, w http.ResponseWriter, r *http.Request) error {
	userCredential, err := repositories.GetUserCredential(userName)
	if err != nil {
		return fmt.Errorf("failed to get user credential for %s: %w", userName, err)
	}

	cookie, err := r.Cookie("registration")
	if err != nil {
		return fmt.Errorf("failed to get registration cookie: %w", err)
	}

	sessionData, err := repositories.GetSession(cookie.Value)
	if err != nil {
		return fmt.Errorf("failed to get session for cookie %s: %w", cookie.Value, err)
	}

	credential, err := Wc.FinishRegistration(userCredential, *sessionData, r)
	if err != nil {
		return fmt.Errorf("failed to finish registration: %w", err)
	}

	userCredential.Credentials = append(userCredential.Credentials, *credential)
	err = repositories.UpdateUserCredential(userCredential)
	if err != nil {
		return fmt.Errorf("failed to update user credential for %s: %w", userName, err)
	}

	err = repositories.DeleteSession(cookie.Value)
	if err != nil {
		return fmt.Errorf("failed to delete session for cookie %s: %w", cookie.Value, err)
	}
	log.Printf("User %s finished registration successfully", userName)

	return nil
}

func getDynamicRPID(r *http.Request) string {
	host := r.Host
	if strings.Contains(host, "localhost") || strings.Contains(host, "127.0.0.1") {
		return "localhost"
	}
	// ngrokのドメインを処理
	if strings.HasSuffix(host, ".ngrok.io") {
		parts := strings.Split(host, ".")
		return strings.Join(parts[len(parts)-3:], ".")
	}
	return host
}

func getOrigin(r *http.Request) string {
	scheme := "https"
	if r.TLS == nil {
		scheme = "http"
	}
	return scheme + "://" + r.Host
}

func credentialExcludeList(userCredential model.UserCredential) []protocol.CredentialDescriptor {
	credentialExcludeList := []protocol.CredentialDescriptor{}
	for _, cred := range userCredential.Credentials {
		descriptor := protocol.CredentialDescriptor{
			Type:         protocol.PublicKeyCredentialType,
			CredentialID: cred.ID,
		}
		credentialExcludeList = append(credentialExcludeList, descriptor)
	}

	return credentialExcludeList
}
