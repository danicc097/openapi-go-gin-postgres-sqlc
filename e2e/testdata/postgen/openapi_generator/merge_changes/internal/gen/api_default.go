package handlers

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// Default handles routes with the 'Default_' tag.
type Default struct {
	svc services.Default
	// add or remove services, etc. as required
}

// NewDefault returns a new handler for the 'Default_' route group.
// Edit as required.
func NewDefault(svc services.Default) *Default {
	return &Default{
		svc: svc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *Default) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "Ping",
			Method:      http.MethodGet,
			Pattern:     "/ping",
			HandlerFunc: h.Ping,
			Middlewares: h.middlewares("Ping"),
		},
	}
	rest.RegisterRoutes(r, routes, "/default", mws)
}

// middlewares returns individual route middleware per operation id.
// Edit as required.
func (h *Default) middlewares(opID string) []gin.HandlerFunc {
	switch opID {
	default:
		return []gin.HandlerFunc{}
	}
}

// Ping ping pongs.
func (h *Default) Ping(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
