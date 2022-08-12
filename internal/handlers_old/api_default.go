package handlers

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	"github.com/gin-gonic/gin"
)

// Ping ping pongs.
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// OpenapiYamlGet returns this very OpenAPI spec.
func OpenapiYamlGet(c *gin.Context) {
	oas, err := static.SwaggerUI.ReadFile("swagger-ui/openapi.yaml")
	if err != nil {
		panic("openapi spec not found")
	}

	c.Data(http.StatusOK, gin.MIMEYAML, oas)
}
