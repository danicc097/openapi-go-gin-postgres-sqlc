package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// TODO postgen generated map entries for rest/models.go as well
type UserResponse struct {
	Role models.Role `json:"role"`
	db.UserPublic
}
