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
			Name:        string(addPet),
			Method:      http.MethodPost,
			Pattern:     "/pet",
			HandlerFunc: h.addPet,
			Middlewares: h.middlewares(addPet),
		},
		{
			Name:        string(deletePet),
			Method:      http.MethodDelete,
			Pattern:     "/pet/:petId",
			HandlerFunc: h.deletePet,
			Middlewares: h.middlewares(deletePet),
		},
		{
			Name:        string(findPetsByStatus),
			Method:      http.MethodGet,
			Pattern:     "/pet/findByStatus",
			HandlerFunc: h.findPetsByStatus,
			Middlewares: h.middlewares(findPetsByStatus),
		},
		{
			Name:        string(findPetsByTags),
			Method:      http.MethodGet,
			Pattern:     "/pet/findByTags",
			HandlerFunc: h.findPetsByTags,
			Middlewares: h.middlewares(findPetsByTags),
		},
		{
			Name:        string(getPetById),
			Method:      http.MethodGet,
			Pattern:     "/pet/:petId",
			HandlerFunc: h.getPetById,
			Middlewares: h.middlewares(getPetById),
		},
		{
			Name:        string(updatePet),
			Method:      http.MethodPut,
			Pattern:     "/pet",
			HandlerFunc: h.updatePet,
			Middlewares: h.middlewares(updatePet),
		},
		{
			Name:        string(updatePetWithForm),
			Method:      http.MethodPost,
			Pattern:     "/pet/:petId",
			HandlerFunc: h.updatePetWithForm,
			Middlewares: h.middlewares(updatePetWithForm),
		},
		{
			Name:        string(uploadFile),
			Method:      http.MethodPost,
			Pattern:     "/pet/:petId/uploadImage",
			HandlerFunc: h.uploadFile,
			Middlewares: h.middlewares(uploadFile),
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

// addPet add a new pet to the store.
func (h *Pet) addPet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// deletePet deletes a pet.
func (h *Pet) deletePet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// findPetsByStatus finds pets by status.
func (h *Pet) findPetsByStatus(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// findPetsByTags finds pets by tags.
// Deprecated
func (h *Pet) findPetsByTags(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// updatePet update an existing pet.
func (h *Pet) updatePet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// updatePetWithForm updates a pet in the store with form data.
func (h *Pet) updatePetWithForm(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// uploadFile uploads an image.
func (h *Pet) uploadFile(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// getPetById find pet by id.
func (h *Pet) getPetById(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
