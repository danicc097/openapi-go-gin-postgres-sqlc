package tests_test

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/environment"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/vault"
	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewDB(tb testing.TB) *pgxpool.Pool {
	tb.Helper()

	provider, err := vault.New()
	if err != nil {
		tb.Fatalf("Couldn't create provider: %s", err)
	}

	conf := envvar.New(provider)

	os.Setenv("POSTGRES_DB", os.Getenv("POSTGRES_DB")+"_test")
	// TODO use bash script before running tests to create via psql

	pool, err := postgresql.New(conf)
	if err != nil {
		tb.Fatalf("Couldn't create pool: %s", err)
	}

	environment.Pool = pool

	db, err := sql.Open("pgx", pool.Config().ConnString())
	if err != nil {
		tb.Fatalf("Couldn't open DB: %s", err)
	}

	defer db.Close()

	// if err := pool.Retry(func() (err error) {
	// 	return db.Ping()
	// }); err != nil {
	// 	tb.Fatalf("Couldn't ping DB: %s", err)
	// }

	//-

	instance, err := migratepostgres.WithInstance(db, &migratepostgres.Config{})
	if err != nil {
		tb.Fatalf("Couldn't migrate (1): %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://../db/migrations/", "postgres", instance)
	if err != nil {
		tb.Fatalf("Couldn't migrate (2): %s", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		tb.Fatalf("Couldnt' migrate (3): %s", err)
	}

	//-

	dbpool, err := pgxpool.Connect(context.Background(), pool.Config().ConnString())
	if err != nil {
		tb.Fatalf("Couldn't open DB Pool: %s", err)
	}

	tb.Cleanup(func() {
		dbpool.Close()
	})

	return dbpool
}
