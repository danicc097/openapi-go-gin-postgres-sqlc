package main

import (
	"flag"
	"log"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
)

func main() {
	var env, specPath string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.StringVar(&specPath, "spec-path", "openapi.yaml", "OpenAPI specification filepath")
	flag.Parse()

	if err := envvar.Load(env); err != nil {
		log.Fatalf("Couldn't load env vars: %s", err)
	}

	if err := internal.SetupSwaggerUI(internal.BuildAPIURL(specPath), specPath); err != nil {
		log.Fatalf("Couldn't setup Swagger UI: %s", err)
	}
}
