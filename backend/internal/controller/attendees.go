package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAttendeesListHandlerFunc(ctx *gin.Context) {
	repository.PutStatus()
	ctx.String(http.StatusOK, "success")
}