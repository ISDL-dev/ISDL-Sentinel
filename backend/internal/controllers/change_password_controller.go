package controllers

import (
    "net/http"
	"strconv"

    "github.com/gin-gonic/gin"
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/services"
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func PutChangePasswordController(ctx *gin.Context){
    
	userIDStr := ctx.Param("id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
	
	var user schema.PutChangePasswordRequest
    if err := ctx.ShouldBind(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := services.PutChangePasswordService(user, userID); err != nil {
        ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK,"password changed")
}