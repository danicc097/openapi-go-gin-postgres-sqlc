package handlers

import (
	"net/http"

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

// Register connects the handlers to a router with the given middleware.
// Generated method. DO NOT EDIT.
func (t *Pet) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "UpdatePetWithForm",
			Method:      http.MethodPost,
			Pattern:     "/pet/:petId",
			HandlerFunc: t.UpdatePetWithForm,
			Middlewares: t.middlewares("UpdatePetWithForm"),
		},
		{
			Name:        "UploadFile",
			Method:      http.MethodPost,
			Pattern:     "/pet/:petId/uploadImage",
			HandlerFunc: t.UploadFile,
			Middlewares: t.middlewares("UploadFile"),
		},
	}
	rest.RegisterRoutes(r, routes, "/pet", mws)
}

// middlewares returns individual route middleware per operation id.
// Edit as required.
func (t *Pet) middlewares(opId string) []gin.HandlerFunc {
	switch opId {
	default:
		return []gin.HandlerFunc{}
	}
}

// UpdatePetWithForm updates a pet in the store with form data.
func (t *Pet) UpdatePetWithForm(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UploadFile uploads an image.
func (t *Pet) UploadFile(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
