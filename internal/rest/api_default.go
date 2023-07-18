package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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

	ctx := c.Request.Context()
	// TODO: tx could be middleware. no need to check if context tx is undefined
	// because itll always be, else it prematurely renders errors and abort
	// if we forget to add mw tests just fail since all routes are tested...
	// easiest would be to by default have tx mw in all routes, but the option
	// to exclude an array of opIDs that turn the mw into a noop. (auth provider login, etc)
	tx, err := h.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		renderErrorResponse(c, "Database error", internal.WrapErrorf(err, models.ErrorCodePrivate, "could not begin tx"))

		return
	}
	defer tx.Rollback(ctx)

	c.String(http.StatusOK, "pong")
}
