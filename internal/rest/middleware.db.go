package rest

import (
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// dbMiddleware handles authentication and authorization middleware.
type dbMiddleware struct {
	logger *zap.SugaredLogger
	pool   *pgxpool.Pool
}

func newDBMiddleware(
	logger *zap.SugaredLogger, pool *pgxpool.Pool,
) *dbMiddleware {
	return &dbMiddleware{
		logger: logger,
		pool:   pool,
	}
}

func (m *dbMiddleware) BeginTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		tx, err := m.pool.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			renderErrorResponse(c, "", internal.WrapErrorf(err, models.ErrorCodePrivate, "could not begin transaction"))
			c.Abort()

			return
		}

		ctxWithTx(c, tx)

		c.Next()

		if err := tx.Commit(ctx); err != nil {
			msg := "could not commit transaction"
			if err := tx.Rollback(ctx); err != nil {
				msg += fmt.Sprintf(" (rollback error: %s)", err)
			}
			// should not render this response at all
			m.logger.Infof("tx.Commit: %v: %v", msg, err)

			return
		}
	}
}
