package handlers

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	"github.com/gin-gonic/gin"
)

// Default handles routes with the 'Default_' tag.
type Default struct {
	svc services.Default
	// add or remove services, etc. as required
}

// NewDefault returns a new handler for the 'Default_' route group.
// Edit as required.
func NewDefault(svc services.Default) *Default {
	return &Default{
		svc: svc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (t *Default) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "OpenapiYamlGet",
			Method:      http.MethodGet,
			Pattern:     "/openapi.yaml",
			HandlerFunc: t.OpenapiYamlGet,
			Middlewares: t.middlewares("OpenapiYamlGet"),
		},
		{
			Name:        "Ping",
			Method:      http.MethodGet,
			Pattern:     "/ping",
			HandlerFunc: t.Ping,
			Middlewares: t.middlewares("Ping"),
		},
	}
	rest.RegisterRoutes(r, routes, "/default", mws)
}

// middlewares returns individual route middleware per operation id.
// Edit as required.
func (t *Default) middlewares(opId string) []gin.HandlerFunc {
	switch opId {
	default:
		return []gin.HandlerFunc{}
	}
}

// OpenapiYamlGet returns this very openapi spec..
func (t *Default) OpenapiYamlGet(c *gin.Context) {

	oas, err := static.SwaggerUI.ReadFile("swagger-ui/openapi.yaml")
	if err != nil {
		panic("openapi spec not found")
	}

	c.Data(http.StatusOK, gin.MIMEYAML, oas)
}

// Ping ping pongs.
func (t *Default) Ping(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
