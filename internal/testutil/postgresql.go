package testutil

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path"
	"runtime"
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
		panic(fmt.Sprintf("Couldn't migrate (migrations): %s\n", err))
	}
	postDriver, err := migratepostgres.WithInstance(sqlpool, &migratepostgres.Config{MigrationsTable: "schema_post_migrations"})
	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate (post-migrations): %s\n", err))
	}
	_, src, _, ok := runtime.Caller(0)
	if !ok {
		panic("No runtime caller information")
	}
	mPostMigrations, err := migrate.NewWithDatabaseInstance("file://"+path.Join(path.Dir(src), "../../db/post-migrations/"), "postgres", driver)
	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate (migrations): %s\n", err))
	}
	_ = mPostMigrations

	mMigrations, err := migrate.NewWithDatabaseInstance("file://"+path.Join(path.Dir(src), "../../db/migrations/"), "postgres", postDriver)
	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate (post-migrations): %s\n", err))
	}

	fmt.Println("RUNNING DOWN MIGRATIONS")
	// ~~~~ NOTE: drop table before tests only externally.
	// ^ now that we use a lock for test migrations with parallel tests, we can run m.Down
	// instead of dropping
	// if err = m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
	// 	panic(fmt.Sprintf("Couldnt' migrate down: %s\n", err))
	// }
	if err = mMigrations.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(fmt.Sprintf("Couldnt' migrate up (migrations): %s\n", err))
	}
	// Down must be a noop (no .down.sql files)
	// Up should be idempotent so they don't break running tests.
	// NOTE: might have deadlocks, ideally should use some marker file like before
	// to prevent both migrations and post-migrations being run after the first
	// test suite that got the lock ran them and released it (next suite wouldn't be aware and run this again,
	// we would need the first suite to run indefinitely until all others end, and only then remove the marker file.)
	// we would still have the issue to remove that file before running go tests out of `project` (vscode, regular shell call...)
	//
	if err = mPostMigrations.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(fmt.Sprintf("Couldnt' migrate down (post-migrations): %s\n", err))
	}
	if err = mPostMigrations.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(fmt.Sprintf("Couldnt' migrate up (post-migrations): %s\n", err))
	}

	return pool, sqlpool, nil
}
