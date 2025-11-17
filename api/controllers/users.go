// controllers/ordercontroller.go

//Will create possible token creation and validation

package controllers

import (
	"github.com/Acstrayer/TESCSE-Ecom/api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Hashes and salts password for DB storage
func HashPassword(pass string) (string, error) {
	bytepass := []byte(pass)
	hash, err := bcrypt.GenerateFromPassword(bytepass, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// User creation endpoint
func CreateUser(c *gin.Context) {
	input := new(models.UserInput)
	//checks value extraction from json
	if err := c.ShouldBindJSON(input); err != nil || input.Users_name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashed_password, err := HashPassword(input.Password)
	//checks whether hash and salt failed
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}
	//creates DB model object
	user := models.User{
    Users_name:        input.Users_name,
		Permissions_level: 0,
		Password_hash:     hashed_password,
		Email:             input.Email,
  }
	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// User login endpoint
func UserLogin(c *gin.Context) {
	input := new(models.UserInput)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
  user := models.User{}
  models.DB.Where("email = ?", input.Email).First(&user)
  if user == (models.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
  }
  if err := bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(input.Password)); err == nil {
    c.JSON(http.StatusOK, gin.H{"access": "granted"})
  } else {
    c.JSON(http.StatusForbidden, gin.H{"access": "denied"})
  }
}
