package handlers

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
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

// Register connects the handlers to a router with the given middleware.
// Generated method. DO NOT EDIT.
func (t *Newapi) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "NewApiEndpoint",
			Method:      http.MethodPost,
			Pattern:     "/newapi/endpoint",
			HandlerFunc: t.NewApiEndpoint,
			Middlewares: t.middlewares("NewApiEndpoint"),
		},
	}
	rest.RegisterRoutes(r, routes, "/newapi", mws)
}

// middlewares returns individual route middleware per operation id.
// Edit as required.
func (t *Newapi) middlewares(opId string) []gin.HandlerFunc {
	switch opId {
	default:
		return []gin.HandlerFunc{}
	}
}

// NewApiEndpoint a new endpoint added to the spec.
func (t *Newapi) NewApiEndpoint(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
