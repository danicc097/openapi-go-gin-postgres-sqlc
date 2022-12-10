package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pregen"
)

func main() {
	var env, cacheDir, spec, opIDAuthPath string
	var validateSpecOnly bool
	var stderr bytes.Buffer

	flag.StringVar(&opIDAuthPath, "op-id-auth", "", "JSON file with authorization information per operation ID")
	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.StringVar(&cacheDir, "cachedir", ".postgen.cache", "Cache dir")
	flag.StringVar(&spec, "spec", "openapi.yaml", "OpenAPI specification")
	flag.BoolVar(&validateSpecOnly, "validate-spec-only", false, "OpenAPI specification")
	flag.Parse()

	if err := envvar.Load(env); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}

	if err := internal.NewAppConfig(); err != nil {
		log.Fatalf("internal.NewAppConfig: %s\n", err)
	}

	if validateSpecOnly {
		pg := pregen.New(&stderr, spec, opIDAuthPath)
		if err := pg.ValidateProjectSpec(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintln(os.Stderr, stderr.String())
			os.Exit(1)
		}
		os.Exit(0)
	}

	if opIDAuthPath == "" {
		log.Fatal("op-id-auth flag is required")
	}

	pg := pregen.New(&stderr, spec, opIDAuthPath)
	if err := pg.Generate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, stderr.String())
		os.Exit(1)
	}
}
