package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pregen"
)

func main() {
	var env, cacheDir, spec, opIDAuthPath string

	flag.StringVar(&opIDAuthPath, "op-id-auth", "", "JSON file with authorization information per operation ID")
	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.StringVar(&cacheDir, "cachedir", ".postgen.cache", "Cache dir")
	flag.StringVar(&spec, "spec", "openapi.yaml", "OpenAPI specification")
	flag.Parse()

	if opIDAuthPath == "" {
		log.Fatal("op-id-auth flag is required")
	}

	if err := envvar.Load(env); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}

	var stderr bytes.Buffer
	pg := pregen.New(&stderr, spec, opIDAuthPath)
	if err := pg.Generate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, stderr.String())
		os.Exit(1)
	}
}
