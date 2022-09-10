package rest

import (
	"net/http"

	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// Newapi handles routes with the 'newapi' tag.
type Newapi struct {
	svc services.Newapi
	// add or remove services, etc. as required
}

// NewNewapi returns a new handler for the 'newapi' route group.
// Edit as required.
func NewNewapi(svc services.Newapi) *Newapi {
	return &Newapi{
		svc: svc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *Newapi) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []route{
		{
			Name:        "NewApiEndpoint",
			Method:      http.MethodPost,
			Pattern:     "/newapi/endpoint",
			HandlerFunc: h.NewApiEndpoint,
			Middlewares: h.middlewares("NewApiEndpoint"),
		},
	}

	registerRoutes(r, routes, "/newapi", mws)
}

// middlewares returns individual route middleware per operation id.
// Edit as required.
func (h *Newapi) middlewares(opID string) []gin.HandlerFunc {
	switch opID {
	default:
		return []gin.HandlerFunc{}
	}
}

// NewApiEndpoint a new endpoint added to the spec.
func (h *Newapi) NewApiEndpoint(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
