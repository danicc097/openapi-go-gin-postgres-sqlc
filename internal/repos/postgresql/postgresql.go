// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	zapadapter "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	// to open with "pgx" driver.
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v5/tracelog"
	"go.uber.org/zap"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
)

var pgxAfterConnectLock = sync.Mutex{}

type DBOptions struct {
	DBName string
}

type Option func(*DBOptions)

// WithDBName sets the postgres database to connect to.
func WithDBName(db string) Option {
	return func(opt *DBOptions) {
		opt.DBName = db
	}
}

// New instantiates the PostgreSQL database using configuration defined in environment variables.
func New(logger *zap.SugaredLogger, options ...Option) (*pgxpool.Pool, *sql.DB, error) {
	cfg := internal.Config

	dbOptions := &DBOptions{}
	for _, option := range options {
		option(dbOptions)
	}

	if dbOptions.DBName == "" {
		dbOptions.DBName = cfg.Postgres.DB
	}

	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.Postgres.User, cfg.Postgres.Password),
		Host:   fmt.Sprintf("%s:%s", cfg.Postgres.Server, strconv.Itoa(cfg.Postgres.Port)),
		Path:   dbOptions.DBName,
	}

	q := dsn.Query()
	q.Add("sslmode", os.Getenv("DATABASE_SSLMODE"))

	dsn.RawQuery = q.Encode()

	poolConfig, err := pgxpool.ParseConfig(dsn.String())
	if err != nil {
		return nil, nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "pgx.ParseConfig")
	}

	poolConfig.ConnConfig.OnNotice = func(pc *pgconn.PgConn, n *pgconn.Notice) {
		logger.Infof("Postgres notice: %+v", *n)
	}

	if cfg.Postgres.TraceEnabled {
		poolConfig.ConnConfig.Tracer = &tracelog.TraceLog{
			Logger:   zapadapter.NewLogger(logger.Desugar()),
			LogLevel: tracelog.LogLevelTrace,
		}
	}

	var atLeastOneConnInPool atomic.Bool

	poolConfig.MinConns = 4
	// NOTE: CI fails using default of 4
	poolConfig.MaxConns = 20
	if os.Getenv("IS_TESTING") != "" {
		poolConfig.ConnConfig.RuntimeParams["statement_timeout"] = "60s"
		poolConfig.MaxConns = 20
	}

	// called after a connection is established, but before it is added to the pool.
	// Will run once.
	poolConfig.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
		// DATA RACE (when using shared typeNames)
		// pgxAfterConnectLock.Lock()
		// defer pgxAfterConnectLock.Unlock()
		var err error

		searchPaths := []string{"public", "xo_tests"}
		typeNames, err := queryDatabaseTypeNames(c, logger, searchPaths...)
		if err != nil {
			return internal.WrapErrorf(err, models.ErrorCodeUnknown, "could not query database types")
		}

		if err := registerDataTypes(ctx, c, typeNames); err != nil {
			return internal.WrapErrorf(err, models.ErrorCodeUnknown, "could not register data types")
		}

		atLeastOneConnInPool.Store(true)

		return nil
	}

	pgxPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "pgxpool.NewWithConfig")
	}

	if err := pgxPool.Ping(context.Background()); err != nil {
		return nil, nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "pgxPool.Ping")
	}

	sqlPool, err := sql.Open("pgx", pgxPool.Config().ConnString())
	if err != nil {
		return nil, nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "sql.Open")
	}

	for !atLeastOneConnInPool.Load() {
		time.Sleep(50 * time.Millisecond)
	}

	return pgxPool, sqlPool, nil
}

// registerDataTypes automatically registers all enums and tables types
// for proper encoding/decoding in pgx.
// See https://pkg.go.dev/github.com/jackc/pgx/v5@v5.3.1/pgtype#hdr-New_PostgreSQL_Type_Support
func registerDataTypes(ctx context.Context, conn *pgx.Conn, typeNames []string) error {
	for _, typeName := range typeNames {
		dataType, err := conn.LoadType(ctx, typeName)
		if err != nil {
			return err
		}
		conn.TypeMap().RegisterType(dataType)
	}

	return nil
}

func queryDatabaseTypeNames(conn *pgx.Conn, logger *zap.SugaredLogger, searchPaths ...string) ([]string, error) {
	var typeNames []string
	for _, sp := range searchPaths {
		query := fmt.Sprintf(`SELECT table_name
	FROM information_schema.tables
	WHERE table_schema IN ('%s')`, sp)
		tableTypes, err := queryTypeNames(conn, query, sp)
		if err != nil {
			return []string{}, fmt.Errorf("querying table names: %w", err)
		}

		query = fmt.Sprintf(`SELECT t.typname AS enum_name
	FROM pg_type t
	INNER JOIN pg_namespace n ON n.oid = t.typnamespace
	WHERE t.typtype = 'e' AND n.nspname IN ('%s');`, sp)
		enumTypes, err := queryTypeNames(conn, query, sp)
		if err != nil {
			return []string{}, fmt.Errorf("querying enum names: %w", err)
		}

		// register enumTypes first, in case they're used in tables
		typeNames = append(typeNames, append(enumTypes, tableTypes...)...)
	}

	if len(typeNames) == 0 {
		logger.Warn("database typenames not found - make sure migrations have been run")
	}

	return typeNames, nil
}

func queryTypeNames(conn *pgx.Conn, query string, searchPath string) ([]string, error) {
	names := []string{}
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return []string{}, fmt.Errorf("conn.Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var enumName string

		if err := rows.Scan(&enumName); err != nil {
			return []string{}, fmt.Errorf("rows.Scan: %w", err)
		}
		names = append(names, searchPath+"."+enumName)
		names = append(names, searchPath+"."+"_"+enumName) // postgres internal array type automatically created
	}
	if err = rows.Err(); err != nil {
		return []string{}, fmt.Errorf("rows.Next: %w", err)
	}
	return names, err
}
