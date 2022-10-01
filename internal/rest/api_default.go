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
func NewDefault() *Default {
	return &Default{}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *Default) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []route{
		{
			Name:        string(openapiYamlGet),
			Method:      http.MethodGet,
			Pattern:     "/openapi.yaml",
			HandlerFunc: h.openapiYamlGet,
			Middlewares: h.middlewares(openapiYamlGet),
		},
		{
			Name:        string(ping),
			Method:      http.MethodGet,
			Pattern:     "/ping",
			HandlerFunc: h.ping,
			Middlewares: h.middlewares(ping),
		},
	}

	registerRoutes(r, routes, "/default", mws)
}

// middlewares returns individual route middleware per operation id.
func (h *Default) middlewares(opID defaultOpID) []gin.HandlerFunc {
	switch opID {
	default:
		return []gin.HandlerFunc{}
	}
}

// openapiYamlGet returns this very openapi spec..
func (h *Default) openapiYamlGet(c *gin.Context) {

	oas, err := static.SwaggerUI.ReadFile("swagger-ui/openapi.yaml")
	if err != nil {
		panic("openapi spec not found")
	}

	c.Data(http.StatusOK, gin.MIMEYAML, oas)
}

// ping ping pongs.
func (h *Default) ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
