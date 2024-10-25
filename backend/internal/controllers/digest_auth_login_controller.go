package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func DigestLoginController(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")

	if auth == "" {
		log.Println("challenge")

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Expose-Headers", "WWW-Authenticate")

		nonce := services.GenerateNonce()
		wwwAuthenticateHeader := services.CreateWWWAuthenticateHeader(nonce)
		ctx.Header("X-WWW-Authenticate", wwwAuthenticateHeader)
		ctx.JSON(http.StatusUnauthorized, schema.Error{
			Code:    http.StatusUnauthorized,
			Message: "Authentication required",
		})
		return
	}

	userName, err := services.ValidateDigestAuth(auth, ctx.Request.Method, ctx.Request.URL.RequestURI())
	if err != nil {
		log.Println(fmt.Errorf("Authentication failed:%w", err))
		ctx.JSON(http.StatusUnauthorized, schema.Error{
			Code:    http.StatusUnauthorized,
			Message: "Authentication failed",
		})
		return
	}

	loginUserInfo, err := services.GetLoginUserInfoService(userName)
	if err != nil {
		log.Println(fmt.Errorf("Failed to get user information:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get user information",
		})
		return
	}

	ctx.JSON(http.StatusOK, loginUserInfo)
}
