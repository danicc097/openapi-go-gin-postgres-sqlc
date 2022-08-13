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
	// add necessary services, etc. as required
}

// NewFake returns a new handler for the 'fake' route group.
// Edit as required.
func NewFake(svc services.Fake) *Fake {
	return &Fake{
		svc: svc,
	}
}

// Register connects the handlers to a router with the given middleware.
// GENERATED METHOD. Only Middlewares will be saved between runs.
func (t *Fake) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "FakeDataFile",
			Method:      http.MethodGet,
			Pattern:     "/fake/data_file",
			HandlerFunc: t.FakeDataFile,
			Middlewares: []gin.HandlerFunc{},
		},
	}

	rest.RegisterRoutes(r, routes, "/fake", mws)
}

// FakeDataFile test data_file to ensure it's escaped correctly.
func (t *Fake) FakeDataFile(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
