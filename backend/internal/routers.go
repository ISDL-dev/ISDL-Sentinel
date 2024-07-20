package internal

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ISDL-dev/ISDL_Sentinel/backend/internal/controller"
)

func SetRoutes(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4000"},
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))


	public := router.Group("/")
	{
		public.POST("/auth/sign-in", controller.DigestAuthSignIn)
	}

	authenticated := router.Group("/")
	authenticated.Use(controller.DigestAuthMiddleware())
	{
		// authenticated.GET("/profile", controller.GetProfile)ßß
	}
}