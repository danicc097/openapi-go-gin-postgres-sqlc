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

	ctx := context.Background()

	conn, err := pool.Acquire(ctx)
	if err != nil {
		panic("could not acquire connection")
	}

	mustMigrate, err := tryAdvisoryLock(ctx, conn, pool, migrationsLockID)
	if err != nil {
		panic(fmt.Sprintf("advisory lock: %s\n", err))
	}
	fmt.Printf("mustMigrate: %v\n", mustMigrate)
	if !mustMigrate {
		return pool, sqlpool, nil
	}
	defer func() {
		if _, err := conn.Exec(ctx, `SELECT pg_advisory_unlock($1)`, migrationsLockID); err != nil {
			panic(fmt.Sprintf("Error in advisory lock cleanup: %s\n", err))
		}
		conn.Release()
	}()

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

// Returns whether we must migrate or not, and a cleanup func.
// If we must migrate, it acquires a lock which should be released afterwards
// If we must not migrate, it waits until lock is released, meaning migrations
// have been run by a concurrent process.
func tryAdvisoryLock(ctx context.Context, conn *pgxpool.Conn, pool *pgxpool.Pool, lockID int) (bool, error) {
	var lockExists bool

	// must use this same conn for release later if successful
	row := conn.QueryRow(ctx, `SELECT pg_try_advisory_lock($1)`, lockID)
	if err := row.Scan(&lockExists); err != nil {
		return false, fmt.Errorf("lock query: %w", err)
	}

	if !lockExists {
		return true, nil
	}

	checkLockQuery := `
	SELECT EXISTS (
			SELECT 1
			FROM pg_locks
			JOIN pg_stat_activity USING (pid)
			WHERE locktype = 'advisory' AND objid = $1
	) AS lock_acquired;
`
	for i := 0; i < 30; i++ {
		// Check if the lock is acquired using a separate connection
		row := pool.QueryRow(ctx, checkLockQuery, lockID)
		if err := row.Scan(&lockExists); err != nil {
			return false, fmt.Errorf("query: %w", err)
		}

		if !lockExists {
			return false, nil
		}

		time.Sleep(200 * time.Millisecond)
	}

	return false, fmt.Errorf("timeout waiting for lock release with objid %d", lockID)
}
