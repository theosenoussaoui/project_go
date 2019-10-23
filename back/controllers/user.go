package controllers

import (
	"net/http"

	"github.com/theosenoussaoui/project_go/back/db"
	"github.com/theosenoussaoui/project_go/back/models"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

// GetUsers g
func GetUsers(c *gin.Context) {
	var users []models.User
	db := db.GetDB()
	db.Find(&users)
	c.JSON(200, users)
}

// CreateUser c
func CreateUser(c *gin.Context) {
	var user models.User
	var db = db.GetDB()

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check admin right creation
	if user.IsAdmin() {
		// Retrive user information
		jwtClaims := jwt.ExtractClaims(c)
		authUserAccessLevel := jwtClaims["access_level"].(float64)
		if authUserAccessLevel != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Sorry but you can't created admin user",
			})
			return
		}
	}

	// Create user if no errors throws by Validate method in user model
	if err := db.Create(&user); err.Error != nil {

		// convert array of errors to JSON
		errs := err.GetErrors()
		strErrors := make([]string, len(errs))
		for i, err := range errs {
			strErrors[i] = err.Error()
		}

		// return errors
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": strErrors,
		})

		return
	}

	c.JSON(http.StatusOK, &user)
}

// UpdateUser u
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(http.StatusOK, &user)
}

// Deleting User u [WIP]
// func DeleteUser(c *gin.Context) {
// 	uuid := c.Params.ByName("uuid")
// 	var user models.User
// 	db := db.GetDB()
// 	if uuid != "" {
// 	}
// }
