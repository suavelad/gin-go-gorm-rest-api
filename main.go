package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/suavelad/gin-gorm-rest/config"

	_ "github.com/suavelad/gin-gorm-rest/docs"
	"github.com/suavelad/gin-gorm-rest/initializer"
	"github.com/suavelad/gin-gorm-rest/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializer.LoadEnvVariables()

}

// @title User Gin App Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @host localhost:8888
// @contact.name API Support
// @contact.email sunnexajayi@gmail.com

//@BasePath /

func main() {
	router := gin.New()
	config.Connect()

	// Unprotected Routes
	routes.AuthRouter(router)
	routes.HealthRouter(router)

	swaggerURL := fmt.Sprintf("%s/swagger/doc.json", os.Getenv("DOMAIN_HOST"))

	// Use ginSwagger.WrapHandler to serve Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(swaggerURL)))

	//For Protected Routes
	routes.UserRouter(router)

	router.Run(":" + os.Getenv("PORT"))
}
