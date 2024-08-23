package internal

import (
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/controllers"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
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

	v1 := router.Group("/v1")
	{
		v1.GET("/users/:user_id", controllers.GetUsersByIdController)
		v1.POST("/avatar", controllers.PostAvatarController)
		v1.PUT("/avatar", controllers.PutAvatarController)
		v1.DELETE("/avatar", controllers.DeleteAvatarController)
		v1.GET("/attendees-list", controllers.GetAttendeesListController)
		v1.PUT("/status", controllers.PutStatusController)
		v1.GET("/access-history/:date", controllers.GetAccessHistoryController)
		v1.GET("/ranking", controllers.GetRankingController)
		v1.GET("/lab-assistant-member", controllers.GetLabAssistantMemberController)
		v1.GET("/lab-assistant/:month", controllers.GetLabAssistantScheduleController)
		v1.POST("/lab-assistant/:month", controllers.PostLabAssistantScheduleController)

		oauthn := v1.Group("/oauthn")
		{
			oauthn.GET("/callback", infrastructures.Callback)
		}

		webauthn := v1.Group("/webauthn")
		{
			webauthn.GET("/register-begin/:user_name", controllers.GetBeginRegistrationController)
			webauthn.POST("/register-finish/:user_name", controllers.GetFinishRegistrationController)
			webauthn.GET("/login-begin/:user_name", controllers.GetBeginLoginController)
			webauthn.POST("/login-finish/:user_name", controllers.GetFinishLoginController)
		}
	}
}
