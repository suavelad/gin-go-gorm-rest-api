package utils

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/suavelad/gin-gorm-rest/config"
	_ "github.com/suavelad/gin-gorm-rest/docs"
	"github.com/suavelad/gin-gorm-rest/models"
)

var SECRET string = os.Getenv("SECRET")

func GetJWTSecretByte() []byte {

	key := []byte(SECRET)

	return key
}

func ValidateJWTToken(c *gin.Context, access_token string) (*models.User, string) {

	token, err := jwt.Parse(access_token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return GetJWTSecretByte(), nil
	})

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil, "Invalid access token"
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		// Check if token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			// error_message := "Invalid payload used"
			// utils.ErrorJSONResponse(400, c, error_message)
			return nil, "Token has expired"

		}
		// check if user exist
		var user models.User
		config.DB.First(&user, claims["sub"])

		if user.Id == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return nil, "Invalid access token"
		}
		return &models.User{}, ""
	}
	return nil, "Invalid access token"
}

// Generate jwt token
func GenerateJwtToken(c *gin.Context, user models.User) (bool, map[string]string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	t, err := token.SignedString(GetJWTSecretByte())
	if err != nil {
		fmt.Println(err)
		error_message := "Invalid authentication payload "
		ErrorJSONResponse(401, c, error_message)
		return false, nil
	}
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = user.Id
	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	rt, err := refreshToken.SignedString(GetJWTSecretByte())
	if err != nil {
		fmt.Println(err)
		error_message := "Invalid authentication payload "
		ErrorJSONResponse(401, c, error_message)
		return false, nil
	}

	return true, map[string]string{
		"access_token":  t,
		"refresh_token": rt,
		// "expiration":    string(time.Now().Add(time.Hour * 24 * 7).Unix()),
	}

}

func GenerateAccessTokenFromRefreshToken(c *gin.Context, user models.User, refresh_token string) (bool, map[string]string) {

	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify
	// which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(refresh_token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return GetJWTSecretByte(), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		if int(claims["sub"].(float64)) == user.Id {

			status, token_pair := GenerateJwtToken(c, user)
			if status != true {
				return false, nil
			}

			return true, token_pair
		}

		return false, nil
	}
	if err != nil {
		fmt.Println(err)
		return false, nil
	}

	return false, nil
}
