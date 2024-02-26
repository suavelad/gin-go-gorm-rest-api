package controller

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/suavelad/gin-gorm-rest/config"
	_ "github.com/suavelad/gin-gorm-rest/docs"
	"github.com/suavelad/gin-gorm-rest/models"
	"github.com/suavelad/gin-gorm-rest/serializers"
	"github.com/suavelad/gin-gorm-rest/utils"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func PingController(c *gin.Context) {
	c.String(200, "Hello World")
}

// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {array} models.NewUser
// @Failure 404 {object} object
// @Router /users/ [get]
// @Security ApiKeyAuth

func GetUsers(c *gin.Context) {
	users := []models.User{}

	config.DB.Find(&users)

	message := "Users Data Fetched Successfully "
	utils.SuccessJSONResponse(200, c, message, serializers.SerializeUsers(users))
}

// @Summary Get one user
// @Description get user by ID
// @Tags Users
// @Param id path string true "User ID"
// @Success 200 {object} models.NewUser
// @Failure 400,404 {object} object
// @Router /user/{id} [get]

func GetUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id=?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	message := "User Data Fetched Successfully "
	utils.SuccessJSONResponse(200, c, message, serializers.SerializeUser(user))

}

// @Summary Create new user based on paramters
// @Description Create new user
// @Tags Users
// @Accept json
// @Param user body models.NewUser true "User Data"
// @Success 200 {object} object
// @Failure 400,500 {object} object
// @Router /user/ [post]
// @Security ApiKeyAuth

func CreateUser(c *gin.Context) {
	var user models.User

	if c.Bind(&user) != nil {
		error_message := "Invalid payload used"
		utils.ErrorJSONResponse(400, c, error_message)
		return

	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		error_message := "Error when creating user"
		utils.ErrorJSONResponse(400, c, error_message)
		return
	}
	new_user := models.User{Email: user.Email, Password: string(hash), Name: user.Name}

	result := config.DB.Create(&new_user)
	if result.Error != nil {
		error_message := "An error occurred when creating this user"
		utils.ErrorJSONResponse(400, c, error_message)
		return

	}
	message := "User Created Successfully "
	utils.SuccessJSONResponse(201, c, message, serializers.SerializeUser(user))

}

// @Summary Create new user based on parameters
// @Description Create new user
// @Tags Users
// @Accept json
// @Param user body models.NewUser true "User Data"
// @Success 200 {object} object
// @Failure 400,500 {object} object
// @Router /login [post]
// @Security ApiKeyAuth

func Login(c *gin.Context) {
	var user models.User
	var login_user models.NewUser

	if c.Bind(&login_user) != nil {
		error_message := "Email or Password is invalid"
		utils.ErrorJSONResponse(400, c, error_message)
		return

	}
	// config.DB.Where("id=?", c.Param("id")).First(&user)
	config.DB.First(&user, "email= ?", login_user.Email)
	if user.Id == 0 {
		error_message := "Email or Password is incorrect"
		utils.ErrorJSONResponse(400, c, error_message)
		return

	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login_user.Password))

	if err != nil {
		error_message := "Email or Password is incorrect !"
		utils.ErrorJSONResponse(400, c, error_message)
		return

	}

	status, token := utils.GenerateJwtToken(c, user)

	if status != true {
		fmt.Println(err)
		error_message := "Failed to create token"
		utils.ErrorJSONResponse(400, c, error_message)
		return

	}
	message := "User Login Successful"
	utils.SuccessTokenJSONResponse(200, c, message, serializers.SerializeUser(user), token)

}

func GetNewAccessToken(c *gin.Context) {
	var refresh_details models.RefreshToken

	c.Bind(&refresh_details)

	var user models.User
	config.DB.First(&user, refresh_details.Id)

	status, token_pair := utils.GenerateAccessTokenFromRefreshToken(c, user, refresh_details.RefreshToken)

	if status == false {
		error_message := "Failed to create token"
		utils.ErrorJSONResponse(400, c, error_message)
		return

	} else {
		message := "New Token Gotten Successfully"
		utils.SuccessTokenJSONResponse(200, c, message, serializers.SerializeUser(user), token_pair)
	}
}

// @Summary Get one user
// @Description get user by ID
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.NewUser
// @Failure 400,404 {object} object
// @Router /user/{id} [delete]
// @Security ApiKeyAuth

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(204, &user)
	message := "User Deleted Successfully "
	utils.SuccessJSONResponse(204, c, message, "")
}

// @Summary Get one user
// @Description get user by ID
// @Produce json
// @Tags Users
// @Param id path string true "User ID"
// @Success 200 {object} ""
// @Failure 400,404 {object} object
// @Router /upload [post]
// @Security ApiKeyAuth

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	message := "User Data Updated Successfully "
	utils.SuccessJSONResponse(200, c, message, serializers.SerializeUser(user))

}

// @Summary Create new upload based on paramters
// @Description Create new file upload
// @Tags Users
// @Produce json
// @Accept json
// @Param user body "profile_images" true "User Data"
// @Success 200 {object} object
// @Failure 400,500 {object} object
// @Router /user/ [post]

func UploadUserProfileImage(c *gin.Context) {
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		error_message := fmt.Sprintf("Error parsing form: %s", err.Error())
		utils.ErrorJSONResponse(400, c, error_message)
		return
	}

	files := form.File["profile_images"]

	// Specify the destination path (adjust accordingly)
	uploadPath := "./uploads/"
	log.Println(files)

	if len(files) < 1 {
		error_message := fmt.Sprintf("No file was selected")
		utils.ErrorJSONResponse(400, c, error_message)
		return

	}

	for _, file := range files {
		log.Println("Uploading file:", file.Filename)

		// Construct the full path where the file will be saved
		filePath := filepath.Join(uploadPath, file.Filename)

		// Upload the file to the specified destination path
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			error_message := fmt.Sprintf("Error saving file: %s", err.Error())
			utils.ErrorJSONResponse(500, c, error_message)
			return
		}
	}

	// Provide a success message with the full path
	message := fmt.Sprintf("%d file(s) uploaded ", len(files))
	utils.SuccessJSONResponse(200, c, message, "")
}
