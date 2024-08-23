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

func PostAvatarController(ctx *gin.Context) {
	// Get user_id from form data and convert it to an int
	userIdStr := ctx.PostForm("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println(fmt.Errorf("invalid user_id: %w", err))
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Get avatar_file from form data
	avatarFile, err := ctx.FormFile("avatar_file")
	if err != nil {
		log.Println(fmt.Errorf("failed to get avatar file: %w", err))
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Call the service to handle the avatar upload
	err = services.PostAvatarService(userId, avatarFile)
	if err != nil {
		log.Println(fmt.Errorf("failed to upload avatar: %w", err))
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	} else {
		ctx.Status(http.StatusOK)
	}
}

func PutAvatarController(ctx *gin.Context) {
	var avatarRequest schema.Avatar
	if err := ctx.BindJSON(&avatarRequest); err != nil {
		log.Printf("Internal Server Error: failed to bind a request body with a struct: %v\n", err)
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	avatarResponse, err := services.PutAvatarService(avatarRequest)
	if err != nil {
		log.Println(fmt.Errorf("failed to put avatar:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, avatarResponse)
	}
}

func DeleteAvatarController(ctx *gin.Context) {
	var avatarRequest schema.Avatar
	if err := ctx.BindJSON(&avatarRequest); err != nil {
		log.Printf("Internal Server Error: failed to bind a request body with a struct: %v\n", err)
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	if avatarRequest.AvatarId == 1 {
		log.Printf("the default avatar cannot be deleted")
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: "the default avatar cannot be deleted",
		})
		return
	}

	err := services.DeleteAvatarService(avatarRequest)
	if err != nil {
		log.Println(fmt.Errorf("failed to delete avatar:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.Status(http.StatusOK)
	}
}
