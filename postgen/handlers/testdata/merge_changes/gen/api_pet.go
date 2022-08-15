package handlers

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// Pet handles routes with the pet tag.
type Pet struct {
	svc services.Pet
	// add or remove services, etc. as required
}

// NewPet returns a new handler for pet.
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
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "NewHandlerPost",
			Method:      http.MethodPost,
			Pattern:     "/pet/:petId/NewHandlerPost",
			HandlerFunc: t.NewHandlerPost,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "UploadFile",
			Method:      http.MethodPost,
			Pattern:     "/pet/:petId/uploadImage",
			HandlerFunc: t.UploadFile,
			Middlewares: []gin.HandlerFunc{},
		},
	}

	rest.RegisterRoutes(r, routes, "/pet", mws)
}

// UpdatePetWithForm updates a pet in the store with form data.
func (t *Pet) UpdatePetWithForm(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UploadFile uploads an image.
func (t *Pet) UploadFile(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// NewHandlerPost is a newly generated handler.
func (t *Pet) NewHandlerPost(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
