package main

import (
	"os"

	postgen "github.com/danicc097/openapi-go-gin-postgres-sqlc/postgen"
)

func main() {
	// TODO generate api_*.go, parse them and extract handlers then delete all  --> get list of operation ids and delete
	// parse handler funcs the same way in handlers/*.go -->  get list of operation ids
	// op ids not there --> generate not_implemented.go if not exists with handler funcs, else parse again and append the ones that dont exist.
	// all return notimplemented status, implement at dev's discretion

	// handlers should contain all handlers that do not exist in api_*.go files but exist in handlers/*.go (excluding not_implemented.go)
	handlers := []postgen.Handler{
		{
			OperationId: "MyGeneratedOperationId1",
			Comment:     "MyGeneratedOperationId1 has this cool comment.",
		},
		{
			OperationId: "MyGeneratedOperationId2",
			Comment:     "MyGeneratedOperationId2 has this cool comment.",
		},
	}
	// out to handlers/not_implemented.go and overwrite
	postgen.GenNotImpHandlers(handlers, os.Stdout)
}
