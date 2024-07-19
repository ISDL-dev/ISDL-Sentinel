package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ISDL-dev/ISDL_Sentinel/backend/internal/repository"
	"github.com/ISDL-dev/ISDL_Sentinel/backend/internal/schema"
	"github.com/gin-gonic/gin"
)

func PutStatusHandlerFunc(ctx *gin.Context) {
	var status schema.Status
	user, err := repository.PutStatus(status)
	if err != nil {
		log.Println(fmt.Errorf("failed to send status to the client:%w", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("failed to send status to the client:%w", err),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"content": user,
		})
	}
}