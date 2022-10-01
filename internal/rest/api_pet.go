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
			Name:        string(AddPet),
			Method:      http.MethodPost,
			Pattern:     "/pet",
			HandlerFunc: h.AddPet,
			Middlewares: h.middlewares(AddPet),
		},
		{
			Name:        string(DeletePet),
			Method:      http.MethodDelete,
			Pattern:     "/pet/:petId",
			HandlerFunc: h.DeletePet,
			Middlewares: h.middlewares(DeletePet),
		},
		{
			Name:        string(FindPetsByStatus),
			Method:      http.MethodGet,
			Pattern:     "/pet/findByStatus",
			HandlerFunc: h.FindPetsByStatus,
			Middlewares: h.middlewares(FindPetsByStatus),
		},
		{
			Name:        string(FindPetsByTags),
			Method:      http.MethodGet,
			Pattern:     "/pet/findByTags",
			HandlerFunc: h.FindPetsByTags,
			Middlewares: h.middlewares(FindPetsByTags),
		},
		{
			Name:        string(GetPetById),
			Method:      http.MethodGet,
			Pattern:     "/pet/:petId",
			HandlerFunc: h.GetPetById,
			Middlewares: h.middlewares(GetPetById),
		},
		{
			Name:        string(UpdatePet),
			Method:      http.MethodPut,
			Pattern:     "/pet",
			HandlerFunc: h.UpdatePet,
			Middlewares: h.middlewares(UpdatePet),
		},
		{
			Name:        string(UpdatePetWithForm),
			Method:      http.MethodPost,
			Pattern:     "/pet/:petId",
			HandlerFunc: h.UpdatePetWithForm,
			Middlewares: h.middlewares(UpdatePetWithForm),
		},
		{
			Name:        string(UploadFile),
			Method:      http.MethodPost,
			Pattern:     "/pet/:petId/uploadImage",
			HandlerFunc: h.UploadFile,
			Middlewares: h.middlewares(UploadFile),
		},
	}

	registerRoutes(r, routes, "/pet", mws)
}

// middlewares returns individual route middleware per operation id.
func (h *Pet) middlewares(opID petOpID) []gin.HandlerFunc {
	switch opID {
	default:
		return []gin.HandlerFunc{}
	}
}

// AddPet add a new pet to the store.
func (h *Pet) AddPet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// DeletePet deletes a pet.
func (h *Pet) DeletePet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// FindPetsByStatus finds pets by status.
func (h *Pet) FindPetsByStatus(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// FindPetsByTags finds pets by tags.
// Deprecated
func (h *Pet) FindPetsByTags(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UpdatePet update an existing pet.
func (h *Pet) UpdatePet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UpdatePetWithForm updates a pet in the store with form data.
func (h *Pet) UpdatePetWithForm(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UploadFile uploads an image.
func (h *Pet) UploadFile(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// GetPetById find pet by id.
func (h *Pet) GetPetById(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
