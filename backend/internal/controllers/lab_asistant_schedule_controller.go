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
	month := ctx.Param("month")

	labAsistantSchedule, err := services.GetLabAsistantScheduleService(month)
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

func PostLabAsistantScheduleController(ctx *gin.Context) {
	var labAsistantScheduleRequest []schema.PostLabAsistantScheduleRequestInner
	if err := ctx.BindJSON(&labAsistantScheduleRequest); err != nil {
		log.Printf("Internal Server Error: failed to bind a request body with a struct: %v\n", err)
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	month := ctx.Param("month")

	labAsistantSchedule, err := services.PostLabAsistantScheduleService(month, labAsistantScheduleRequest)
	if err != nil {
		log.Println(fmt.Errorf("failed to post lab asistant schedule:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, labAsistantSchedule)
	}
}
