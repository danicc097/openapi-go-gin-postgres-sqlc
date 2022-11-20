package testutil

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// NewDB returns a new testing Postgres pool.
func NewDB() (*pgxpool.Pool, error) {
	conf := envvar.New()
	pool, err := postgresql.New(conf)
	if err != nil {
		fmt.Printf("Couldn't create pool: %s", err)
		return nil, err
	}

	db, err := sql.Open("pgx", pool.Config().ConnString())
	if err != nil {
		fmt.Printf("Couldn't open Pool: %s", err)
		return nil, err
	}

	defer db.Close()

	instance, err := migratepostgres.WithInstance(db, &migratepostgres.Config{})
	if err != nil {
		fmt.Printf("Couldn't migrate (1): %s", err)
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../db/migrations/", "postgres", instance)
	if err != nil {
		fmt.Printf("Couldn't migrate (2): %s", err)
		return nil, err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		fmt.Printf("Couldnt' migrate (3): %s", err)
		return nil, err
	}

	testpool, err := pgxpool.Connect(context.Background(), pool.Config().ConnString())
	if err != nil {
		fmt.Printf("Couldn't open Pool: %s", err)
		return nil, err
	}

	return testpool, nil
}
