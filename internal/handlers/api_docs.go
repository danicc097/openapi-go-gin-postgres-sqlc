package handlers

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	"github.com/gin-gonic/gin"
)

// OpenapiYamlGet returns this very OpenAPI spec.
func OpenapiYamlGet(c *gin.Context) {
	c.Header("Content-Type", "application/x-yaml")

	oas, err := static.SwaggerUI.ReadFile("swagger-ui/openapi.yaml")
	if err != nil {
		panic("openapi spec not found")
	}

	c.String(http.StatusOK, string(oas))
}
