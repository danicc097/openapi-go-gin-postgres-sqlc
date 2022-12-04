// Package rest provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
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

	// (GET /events)
	Events(c *gin.Context)
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
	// updates user role and scopes by id
	// (PATCH /user/{id}/authorization)
	UpdateUserAuthorization(c *gin.Context, id externalRef0.UserID)

	middlewares(opID OperationID) []gin.HandlerFunc
	authMiddlewares(opID OperationID) []gin.HandlerFunc
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

	siw.Handler.AdminPing(c)
}

// Events operation with its own middleware.
func (siw *ServerInterfaceWrapper) Events(c *gin.Context) {

	siw.Handler.Events(c)
}

// OpenapiYamlGet operation with its own middleware.
func (siw *ServerInterfaceWrapper) OpenapiYamlGet(c *gin.Context) {

	siw.Handler.OpenapiYamlGet(c)
}

// Ping operation with its own middleware.
func (siw *ServerInterfaceWrapper) Ping(c *gin.Context) {

	siw.Handler.Ping(c)
}

// GetCurrentUser operation with its own middleware.
func (siw *ServerInterfaceWrapper) GetCurrentUser(c *gin.Context) {

	c.Set(externalRef0.Bearer_authScopes, []string{""})

	c.Set(externalRef0.Api_keyScopes, []string{""})

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

	siw.Handler.UpdateUser(c, id)
}

// UpdateUserAuthorization operation with its own middleware.
func (siw *ServerInterfaceWrapper) UpdateUserAuthorization(c *gin.Context) {

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

	siw.Handler.UpdateUserAuthorization(c, id)
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

	// calling mw(c) directly has unexpected consequences: closed channels, etc.
	router.GET(options.BaseURL+"/admin/ping", append(
		wrapper.Handler.authMiddlewares(AdminPing),
		append(wrapper.Handler.middlewares(AdminPing), wrapper.AdminPing)...,
	)...)

	router.GET(options.BaseURL+"/events", append(
		wrapper.Handler.authMiddlewares(Events),
		append(wrapper.Handler.middlewares(Events), wrapper.Events)...,
	)...)

	router.GET(options.BaseURL+"/openapi.yaml", append(
		wrapper.Handler.authMiddlewares(OpenapiYamlGet),
		append(wrapper.Handler.middlewares(OpenapiYamlGet), wrapper.OpenapiYamlGet)...,
	)...)

	router.GET(options.BaseURL+"/ping", append(
		wrapper.Handler.authMiddlewares(Ping),
		append(wrapper.Handler.middlewares(Ping), wrapper.Ping)...,
	)...)

	router.GET(options.BaseURL+"/user/me", append(
		wrapper.Handler.authMiddlewares(GetCurrentUser),
		append(wrapper.Handler.middlewares(GetCurrentUser), wrapper.GetCurrentUser)...,
	)...)

	router.DELETE(options.BaseURL+"/user/:id", append(
		wrapper.Handler.authMiddlewares(DeleteUser),
		append(wrapper.Handler.middlewares(DeleteUser), wrapper.DeleteUser)...,
	)...)

	router.PATCH(options.BaseURL+"/user/:id", append(
		wrapper.Handler.authMiddlewares(UpdateUser),
		append(wrapper.Handler.middlewares(UpdateUser), wrapper.UpdateUser)...,
	)...)

	router.PATCH(options.BaseURL+"/user/:id/authorization", append(
		wrapper.Handler.authMiddlewares(UpdateUserAuthorization),
		append(wrapper.Handler.middlewares(UpdateUserAuthorization), wrapper.UpdateUserAuthorization)...,
	)...)

	return router
}
