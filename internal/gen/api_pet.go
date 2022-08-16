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

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (t *Pet) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "AddPet",
			Method:      http.MethodPost,
			Pattern:     "/pet",
			HandlerFunc: t.AddPet,
			Middlewares: t.middlewares("AddPet"),
		},
		{
			Name:        "DeletePet",
			Method:      http.MethodDelete,
			Pattern:     "/pet/:petId",
			HandlerFunc: t.DeletePet,
			Middlewares: t.middlewares("DeletePet"),
		},
		{
			Name:        "FindPetsByStatus",
			Method:      http.MethodGet,
			Pattern:     "/pet/findByStatus",
			HandlerFunc: t.FindPetsByStatus,
			Middlewares: t.middlewares("FindPetsByStatus"),
		},
		{
			Name:        "FindPetsByTags",
			Method:      http.MethodGet,
			Pattern:     "/pet/findByTags",
			HandlerFunc: t.FindPetsByTags,
			Middlewares: t.middlewares("FindPetsByTags"),
		},
		{
			Name:        "GetPetById",
			Method:      http.MethodGet,
			Pattern:     "/pet/:petId",
			HandlerFunc: t.GetPetById,
			Middlewares: t.middlewares("GetPetById"),
		},
		{
			Name:        "UpdatePet",
			Method:      http.MethodPut,
			Pattern:     "/pet",
			HandlerFunc: t.UpdatePet,
			Middlewares: t.middlewares("UpdatePet"),
		},
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

// AddPet add a new pet to the store.
func (t *Pet) AddPet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// DeletePet deletes a pet.
func (t *Pet) DeletePet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// FindPetsByStatus finds pets by status.
func (t *Pet) FindPetsByStatus(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// FindPetsByTags finds pets by tags.
// Deprecated
func (t *Pet) FindPetsByTags(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// GetPetById find pet by id.
func (t *Pet) GetPetById(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UpdatePet update an existing pet.
func (t *Pet) UpdatePet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UpdatePetWithForm updates a pet in the store with form data.
func (t *Pet) UpdatePetWithForm(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UploadFile uploads an image.
func (t *Pet) UploadFile(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
