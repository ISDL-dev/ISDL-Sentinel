package controllers

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
)

func DigestAuthController(ctx *gin.Context) {
    auth := ctx.GetHeader("Authorization")
    if auth == "" {
        services.Challenge(ctx)
        return
    }

	log.Println("Authentication attempt")
    if !services.ValidateDigestAuth(auth, ctx.Request.Method, ctx.Request.URL.Path) {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }
	log.Println("Authentication successful")

	ctx.JSON(http.StatusOK, gin.H{"message": "Authentication successful"})
}

func DigestAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			services.Challenge(ctx)
			ctx.Abort()
			return
		}

		if !services.ValidateDigestAuth(auth, ctx.Request.Method, ctx.Request.URL.Path) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
