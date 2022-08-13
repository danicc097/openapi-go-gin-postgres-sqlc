package handlers

import (
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// Pet handles routes with the pet tag.
type Pet struct {
	svc services.Pet
	// add necessary services, etc. as required
}

// NewPet returns a new handler for pet.
// Edit as required
// TODO rewriting handler methods based on current postgen:
// see https://eli.thegreenplace.net/2021/rewriting-go-source-code-with-ast-tooling/
// simpler solutions based on drawbacks (complicated, comments not attached to nodes):
// - https://github.com/dave/dst
// - https://github.com/uber-go/gopatch
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
		{
			Name:        "UploadFile",
			Method:      http.MethodPost,
			Pattern:     "/pet/:petId/uploadImage",
			HandlerFunc: t.UploadFile,
			Middlewares: []gin.HandlerFunc{},
		},
		// this is a new handler added by hand.
		// I wouldnt care that much if this comment is deleted.
		// Order is not important
		{
			Name:        "NewHandlerGet",
			Method:      http.MethodGet,
			Pattern:     "/pet/:petId/NewHandlerGet",
			HandlerFunc: t.NewHandlerGet,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "NewHandlerPost",
			Method:      http.MethodPost,
			Pattern:     "/pet/:petId/NewHandlerPost",
			HandlerFunc: t.NewHandlerPost,
			Middlewares: []gin.HandlerFunc{},
		},
	}

	rest.RegisterRoutes(r, routes, "/pet", mws)
}

// I added some important comments here

/*
and here as well */

// AddPet add a new pet to the store.
func (t *Pet) AddPet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeletePet deletes a pet.
func (t *Pet) DeletePet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindPetsByStatus finds pets by status.
func (t *Pet) FindPetsByStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FindPetsByTags finds pets by tags.
// Deprecated
func (t *Pet) FindPetsByTags(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetPetById find pet by id.
func (t *Pet) GetPetById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdatePet update an existing pet.
func (t *Pet) UpdatePet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdatePetWithForm updates a pet in the store with form data.
func (t *Pet) UpdatePetWithForm(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UploadFile uploads an image.
func (t *Pet) UploadFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// AddPet add a new pet to the store.
func (t *Pet) NewHandlerPost(c *gin.Context) {
	fmt.Println("this is the implementation for NewHandlerPost")
}

// NewHandlerGet was added by hand.
// This shouldn't be overriden/deleted in any case.
func (t *Pet) NewHandlerGet(c *gin.Context) {
	fmt.Println("this is the implementation for NewHandlerGet")
}

// This is an unused method. should not be deleted.
func (t *Pet) AnUnusedHandler(c *gin.Context) {
	fmt.Println("this is the implementation for anUnusedHandler not used by any route")
}
