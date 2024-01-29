package postgresql_test

import (
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
	pool, sqlPool, err = postgresql.New(logger.Sugar())
	if err != nil {
		panic(fmt.Sprintf("Couldn't create pool: %s\n", err))
	}

	// if it blocks here, we are not properly releasing connections in tests
	// leading to max conn count reached
	defer pool.Close()

	return m.Run()
}
