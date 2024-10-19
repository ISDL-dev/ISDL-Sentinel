package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func GetRankingController(ctx *gin.Context) {
	term := ctx.Param("term")
	rankingList, err := services.GetRankingService(term)
	if err != nil {
		log.Println(fmt.Errorf("failed to get ranking list:%w", err))
		ctx.JSON(http.StatusInternalServerError, schema.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, rankingList)
	}
}
