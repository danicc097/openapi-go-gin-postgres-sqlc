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
	// add or remove services, etc. as required
}

// NewUser returns a new handler for user.
// Edit as required.
func NewUser(svc services.User) *User {
	return &User{
		svc: svc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *User) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "CreateUser",
			Method:      http.MethodPost,
			Pattern:     "/user",
			HandlerFunc: h.CreateUser,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "CreateUsersWithArrayInput",
			Method:      http.MethodPost,
			Pattern:     "/user/createWithArray",
			HandlerFunc: h.CreateUsersWithArrayInput,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "CreateUsersWithListInput",
			Method:      http.MethodPost,
			Pattern:     "/user/createWithList",
			HandlerFunc: h.CreateUsersWithListInput,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "DeleteUser",
			Method:      http.MethodDelete,
			Pattern:     "/user/:username",
			HandlerFunc: h.DeleteUser,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "GetUserByName",
			Method:      http.MethodGet,
			Pattern:     "/user/:username",
			HandlerFunc: h.GetUserByName,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "LoginUser",
			Method:      http.MethodGet,
			Pattern:     "/user/login",
			HandlerFunc: h.LoginUser,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "LogoutUser",
			Method:      http.MethodGet,
			Pattern:     "/user/logout",
			HandlerFunc: h.LogoutUser,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "UpdateUser",
			Method:      http.MethodPut,
			Pattern:     "/user/:username",
			HandlerFunc: h.UpdateUser,
			Middlewares: []gin.HandlerFunc{},
		},
	}

	rest.RegisterRoutes(r, routes, "/user", mws)
}

// CreateUser creates a new user.
func (h *User) CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// CreateUsersWithArrayInput creates list of users with given input array.
func (h *User) CreateUsersWithArrayInput(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// CreateUsersWithListInput creates list of users with given input array.
func (h *User) CreateUsersWithListInput(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// DeleteUser delete user.
func (h *User) DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// GetUserByName get user by user name.
func (h *User) GetUserByName(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// LoginUser logs user into the system.
func (h *User) LoginUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// LogoutUser logs out current logged in user session.
func (h *User) LogoutUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UpdateUser updated user.
func (h *User) UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
