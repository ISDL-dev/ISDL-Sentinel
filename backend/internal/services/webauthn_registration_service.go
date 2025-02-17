package services

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

var (
	Wc  *webauthn.WebAuthn
	err error
)

func GetBeginRegistrationService(userName string, w http.ResponseWriter, r *http.Request) (*protocol.CredentialCreation, error) {
	serverHost := os.Getenv("SERVER_HOST")
	envType := os.Getenv("ENV_TYPE")

	if serverHost == "" {
		return nil, fmt.Errorf("SERVER_HOST environment variable is not set")
	}

	rpID, rpOrigin := serverHost, "http://"+serverHost
	if envType == "prod" {
		rpID = "www.isdl-sentinel.com"
		rpOrigin = "https://" + rpID
	}

	Wc, err = webauthn.New(&webauthn.Config{
		RPDisplayName: "ISDL-Sentinel",
		RPID:          rpID,
		RPOrigin:      rpOrigin,
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

	cookie := &http.Cookie{
		Name:     "registration",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
	}

	// ローカル環境では Secure フラグを設定しない
	if !strings.Contains(r.Host, "localhost") && !strings.Contains(r.Host, "127.0.0.1") {
		cookie.Secure = true
	}

	http.SetCookie(w, cookie)

	log.Printf("Set registration cookie: %v", sessionID)

	return options, nil
}

func GetFinishRegistrationService(userName string, w http.ResponseWriter, r *http.Request) (loginUserInfo schema.PostSignIn200Response, err error) {
	userCredential, err := repositories.GetUserCredential(userName)
	if err != nil {
		return schema.PostSignIn200Response{}, fmt.Errorf("failed to get user credential for %s: %w", userName, err)
	}

	cookie, err := r.Cookie("registration")
	if err != nil {
		return schema.PostSignIn200Response{}, fmt.Errorf("failed to get registration cookie: %w", err)
	}

	sessionData, err := repositories.GetSession(cookie.Value)
	if err != nil {
		return schema.PostSignIn200Response{}, fmt.Errorf("failed to get session for cookie %s: %w", cookie.Value, err)
	}

	credential, err := Wc.FinishRegistration(userCredential, *sessionData, r)
	if err != nil {
		return schema.PostSignIn200Response{}, fmt.Errorf("failed to finish registration: %w", err)
	}

	userCredential.Credentials = append(userCredential.Credentials, *credential)
	err = repositories.UpdateUserCredential(userCredential)
	if err != nil {
		return schema.PostSignIn200Response{}, fmt.Errorf("failed to update user credential for %s: %w", userName, err)
	}

	err = repositories.DeleteSession(cookie.Value)
	if err != nil {
		return schema.PostSignIn200Response{}, fmt.Errorf("failed to delete session for cookie %s: %w", cookie.Value, err)
	}
	log.Printf("User %s finished registration successfully", userName)

	loginUserInfo, err = repositories.GetLoginUserInfo(userName)
	if err != nil {
		return schema.PostSignIn200Response{}, fmt.Errorf("failed to get login user info %s: %w", cookie.Value, err)
	}

	return loginUserInfo, nil
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
