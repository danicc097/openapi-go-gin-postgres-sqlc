package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	"github.com/gin-gonic/gin"
)

// Events represents server events.
func (h *Handlers) Events(c *gin.Context) {
	oas, err := static.SwaggerUI.ReadFile("swagger-ui/openapi.yaml")
	if err != nil {
		panic("openapi spec not found")
	}

	c.Data(http.StatusOK, gin.MIMEYAML, oas)
}
