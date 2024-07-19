package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func GetUsersByIdControlller(ctx *gin.Context) {
	userIdStr := ctx.Param("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println(fmt.Errorf("invalid user_id '%s': %w", userIdStr, err))
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userInformation, err := services.GetUsersByIdService(userId)
	if err != nil {
		log.Println(fmt.Errorf("failed to get user information:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, userInformation)
	}
}
