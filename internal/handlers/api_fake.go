// Code generated by openapi-generator. DO NOT EDIT.

package handlers

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// Fake handles routes with the fake tag.
type Fake struct {
	svc services.Fake
	// add your own services, etc. as required
}

// NewFake returns a new handler for fake.
// Edit as required
// TODO rewriting handler methods based on current postgen:
// see https://eli.thegreenplace.net/2021/rewriting-go-source-code-with-ast-tooling/
// simpler solutions based on drawbacks (complicated, comments not attached to nodes):
// - https://github.com/dave/dst
// - https://github.com/uber-go/gopatch
func NewFake(svc services.Fake) *Fake {
	return &Fake{
		svc: svc,
	}
}

// Register connects the handlers to a router with the given middleware.
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
	c.JSON(http.StatusOK, gin.H{})
}
