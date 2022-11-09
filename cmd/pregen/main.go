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
	var env, cacheDir, spec string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.StringVar(&cacheDir, "cachedir", ".postgen.cache", "Cache dir")
	flag.StringVar(&spec, "spec", "openapi.yaml", "OpenAPI specification")
	flag.Parse()

	if err := envvar.Load(env); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}

	// TODO read openapi spec, look for x-db-enum or x-db-tables vendor ext.
	// and replace enum values with sql query output
	// TODO	messy, could all be done with docker exec and yq...

	var stderr bytes.Buffer
	pg := pregen.New(&stderr, spec)
	if err := pg.Generate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, stderr.String())
		os.Exit(1)
	}
}
