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
			Name:        string(createUser),
			Method:      http.MethodPost,
			Pattern:     "/user",
			HandlerFunc: h.createUser,
			Middlewares: h.middlewares(createUser),
		},
	}

	registerRoutes(r, routes, "/user", mws)
}

// middlewares returns individual route middleware per operation id.
func (h *User) middlewares(opID userOpID) []gin.HandlerFunc {
	switch opID {
	default:
		return []gin.HandlerFunc{}
	}
}

// createUser creates a new user.
func (h *User) createUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
