package rest

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// OpenapiYamlGet returns this very openapi spec.
func (h *Handlers) OpenapiYamlGet(c *gin.Context) {
	c.Set(skipResponseValidation, true)

	oas, err := os.ReadFile(h.specPath)
	if err != nil {
		panic("openapi spec not found")
	}

	c.Data(http.StatusOK, gin.MIMEYAML, oas)
}

// Ping ping pongs.
func (h *Handlers) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
