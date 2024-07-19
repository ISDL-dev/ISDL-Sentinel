package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repository"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/gin-gonic/gin"
)

func GetUsersHandlerFunc(ctx *gin.Context) {
	userIdStr := ctx.Param("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println(fmt.Errorf("invalid user_id '%s': %w", userIdStr, err))
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	now := time.Now()
	date := now.Format("2006-01")

	userInformation, err := repository.GetUsers(userId, date)
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
