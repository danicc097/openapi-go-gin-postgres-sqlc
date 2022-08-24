package tests

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	internaldomain "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/redis"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/server"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/vault"
	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap/zaptest"
)

// GetStderr returns the contents of stderr.txt in dir.
func GetStderr(t *testing.T, dir string) string {
	t.Helper()
	path := filepath.Join(dir, "stderr.txt")

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		blob, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}

		return string(blob)
	}

	return ""
}

// Run configures a test server and underlying services with the given configuration.
func Run(tb testing.TB, env, address string) (*http.Server, error) {
	if err := envvar.Load(env); err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "envvar.Load")
	}

	provider, err := vault.New()
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewVaultProvider")
	}

	conf := envvar.New(provider)

	pool := NewDB(tb)

	rdb, err := redis.New(conf)
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewRedis")
	}

	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.zapNew")
	}

	srv, err := server.New(server.Config{
		Address: address,
		DB:      pool,
		Redis:   rdb,
		Logger:  zaptest.NewLogger(tb),
	})
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "New")
	}

	return srv, nil
}

// NewServer returns a new test server.
func NewServer(t *testing.T) *http.Server {
	t.Helper()

	srv, err := Run(t, "../.env", ":8099")
	if err != nil {
		t.Fatalf("Couldn't run test server: %s", err)
	}

	return srv
}

// NewDB returns a new testing Postgres pool.
func NewDB(tb testing.TB) *pgxpool.Pool {
	tb.Helper()

	provider, err := vault.New()
	if err != nil {
		tb.Fatalf("Couldn't create provider: %s", err)
	}

	conf := envvar.New(provider)

	tb.Setenv("POSTGRES_DB", "postgres_test")

	pool, err := postgresql.New(conf)
	if err != nil {
		tb.Fatalf("Couldn't create pool: %s", err)
	}

	db, err := sql.Open("pgx", pool.Config().ConnString())
	if err != nil {
		tb.Fatalf("Couldn't open DB: %s", err)
	}

	defer db.Close()

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

	dbpool, err := pgxpool.Connect(context.Background(), pool.Config().ConnString())
	if err != nil {
		tb.Fatalf("Couldn't open DB Pool: %s", err)
	}

	tb.Cleanup(func() {
		dbpool.Close()
	})

	return dbpool
}
