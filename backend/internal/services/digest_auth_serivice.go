package services

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"


	"github.com/gin-gonic/gin"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

var realm = "@mikilab.doshisha.ac.jp"
var users map[string]schema.PostSignInRequest

func init() {
	var err error
	users, err = GetAuthInfoService()
	if err != nil {
		log.Fatalf("Failed to get auth user from database: %v", err)
	}
}

func ValidateDigestAuth(auth, method, uri string) bool {
	params := ParseDigestAuth(auth)

	mailAddress := params["username"]
	user, ok := users[mailAddress]
	if !ok {
		log.Println("not found user")
		return false
	}

	ha1 := MD5Hash(fmt.Sprintf("%s:%s:%s", mailAddress, realm, user.Password))
	ha2 := MD5Hash(fmt.Sprintf("%s:%s", method, uri))

	expectedResponse := MD5Hash(fmt.Sprintf("%s:%s:%s:%s:%s:%s",
		ha1, params["nonce"], params["nc"], params["cnonce"], params["qop"], ha2))
	log.Println(params["response"])
	log.Println(expectedResponse)
	return expectedResponse == params["response"]
}

func Challenge(ctx *gin.Context) {
	nonce := GenerateNonce()
	ctx.Header("WWW-Authenticate", fmt.Sprintf(`Digest realm="%s",qop="auth",nonce="%s"`, realm, nonce))
	ctx.AbortWithStatus(http.StatusUnauthorized)
}

func GenerateNonce() string {
	nonceBytes := make([]byte, 16)
	rand.Read(nonceBytes)
	return fmt.Sprintf("%d-%s", time.Now().UnixNano(), hex.EncodeToString(nonceBytes))
}

func ParseDigestAuth(auth string) map[string]string {
	params := make(map[string]string)
	re := regexp.MustCompile(`(\w+)=("([^"]+)"|\w+)`)
	matches := re.FindAllStringSubmatch(auth, -1)
	
	for _, match := range matches {
		key := strings.ToLower(match[1])
		value := strings.Trim(match[2], `"`)
		params[key] = value
	}
	
	return params
}

func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetAuthInfoService() (map[string]schema.PostSignInRequest, error) {
	authUsers := make(map[string]schema.PostSignInRequest)
	
	rows, err := repositories.GetAuthInfoRepository()
	if err != nil {
		return authUsers, fmt.Errorf("failed to execute query to get auth user information: %v", err)
	}

	for rows.Next() {
        var mailAddress, password string
        err := rows.Scan(&mailAddress, &password)
        if err != nil {
            return nil, err
        }

        authUsers[mailAddress] = schema.PostSignInRequest{
            MailAddress: mailAddress,
            Password:    password,
        }
    }

	return authUsers, nil
}