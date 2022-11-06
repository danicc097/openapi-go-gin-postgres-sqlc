// Package rest provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/danicc097/openapi-go-gin-postgres-sqlc version (devel) DO NOT EDIT.
package rest

import (
	"fmt"
	"net/http"

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
	DeleteUser(c *gin.Context, id string)
	// updates the user by id
	// (PUT /user/{id})
	UpdateUser(c *gin.Context, id string)

	middlewares(opID operationID) []gin.HandlerFunc
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc func(c *gin.Context)

// AdminPing operation middleware.
func (siw *ServerInterfaceWrapper) AdminPing(c *gin.Context) {

	c.Set(Bearer_authScopes, []string{""})

	c.Set(Api_keyScopes, []string{""})

	// apply middlewares for operation "AdminPing".
	for _, mw := range siw.Handler.middlewares(AdminPing) {
		mw(c)
	}

	siw.Handler.AdminPing(c)
}

// OpenapiYamlGet operation middleware.
func (siw *ServerInterfaceWrapper) OpenapiYamlGet(c *gin.Context) {

	// apply middlewares for operation "OpenapiYamlGet".
	for _, mw := range siw.Handler.middlewares(OpenapiYamlGet) {
		mw(c)
	}

	siw.Handler.OpenapiYamlGet(c)
}

// Ping operation middleware.
func (siw *ServerInterfaceWrapper) Ping(c *gin.Context) {

	// apply middlewares for operation "Ping".
	for _, mw := range siw.Handler.middlewares(Ping) {
		mw(c)
	}

	siw.Handler.Ping(c)
}

// GetCurrentUser operation middleware.
func (siw *ServerInterfaceWrapper) GetCurrentUser(c *gin.Context) {

	c.Set(Bearer_authScopes, []string{""})

	c.Set(Api_keyScopes, []string{""})

	// apply middlewares for operation "GetCurrentUser".
	for _, mw := range siw.Handler.middlewares(GetCurrentUser) {
		mw(c)
	}

	siw.Handler.GetCurrentUser(c)
}

// DeleteUser operation middleware.
func (siw *ServerInterfaceWrapper) DeleteUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	c.Set(Bearer_authScopes, []string{""})

	c.Set(Api_keyScopes, []string{""})

	// apply middlewares for operation "DeleteUser".
	for _, mw := range siw.Handler.middlewares(DeleteUser) {
		mw(c)
	}

	siw.Handler.DeleteUser(c, id)
}

// UpdateUser operation middleware.
func (siw *ServerInterfaceWrapper) UpdateUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	c.Set(Bearer_authScopes, []string{""})

	c.Set(Api_keyScopes, []string{""})

	// apply middlewares for operation "UpdateUser".
	for _, mw := range siw.Handler.middlewares(UpdateUser) {
		mw(c)
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

	router.PUT(options.BaseURL+"/user/:id", wrapper.UpdateUser)

	return router
}
