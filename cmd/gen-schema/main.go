package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	// kinopenapi3 "github.com/getkin/kin-openapi/openapi3"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgen"
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
	var structNamesList string

	flag.StringVar(&structNamesList, "struct-names", "", "comma-delimited db package struct names to generate an OpenAPI schema for")
	flag.Parse()

	structNames := strings.Split(structNamesList, ",")
	for i := range structNames {
		structNames[i] = strings.TrimSpace(structNames[i])
	}

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
	// IMPORTANT:
	// we only need to get Db** generated structs and replace in our spec. (we already have a reference in our openapi.yaml operation, leading to an empty schema - else spec wont compile - so all thats left is to replace the schema with the generated one.)
	// we can use yq for this (plain replace schema by name) and forget about an openapi.gen.yaml that messes things up.
	// gen-schema cli should generate a new yaml file with dummy operations (/dummy-$i), i++ while reflector is updated, as we do now (openapi.test.gen.yaml), so **yq reads the schema there, and replaces it in our openapi.yaml**
	// we dont have to do anything else!
	for i, sn := range structNames {
		dummyOp := openapi3.Operation{}
		st, ok := postgen.DbStructs[sn]
		if !ok {
			log.Fatalf("struct-name %s does not exist in db package", sn)
		}
		handleError(reflector.SetJSONResponse(&dummyOp, st, http.StatusTeapot))
		handleError(reflector.Spec.AddOperation(http.MethodGet, "/dummy-op-"+strconv.Itoa(i), dummyOp))
		// reflector.Spec.Paths.MapOfPathItemValues["mypath"].MapOfOperationValues["method"].
	}
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
