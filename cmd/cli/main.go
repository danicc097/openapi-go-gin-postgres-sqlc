package main

import (
	"context"
	"flag"
	"log"

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

	conf := envvar.New()
	pool, err := postgresql.New(conf)
	if err != nil {
		log.Fatalf("postgresql.New: %s\n", err)
	}

	username := "user_1"
	// username := "doesntexist" // User should be nil
	// username := "superadmin"
	user, err := db.UserByUsername(context.Background(), pool, username,
		db.UserWithJoin(db.UserJoins{
			TimeEntries: true,
			UserAPIKey:  true,
			WorkItems:   false,
			Teams:       false,
		}))
	if err != nil {
		log.Fatalf("db.UserByUsername: %s\n", err)
	}
	format.PrintJSON(user)
}
