package handlers

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// User handles routes with the 'user' tag.
type User struct {
	svc services.User
	// add or remove services, etc. as required
}

// NewUser returns a new handler for the 'user' route group.
// Edit as required.
func NewUser(svc services.User) *User {
	return &User{
		svc: svc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (t *User) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "CreateUser",
			Method:      http.MethodPost,
			Pattern:     "/user",
			HandlerFunc: t.CreateUser,
			Middlewares: t.middlewares("CreateUser"),
		},
		{
			Name:        "CreateUsersWithArrayInput",
			Method:      http.MethodPost,
			Pattern:     "/user/createWithArray",
			HandlerFunc: t.CreateUsersWithArrayInput,
			Middlewares: t.middlewares("CreateUsersWithArrayInput"),
		},
		{
			Name:        "CreateUsersWithListInput",
			Method:      http.MethodPost,
			Pattern:     "/user/createWithList",
			HandlerFunc: t.CreateUsersWithListInput,
			Middlewares: t.middlewares("CreateUsersWithListInput"),
		},
		{
			Name:        "DeleteUser",
			Method:      http.MethodDelete,
			Pattern:     "/user/:username",
			HandlerFunc: t.DeleteUser,
			Middlewares: t.middlewares("DeleteUser"),
		},
		{
			Name:        "GetUserByName",
			Method:      http.MethodGet,
			Pattern:     "/user/:username",
			HandlerFunc: t.GetUserByName,
			Middlewares: t.middlewares("GetUserByName"),
		},
		{
			Name:        "LoginUser",
			Method:      http.MethodGet,
			Pattern:     "/user/login",
			HandlerFunc: t.LoginUser,
			Middlewares: t.middlewares("LoginUser"),
		},
		{
			Name:        "LogoutUser",
			Method:      http.MethodGet,
			Pattern:     "/user/logout",
			HandlerFunc: t.LogoutUser,
			Middlewares: t.middlewares("LogoutUser"),
		},
		{
			Name:        "UpdateUser",
			Method:      http.MethodPut,
			Pattern:     "/user/:username",
			HandlerFunc: t.UpdateUser,
			Middlewares: t.middlewares("UpdateUser"),
		},
	}
	rest.RegisterRoutes(r, routes, "/user", mws)
}

// middlewares returns individual route middleware per operation id.
// Edit as required.
func (t *User) middlewares(opId string) []gin.HandlerFunc {
	switch opId {
	default:
		return []gin.HandlerFunc{}
	}
}

// CreateUser creates a new user.
func (t *User) CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// CreateUsersWithArrayInput creates list of users with given input array.
func (t *User) CreateUsersWithArrayInput(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// CreateUsersWithListInput creates list of users with given input array.
func (t *User) CreateUsersWithListInput(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// DeleteUser delete user.
func (t *User) DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// GetUserByName get user by user name.
func (t *User) GetUserByName(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// LoginUser logs user into the system.
func (t *User) LoginUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// LogoutUser logs out current logged in user session.
func (t *User) LogoutUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UpdateUser updated user.
func (t *User) UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
