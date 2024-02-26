package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type NewUser struct {
	Id       int
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
	Id           int    `json:"id"`
}
