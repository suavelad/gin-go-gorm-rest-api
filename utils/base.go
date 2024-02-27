package utils

import (
	"github.com/gin-gonic/gin"
)

func SuccessJSONResponse(status int, c *gin.Context, message string, data interface{}) {

	c.JSON(status, gin.H{
		"code":    status,
		"message": message,
		"data":    data,
	})

}

func SuccessTokenJSONResponse(status int, c *gin.Context, message string, data interface{}, token map[string]string) {

	c.JSON(status, gin.H{
		"code":    status,
		"message": message,
		"data":    data,
		"token":   token,
	})

}

func ErrorJSONResponse(status int, c *gin.Context, message string) {

	c.JSON(status, gin.H{
		"code":    status,
		"message": message,
	})

}
