package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	"github.com/gin-gonic/gin"
)

// Default handles routes with the 'Default_' tag.
type Default struct {
	// add or remove services, etc. as required
}

// NewDefault returns a new handler for the 'Default_' route group.
// Edit as required.
func NewDefault() *Default {
	return &Default{}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *Default) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []route{
		{
			Name:        "OpenapiYamlGet",
			Method:      http.MethodGet,
			Pattern:     "/openapi.yaml",
			HandlerFunc: h.OpenapiYamlGet,
			Middlewares: h.middlewares("OpenapiYamlGet"),
		},
		{
			Name:        "Ping",
			Method:      http.MethodGet,
			Pattern:     "/ping",
			HandlerFunc: h.Ping,
			Middlewares: h.middlewares("Ping"),
		},
	}

	registerRoutes(r, routes, "/default", mws)
}

// middlewares returns individual route middleware per operation id.
// Edit as required.
func (h *Default) middlewares(opID string) []gin.HandlerFunc {
	switch opID {
	default:
		return []gin.HandlerFunc{}
	}
}

// OpenapiYamlGet returns this very openapi spec..
func (h *Default) OpenapiYamlGet(c *gin.Context) {

	oas, err := static.SwaggerUI.ReadFile("swagger-ui/openapi.yaml")
	if err != nil {
		panic("openapi spec not found")
	}

	c.Data(http.StatusOK, gin.MIMEYAML, oas)
}

// Ping ping pongs.
func (h *Default) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
