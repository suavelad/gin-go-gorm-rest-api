package controller

import (
	"fmt"
	"log"
	"path/filepath"

	_ "github.com/suavelad/gin-gorm-rest/docs"
	"github.com/suavelad/gin-gorm-rest/utils"

	"github.com/gin-gonic/gin"
)

// swagger:route POST /user/upload user UploadUserProfile
// @Summary Upload User Profile
// @Description upload user profile
// @Produce json
// @Tags Users
// @Param profile_image formData file true "Image file"
// @Success 200 {array} structure.NewUser
// @Failure 400,401,404 {object} object
// @Router /user/upload [post]
// @Security oauth2
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
