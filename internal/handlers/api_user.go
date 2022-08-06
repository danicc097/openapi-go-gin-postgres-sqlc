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

func CreateUsersWithListInput(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
