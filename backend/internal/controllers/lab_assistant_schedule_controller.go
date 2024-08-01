package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func GetLabAssistantScheduleController(ctx *gin.Context) {
	month := ctx.Param("month")

	labAssistantSchedule, err := services.GetLabAssistantScheduleService(month)
	if err != nil {
		log.Println(fmt.Errorf("failed to get lab assistant schedule:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, labAssistantSchedule)
	}
}

func PostLabAssistantScheduleController(ctx *gin.Context) {
	var labAssistantScheduleRequest []schema.PostLabAssistantScheduleRequestInner
	if err := ctx.BindJSON(&labAssistantScheduleRequest); err != nil {
		log.Printf("Internal Server Error: failed to bind a request body with a struct: %v\n", err)
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	month := ctx.Param("month")

	labAssistantSchedule, err := services.PostLabAssistantScheduleService(month, labAssistantScheduleRequest)
	if err != nil {
		log.Println(fmt.Errorf("failed to post lab assistant schedule:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, labAssistantSchedule)
	}
}
