package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suavelad/gin-gorm-rest/controller"
	"github.com/suavelad/gin-gorm-rest/middleware"
)

func UserRouter(router *gin.Engine) {

	// Define routes without RequireAuth middleware
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/signup", controller.CreateUser)
		authRoutes.POST("/login", controller.Login)
		authRoutes.POST("/refresh", middleware.RequireAuth, controller.GetNewAccessToken)

	}

	// Define routes with RequireAuth middleware

	// Apply global middleware
	router.Use(middleware.RequireAuth)

	protectedRoutes := router.Group("/")
	// You can add additional routes to the protected group if needed
	{
		protectedRoutes.GET("/", controller.PingController)
		protectedRoutes.GET("/users", controller.GetUsers)
		protectedRoutes.GET("/user/:id", controller.GetUser)
		protectedRoutes.DELETE("/user/:id", controller.DeleteUser)
		protectedRoutes.PUT("/user/:id", controller.UpdateUser)
		protectedRoutes.POST("/user/upload", controller.UploadUserProfileImage)
	}
}
