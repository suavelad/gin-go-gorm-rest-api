package controller

import (
	_ "github.com/suavelad/gin-gorm-rest/docs"

	"github.com/gin-gonic/gin"
)

// swagger:route GET /health  checkHealth
// @Summary Check Health
// @Description Check Health
// @Tags Health
// @Success 200
// @Failure 400,401,404 {object} object
// @Router /health [get]
func PingController(c *gin.Context) {
	c.String(200, "Hello World")
}
