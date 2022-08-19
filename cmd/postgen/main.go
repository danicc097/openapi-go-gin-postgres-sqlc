package main

import (
	"bytes"
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
	og.Generate()
}
