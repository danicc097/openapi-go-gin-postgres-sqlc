package rest

import (
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	"github.com/gin-gonic/gin"
)

// OpenapiYamlGet returns this very openapi spec.
func (h *Handlers) OpenapiYamlGet(c *gin.Context) {
	oas, err := static.SwaggerUI.ReadFile("swagger-ui/openapi.yaml")
	if err != nil {
		panic("openapi spec not found")
	}

	c.String(http.StatusOK, string(oas))
}

// Ping ping pongs.
func (h *Handlers) Ping(c *gin.Context) {
	fmt.Printf("internal.Config.AppEnv: %v\n", internal.Config.AppEnv)

	c.String(http.StatusOK, "pong")
}
