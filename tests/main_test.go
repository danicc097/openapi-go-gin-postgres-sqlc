package tests

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool
var srv *http.Server

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	os.Setenv("POSTGRES_DB", "postgres_test")
	os.Setenv("IS_TESTING", "1")
	// call flag.Parse() here if TestMain uses flags
	var err error

	pool, err = newDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create pool: %s\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	envFile := fmt.Sprintf("../.env.%s", os.Getenv("APP_ENV"))
	spec := "../openapi.yaml"
	srv, err = run(envFile, ":0", spec, pool)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't run test server: %s\n", err)
		os.Exit(1)
	}
	defer srv.Close()

	return m.Run()
}
