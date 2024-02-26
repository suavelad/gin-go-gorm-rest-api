package serializers

import (
	"github.com/suavelad/gin-gorm-rest/models"

	"github.com/gin-gonic/gin"
)

func SerializeUser(user models.User) gin.H {
	return gin.H{
		"id":    user.Id,
		"name":  user.Name,
		"email": user.Email,
	}
}

func SerializeUsers(users []models.User) []gin.H {
	var serializedUsers []gin.H
	for _, user := range users {
		serializedUser := SerializeUser(user)
		serializedUsers = append(serializedUsers, serializedUser)
	}
	return serializedUsers
}
