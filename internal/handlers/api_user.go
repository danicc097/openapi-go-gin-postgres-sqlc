package handlers

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateUser - Create user
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	c.String(http.StatusOK, "Would have created a user for %v", user.Username)
}

// CreateUsersWithArrayInput - Creates list of users with given input array
func CreateUsersWithArrayInput(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// CreateUsersWithListInput - Creates list of users with given input array
func CreateUsersWithListInput(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteUser - Delete user
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetUserByName - Get user by user name
func GetUserByName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// LoginUser - Logs user into the system
func LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// LogoutUser - Logs out current logged in user session
func LogoutUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateUser - Updated user
func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
