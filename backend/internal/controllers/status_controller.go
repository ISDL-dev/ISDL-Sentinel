package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func PutStatusController(ctx *gin.Context) {
	var status schema.Status

	if err := ctx.BindJSON(&status); err != nil {
		log.Printf("Internal Server Error: failed to bind a request body!: %v\n", err)
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	envType := os.Getenv("ENV_TYPE")
	if envType == "prod" {
		xForwardedFor := ctx.Request.Header.Get("X-Forwarded-For")
		allowedIP := os.Getenv("LAB_NETWORK_IP")

		if status.Status == model.IN_ROOM && xForwardedFor != allowedIP {
			log.Printf("Unauthorized access attempt from IP: %s", xForwardedFor)
			ctx.JSON(http.StatusUnauthorized, schema.Error{
				Code:    http.StatusUnauthorized,
				Message: "研究室のNetworkからアクセスしてください.",
			})
			return
		}
	}

	user, err := services.PutStatusService(status)
	if err != nil {
		log.Println(fmt.Errorf("failed to get user status: %w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: "failed to get user status",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
