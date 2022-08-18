package handlers

import (
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// Pet handles routes with the 'pet' tag.
type Pet struct {
	svc services.Pet
	// add or remove services, etc. as required
}

// NewPet returns a new handler for the 'pet' route group.
// Edit as required.
func NewPet(svc services.Pet) *Pet {
	return &Pet{
		svc: svc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *Pet) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{}
	rest.RegisterRoutes(r, routes, "/pet", mws)
}

// middlewares returns individual route middleware per operation id.
// Edit as required.
func (h *Pet) middlewares(opId string) []gin.HandlerFunc {
	switch opId {
	default:
		return []gin.HandlerFunc{}
	}
}

// ConflictEndpointPet will clash with a generated operation id.
func (h *Pet) ConflictEndpointPet(param1 string, param2 string) {
	fmt.Println("this method will clash with a generated operation id")
}
