package main

import (
	"path"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgen"
)

func main() {
	var (
		baseDir = "internal"
		conf    = &postgen.Conf{
			CurrentHandlersDir: path.Join(baseDir, "handlers"),
			GenHandlersDir:     path.Join(baseDir, "gen"),
			OutHandlersDir:     path.Join(baseDir, "handlers"),
			OutServicesDir:     path.Join(baseDir, "services"),
		}
	)

	og := postgen.NewOpenapiGenerator(conf)
	og.Generate()
}
