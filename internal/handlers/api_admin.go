package handlers

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/gin-gonic/gin"
)

// Admin handles routes with the 'admin' tag.
type Admin struct {
	svc UserService
	// add or remove services, etc. as required
}

// NewAdmin returns a new handler for the 'admin' route group.
// Edit as required.
func NewAdmin(svc UserService) *Admin {
	return &Admin{
		svc: svc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *Admin) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "AdminPing",
			Method:      http.MethodGet,
			Pattern:     "/admin/ping",
			HandlerFunc: h.AdminPing,
			Middlewares: h.middlewares("AdminPing"),
		},
	}

	rest.RegisterRoutes(r, routes, "/admin", mws)
}

// middlewares returns individual route middleware per operation id.
// Edit as required.
func (h *Admin) middlewares(opID string) []gin.HandlerFunc {
	switch opID {
	default:
		return []gin.HandlerFunc{}
	}
}

// AdminPing ping pongs.
func (h *Admin) AdminPing(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
