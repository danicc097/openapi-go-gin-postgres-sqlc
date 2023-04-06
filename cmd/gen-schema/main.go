/*
gen-schema generates OpenAPI v3 schema portions from code.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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

	reflector.InterceptDefName(func(t reflect.Type, defaultDefName string) string {
		// default name comes from package directory, not the given import alias
		// e.g. repomodels -/-> Repomodels, its the last dir (models)

		return defaultDefName
	})

	for i, sn := range structNames {
		dummyOp := openapi3.Operation{}
		st, ok := postgen.PublicStructs[sn]
		if !ok {
			log.Fatalf("struct-name %s does not exist in PublicStructs", sn)
		}
		if !hasJSONTag(st) {
			log.Fatalf("struct %s: ensure there is at least a JSON tag set", sn)
		}
		handleError(reflector.SetJSONResponse(&dummyOp, st, http.StatusTeapot))
		handleError(reflector.Spec.AddOperation(http.MethodGet, "/dummy-op-"+strconv.Itoa(i), dummyOp))

		// IMPORTANT: ensure structs are public
		reflector.Spec.Components.Schemas.MapOfSchemaOrRefValues[sn].Schema.MapOfAnything = map[string]interface{}{"x-postgen-struct": sn}
	}
	s, err := reflector.Spec.MarshalYAML()
	handleError(err)

	fmt.Println(string(s))
}

func hasJSONTag(input interface{}) bool {
	t := reflect.TypeOf(input)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if _, ok := field.Tag.Lookup("json"); ok {
			return true
		}

		// Check embedded structs
		if field.Type.Kind() == reflect.Struct && field.Anonymous {
			if hasJSONTag(reflect.New(field.Type).Elem().Interface()) {
				return true
			}
		}
	}

	return false
}
