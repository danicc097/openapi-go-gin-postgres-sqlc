package rest

import (
	"net/http"

	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// User handles routes with the 'user' tag.
type User struct {
	svc services.User
	// add or remove services, etc. as required
}

// NewUser returns a new handler for the 'user' route group.
// Edit as required.
func NewUser(svc services.User) *User {
	return &User{
		svc: svc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *User) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []route{
		{
			Name:        "CreateUser",
			Method:      http.MethodPost,
			Pattern:     "/user",
			HandlerFunc: h.CreateUser,
			Middlewares: h.middlewares("CreateUser"),
		},
	}

	registerRoutes(r, routes, "/user", mws)
}

// middlewares returns individual route middleware per operation id.
// Edit as required.
func (h *User) middlewares(opID string) []gin.HandlerFunc {
	switch opID {
	default:
		return []gin.HandlerFunc{}
	}
}

// CreateUser creates a new user.
func (h *User) CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
