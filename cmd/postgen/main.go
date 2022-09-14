package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/format"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgen"
)

func main() {
	var env, cacheDir string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.StringVar(&cacheDir, "cachedir", ".postgen.cache", "Cache dir")
	flag.Parse()

	if err := envvar.Load(env); err != nil {
		log.Fatalf("envvar.Load: %s", err)
	}

	const baseDir = "internal"
	conf := &postgen.Conf{
		CurrentHandlersDir: path.Join(baseDir, "rest"),
		GenHandlersDir:     path.Join(baseDir, "gen"),
		OutHandlersDir:     path.Join(baseDir, "rest"),
		OutServicesDir:     path.Join(baseDir, "services"),
	}

	var stderr bytes.Buffer
	og := postgen.NewOpenapiGenerator(conf, &stderr, cacheDir)

	if err := og.Generate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, stderr.String())
		os.Exit(1)
	}

	url := format.BuildBackendURL("openapi.yaml")

	postgen.SetupSwaggerUI(url)
}
