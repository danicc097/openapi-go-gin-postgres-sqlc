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

// middlewares returns individual route middleware per operation id.
// Edit as required.
func (t *Pet) middlewares(opId string) []gin.HandlerFunc {
	switch opId {
	case "UploadFile":
		return []gin.HandlerFunc{rest.NewAuthMiddleware(t.svc.Logger).EnsureAuthenticated()}
	default:
		return []gin.HandlerFunc{}
	}
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
