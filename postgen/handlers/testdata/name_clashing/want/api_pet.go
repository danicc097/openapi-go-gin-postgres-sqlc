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

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (t *Pet) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "UpdatePetWithForm",
			Method:      http.MethodPost,
			Pattern:     "/pet/:petId",
			HandlerFunc: t.UpdatePetWithForm,
			Middlewares: []gin.HandlerFunc{rest.AuthMiddleware()},
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

// NewHandlerPost is a newly generated handler.
func (t *Pet) NewHandlerPost(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UploadFile uploads an image.
func (t *Pet) UploadFile(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
