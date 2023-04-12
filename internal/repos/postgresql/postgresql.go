// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"

	zapadapter "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v5/tracelog"
	"go.uber.org/zap"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
)

// New instantiates the PostgreSQL database using configuration defined in environment variables.
func New(logger *zap.Logger) (*pgxpool.Pool, *sql.DB, error) {
	cfg := internal.Config
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.Postgres.User, cfg.Postgres.Password),
		Host:   fmt.Sprintf("%s:%s", cfg.Postgres.Server, fmt.Sprint(cfg.Postgres.Port)),
		Path:   cfg.Postgres.DB,
	}

	q := dsn.Query()
	q.Add("sslmode", os.Getenv("DATABASE_SSLMODE"))

	dsn.RawQuery = q.Encode()

	poolConfig, err := pgxpool.ParseConfig(dsn.String())
	if err != nil {
		return nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "pgx.ParseConfig")
	}

	if internal.Config.Postgres.TraceEnabled {
		poolConfig.ConnConfig.Tracer = &tracelog.TraceLog{
			Logger:   zapadapter.NewLogger(logger),
			LogLevel: tracelog.LogLevelTrace,
		}
	}

	pgxPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "pgxpool.New")
	}

	if err := pgxPool.Ping(context.Background()); err != nil {
		return nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "db.Ping")
	}

	sqlPool, err := sql.Open("pgx", pgxPool.Config().ConnString())
	if err != nil {
		return nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "sql.Open")
	}

	return pgxPool, sqlPool, nil
}
