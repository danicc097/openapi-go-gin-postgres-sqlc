// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package main

import (
	"flag"
	"log"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/server"
)

func main() {
	var env, address, specPath string

	flag.StringVar(&env, "env", "", "Environment Variables filename")
	flag.StringVar(&address, "address", ":8090", "HTTP Server Address")
	flag.StringVar(&specPath, "spec-path", "openapi.yaml", "OpenAPI specification filepath")
	flag.Parse()

	errC, err := server.Run(env, address, specPath)
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}
