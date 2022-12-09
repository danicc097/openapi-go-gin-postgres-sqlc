// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	zapadapter "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v5/tracelog"
	"go.uber.org/zap"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
)

// New instantiates the PostgreSQL database using configuration defined in environment variables.
func New(conf *envvar.Configuration, logger *zap.Logger) (*pgxpool.Pool, *sql.DB, error) {
	get := func(v string) string {
		res, err := conf.Get(v)
		if err != nil {
			log.Fatalf("Couldn't get configuration value for %s: %s", v, err)
		}

		return res
	}

	// XXX: We will revisit this code in future episodes replacing it with another solution
	databaseHost := get("POSTGRES_SERVER")
	databaseUsername := get("POSTGRES_USER")
	databasePassword := get("POSTGRES_PASSWORD")
	databaseName := get("POSTGRES_DB")
	databaseSSLMode := get("DATABASE_SSLMODE")
	// XXX: -

	var databasePort string
	switch env := os.Getenv("APP_ENV"); env {
	case "prod":
		databasePort = get("POSTGRES_PORT") // container
	default:
		databasePort = get("DB_PORT")
	}

	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(databaseUsername, databasePassword),
		Host:   fmt.Sprintf("%s:%s", databaseHost, databasePort),
		Path:   databaseName,
	}

	q := dsn.Query()
	q.Add("sslmode", databaseSSLMode)

	dsn.RawQuery = q.Encode()

	poolConfig, err := pgxpool.ParseConfig(dsn.String())
	if err != nil {
		return nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "pgx.ParseConfig")
	}
	poolConfig.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   zapadapter.NewLogger(logger),
		LogLevel: tracelog.LogLevelTrace,
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
