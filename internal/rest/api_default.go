package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/gin-gonic/gin"
)

// OpenapiYamlGet returns this very openapi spec.
func (h *Handlers) OpenapiYamlGet(c *gin.Context) {
	oas, err := os.ReadFile(h.specPath)
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
