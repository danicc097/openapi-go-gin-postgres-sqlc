package testutil

import (
	"database/sql"
	"errors"
	"fmt"
	"path"
	"runtime"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// NewDB returns a new (shared) testing Postgres pool with up-to-date migrations.
func NewDB() (*pgxpool.Pool, *sql.DB, error) {
	logger, _ := zap.NewDevelopment()

	pool, sqlpool, err := postgresql.New(logger.Sugar())
	if err != nil {
		fmt.Printf("Couldn't create pool: %s\n", err)
		return nil, nil, err
	}

	instance, err := migratepostgres.WithInstance(sqlpool, &migratepostgres.Config{})
	if err != nil {
		fmt.Printf("Couldn't migrate (1): %s\n", err)
		return nil, nil, err
	}

	_, src, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+path.Join(path.Dir(src), "../../db/migrations/"), "postgres", instance)
	if err != nil {
		fmt.Printf("Couldn't migrate (2): %s\n", err)
		return nil, nil, err
	}

	// migrate down before tests externally if needed. This function will be called
	// by any test package that needs a db and `up` will be a no-op.
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		fmt.Printf("Couldnt' migrate (3): %s\n", err)
		return nil, nil, err
	}

	return pool, sqlpool, nil
}
