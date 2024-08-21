package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func DigestAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        auth := c.GetHeader("Authorization")
        if auth == "" {
            nonce := services.GenerateNonce()
            c.Header("WWW-Authenticate", services.CreateWWWAuthenticateHeader(nonce))
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        username, err := services.ValidateDigestAuth(auth, c.Request.Method, c.Request.URL.RequestURI())
        if err != nil {
            c.Header("WWW-Authenticate", services.CreateWWWAuthenticateHeader(services.GenerateNonce()))
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        c.Set("username", username)
        c.Next()
    }
}

func DigestLoginController(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	if auth == "" {
		nonce := services.GenerateNonce()
		wwwAuthenticateHeader := services.CreateWWWAuthenticateHeader(nonce)
		c.Header("WWW-Authenticate", wwwAuthenticateHeader)
		c.JSON(http.StatusUnauthorized, schema.Error{
			Code:    http.StatusUnauthorized,
			Message: "Authentication required",
		})
		return
	}

	username, err := services.ValidateDigestAuth(auth, c.Request.Method, c.Request.URL.RequestURI())
	if err != nil {
		c.JSON(http.StatusUnauthorized, schema.Error{
			Code:    http.StatusUnauthorized,
			Message: "Authentication failed",
		})
		return
	}

	loginUserInfo, err := repositories.GetDigestCredential(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get user information",
		})
		return
	}

	c.JSON(http.StatusOK, loginUserInfo)
}