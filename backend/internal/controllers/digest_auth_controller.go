package controllers

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"regexp"
	"time"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
)

var realm = "@mikilab.doshisha.ac.jp"
var users map[string]schema.PostSignInRequest

func init() {
	var err error
	users, err = services.GetAuthUserInfoService()
	if err != nil {
		log.Fatalf("Failed to get auth user from database: %v", err)
	}
}

func DigestAuthController(ctx *gin.Context) {
    auth := ctx.GetHeader("Authorization")
    if auth == "" {
        challenge(ctx)
        return
    }

    if !validateDigestAuth(auth, ctx.Request.Method, ctx.Request.URL.Path) {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

	ctx.JSON(http.StatusOK, gin.H{"message": "Authentication successful"})
}

func validateDigestAuth(auth, method, uri string) bool {
	params := parseDigestAuth(auth)

	mailAddress := params["username"] // ダイジェスト認証ではusernameフィールドを使用
	user, ok := users[mailAddress]
	if !ok {
		return false
	}

	ha1 := md5Hash(fmt.Sprintf("%s:%s:%s", mailAddress, realm, user.Password))
	ha2 := md5Hash(fmt.Sprintf("%s:%s", method, uri))

	expectedResponse := md5Hash(fmt.Sprintf("%s:%s:%s:%s:%s:%s",
		ha1, params["nonce"], params["nc"], params["cnonce"], params["qop"], ha2))

	return expectedResponse == params["response"]
}

func challenge(ctx *gin.Context) {
	nonce := generateNonce()
	ctx.Header("WWW-Authenticate", fmt.Sprintf(`Digest realm="%s",qop="auth",nonce="%s"`, realm, nonce))
	ctx.AbortWithStatus(http.StatusUnauthorized)
}

func generateNonce() string {
	nonceBytes := make([]byte, 16)
	rand.Read(nonceBytes)
	return fmt.Sprintf("%d-%s", time.Now().UnixNano(), hex.EncodeToString(nonceBytes))
}

func parseDigestAuth(auth string) map[string]string {
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

func md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func DigestAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			challenge(ctx)
			ctx.Abort()
			return
		}

		if !validateDigestAuth(auth, ctx.Request.Method, ctx.Request.URL.Path) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
