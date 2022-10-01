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
// Generated. DO NOT EDIT.
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
			Name:        string(UpdatePet),
			Method:      http.MethodPut,
			Pattern:     "/pet",
			HandlerFunc: h.UpdatePet,
			Middlewares: h.middlewares(UpdatePet),
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
	fmt.Println("new logic for DeletePet")
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

// UpdatePet update an existing pet.
func (h *Pet) UpdatePet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
