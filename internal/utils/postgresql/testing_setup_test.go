package postgresql_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

const (
	errNoRows                  = "no rows in result set"
	errViolatesCheckConstraint = "violates check constraint"
)

var (
	pool    *pgxpool.Pool
	sqlPool *sql.DB // for jet, use .Sql() to use pgx directly
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	testutil.Setup()

	// call flag.Parse() here if TestMain uses flags
	var err error

	logger, _ := zap.NewDevelopment()
	defaultPool, _, err := postgresql.New(logger.Sugar())
	if err != nil {
		panic(fmt.Sprintf("Couldn't create default pool: %s\n", err))
	}
	defer defaultPool.Close()

	dbName := "postgres_test"
	if _, err := defaultPool.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %s;", dbName)); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			panic(fmt.Sprintf("Couldn't create database: %s\n", err))
		}
	}

	pool, sqlPool, err = postgresql.New(logger.Sugar(), postgresql.WithDBName(dbName))
	if err != nil {
		panic(fmt.Sprintf("Couldn't create pool: %s\n", err))
	}

	// if it blocks here, we are not properly releasing connections in tests
	// leading to max conn count reached
	defer pool.Close()

	return m.Run()
}
