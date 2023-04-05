/*
gen-schema generates OpenAPI v3 schema portions from code.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	// kinopenapi3 "github.com/getkin/kin-openapi/openapi3"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgen"
	"github.com/swaggest/openapi-go/openapi3"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
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

	// update when adding new packages to gen structs map
	reflector.InterceptDefName(func(t reflect.Type, defaultDefName string) string {
		// default name comes from package directory, not the given import alias
		// e.g. repomodels -/-> Repomodels, its the last dir (models)

		return defaultDefName
	})

	fmt.Fprintf(os.Stderr, "postgen.PublicStructs: %v\n", postgen.PublicStructs)

	keys := make([]string, 0, len(reflector.Spec.Components.Schemas.MapOfSchemaOrRefValues))
	for k := range reflector.Spec.Components.Schemas.MapOfSchemaOrRefValues {
		keys = append(keys, k)
	}
	fmt.Fprintf(os.Stderr, "keys: %v\n", keys)

	for i, sn := range structNames {
		dummyOp := openapi3.Operation{}
		st, ok := postgen.PublicStructs[sn]
		if !ok {
			log.Fatalf("struct-name %s does not exist in PublicStructs", sn)
		}
		fmt.Fprintf(os.Stderr, "sn: %v\n", sn)
		handleError(reflector.SetJSONResponse(&dummyOp, st, http.StatusTeapot))
		// ensure json tags are set for all top level fields in the struct. Either skip "-" or set fields.
		reflector.Spec.Components.Schemas.MapOfSchemaOrRefValues[sn].Schema.MapOfAnything = map[string]interface{}{"x-postgen-struct": sn}
		handleError(reflector.Spec.AddOperation(http.MethodGet, "/dummy-op-"+strconv.Itoa(i), dummyOp))
		// reflector.Spec.Paths.MapOfPathItemValues["mypath"].MapOfOperationValues["method"].
	}
	s, err := reflector.Spec.MarshalYAML()
	handleError(err)

	fmt.Println(string(s))
}
