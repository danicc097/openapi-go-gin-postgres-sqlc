package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/format"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

func main() {
	var env string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.Parse()

	if err := envvar.Load(env); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}

	appEnv := envvar.GetEnv("APP_ENV", "dev")
	if err := envvar.Load(path.Join(".env." + appEnv)); err != nil {
		fmt.Fprintf(os.Stderr, "envvar.Load: %s\n", err)
		os.Exit(1)
	}

	cmd := exec.Command(
		"bash", "-c",
		"source .envrc",
	)
	cmd.Dir = "."
	if out, err := cmd.CombinedOutput(); err != nil {
		errAndExit(out, err)
	}

	cmd = exec.Command(
		"bash", "-c",
		"project db.initial-data",
	)
	cmd.Dir = "."
	if out, err := cmd.CombinedOutput(); err != nil {
		errAndExit(out, err)
	}

	conf := envvar.New()
	pool, err := postgresql.New(conf)
	if err != nil {
		log.Fatalf("postgresql.New: %s\n", err)
	}

	username := "user_2"
	// username := "doesntexist" // User should be nil
	// username := "superadmin"
	user, err := db.UserByUsername(context.Background(), pool, username,
		db.UserWithJoin(db.UserJoins{
			TimeEntries: true,
			UserAPIKey:  true,
			WorkItems:   true,
			Teams:       true,
		}))
	if err != nil {
		log.Fatalf("db.UserByUsername: %s\n", err)
	}
	format.PrintJSON(user)
}

func errAndExit(out []byte, err error) {
	fmt.Fprintf(os.Stderr, "combined out:\n%s\n", string(out))
	fmt.Fprintf(os.Stderr, "cmd.Run() failed with %s\n", err)
	os.Exit(1)
}
