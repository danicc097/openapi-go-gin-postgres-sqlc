package handlers

import (
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// some struct added later. Shouldn't be removed
type PetThing struct {
}

// Pet handles routes with the pet tag.
type Pet struct {
	svc services.Pet
	// add necessary services, etc. as required
}

// NewPet returns a new handler for pet.
// Edit as required.
func NewPet(svc services.Pet) *Pet {
	return &Pet{
		svc: svc,
	}
}

// Register connects the handlers to a router with the given middleware.
// GENERATED METHOD. Only Middlewares will be saved between runs.
func (t *Pet) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "UpdatePetWithForm",
			Method:      http.MethodPost,
			Pattern:     "/pet/:petId",
			HandlerFunc: t.UpdatePetWithForm,
			// added middleware, would not want to lose it.
			Middlewares: []gin.HandlerFunc{rest.AuthMiddleware()},
		},
		// this is a new handler added by hand.
		// This will be overriden by generated routes.
		{
			Name:        "NewHandlerGet",
			Method:      http.MethodGet,
			Pattern:     "/pet/:petId/NewHandlerGet",
			HandlerFunc: t.NewHandlerGet,
			Middlewares: []gin.HandlerFunc{},
		},
		// UploadFile was deleted for some reason
	}

	rest.RegisterRoutes(r, routes, "/pet", mws)
}

// I added some important comments here

/*
and here as well */

// UpdatePetWithForm updates a pet in the store with form data.
func (t *Pet) UpdatePetWithForm(c *gin.Context) {
	fmt.Println("would have run logic for UpdatePetWithForm")
	c.JSON(http.StatusOK, gin.H{})
}

// UploadFile was deleted for some reason

// newFunction was added by hand.
// This shouldn't be overriden/deleted in any case.
func (t *Pet) newFunction(c *gin.Context) {
	fmt.Println("this is some random helper newFunction")
}

// NewHandlerPost is a newly generated handler.
func (t *Pet) NewHandlerPost(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UploadFile uploads an image.
func (t *Pet) UploadFile(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
