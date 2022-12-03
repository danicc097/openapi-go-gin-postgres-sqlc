package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgen"
)

func main() {
	var env, cacheDir, specPath string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.StringVar(&cacheDir, "cachedir", ".postgen.cache", "Cache dir")
	flag.StringVar(&specPath, "spec", "openapi.yaml", "OpenAPI specification")
	flag.Parse()

	if err := envvar.Load(env); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}

	var stderr bytes.Buffer

	url := internal.BuildAPIURL(specPath)

	if err := postgen.SetupSwaggerUI(url, specPath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, stderr.String())
		os.Exit(1)
	}
}
