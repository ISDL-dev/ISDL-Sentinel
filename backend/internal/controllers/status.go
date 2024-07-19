package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func PutStatusHandlerFunc(ctx *gin.Context) {
	var status schema.Status
	if err := ctx.BindJSON(&status); err != nil {
		log.Printf("Internal Server Error: failed to bind a request body!: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := services.PutStatusService(schema.Status{UserId: status.UserId, Status: status.Status})
	if err != nil {
		log.Println(fmt.Errorf("failed to get user status:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}