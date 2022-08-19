package main

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgen"
)

func main() {
	const baseDir = "internal"
	conf := &postgen.Conf{
		CurrentHandlersDir: path.Join(baseDir, "handlers"),
		GenHandlersDir:     path.Join(baseDir, "gen"),
		OutHandlersDir:     path.Join(baseDir, "handlers"),
		OutServicesDir:     path.Join(baseDir, "services"),
	}

	var stderr bytes.Buffer
	og := postgen.NewOpenapiGenerator(conf, &stderr)

	if err := og.Generate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, stderr.String())
		os.Exit(1)
	}
}
