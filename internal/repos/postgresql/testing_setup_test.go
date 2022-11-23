package postgresql_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testpool *pgxpool.Pool

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	testutil.Setup()

	// call flag.Parse() here if TestMain uses flags
	var err error

	testpool, err = testutil.NewDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create testpool: %s\n", err)
		os.Exit(1)
	}
	defer testpool.Close()

	return m.Run()
}
