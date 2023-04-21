// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strings"

	zapadapter "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v5/tracelog"
	"go.uber.org/zap"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
)

// New instantiates the PostgreSQL database using configuration defined in environment variables.
func New(logger *zap.Logger) (*pgxpool.Pool, *sql.DB, error) {
	cfg := internal.Config()
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

	if cfg.Postgres.TraceEnabled {
		poolConfig.ConnConfig.Tracer = &tracelog.TraceLog{
			Logger:   zapadapter.NewLogger(logger),
			LogLevel: tracelog.LogLevelTrace,
		}
	}

	poolConfig.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
		searchPaths := []string{"public"}
		typeNames, err := QueryDatabaseTypeNames(context.Background(), c, searchPaths...)
		if err != nil {
			return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "could not query database types")
		}

		err = RegisterDataTypes(context.Background(), c, typeNames)
		if err != nil {
			return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "could not register data types")
		}

		return nil
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

// RegisterDataTypes automatically registers all enums and tables types
// for proper encoding/decoding in pgx.
// See https://pkg.go.dev/github.com/jackc/pgx/v5@v5.3.1/pgtype#hdr-New_PostgreSQL_Type_Support
func RegisterDataTypes(ctx context.Context, conn *pgx.Conn, typeNames []string) error {
	for _, typeName := range typeNames {
		dataType, err := conn.LoadType(ctx, typeName)
		if err != nil {
			return err
		}
		conn.TypeMap().RegisterType(dataType)
	}

	return nil
}

func QueryDatabaseTypeNames(ctx context.Context, conn *pgx.Conn, searchPaths ...string) ([]string, error) {
	query := fmt.Sprintf(`SELECT table_name
	FROM information_schema.tables
	WHERE table_schema IN ('%s')`, strings.Join(searchPaths, "', '"))
	tableTypes, err := queryTypeNames(conn, query)
	if err != nil {
		return []string{}, fmt.Errorf("querying table names: %w", err)
	}

	query = fmt.Sprintf(`SELECT t.typname AS enum_name
	FROM pg_type t
	INNER JOIN pg_namespace n ON n.oid = t.typnamespace
	WHERE t.typtype = 'e' AND n.nspname IN ('%s');`, strings.Join(searchPaths, "', '"))
	enumTypes, err := queryTypeNames(conn, query)
	if err != nil {
		return []string{}, fmt.Errorf("querying enum names: %w", err)
	}

	// register enumTypes first, in case they're used in tables
	typeNames := append(enumTypes, tableTypes...)

	return typeNames, err
}

func queryTypeNames(conn *pgx.Conn, query string) ([]string, error) {
	names := []string{}
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return []string{}, fmt.Errorf("conn.Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var enumName string
		err = rows.Scan(&enumName)
		if err != nil {
			return []string{}, fmt.Errorf("rows.Scan: %w", err)
		}
		names = append(names, enumName)
		names = append(names, "_"+enumName) // postgres internal array type automatically created
	}
	if err = rows.Err(); err != nil {
		return []string{}, fmt.Errorf("rows.Next: %w", err)
	}
	return names, err
}
