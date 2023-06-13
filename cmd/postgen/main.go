package main

import (
	"flag"
	"log"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
)

func main() {
	var env, specPath string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.StringVar(&specPath, "spec", "openapi.yaml", "OpenAPI specification")
	flag.Parse()

	if err := envvar.Load(env); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}
}
