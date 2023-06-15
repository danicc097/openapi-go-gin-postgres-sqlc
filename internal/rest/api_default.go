package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

// OpenapiYamlGet returns this very openapi spec.
func (h *Handlers) OpenapiYamlGet(c *gin.Context) {
	c.Set(skipResponseValidation, true)

	oas, err := os.ReadFile(h.specPath)
	if err != nil {
		panic("openapi spec not found")
	}

	c.String(http.StatusOK, string(oas))
}

// Ping ping pongs.
func (h *Handlers) Ping(c *gin.Context) {
	fmt.Printf("internal.Config.AppEnv: %v\n", internal.Config.AppEnv)

	ctx := c.Request.Context()
	tx, err := h.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		renderErrorResponse(c, "database error", internal.WrapErrorf(err, internal.ErrorCodePrivate, "could not being tx"))

		return
	}
	defer tx.Rollback(ctx)

	c.String(http.StatusOK, "pong")
}
