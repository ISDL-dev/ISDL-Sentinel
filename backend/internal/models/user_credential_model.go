package model

import (
	"github.com/go-webauthn/webauthn/webauthn"
)

type UserCredential struct {
	Id          uint64
	Name        string
	DisplayName string
	Credentials []webauthn.Credential
}

func (u UserCredential) WebAuthnID() []byte {
	return []byte(u.Name) // IDとしてNameを使用します
}

func (u UserCredential) WebAuthnName() string {
	return u.Name
}

func (u UserCredential) WebAuthnDisplayName() string {
	return u.DisplayName
}

func (u UserCredential) WebAuthnIcon() string {
	return "" // アイコンがない場合は空文字を返します
}

func (u UserCredential) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

func (userCredential *UserCredential) AddCredential(credential webauthn.Credential) {
	userCredential.Credentials = append(userCredential.Credentials, credential)
}
