package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/environment"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgresql"
	"github.com/gin-gonic/gin"
)

// CreateUser creates a new user.
func CreateUser(c *gin.Context) {
	var user models.CreateUserRequest

	c.BindJSON(&user)

	usersService := postgresql.NewUser(environment.Pool)

	res, err := usersService.Create(context.Background(), user)
	if err != nil {
		// TODO  equivalent of Python exception handler context manager:
		// https://stackoverflow.com/questions/69948784/how-to-handle-errors-in-gin-middleware
		c.JSON(http.StatusInternalServerError, models.ValidationError{Msg: err.Error()})

		return
	}

	fmt.Printf("Res %#v\n", res)
	c.JSON(http.StatusOK, res)
}

func CreateUsersWithListInput(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
