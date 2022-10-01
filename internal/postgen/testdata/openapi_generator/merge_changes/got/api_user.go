package rest

import (
	"net/http"

	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// User handles routes with the user tag.
type User struct {
	svc services.User
	// add or remove services, etc. as required
}

// NewUser returns a new handler for user.
func NewUser(svc services.User) *User {
	return &User{
		svc: svc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *User) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []route{
		{
			Name:        string(createUser),
			Method:      http.MethodPost,
			Pattern:     "/user",
			HandlerFunc: h.createUser,
			Middlewares: h.middlewares(createUser),
		},
	}

	registerRoutes(r, routes, "/user", mws)
}

// createUser creates a new user.
func (h *User) createUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// createUsersWithArrayInput creates list of users with given input array.
func (h *User) createUsersWithArrayInput(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// createUsersWithListInput creates list of users with given input array.
func (h *User) createUsersWithListInput(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// deleteUser delete user.
func (h *User) deleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// getUserByName get user by user name.
func (h *User) getUserByName(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// loginUser logs user into the system.
func (h *User) loginUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// logoutUser logs out current logged in user session.
func (h *User) logoutUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// updateUser updated user.
func (h *User) updateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// middlewares returns individual route middleware per operation id.
func (h *User) middlewares(opID userOpID) []gin.HandlerFunc {
	switch opID {
	default:
		return []gin.HandlerFunc{}
	}
}
