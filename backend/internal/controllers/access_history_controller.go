package controllers

import (
	"log"
	"net/http"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func GetAccessHistoryController(ctx *gin.Context) {
	var accessHistrory []schema.GetAccessHistory200ResponseInner
	date := ctx.Param("date")

	accessHistrory, err := services.GetAccessHistoryService(date)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, accessHistrory)
	}
}
