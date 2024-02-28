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

	tempTestPool, _, err := postgresql.New(logger.Sugar())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create temporary pool: %s\n", err)
		os.Exit(1)
	}
	defer tempTestPool.Close()

	schemaPath := "xotests_schema.sql"
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't read %s: %s\n", schemaPath, err)
		return 1
	}

	_, err = tempTestPool.Exec(context.Background(), string(schema))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't execute %s: %s\n", schemaPath, err)
		return 1
	}

	// refresh pgxpool types now that xotests_schema.sql is loaded.
	// maybe a postgresl.New postgresql.WithSchemas(schemas...) executed in order is worth it
	// to avoid this if it's in demand
	testPool, _, err = postgresql.New(logger.Sugar())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create testPool: %s\n", err)
		return 1
	}
	defer testPool.Close()

	return m.Run()
}
