package controller

import (
	"fmt"

	"github.com/suavelad/gin-gorm-rest/config"
	_ "github.com/suavelad/gin-gorm-rest/docs"
	"github.com/suavelad/gin-gorm-rest/jobs"
	"github.com/suavelad/gin-gorm-rest/models"
	"github.com/suavelad/gin-gorm-rest/serializers"
	"github.com/suavelad/gin-gorm-rest/structure"
	"github.com/suavelad/gin-gorm-rest/utils"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

// swagger:route POST /auth/signup user SignUp
// @Summary Create new user based on parameters
// @Description Create new user
// @Tags Auth
// @Accept json
// @Param user body structure.NewUser true "User Data"
// @Success 200 {object} object
// @Failure 400,401,500 {object} object
// @Router /auth/signup/ [post]
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
	config.DB.Where("email=?", user.Email).First(&user)

	if user.Id != 0 {
		error_message := "User with this email already exist"
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

	fmt.Println("Sent to Email Send Job")
	task := jobs.Task{
		Function: utils.SendEmailTask,
		Input:    []interface{}{user.Email, "", "First Go Mail", "New User SignUp"},
		Result:   make(chan interface{}),
		Error:    make(chan error),
	}

	// Run the task in the background using a goroutine
	go jobs.ExecuteTask(task)

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	jobs.ExecuteTask(task)
	// }()

	// // ... continue with your main code ...

	// // Wait for all background jobs to finish
	// wg.Wait()

	message := "User Created Successfully "
	utils.SuccessJSONResponse(201, c, message, serializers.SerializeUser(user))

}

// swagger:route POST /auth/login  user login
// @Summary User Login based on parameters
// @Description User Login
// @Tags Auth
// @Accept json
// @Param user body structure.LoginUser true "User Login"
// @Success 200 {object} object
// @Failure 400,500 {object} object
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var user models.User
	var login_user structure.NewUser

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

// swagger:route POST /auth/refresh  user refresh
// @Summary Generate Access token using Refresh Token based on parameters
// @Description Generate Access token using Refresh Token based on parameters
// @Tags Auth
// @Accept json
// @Param user body structure.RefreshToken true "User Refresh"
// @Success 200 {object} object
// @Failure 400,500 {object} object
// @Router /auth/refresh [post]
func GetNewAccessToken(c *gin.Context) {
	var refresh_details structure.RefreshToken

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
