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
	uuid "github.com/google/uuid"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Ping pongs
	// (GET /admin/ping)
	AdminPing(c *gin.Context)

	// (GET /auth/myprovider/callback)
	MyProviderCallback(c *gin.Context)

	// (GET /auth/myprovider/login)
	MyProviderLogin(c *gin.Context)

	// (GET /events)
	Events(c *gin.Context, params externalRef0.EventsParams)
	// Get paginated user notifications
	// (GET /notifications/user/page)
	GetPaginatedNotifications(c *gin.Context, params externalRef0.GetPaginatedNotificationsParams)
	// Returns this very OpenAPI spec.
	// (GET /openapi.yaml)
	OpenapiYamlGet(c *gin.Context)
	// Ping pongs
	// (GET /ping)
	Ping(c *gin.Context)
	// returns board data for a project
	// (GET /project/{projectName}/)
	GetProject(c *gin.Context, projectName externalRef0.ProjectName)
	// returns board data for a project
	// (GET /project/{projectName}/board)
	GetProjectBoard(c *gin.Context, projectName externalRef0.ProjectName)
	// returns the project configuration
	// (GET /project/{projectName}/config)
	GetProjectConfig(c *gin.Context, projectName externalRef0.ProjectName)
	// updates the project configuration
	// (PUT /project/{projectName}/config)
	UpdateProjectConfig(c *gin.Context, projectName externalRef0.ProjectName)
	// creates initial data (teams, tags...) for a new project
	// (POST /project/{projectName}/initialize)
	InitializeProject(c *gin.Context, projectName externalRef0.ProjectName)
	// create workitem tag
	// (POST /project/{projectName}/tag/)
	CreateWorkitemTag(c *gin.Context, projectName externalRef0.ProjectName)
	// returns workitems for a project
	// (GET /project/{projectName}/workitems)
	GetProjectWorkitems(c *gin.Context, projectName externalRef0.ProjectName, params externalRef0.GetProjectWorkitemsParams)
	// returns the logged in user
	// (GET /user/me)
	GetCurrentUser(c *gin.Context)
	// deletes the user by id
	// (DELETE /user/{id})
	DeleteUser(c *gin.Context, id uuid.UUID)
	// updates the user by id
	// (PATCH /user/{id})
	UpdateUser(c *gin.Context, id uuid.UUID)
	// updates user role and scopes by id
	// (PATCH /user/{id}/authorization)
	UpdateUserAuthorization(c *gin.Context, id uuid.UUID)
	// create workitem
	// (POST /workitem/)
	CreateWorkitem(c *gin.Context)
	// delete workitem
	// (DELETE /workitem/{id}/)
	DeleteWorkitem(c *gin.Context, id externalRef0.Serial)
	// get workitem
	// (GET /workitem/{id}/)
	GetWorkitem(c *gin.Context, id externalRef0.Serial)
	// update workitem
	// (PATCH /workitem/{id}/)
	UpdateWorkitem(c *gin.Context, id externalRef0.Serial)
	// create workitem comment
	// (POST /workitem/{id}/comments/)
	CreateWorkitemComment(c *gin.Context, id externalRef0.Serial)

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

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.AdminPing(c)
}

// MyProviderCallback operation with its own middleware.
func (siw *ServerInterfaceWrapper) MyProviderCallback(c *gin.Context) {

	siw.Handler.MyProviderCallback(c)
}

// MyProviderLogin operation with its own middleware.
func (siw *ServerInterfaceWrapper) MyProviderLogin(c *gin.Context) {

	siw.Handler.MyProviderLogin(c)
}

// Events operation with its own middleware.
func (siw *ServerInterfaceWrapper) Events(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params externalRef0.EventsParams

	// ------------- Required query parameter "projectName" -------------

	if paramValue := c.Query("projectName"); paramValue != "" {

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Query argument projectName is required, but not found"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "projectName", c.Request.URL.Query(), &params.ProjectName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter projectName: %s", err)})
		return
	}

	siw.Handler.Events(c, params)
}

// GetPaginatedNotifications operation with its own middleware.
func (siw *ServerInterfaceWrapper) GetPaginatedNotifications(c *gin.Context) {

	var err error

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params externalRef0.GetPaginatedNotificationsParams

	// ------------- Required query parameter "limit" -------------

	if paramValue := c.Query("limit"); paramValue != "" {

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Query argument limit is required, but not found"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", c.Request.URL.Query(), &params.Limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter limit: %s", err)})
		return
	}

	// ------------- Required query parameter "direction" -------------

	if paramValue := c.Query("direction"); paramValue != "" {

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Query argument direction is required, but not found"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "direction", c.Request.URL.Query(), &params.Direction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter direction: %s", err)})
		return
	}

	// ------------- Required query parameter "cursor" -------------

	if paramValue := c.Query("cursor"); paramValue != "" {

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Query argument cursor is required, but not found"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "cursor", c.Request.URL.Query(), &params.Cursor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter cursor: %s", err)})
		return
	}

	siw.Handler.GetPaginatedNotifications(c, params)
}

