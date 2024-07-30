package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func GetLabAsistantScheduleController(ctx *gin.Context) {
	date := ctx.Param("date")

	labAsistantSchedule, err := services.GetLabAsistantScheduleService(date)
	if err != nil {
		log.Println(fmt.Errorf("failed to get lab asistant schedule:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, labAsistantSchedule)
	}
}
