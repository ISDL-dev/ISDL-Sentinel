package internal

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/controllers"
)

func SetRoutes(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4000"},
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	v1 := router.Group("/v1")
	{
		v1.GET("/users/:user_id", controllers.GetUsersByIdController)
		v1.PUT("/avatar", controllers.PutAvatarController)
		v1.GET("/attendees-list", controllers.GetAttendeesListController)
		v1.PUT("/status", controllers.PutStatusController)
	}

	public := router.Group("/")
	{
		public.POST("/auth/sign-in", controllers.DigestAuthController)
	}

	authenticated := router.Group("/")
	authenticated.Use(controllers.DigestAuthMiddleware())
	{
	}
}