package controller

import (
	"github.com/suavelad/gin-gorm-rest/config"
	_ "github.com/suavelad/gin-gorm-rest/docs"
	"github.com/suavelad/gin-gorm-rest/models"
	"github.com/suavelad/gin-gorm-rest/serializers"
	"github.com/suavelad/gin-gorm-rest/utils"

	"github.com/gin-gonic/gin"
)

// swagger:route GET /users user listUsers
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {array} structure.NewUser
// @Failure 400,401,404 {object} object
// @Router /users/ [get]
// @Security oauth2
func GetUsers(c *gin.Context) {
	users := []models.User{}

	config.DB.Find(&users)

	message := "Users Data Fetched Successfully "
	utils.SuccessJSONResponse(200, c, message, serializers.SerializeUsers(users))
}

// swagger:route GET /users user getUser
// @Summary Get a user
// @Description get a user
// @Tags Users
// @Success 200 {array} structure.NewUser
// @Failure 400,401,404 {object} object
// @Router /users/{id} [get]
// @Security oauth2
func GetUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id=?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	message := "User Data Fetched Successfully "
	utils.SuccessJSONResponse(200, c, message, serializers.SerializeUser(user))

}

// swagger:route DELETE /users user DeleteUser
// @Summary Delete a user
// @Description delete a user
// @Tags Users
// @Success 204
// @Failure 400,401,404 {object} object
// @Router /users/{id} [delete]
// @Security oauth2
func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(204, &user)
	message := "User Deleted Successfully "
	utils.SuccessJSONResponse(204, c, message, "")
}

// swagger:route PUT /users user UpdateUser
// @Summary Update one user by ID
// @Description Update one user by ID
// @Tags Users
// @Success 200 {array} structure.NewUser
// @Failure 400,401,404 {object} object
// @Router /users/{id} [put]
// @Security oauth2
func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	message := "User Data Updated Successfully "
	utils.SuccessJSONResponse(200, c, message, serializers.SerializeUser(user))

}
