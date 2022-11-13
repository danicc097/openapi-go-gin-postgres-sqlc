package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	kinopenapi3 "github.com/getkin/kin-openapi/openapi3"
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
	// and if "x-db-struct" found in an OPERATION.
	// NOTE: comments are gone. should print result to openapi.gen.yaml and use that since this is in gen/postgen step
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

	var s openapi3.Spec

	schemaBlob, err := os.ReadFile("openapi.yaml")
	if err != nil {
		log.Fatalf("error opening schema file: %s", err)
	}

	oas, err := rest.ReadOpenAPI("openapi.yaml")
	if err != nil {
		log.Fatalf("ReadOpenAPI: %s", err)
	}
	oas.Components.SecuritySchemes = kinopenapi3.SecuritySchemes{} // error

	fmt.Println(string(schemaBlob))
	specWithoutSec, err := oas.MarshalJSON()
	if err != nil {
		log.Fatalf("oas.MarshalJSON: %s", err)
	}
	fmt.Println(string(specWithoutSec))

	if err := s.UnmarshalYAML(specWithoutSec); err != nil {
		log.Fatal(err)
	}

	fmt.Println(s.Info.Title)
	// fmt.Println(s.Info.Title)
	// fmt.Println(s.Components.Schemas.MapOfSchemaOrRefValues["Error"].Schema.Properties["code"].Schema.MapOfAnything["x-foo"])
}
