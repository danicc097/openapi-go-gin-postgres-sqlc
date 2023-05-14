package tests

import (
	"context"
	"database/sql"
	"fmt"
	"os"
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
	testPool    *pgxpool.Pool
	testSQLPool *sql.DB // for jet, use .Sql() to use pgx directly
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	testutil.Setup()

	// call flag.Parse() here if TestMain uses flags
	var err error

	logger, _ := zap.NewDevelopment()

	testPool, _, err = postgresql.New(logger.Sugar())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create testPool: %s\n", err)
		os.Exit(1)
	}
	defer testPool.Close()

	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't read schema.sql: %s\n", err)
		return 1
	}

	_, err = testPool.Exec(context.Background(), string(schema))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't read schema.sql: %s\n", err)
		return 1
	}

	return m.Run()
}
