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
	"strings"

	// kinopenapi3 "github.com/getkin/kin-openapi/openapi3"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/google/uuid"
	"github.com/swaggest/jsonschema-go"
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

		// can intercept generated name full package path
		// fmt.Fprintf(os.Stderr, "t.PkgPath(): %v\n", t.PkgPath())

		return defaultDefName
	})

	// see https://github.com/swaggest/openapi-go/discussions/62#discussioncomment-5710581
	reflector.DefaultOptions = append(reflector.DefaultOptions,
		jsonschema.InterceptProp(func(params jsonschema.InterceptPropParams) error {
			if params.Field.Tag.Get("openapi-go") == "ignore" {
				return jsonschema.ErrSkipProperty
			}

			// reproduce: gen-schema --struct-names DbProject | yq 'with_entries(select(.key == "components"))'
			if shouldSkipType(params.Field.Type) {
				fmt.Fprintf(os.Stderr, "skipping schema: %s", params.Name)

				return jsonschema.ErrSkipProperty
			}

			// NOTE: forget about this, remove extra schemas manually that are referenced in db models
			// (which are themselves already generated from the spec, most likely), since
			// gen-schema doesn't know about external components we can't skip them beforehand.
			// if params.PropertySchema != nil {
			// 	if params.PropertySchema.Ref != nil {
			// 		fmt.Fprintf(os.Stderr, "params.PropertySchema.Ref: %v\n", *params.PropertySchema.Ref)
			// 		// if we ErrSkipProperty, we don't get the property. we just want to skip
			// 		// schema generation.
			// 		// ideally, it would try to generate, and skip if ref already exists
			// 	}
			// }

			return nil
		}),
		jsonschema.InterceptSchema(func(params jsonschema.InterceptSchemaParams) (stop bool, err error) {
			if params.Schema.ReflectType == reflect.TypeOf(uuid.New()) {
				params.Schema.Type = &jsonschema.Type{SimpleTypes: pointers.New(jsonschema.String)}
				params.Schema.Items = &jsonschema.Items{}
			}

			return false, nil
		}),
	)

	for _, sn := range structNames {
		dummyOp := openapi3.Operation{}
		// TODO we need to compile gen-schema right after PublicStructs file is updated
		// cannot import packages at runtime
		// if we have an uncompilable state then we need workarounds to compile regardless
		// UDPATE: use https://pkg.go.dev/plugin
		st, ok := postgen.PublicStructs[sn]
		if !ok {
			log.Fatalf("struct-name %s does not exist in PublicStructs", sn)
		}
		if !hasJSONTag(st) {
			log.Fatalf("struct %s: ensure there is at least a JSON tag set", sn)
		}

		// st = cloneStructWithoutIgnoredFields(st)

		handleError(reflector.SetJSONResponse(&dummyOp, st, http.StatusTeapot))
		// not really needed
		// handleError(reflector.Spec.AddOperation(http.MethodGet, "/dummy-op-"+strconv.Itoa(i), dummyOp))

		// IMPORTANT: ensure structs are public
		reflector.Spec.Components.Schemas.MapOfSchemaOrRefValues[sn].Schema.MapOfAnything = map[string]any{"x-postgen-struct": sn}
	}
	s, err := reflector.Spec.MarshalYAML()
	handleError(err)

	fmt.Println(string(s))
}

func hasJSONTag(input any) bool {
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

func cloneStructWithoutIgnoredFields(src any) any {
	srcType := reflect.TypeOf(src)

	dstType := reflect.StructOf(filterIgnoredFields(srcType))
	// dstType = reflect.StructOf([]reflect.StructField{{
	// 	Name:    srcType.Name(),
	// 	PkgPath: srcType.PkgPath(),
	// 	Type:    dstType,
	// 	Tag:     `json:"activity" required:"true"`,
	// }})

	dst := reflect.New(dstType).Elem()

	return dst.Interface()
}

// filterIgnoredFields returns all fields that do not have
// the "openapi-go" tag set to "ignore".
func filterIgnoredFields(t reflect.Type) []reflect.StructField {
	var filteredFields []reflect.StructField
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if val, ok := field.Tag.Lookup("openapi-go"); !ok || val != "ignore" {
			if field.Type.Kind() == reflect.Struct {
				subFields := filterIgnoredFields(field.Type)
				field.Type = reflect.StructOf(subFields)
			}
			filteredFields = append(filteredFields, field)
		}
	}

	return filteredFields
}

func shouldSkipType(typ reflect.Type) bool {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Array || typ.Kind() == reflect.Slice {
		return shouldSkipType(typ.Elem())
	}

	// return strings.HasSuffix(typ.PkgPath(), "/internal/models")

	return false
}
