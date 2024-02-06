package testutil

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	postgresqlutils "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/postgresql"
	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/iancoleman/strcase"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var once sync.Once

type TestDBOptions struct {
	WithMigrations bool
}

type TestDBOption func(*TestDBOptions)

// WithMigrations runs migrations on top of the database.
func WithMigrations() TestDBOption {
	return func(opt *TestDBOptions) {
		opt.WithMigrations = true
	}
}

// NewDB returns a new testing Postgres pool with up-to-date migrations.
// It is shared within the test suite package.
// Panics on any error encountered.
func NewDB(options ...TestDBOption) (*pgxpool.Pool, *sql.DB, error) {
	_, b, _, _ := runtime.Caller(1)
	dir := path.Join(path.Dir(b))

	opts := &TestDBOptions{}
	for _, option := range options {
		option(opts)
	}

	pre := "postgres_test_"
	d := strcase.ToSnake(strcase.ToCamel(dir))
	dbName := pre + d
	if len(dbName) > 63 {
		dbName = pre + d[len(d)-63+len(pre):] // max postgres identifier length
	}

	logger, _ := zap.NewDevelopment()

	DropAndRecreateDB(dbName)

	pool, sqlpool, err := postgresql.New(logger.Sugar(), postgresql.WithDBName(dbName))
	if err != nil {
		panic(fmt.Sprintf("Couldn't create pool: %s\n", err))
	}

	if !opts.WithMigrations {
		return pool, sqlpool, nil
	}

	migrationsLockID, _ := strconv.ParseInt(dbName, 10, 32)

	lock, err := postgresqlutils.NewAdvisoryLock(pool, int(migrationsLockID))
	if err != nil {
		panic(fmt.Sprintf("NewAdvisoryLock: %s\n", err))
	}

	acquired, err := lock.TryLock(context.Background())
	if err != nil {
		panic(fmt.Sprintf("lock.TryLock: %s\n", err))
	}
	if !acquired { // this probably won't be a path anymore since using test suite name for db name
		fmt.Println("Waiting for migrations to run in test suite: " + dir)
		if err := lock.WaitForRelease(50, 200*time.Millisecond); err != nil {
			panic(fmt.Sprintf("lock.WaitForRelease: %s\n", err))
		}

		return pool, sqlpool, nil
	}

	fmt.Println("Running migrations in test suite: " + dir)

	defer func() {
		unlockSuccess := lock.Release()
		for i := 0; !unlockSuccess && lock.IsLocked() && i < 10; i++ {
			unlockSuccess = lock.Release()
		}
		lock.ReleaseConn()
		// if lock.IsLocked() {
		// 	// FIXME: if sharing db for all test suices: race condition -> lock.IsLocked() was false right above when releasing,
		// 	// but then a new test suite came in and grabbed it.
		// 	// should use transactions internally for Release(), and instead of unlockSuccess
		// 	// it should check that indeed the lock was released (at the time)
		// 	panic(fmt.Sprintf("advisory lock was not released\n"))
		// }
	}()

	if internal.Config.AppEnv == internal.AppEnvCI {
		printMigrationsState(pool)
	}

	driver, err := migratepostgres.WithInstance(sqlpool, &migratepostgres.Config{MigrationsTable: "schema_migrations"})
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
	mPostMigrations, err := migrate.NewWithDatabaseInstance("file://"+path.Join(path.Dir(src), "../../db/post-migrations/"), "postgres", postDriver)
	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate (migrations): %s\n", err))
	}

	mMigrations, err := migrate.NewWithDatabaseInstance("file://"+path.Join(path.Dir(src), "../../db/migrations/"), "postgres", driver)
	if err != nil {
		panic(fmt.Sprintf("Couldn't migrate (post-migrations): %s\n", err))
	}

	if err = mMigrations.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(fmt.Sprintf("Couldnt' migrate down: %s\n", err))
	}
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
	// TODO: may use "go.testEnvFile": null,
	// which will be read by all .env will
	// TODO: should just be done once, like regular migrations table down
	// if err = mPostMigrations.Force(1); err != nil && !errors.Is(err, migrate.ErrNoChange) { // no down.sql files on purpose
	// 	panic(fmt.Sprintf("Couldnt' force migrate down to 1 (post-migrations): %s\n", err))
	// }

	if err = mPostMigrations.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(fmt.Sprintf("Couldnt' migrate up (post-migrations): %s\n", err))
	}

	return pool, sqlpool, nil
}

// DropAndRecreateDB drops and recreates a given database.
func DropAndRecreateDB(dbName string) {
	defaultDb := "postgres"
	if dbName == defaultDb {
		panic(fmt.Sprintf("cannot drop default %q database", defaultDb))
	}

	cfg := internal.Config
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.Postgres.User, cfg.Postgres.Password),
		Host:   fmt.Sprintf("%s:%s", cfg.Postgres.Server, strconv.Itoa(cfg.Postgres.Port)),
		Path:   defaultDb,
	}

	q := dsn.Query()
	q.Add("sslmode", os.Getenv("DATABASE_SSLMODE"))
	dsn.RawQuery = q.Encode()

	conn, err := pgx.Connect(context.Background(), dsn.String())
	if err != nil {
		panic(fmt.Sprintf("Couldn't connect to the database: %s\n", err))
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName))
	if err != nil {
		panic(fmt.Sprintf("Couldn't drop database %q: %s\n", dbName, err))
	}

	_, err = conn.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %s;", dbName))
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			panic(fmt.Sprintf("Couldn't create database %q: %s\n", dbName, err))
		}
	}
}

func printMigrationsState(pool *pgxpool.Pool) {
	query := `
	select
		row_to_json(schema_migrations.*) as schema_migrations,
		row_to_json(schema_post_migrations.*) as schema_post_migrations
	from
		schema_migrations,
		schema_post_migrations
`
	res, err := postgresql.DynamicQuery(pool, query)
	if err != nil {
		fmt.Printf("postgresql.DynamicQuery error: %s\n", err)
	}

	fmt.Printf("Current migrations:\n%s\n", res)
}
