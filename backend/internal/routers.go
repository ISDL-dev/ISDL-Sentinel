package internal

import (
	"github.com/ISDL-dev/ISDL_Sentinel/backend/internal/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:4000",
		},
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))

	router.GET("/v1/attendees-list", controller.GetAttendeesListHandlerFunc)
}
