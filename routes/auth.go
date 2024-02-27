package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suavelad/gin-gorm-rest/controller"
	"github.com/suavelad/gin-gorm-rest/middleware"
)

func AuthRouter(router *gin.Engine) {

	// Define routes without RequireAuth middleware
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/signup", controller.CreateUser)
		authRoutes.POST("/login", controller.Login)
		authRoutes.POST("/refresh", middleware.RequireAuth, controller.GetNewAccessToken)

	}

}
