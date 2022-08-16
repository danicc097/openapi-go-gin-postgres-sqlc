package handlers

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// Fake handles routes with the 'fake' tag.
type Fake struct {
	svc services.Fake
	// add or remove services, etc. as required
}

// NewFake returns a new handler for the 'fake' route group.
// Edit as required.
func NewFake(svc services.Fake) *Fake {
	return &Fake{
		svc: svc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (t *Fake) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "FakeDataFile",
			Method:      http.MethodGet,
			Pattern:     "/fake/data_file",
			HandlerFunc: t.FakeDataFile,
			Middlewares: t.middlewares("FakeDataFile"),
		},
	}
	rest.RegisterRoutes(r, routes, "/fake", mws)
}

// middlewares returns individual route middleware per operation id.
// Edit as required.
func (t *Fake) middlewares(opId string) []gin.HandlerFunc {
	switch opId {
	default:
		return []gin.HandlerFunc{}
	}
}

// FakeDataFile test data_file to ensure it's escaped correctly.
func (t *Fake) FakeDataFile(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
