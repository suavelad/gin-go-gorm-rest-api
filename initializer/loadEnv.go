package initializer

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env files")
	}

	fmt.Println("🚀 🚀 🚀  Successfully Loaded the .env files   🚀 🚀 🚀")

}
