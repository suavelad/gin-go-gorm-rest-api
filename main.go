package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suavelad/gin-gorm-rest/config"
	_ "github.com/suavelad/gin-gorm-rest/docs"
	"github.com/suavelad/gin-gorm-rest/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// func init() {
// 	initializer.LoadEnvVariables()
// }

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
	routes.UserRouter(router)
	// ginSwagger.WrapHandler(swaggerFiles.Handler,
	// 	ginSwagger.URL("http://localhost:8888/swagger/doc.json"),
	// 	ginSwagger.DefaultModelsExpandDepth(-1))
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8888/swagger/doc.json")))
	// PORT := os.Getenv("PORT")
	PORT := "8888"
	router.Run(":" + PORT)

}
