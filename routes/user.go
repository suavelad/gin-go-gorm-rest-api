package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suavelad/gin-gorm-rest/controller"
	"github.com/suavelad/gin-gorm-rest/middleware"
)

func UserRouter(protectedRouter *gin.Engine) {

	// Define routes without RequireAuth middleware
	// authRoutes := router.Group("/auth")
	// {
	// 	authRoutes.POST("/signup", controller.CreateUser)
	// 	authRoutes.POST("/login", controller.Login)
	// 	authRoutes.POST("/refresh", middleware.RequireAuth, controller.GetNewAccessToken)

	// }

	// Define routes with RequireAuth middleware

	// Apply global middleware
	protectedRouter.Use(middleware.RequireAuth)

	protectedRoutes := protectedRouter.Group("/users")
	// You can add additional routes to the protected group if needed
	{
		protectedRoutes.GET("/", controller.GetUsers)
		protectedRoutes.GET("/:id", controller.GetUser)
		protectedRoutes.DELETE("/:id", controller.DeleteUser)
		protectedRoutes.PUT("/:id", controller.UpdateUser)
		protectedRoutes.POST("/upload", controller.UploadUserProfileImage)
	}
}
