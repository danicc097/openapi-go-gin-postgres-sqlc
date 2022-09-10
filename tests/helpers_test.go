package tests

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	internaldomain "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/redis"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/vault"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
)

// getStderr returns the contents of stderr.txt in dir.
func getStderr(t *testing.T, dir string) string {
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

// run configures a test server and underlying services with the given configuration.
func run(env, address, specPath string, pool *pgxpool.Pool) (*http.Server, error) {

	if err := envvar.Load(env); err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "envvar.Load")
	}

	provider, err := vault.New()
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewVaultProvider")
	}

	conf := envvar.New(provider)

	rdb, err := redis.New(conf)
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewRedis")
	}

	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.zapNew")
	}

	logger, _ := zap.NewDevelopment()

	// TODO validation middleware
	_, err = openapi3.NewLoader().LoadFromFile("../openapi.yaml")
	if err != nil {
		panic(err)
	}

	srv, err := rest.NewServer(rest.Config{
		Address:  address,
		Pool:     pool,
		Redis:    rdb,
		Logger:   logger,
		SpecPath: specPath,
	})
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "New")
	}

	return srv, nil
}

// newDB returns a new testing Postgres pool.
func newDB() (*pgxpool.Pool, error) {
	provider, err := vault.New()
	if err != nil {
		fmt.Printf("Couldn't create provider: %s", err)
		return nil, err
	}

	conf := envvar.New(provider)
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

	m, err := migrate.NewWithDatabaseInstance("file://../db/migrations/", "postgres", instance)
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

	return testpool, err
}
