package testutil

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	once           sync.Once
	markerFilePath = "/tmp/migration_marker"
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

	once.Do(func() {
		m, err := migrate.NewWithDatabaseInstance("file://"+path.Join(path.Dir(src), "../../db/migrations/"), "postgres", instance)
		if err != nil {
			fmt.Printf("Couldn't migrate (2): %s\n", err)
			return
		}

		// NOTE: migrate down before tests only externally.
		if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			fmt.Printf("Couldnt' migrate (3): %s\n", err)
			return
		}

		// post-migration scripts may not be idempotent like up migration command
		if _, err := os.Stat(markerFilePath); os.IsExist(err) {
			fmt.Println("Marker file exists, skipping post-migrations.")
			return
		}

		markerFile, err := os.Create(markerFilePath)
		if err != nil {
			fmt.Printf("Error creating marker file: %s\n", err)
			return
		}
		defer markerFile.Close()

		postMigrationPath := path.Join(path.Dir(src), "../../db/post-migration/")
		files, err := os.ReadDir(postMigrationPath)
		if err != nil {
			fmt.Printf("Error reading post-migration directory: %s\n", err)
			return
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			filePath := path.Join(postMigrationPath, file.Name())
			script, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Error reading post-migration script %s: %s\n", file.Name(), err)
				return
			}

			_, err = sqlpool.Exec(string(script))
			if err != nil {
				fmt.Printf("Error executing post-migration script %s: %s\n", file.Name(), err)
				return
			}

			fmt.Printf("Run post-migration script: %s\n", file.Name())
		}
	})

	return pool, sqlpool, nil
}
