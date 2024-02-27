package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/suavelad/gin-gorm-rest/utils"
)

func RequireAuth(c *gin.Context) {

	authHeader, ok := c.Request.Header["Authorization"]
	if !ok || len(authHeader) == 0 {
		utils.ErrorJSONResponse(401, c, "Authorization header missing")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Get Token from  header
	access_token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	fmt.Println((access_token))

	// Decode/validate it
	user, err := utils.ValidateJWTToken(c, access_token)

	if user != nil {
		c.Set("user", user)
		c.Next()

	} else {
		fmt.Println(user)
		error_message := string(err)
		utils.ErrorJSONResponse(401, c, error_message)
		c.AbortWithStatus(http.StatusUnauthorized)
		return

	}

}
