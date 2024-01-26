package testutil

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	postgresqlutils "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/postgresql"
	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var once sync.Once

const migrationsLockID = 12341234

// NewDB returns a new (shared) testing Postgres pool with up-to-date migrations.
// Panics on error.
func NewDB() (*pgxpool.Pool, *sql.DB, error) {
	logger, _ := zap.NewDevelopment()

	pool, sqlpool, err := postgresql.New(logger.Sugar())
	if err != nil {
		panic(fmt.Sprintf("Couldn't create pool: %s\n", err))
	}

	advisoryLock, err := postgresqlutils.NewAdvisoryLock(pool, migrationsLockID)
	if err != nil {
		panic(fmt.Sprintf("NewAdvisoryLock: %s\n", err))
	}

	mustMigrate, err := advisoryLock.TryLock(context.Background())
	if err != nil {
		panic(fmt.Sprintf("advisoryLock.TryLock: %s\n", err))
	}
	fmt.Printf("mustMigrate: %v\n", mustMigrate)
	if !mustMigrate {
		return pool, sqlpool, nil
	}
	defer func() {
		if err := advisoryLock.Release(context.Background()); err != nil {
			panic(fmt.Sprintf("advisoryLock.Release: %s\n", err))
		}
	}()

	if err := advisoryLock.WaitForRelease(context.Background()); err != nil {
		panic(fmt.Sprintf("advisoryLock.WaitForRelease: %s\n", err))
	}
	instance, err := migratepostgres.WithInstance(sqlpool, &migratepostgres.Config{})
	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate (1): %s\n", err))
	}

	_, src, _, ok := runtime.Caller(0)
	if !ok {
		panic("No runtime caller information")
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+path.Join(path.Dir(src), "../../db/migrations/"), "postgres", instance)
	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate (2): %s\n", err))
	}

	// NOTE: drop table before tests only externally.
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(fmt.Sprintf("Couldnt' migrate (3): %s\n", err))
	}

	postMigrationPath := path.Join(path.Dir(src), "../../db/post-migration/")
	files, err := os.ReadDir(postMigrationPath)
	if err != nil {
		panic(fmt.Sprintf("Error reading post-migration directory: %s\n", err))
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasPrefix(file.Name(), ".sql") {
			continue
		}

		filePath := path.Join(postMigrationPath, file.Name())
		script, err := os.ReadFile(filePath)
		if err != nil {
			panic(fmt.Sprintf("Error reading post-migration script %s: %s\n", file.Name(), err))
		}

		_, err = sqlpool.Exec(string(script))
		if err != nil {
			panic(fmt.Sprintf("Error executing post-migration script %s: %s\n", file.Name(), err))
		}

		fmt.Printf("Run post-migration script: %s\n", file.Name())
	}

	return pool, sqlpool, nil
}
