package handlers

import (
	"context"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/environment"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgresql"
	"github.com/gin-gonic/gin"
)

// CreateUser creates a new user.
func CreateUser(c *gin.Context) {
	var user models.CreateUserRequest

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())

		return
	}

	usersService := postgresql.NewUser(environment.Pool)

	environment.Logger.Sugar().Debugf("CreateUser.user: %v", user)

	res, err := usersService.Create(context.Background(), user)
	if err != nil {
		// TODO  equivalent of Python exception handler context manager:
		// https://stackoverflow.com/questions/69948784/how-to-handle-errors-in-gin-middleware
		c.JSON(http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateUsersWithListInput(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
