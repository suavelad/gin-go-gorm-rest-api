package config

import (
	"log"

	"github.com/suavelad/gin-gorm-rest/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@go-gin-db:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate((&models.User{}))
	DB = db
	log.Println("🚀 🚀 🚀  Connected Successfully to the Database   🚀 🚀 🚀 ")
}
