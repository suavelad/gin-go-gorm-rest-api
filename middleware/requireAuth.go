package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/suavelad/gin-gorm-rest/utils"
)

func RequireAuth(c *gin.Context) {

	// Get Token from  header
	access_token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]

	// Decode/validate it
	user, err := utils.ValidateJWTToken(c, access_token)

	if err == nil {
		c.Set("user", user)
		c.Next()

	} else {
		fmt.Println(user)
		error_message := "Invalid authentication payload "
		utils.ErrorJSONResponse(401, c, error_message)
		c.AbortWithStatus(http.StatusUnauthorized)
		return

	}

}
