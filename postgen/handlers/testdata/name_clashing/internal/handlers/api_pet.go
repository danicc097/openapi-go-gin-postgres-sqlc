package handlers

import (
	"fmt"
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
			Name:        "UpdatePet",
			Method:      http.MethodPut,
			Pattern:     "/pet",
			HandlerFunc: t.UpdatePet,
			Middlewares: t.middlewares("UpdatePet"),
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
	fmt.Println("new logic for DeletePet")
	fmt.Println("new logic for DeletePet")
	fmt.Println("new logic for DeletePet")
	c.JSON(http.StatusOK, gin.H{})
}

// I added some important comments here

/*
and here as well */

// UpdatePet was deleted for some reason

// newFunction was added by hand.
// This shouldn't be overriden/deleted in any case.
func (t *Pet) newFunction(c *gin.Context) {
	fmt.Println("this is some random helper newFunction")
}

// ConflictEndpointPet will clash with a generated operation id.
func (t *Pet) ConflictEndpointPet(param1 string, param2 string) {
	fmt.Println("this method will clash with a generated operation id")
}
