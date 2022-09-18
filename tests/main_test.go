package tests

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
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

	if err := envvar.Load("../.env.dev"); err != nil {
		fmt.Fprintf(os.Stderr, "envvar.Load: %s\n", err)
		os.Exit(1)
	}

	cmd := exec.Command(
		"bash", "-c",
		"source .envrc",
	)
	cmd.Dir = ".."
	if out, err := cmd.CombinedOutput(); err != nil {
		errAndExit(out, err)
	}

	cmd = exec.Command(
		"bin/project",
		"test.backend-setup",
	)
	cmd.Dir = ".."
	if out, err := cmd.CombinedOutput(); err != nil {
		errAndExit(out, err)
	}

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

func errAndExit(out []byte, err error) {
	fmt.Fprintf(os.Stderr, "combined out:\n%s\n", string(out))
	fmt.Fprintf(os.Stderr, "cmd.Run() failed with %s\n", err)
	os.Exit(1)
}
