package services

import (
    "crypto/md5"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "strings"

    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

const (
    realm = "@mikilab.doshisha.ac.jp"
)

func CreateWWWAuthenticateHeader(nonce string) string {
    return fmt.Sprintf(`Digest realm="%s", nonce="%s", qop="auth"`, realm, nonce)
}

func ValidateDigestAuth(auth, method, uri string) (string, error) {
    params, err := ParseDigestAuth(auth)
    if err != nil {
        return "", fmt.Errorf("failed to parse authorization header: %w", err)
    }

    userName := params["username"]

    userInfo, err := repositories.GetDigestCredential(userName)
    if err != nil {
        return "", fmt.Errorf("failed to get user credential: %w", err)
    }

    ha1 := MD5Hash(fmt.Sprintf("%s:%s:%s", userName, realm, userInfo.Password))
    ha2 := MD5Hash(fmt.Sprintf("%s:%s", method, uri))

    expectedResponse := MD5Hash(fmt.Sprintf("%s:%s:%s:%s:%s:%s", ha1, params["nonce"], params["nc"], params["cnonce"], params["qop"], ha2))

    if expectedResponse != params["response"] {
        return "", fmt.Errorf("invalid digest")
    }

    return userInfo.Name, nil
}

func ParseDigestAuth(auth string) (map[string]string, error) {
    if !strings.HasPrefix(auth, "Digest ") {
        return nil, fmt.Errorf("invalid authorization header")
    }

    parts := strings.Split(auth[7:], ",")
    params := make(map[string]string)

    for _, part := range parts {
        kv := strings.SplitN(strings.TrimSpace(part), "=", 2)
        if len(kv) != 2 {
            continue
        }
        params[kv[0]] = strings.Trim(kv[1], "\"")
    }

    return params, nil
}

func GenerateNonce() string {
    b := make([]byte, 16)
    rand.Read(b)
    return hex.EncodeToString(b)
}

func MD5Hash(data string) string {
    hash := md5.Sum([]byte(data))
    return hex.EncodeToString(hash[:])
}

func GetLoginUserInfoService(userName string) (schema.PostSignIn200Response, error) {
    userInfo, err := repositories.GetLoginUserInfo(userName)
    if err != nil {
        return schema.PostSignIn200Response{}, fmt.Errorf("failed to get login user info: %w", err)
    }
    return userInfo, nil
}

