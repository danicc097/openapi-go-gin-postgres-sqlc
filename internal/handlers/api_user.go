// Code generated by openapi-generator. DO NOT EDIT.

package handlers

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// User handles routes with the user tag.
type User struct {
	svc services.User
	// add your own services, etc. as required
}

// NewUser returns a new handler for user.
// Edit as required
// TODO rewriting handler methods based on current postgen:
// see https://eli.thegreenplace.net/2021/rewriting-go-source-code-with-ast-tooling/
// simpler solutions based on drawbacks (complicated, comments not attached to nodes):
// - https://github.com/dave/dst
// - https://github.com/uber-go/gopatch
func NewUser(svc services.User) *User {
	return &User{
		svc: svc,
	}
}

// Register connects the handlers to a router with the given middleware.
func (t *User) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "CreateUser",
			Method:      http.MethodPost,
			Pattern:     "/user",
			HandlerFunc: t.CreateUser,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "CreateUsersWithArrayInput",
			Method:      http.MethodPost,
			Pattern:     "/user/createWithArray",
			HandlerFunc: t.CreateUsersWithArrayInput,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "CreateUsersWithListInput",
			Method:      http.MethodPost,
			Pattern:     "/user/createWithList",
			HandlerFunc: t.CreateUsersWithListInput,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "DeleteUser",
			Method:      http.MethodDelete,
			Pattern:     "/user/:username",
			HandlerFunc: t.DeleteUser,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "GetUserByName",
			Method:      http.MethodGet,
			Pattern:     "/user/:username",
			HandlerFunc: t.GetUserByName,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "LoginUser",
			Method:      http.MethodGet,
			Pattern:     "/user/login",
			HandlerFunc: t.LoginUser,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "LogoutUser",
			Method:      http.MethodGet,
			Pattern:     "/user/logout",
			HandlerFunc: t.LogoutUser,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "UpdateUser",
			Method:      http.MethodPut,
			Pattern:     "/user/:username",
			HandlerFunc: t.UpdateUser,
			Middlewares: []gin.HandlerFunc{},
		},
	}

	rest.RegisterRoutes(r, routes, "/user", mws)
}

// CreateUser creates a new user.
func (t *User) CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// CreateUsersWithArrayInput creates list of users with given input array.
func (t *User) CreateUsersWithArrayInput(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// CreateUsersWithListInput creates list of users with given input array.
func (t *User) CreateUsersWithListInput(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteUser delete user.
func (t *User) DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetUserByName get user by user name.
func (t *User) GetUserByName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// LoginUser logs user into the system.
func (t *User) LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// LogoutUser logs out current logged in user session.
func (t *User) LogoutUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateUser updated user.
func (t *User) UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
