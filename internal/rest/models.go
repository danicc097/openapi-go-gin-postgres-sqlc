package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// UserResponse represents an OpenAPI schema response for a User.
type UserResponse struct {
	Role   models.Role    `json:"role" ref:"#/components/schemas/Role"`
	Scopes []models.Scope `json:"scopes" ref:"#/components/schemas/Scopes"`
	db.UserPublic
}

// WorkItemResponse represents an OpenAPI schema response for a WorkItem.
type WorkItemResponse struct {
	db.WorkItemPublic
}
