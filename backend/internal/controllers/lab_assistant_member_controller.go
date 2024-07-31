package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func GetLabAssistantMemberController(ctx *gin.Context) {
	labAssistantMemberList, err := services.GetLabAssistantMemberService()
	if err != nil {
		log.Println(fmt.Errorf("failed to get lab assistant member:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, labAssistantMemberList)
	}
}
