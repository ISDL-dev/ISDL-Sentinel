package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func DigestLoginController(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")

	if auth == "" {
		nonce := services.GenerateNonce()
		wwwAuthenticateHeader := services.CreateWWWAuthenticateHeader(nonce)
		ctx.Header("WWW-Authenticate", wwwAuthenticateHeader)
		ctx.JSON(http.StatusUnauthorized, schema.Error{
			Code:    http.StatusUnauthorized,
			Message: "Authentication required",
		})
		return
	}

	userName, err := services.ValidateDigestAuth(auth, ctx.Request.Method, ctx.Request.URL.RequestURI())
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, schema.Error{
			Code:    http.StatusUnauthorized,
			Message: "Authentication failed",
		})
		return
	}

	loginUserInfo, err := services.GetLoginUserInfoService(userName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get user information",
		})
		return
	}

	ctx.JSON(http.StatusOK, loginUserInfo)
}