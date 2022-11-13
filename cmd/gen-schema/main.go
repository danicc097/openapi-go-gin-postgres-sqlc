package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/swaggest/openapi-go/openapi3"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// clear && go run cmd/gen-schema/main.go -env .env.dev
func main() {
	reflector := openapi3.Reflector{}

	// we could load the existing spec to reflector.Spec: https://pkg.go.dev/github.com/swaggest/openapi-go/openapi3#example-Spec.UnmarshalYAML
	// and if "x-db-struct" found in an OPERATION
	reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}

	reflector.Spec.Info.
		WithTitle("Things API").
		WithVersion("1.2.3").
		WithDescription("Put something here")

	type req struct {
		ID     string `path:"id" example:"XXX-XXXXX"`
		Locale string `query:"locale" pattern:"^[a-z]{2}-[A-Z]{2}$"`
		Title  string `json:"string"`
		Amount uint   `json:"amount"`
		Items  []struct {
			Count uint   `json:"count"`
			Name  string `json:"name"`
		} `json:"items,omitempty"`
	}

	type resp struct {
		ID     string `json:"id" example:"XXX-XXXXX"`
		Amount uint   `json:"amount"`
		Items  []struct {
			Count uint   `json:"count"`
			Name  string `json:"name"`
		} `json:"items,omitempty"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	// we can edit an existing op by getting all operations with "x-db-struct", trace back to the openapi3.Operation
	putOp := openapi3.Operation{}

	handleError(reflector.SetRequest(&putOp, new(req), http.MethodPut))
	handleError(reflector.SetJSONResponse(&putOp, new(db.User), http.StatusOK))
	handleError(reflector.SetJSONResponse(&putOp, new([]db.User), http.StatusConflict))
	handleError(reflector.Spec.AddOperation(http.MethodPut, "/things/{id}", putOp))

	schema, err := reflector.Spec.MarshalYAML()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(schema))
}
