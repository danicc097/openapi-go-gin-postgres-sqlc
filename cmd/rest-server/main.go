package main

import (
	"io/fs"
	"log"
	"net/http"

	gen "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
)

func main() {
	log.Printf("Server started")

	router := gen.NewRouter()

	// TODO defining static file serving in spec is not supported?
	fsys, _ := fs.Sub(static.SwaggerUI, "swagger-ui")
	router.StaticFS("/v2/docs", http.FS(fsys))

	log.Fatal(router.Run(":8090"))
}
