package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Admin handles routes with the 'admin' tag.
type Admin struct {
	svc UserService
	// add or remove services, etc. as required
}

// NewAdmin returns a new handler for the 'admin' route group.
func NewAdmin(svc UserService) *Admin {
	return &Admin{
		svc: svc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *Admin) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []route{
		{
			Name:        string(AdminPing),
			Method:      http.MethodGet,
			Pattern:     "/admin/ping",
			HandlerFunc: h.AdminPing,
			Middlewares: h.middlewares(AdminPing),
		},
	}

	registerRoutes(r, routes, "/admin", mws)
}

// middlewares returns individual route middleware per operation id.
func (h *Admin) middlewares(opID adminOpID) []gin.HandlerFunc {
	switch opID {
	default:
		return []gin.HandlerFunc{}
	}
}

// AdminPing ping pongs.
func (h *Admin) AdminPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
