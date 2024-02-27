package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suavelad/gin-gorm-rest/controller"
)

func HealthRouter(router *gin.Engine) {

	// Define routes without RequireAuth middleware
	authHealth := router.Group("/health")
	{
		authHealth.GET("/", controller.PingController)

	}

}
