package rest

import (
	"net/http"

	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// Pet handles routes with the 'pet' tag.
type Pet struct {
	svc services.Pet
	// add or remove services, etc. as required
}

// NewPet returns a new handler for the 'pet' route group.
func NewPet(svc services.Pet) *Pet {
	return &Pet{
		svc: svc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *Pet) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []route{
		{
			Name:        "ConflictEndpointPet",
			Method:      http.MethodGet,
			Pattern:     "/pet/ConflictEndpointPet",
			HandlerFunc: h.ConflictEndpointPet,
			Middlewares: h.middlewares("ConflictEndpointPet"),
		},
	}

	registerRoutes(r, routes, "/pet", mws)
}

// middlewares returns individual route middleware per operation id.
func (h *Pet) middlewares(opID string) []gin.HandlerFunc {
	switch opID {
	default:
		return []gin.HandlerFunc{}
	}
}

// ConflictEndpointPet name clashing test.
func (h *Pet) ConflictEndpointPet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
