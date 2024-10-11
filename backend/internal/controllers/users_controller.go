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

func GetUsersByIdController(ctx *gin.Context) {
	userIdStr := ctx.Param("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println(fmt.Errorf("invalid user_id '%s': %w", userIdStr, err))
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
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

func PutUsersByIdController(ctx *gin.Context) {
	var userInformation schema.PutUserByIdRequest
	if err := ctx.BindJSON(&userInformation); err != nil {
		log.Printf("Internal Server Error: failed to bind a request body!: %v\n", err)
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userIdStr := ctx.Param("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println(fmt.Errorf("invalid user_id '%s': %w", userIdStr, err))
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = services.PutUsersByIdService(userId, userInformation)
	if err != nil {
		log.Println(fmt.Errorf("failed to put user information:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.Status(http.StatusOK)
	}
}
