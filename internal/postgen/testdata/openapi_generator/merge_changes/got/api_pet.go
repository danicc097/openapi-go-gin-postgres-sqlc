package rest

import (
	"fmt"
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
			Name:        string(updatePet),
			Method:      http.MethodPut,
			Pattern:     "/pet",
			HandlerFunc: h.updatePet,
			Middlewares: h.middlewares(updatePet),
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
	fmt.Println("new logic for deletePet")
	c.JSON(http.StatusOK, gin.H{})
}

// I added some important comments here

/*
and here as well */

// UpdatePet was deleted for some reason

// newFunction was added by hand.
// This shouldn't be overridden/deleted in any case.
func (h *Pet) newFunction(c *gin.Context) {
	fmt.Println("this is some random helper newFunction")
}

// updatePet update an existing pet.
func (h *Pet) updatePet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
