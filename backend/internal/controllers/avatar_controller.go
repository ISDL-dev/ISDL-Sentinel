package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func PutAvatarControlller(ctx *gin.Context) {
	var avatarRequest schema.PutAvatarRequest
	if err := ctx.BindJSON(&avatarRequest); err != nil {
		log.Printf("Internal Server Error: failed to bind a request body with a struct: %v\n", err)
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err := services.PutAvatarService(avatarRequest)
	if err != nil {
		log.Println(fmt.Errorf("failed to put avatar:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.String(http.StatusOK, "success")
	}
}
