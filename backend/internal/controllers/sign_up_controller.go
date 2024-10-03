package controllers

import (
	"fmt"
    "net/http"
	"log"

    "github.com/gin-gonic/gin"
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)



func PostSignUpController(ctx *gin.Context){
	var user schema.PostUserSignUpRequest

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(fmt.Errorf("Failed to get user information"))

	if err := services.PostSignUpService(user); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	SignUpUserInfo, err := services.GetLoginUserInfoService(user.AuthUserName)
	if err != nil {
		log.Println(fmt.Errorf("Failed to get user information:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get user information",
		})
		return
	}

	ctx.JSON(http.StatusOK, SignUpUserInfo)
}