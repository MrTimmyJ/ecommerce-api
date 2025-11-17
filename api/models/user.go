// models/user.go

package models

import (
	"gorm.io/gorm"
)

type User struct {
  gorm.Model
	Users_name        string `json:"users_name"`
	Permissions_level uint   `json:"permissions_level"`
	Password_hash     string `json:"password_hash"`
	Email             string `json:"email"`
}

type UserInput struct {
	Users_name string `json:"users_name"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"required"`
}
