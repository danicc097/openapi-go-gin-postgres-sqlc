package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	// kinopenapi3 "github.com/getkin/kin-openapi/openapi3"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/swaggest/openapi-go/openapi3"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type req struct {
	ID     string `path:"id" example:"XXX-XXXXX"`
	Locale string `query:"locale" pattern:"^[a-z]{2}-[A-Z]{2}$"`
	Title  string `json:"string"`
	Amount uint   `json:"amount"`
	Items  []struct {
		Count uint   `json:"count"`
		Name  string `json:"name"`
	} `json:"items,omitempty"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
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

func main() {
	var structName string

	flag.StringVar(&structName, "struct-name", "", "db package struct name to generate an OpenAPI schema for")
	flag.Parse()

	reflector := openapi3.Reflector{Spec: &openapi3.Spec{}}

	// we could load the existing spec to reflector.Spec: https://pkg.go.dev/github.com/swaggest/openapi-go/openapi3#example-Spec.UnmarshalYAML
	// and if "x-db-struct" found in an OPERATION.
	// NOTE: comments are gone. should print result to openapi.gen.yaml and use that since this is in gen/postgen step
	schemaBlob, err := os.ReadFile("openapi.yaml")
	if err != nil {
		log.Fatalf("openapi spec: %s", err)
	}
	if err := reflector.Spec.UnmarshalYAML(schemaBlob); err != nil {
		log.Fatalf("Spec.UnmarshalYAML: %s", err)
	}

	// we can edit an existing op by getting all operations with "x-db-struct", trace back to the openapi3.Operation
	putOp := openapi3.Operation{}

	handleError(reflector.SetRequest(&putOp, new(req), http.MethodPut))
	handleError(reflector.SetJSONResponse(&putOp, new(db.User), http.StatusOK))
	st, ok := postgen.DbStructs[structName]
	if !ok {
		log.Fatalf("struct-name %s does not exist in db package", structName)
	}
	handleError(reflector.SetJSONResponse(&putOp, st, http.StatusConflict))
	handleError(reflector.Spec.AddOperation(http.MethodPut, "/things/{id}", putOp))
	// reflector.Spec.Paths.MapOfPathItemValues["mypath"].MapOfOperationValues["method"].
	schema, err := reflector.Spec.MarshalYAML()
	handleError(err)

	// fmt.Println(string(schema))

	os.WriteFile("openapi.test.gen.yaml", schema, 0o777)

	// var s openapi3.Spec
	// have to merge with our own (done easily but comments are lost. output to openapi.gen.yaml, we are in postgen step so doesnt matter)
	// have to change frontend generation, swaggerUI paths
	// schemaBlob, err := os.ReadFile("openapi.yaml")
	// if err != nil {
	// 	log.Fatalf("error opening schema file: %s", err)
	// }

	// oas, err := rest.ReadOpenAPI("openapi.yaml")
	// if err != nil {
	// 	log.Fatalf("ReadOpenAPI: %s", err)
	// }
	// // oas.Components.SecuritySchemes = kinopenapi3.SecuritySchemes{} // error

	// // fmt.Println(string(schemaBlob))
	// specWithoutSec, err := oas.MarshalJSON()
	// if err != nil {
	// 	log.Fatalf("oas.MarshalJSON: %s", err)
	// }
	// // fmt.Println(string(specWithoutSec))

	// if err := s.UnmarshalYAML(specWithoutSec); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(s.Info.Title)
	// fmt.Println(s.Info.Title)
	// fmt.Println(s.Components.Schemas.MapOfSchemaOrRefValues["Error"].Schema.Properties["code"].Schema.MapOfAnything["x-foo"])
}