// OpenapiYamlGet operation with its own middleware.
func (siw *ServerInterfaceWrapper) OpenapiYamlGet(c *gin.Context) {

	siw.Handler.OpenapiYamlGet(c)
}

// Ping operation with its own middleware.
func (siw *ServerInterfaceWrapper) Ping(c *gin.Context) {

	siw.Handler.Ping(c)
}

// GetProject operation with its own middleware.
func (siw *ServerInterfaceWrapper) GetProject(c *gin.Context) {

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName externalRef0.ProjectName

	err = runtime.BindStyledParameter("simple", false, "projectName", c.Param("projectName"), &projectName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter projectName: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.GetProject(c, projectName)
}

// GetProjectBoard operation with its own middleware.
func (siw *ServerInterfaceWrapper) GetProjectBoard(c *gin.Context) {

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName externalRef0.ProjectName

	err = runtime.BindStyledParameter("simple", false, "projectName", c.Param("projectName"), &projectName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter projectName: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.GetProjectBoard(c, projectName)
}

// GetProjectConfig operation with its own middleware.
func (siw *ServerInterfaceWrapper) GetProjectConfig(c *gin.Context) {

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName externalRef0.ProjectName

	err = runtime.BindStyledParameter("simple", false, "projectName", c.Param("projectName"), &projectName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter projectName: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.GetProjectConfig(c, projectName)
}

// UpdateProjectConfig operation with its own middleware.
func (siw *ServerInterfaceWrapper) UpdateProjectConfig(c *gin.Context) {

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName externalRef0.ProjectName

	err = runtime.BindStyledParameter("simple", false, "projectName", c.Param("projectName"), &projectName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter projectName: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.UpdateProjectConfig(c, projectName)
}

// InitializeProject operation with its own middleware.
func (siw *ServerInterfaceWrapper) InitializeProject(c *gin.Context) {

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName externalRef0.ProjectName

	err = runtime.BindStyledParameter("simple", false, "projectName", c.Param("projectName"), &projectName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter projectName: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.InitializeProject(c, projectName)
}

// CreateWorkitemTag operation with its own middleware.
func (siw *ServerInterfaceWrapper) CreateWorkitemTag(c *gin.Context) {

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName externalRef0.ProjectName

	err = runtime.BindStyledParameter("simple", false, "projectName", c.Param("projectName"), &projectName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter projectName: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.CreateWorkitemTag(c, projectName)
}

// GetProjectWorkitems operation with its own middleware.
func (siw *ServerInterfaceWrapper) GetProjectWorkitems(c *gin.Context) {

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName externalRef0.ProjectName

	err = runtime.BindStyledParameter("simple", false, "projectName", c.Param("projectName"), &projectName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter projectName: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params externalRef0.GetProjectWorkitemsParams

	// ------------- Optional query parameter "open" -------------

	err = runtime.BindQueryParameter("form", true, false, "open", c.Request.URL.Query(), &params.Open)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter open: %s", err)})
		return
	}

	// ------------- Optional query parameter "deleted" -------------

	err = runtime.BindQueryParameter("form", true, false, "deleted", c.Request.URL.Query(), &params.Deleted)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter deleted: %s", err)})
		return
	}

	siw.Handler.GetProjectWorkitems(c, projectName, params)
}

// GetCurrentUser operation with its own middleware.
func (siw *ServerInterfaceWrapper) GetCurrentUser(c *gin.Context) {

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.GetCurrentUser(c)
}

// DeleteUser operation with its own middleware.
func (siw *ServerInterfaceWrapper) DeleteUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uuid.UUID

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.DeleteUser(c, id)
}

// UpdateUser operation with its own middleware.
func (siw *ServerInterfaceWrapper) UpdateUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uuid.UUID

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.UpdateUser(c, id)
}

// UpdateUserAuthorization operation with its own middleware.
func (siw *ServerInterfaceWrapper) UpdateUserAuthorization(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uuid.UUID

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.UpdateUserAuthorization(c, id)
}

// CreateWorkitem operation with its own middleware.
func (siw *ServerInterfaceWrapper) CreateWorkitem(c *gin.Context) {

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.CreateWorkitem(c)
}

// DeleteWorkitem operation with its own middleware.
func (siw *ServerInterfaceWrapper) DeleteWorkitem(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id externalRef0.Serial

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.DeleteWorkitem(c, id)
}

// GetWorkitem operation with its own middleware.
func (siw *ServerInterfaceWrapper) GetWorkitem(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id externalRef0.Serial

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.GetWorkitem(c, id)
}

// UpdateWorkitem operation with its own middleware.
func (siw *ServerInterfaceWrapper) UpdateWorkitem(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id externalRef0.Serial

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.UpdateWorkitem(c, id)
}

// CreateWorkitemComment operation with its own middleware.
func (siw *ServerInterfaceWrapper) CreateWorkitemComment(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id externalRef0.Serial

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	c.Set(externalRef0.Bearer_authScopes, []string{})

	c.Set(externalRef0.Api_keyScopes, []string{})

	siw.Handler.CreateWorkitemComment(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL string
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	// calling mw(c) directly has unexpected consequences: closed channels, etc.
	router.GET(options.BaseURL+"/admin/ping", append(
		wrapper.Handler.authMiddlewares(AdminPing),
		append(wrapper.Handler.middlewares(AdminPing), wrapper.AdminPing)...,
	)...)

	router.GET(options.BaseURL+"/auth/myprovider/callback", append(
		wrapper.Handler.authMiddlewares(MyProviderCallback),
		append(wrapper.Handler.middlewares(MyProviderCallback), wrapper.MyProviderCallback)...,
	)...)

	router.GET(options.BaseURL+"/auth/myprovider/login", append(
		wrapper.Handler.authMiddlewares(MyProviderLogin),
		append(wrapper.Handler.middlewares(MyProviderLogin), wrapper.MyProviderLogin)...,
	)...)

	router.GET(options.BaseURL+"/events", append(
		wrapper.Handler.authMiddlewares(Events),
		append(wrapper.Handler.middlewares(Events), wrapper.Events)...,
	)...)

	router.GET(options.BaseURL+"/notifications/user/page", append(
		wrapper.Handler.authMiddlewares(GetPaginatedNotifications),
		append(wrapper.Handler.middlewares(GetPaginatedNotifications), wrapper.GetPaginatedNotifications)...,
	)...)

	router.GET(options.BaseURL+"/openapi.yaml", append(
		wrapper.Handler.authMiddlewares(OpenapiYamlGet),
		append(wrapper.Handler.middlewares(OpenapiYamlGet), wrapper.OpenapiYamlGet)...,
	)...)

	router.GET(options.BaseURL+"/ping", append(
		wrapper.Handler.authMiddlewares(Ping),
		append(wrapper.Handler.middlewares(Ping), wrapper.Ping)...,
	)...)

	router.GET(options.BaseURL+"/project/:projectName/", append(
		wrapper.Handler.authMiddlewares(GetProject),
		append(wrapper.Handler.middlewares(GetProject), wrapper.GetProject)...,
	)...)

	router.GET(options.BaseURL+"/project/:projectName/board", append(
		wrapper.Handler.authMiddlewares(GetProjectBoard),
		append(wrapper.Handler.middlewares(GetProjectBoard), wrapper.GetProjectBoard)...,
	)...)

	router.GET(options.BaseURL+"/project/:projectName/config", append(
		wrapper.Handler.authMiddlewares(GetProjectConfig),
		append(wrapper.Handler.middlewares(GetProjectConfig), wrapper.GetProjectConfig)...,
	)...)

	router.PUT(options.BaseURL+"/project/:projectName/config", append(
		wrapper.Handler.authMiddlewares(UpdateProjectConfig),
		append(wrapper.Handler.middlewares(UpdateProjectConfig), wrapper.UpdateProjectConfig)...,
	)...)

	router.POST(options.BaseURL+"/project/:projectName/initialize", append(
		wrapper.Handler.authMiddlewares(InitializeProject),
		append(wrapper.Handler.middlewares(InitializeProject), wrapper.InitializeProject)...,
	)...)

	router.POST(options.BaseURL+"/project/:projectName/tag/", append(
		wrapper.Handler.authMiddlewares(CreateWorkitemTag),
		append(wrapper.Handler.middlewares(CreateWorkitemTag), wrapper.CreateWorkitemTag)...,
	)...)

	router.GET(options.BaseURL+"/project/:projectName/workitems", append(
		wrapper.Handler.authMiddlewares(GetProjectWorkitems),
		append(wrapper.Handler.middlewares(GetProjectWorkitems), wrapper.GetProjectWorkitems)...,
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

	router.POST(options.BaseURL+"/workitem/", append(
		wrapper.Handler.authMiddlewares(CreateWorkitem),
		append(wrapper.Handler.middlewares(CreateWorkitem), wrapper.CreateWorkitem)...,
	)...)

	router.DELETE(options.BaseURL+"/workitem/:id/", append(
		wrapper.Handler.authMiddlewares(DeleteWorkitem),
		append(wrapper.Handler.middlewares(DeleteWorkitem), wrapper.DeleteWorkitem)...,
	)...)

	router.GET(options.BaseURL+"/workitem/:id/", append(
		wrapper.Handler.authMiddlewares(GetWorkitem),
		append(wrapper.Handler.middlewares(GetWorkitem), wrapper.GetWorkitem)...,
	)...)

	router.PATCH(options.BaseURL+"/workitem/:id/", append(
		wrapper.Handler.authMiddlewares(UpdateWorkitem),
		append(wrapper.Handler.middlewares(UpdateWorkitem), wrapper.UpdateWorkitem)...,
	)...)

	router.POST(options.BaseURL+"/workitem/:id/comments/", append(
		wrapper.Handler.authMiddlewares(CreateWorkitemComment),
		append(wrapper.Handler.middlewares(CreateWorkitemComment), wrapper.CreateWorkitemComment)...,
	)...)
}
