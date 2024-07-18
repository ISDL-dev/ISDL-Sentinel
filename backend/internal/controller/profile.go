package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

func GetUsersHandlerFunc(ctx *gin.Context) {
	var userId int

	userId, _ = strconv.Atoi(ctx.Param("user_id"))
	userInformation, err := repository.GetUsers(userId)
	if err != nil {
		log.Println(fmt.Errorf("failed to get user information:%w", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("failed to get user information:%w", err),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"content": userInformation,
		})
	}
}
