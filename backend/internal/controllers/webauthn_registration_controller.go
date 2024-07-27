package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func GetBeginRegistrationController(ctx *gin.Context) {
	userName := ctx.Param("user_name")

	options, err := services.GetBeginRegistrationService(userName, ctx.Writer, ctx.Request)
	if err != nil {
		log.Println(fmt.Errorf("failed to get public key credential creation options:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, options)
	}
}

func GetFinishRegistrationController(ctx *gin.Context) {
	userName := ctx.Param("user_name")

	loginUserInfo, err := services.GetFinishRegistrationService(userName, ctx.Writer, ctx.Request)
	if err != nil {
		log.Println(fmt.Errorf("failed to get public key credential creation options:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, loginUserInfo)
	}
}
