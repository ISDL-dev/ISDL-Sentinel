package controller

import (
	"net/http"

	"github.com/ISDL-dev/ISDL_Sentinel/backend/internal/repository"
	"github.com/ISDL-dev/ISDL_Sentinel/backend/internal/schema"
	"github.com/gin-gonic/gin"
)

func PutStatusHandlerFunc(ctx *gin.Context) {
	var status schema.PutStatusRequest
	repository.PutStatus(status)
	ctx.String(http.StatusOK, "success")
}