package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func GetLabAsistantMemberController(ctx *gin.Context) {
	labAsistantMemberList, err := services.GetLabAsistantMemberService()
	if err != nil {
		log.Println(fmt.Errorf("failed to get lab asistant member:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, labAsistantMemberList)
	}
}
