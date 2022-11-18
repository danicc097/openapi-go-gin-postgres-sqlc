// Package rest provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/danicc097/openapi-go-gin-postgres-sqlc version (devel) DO NOT EDIT.
package rest

import (
	"fmt"
	"net/http"

	externalRef0 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Ping pongs
	// (GET /admin/ping)
	AdminPing(c *gin.Context)
	// Returns this very OpenAPI spec.
	// (GET /openapi.yaml)
	OpenapiYamlGet(c *gin.Context)
	// Ping pongs
	// (GET /ping)
	Ping(c *gin.Context)
	// returns the logged in user
	// (GET /user/me)
	GetCurrentUser(c *gin.Context)
	// deletes the user by id
	// (DELETE /user/{id})
	DeleteUser(c *gin.Context, id externalRef0.UserID)
	// updates the user by id
	// (PATCH /user/{id})
	UpdateUser(c *gin.Context, id externalRef0.UserID)

	middlewares(opID operationID) []gin.HandlerFunc
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc func(c *gin.Context)

// AdminPing operation with its own middleware.
func (siw *ServerInterfaceWrapper) AdminPing(c *gin.Context) {

	c.Set(externalRef0.Bearer_authScopes, []string{""})

	c.Set(externalRef0.Api_keyScopes, []string{""})

	// apply middlewares for operation "AdminPing".
	for _, mw := range siw.Handler.middlewares(AdminPing) {
		mw(c)

		// should actually call router.<Method> with a slice of mw, last item the actual handler
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AdminPing(c)
}

// OpenapiYamlGet operation with its own middleware.
func (siw *ServerInterfaceWrapper) OpenapiYamlGet(c *gin.Context) {

	// apply middlewares for operation "OpenapiYamlGet".
	for _, mw := range siw.Handler.middlewares(OpenapiYamlGet) {
		mw(c)

		// should actually call router.<Method> with a slice of mw, last item the actual handler
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.OpenapiYamlGet(c)
}

// Ping operation with its own middleware.
func (siw *ServerInterfaceWrapper) Ping(c *gin.Context) {

	// apply middlewares for operation "Ping".
	for _, mw := range siw.Handler.middlewares(Ping) {
		mw(c)

		// should actually call router.<Method> with a slice of mw, last item the actual handler
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.Ping(c)
}

// GetCurrentUser operation with its own middleware.
func (siw *ServerInterfaceWrapper) GetCurrentUser(c *gin.Context) {

	c.Set(externalRef0.Bearer_authScopes, []string{""})

	c.Set(externalRef0.Api_keyScopes, []string{""})

	// apply middlewares for operation "GetCurrentUser".
	for _, mw := range siw.Handler.middlewares(GetCurrentUser) {
		mw(c)

		// should actually call router.<Method> with a slice of mw, last item the actual handler
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetCurrentUser(c)
}

// DeleteUser operation with its own middleware.
func (siw *ServerInterfaceWrapper) DeleteUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id externalRef0.UserID

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{""})

	c.Set(externalRef0.Api_keyScopes, []string{""})

	// apply middlewares for operation "DeleteUser".
	for _, mw := range siw.Handler.middlewares(DeleteUser) {
		mw(c)

		// should actually call router.<Method> with a slice of mw, last item the actual handler
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteUser(c, id)
}

// UpdateUser operation with its own middleware.
func (siw *ServerInterfaceWrapper) UpdateUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id externalRef0.UserID

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{""})

	c.Set(externalRef0.Api_keyScopes, []string{""})

	// apply middlewares for operation "UpdateUser".
	for _, mw := range siw.Handler.middlewares(UpdateUser) {
		mw(c)

		// should actually call router.<Method> with a slice of mw, last item the actual handler
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateUser(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL string
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.RouterGroup, si ServerInterface) *gin.RouterGroup {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.RouterGroup, si ServerInterface, options GinServerOptions) *gin.RouterGroup {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(options.BaseURL+"/admin/ping", wrapper.AdminPing)

	router.GET(options.BaseURL+"/openapi.yaml", wrapper.OpenapiYamlGet)

	router.GET(options.BaseURL+"/ping", wrapper.Ping)

	router.GET(options.BaseURL+"/user/me", wrapper.GetCurrentUser)

	router.DELETE(options.BaseURL+"/user/:id", wrapper.DeleteUser)

	router.PATCH(options.BaseURL+"/user/:id", wrapper.UpdateUser)

	return router
}
