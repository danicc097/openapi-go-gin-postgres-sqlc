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
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	postgresqlutils "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/postgresql"
	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var once sync.Once

// cannot be random since we want to lock parallel test suites.
const migrationsLockID = 12341234

// NewDB returns a new (shared) testing Postgres pool with up-to-date migrations.
// Panics on error.
func NewDB() (*pgxpool.Pool, *sql.DB, error) {
	logger, _ := zap.NewDevelopment()

	pool, sqlpool, err := postgresql.New(logger.Sugar())
	if err != nil {
		panic(fmt.Sprintf("Couldn't create pool: %s\n", err))
	}

	lock, err := postgresqlutils.NewAdvisoryLock(pool, migrationsLockID)
	if err != nil {
		panic(fmt.Sprintf("NewAdvisoryLock: %s\n", err))
	}

	acquired, err := lock.TryLock(context.Background())
	if err != nil {
		panic(fmt.Sprintf("lock.TryLock: %s\n", err))
	}
	if !acquired {
		// wait for migrations
		if err := lock.WaitForRelease(50, 200*time.Millisecond); err != nil {
			panic(fmt.Sprintf("lock.WaitForRelease: %s\n", err))
		}

		return pool, sqlpool, nil
	}

	defer func() {
		unlockSuccess := lock.Release()
		for i := 0; !unlockSuccess && lock.IsLocked() && i < 10; i++ {
			unlockSuccess = lock.Release()
		}
		lock.ReleaseConn()
		if lock.IsLocked() {
			panic(fmt.Sprintf("advisory lock was not released\n"))
		}
	}()

	driver, err := migratepostgres.WithInstance(sqlpool, &migratepostgres.Config{})
	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate (1.1): %s\n", err))
	}
	postDriver, err := migratepostgres.WithInstance(sqlpool, &migratepostgres.Config{MigrationsTable: "schema_post_migrations"})
	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate (1.2): %s\n", err))
	}
	_, src, _, ok := runtime.Caller(0)
	if !ok {
		panic("No runtime caller information")
	}
	mPostMigrations, err := migrate.NewWithDatabaseInstance("file://"+path.Join(path.Dir(src), "../../db/post-migrations/"), "postgres", driver)
	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate (2.1): %s\n", err))
	}
	_ = mPostMigrations

	mMigrations, err := migrate.NewWithDatabaseInstance("file://"+path.Join(path.Dir(src), "../../db/migrations/"), "postgres", postDriver)
	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate (2.2): %s\n", err))
	}

	// FIXME: this is being ran multiple times. refactor post-migrations to instead use x-migrations-table=schema_post_migrations pool connection string
	// since all we care about is post migrations being executed in order after the main ones
	fmt.Println("RUNNING DOWN MIGRATIONS")
	// ~~~~ NOTE: drop table before tests only externally.
	// ^ now that we use a lock for test migrations with parallel tests, we can run m.Down
	// instead of dropping
	// if err = m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
	// 	panic(fmt.Sprintf("Couldnt' migrate down: %s\n", err))
	// }
	if err = mMigrations.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(fmt.Sprintf("Couldnt' migrate up: %s\n", err))
	}

	postMigrationPath := path.Join(path.Dir(src), "../../db/post-migrations/")
	files, err := os.ReadDir(postMigrationPath)
	if err != nil {
		panic(fmt.Sprintf("Error reading post-migrations directory: %s\n", err))
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasPrefix(file.Name(), ".sql") {
			continue
		}

		filePath := path.Join(postMigrationPath, file.Name())
		script, err := os.ReadFile(filePath)
		if err != nil {
			panic(fmt.Sprintf("Error reading post-migrations script %s: %s\n", file.Name(), err))
		}

		if _, err := sqlpool.Exec(string(script)); err != nil {
			panic(fmt.Sprintf("Error executing post-migrations script %s: %s\n", file.Name(), err))
		}

		fmt.Printf("Run post-migrations script: %s\n", file.Name())
	}

	return pool, sqlpool, nil
}
