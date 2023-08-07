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

	// kinopenapi3 "github.com/getkin/kin-openapi/openapi3".

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/codegen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	internalslices "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/slices"
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
		// jsonschema.InterceptNullability(func(params jsonschema.InterceptNullabilityParams) {
		// 	if params.Type.Kind() != reflect.Struct {
		// 		return
		// 	}
		// 	for i := 0; i < params.Type.NumField(); i++ {
		// 		if params.Schema != nil && params.Schema.Type != nil {
		// 			if params.Type.Field(i).Tag.Get("nullable") == "false" {
		// 				if types := params.Schema.Type.SliceOfSimpleTypeValues; len(types) > 0 {
		// 					fmt.Fprintf(os.Stderr, "nullable schema: %s\n", params.Schema.ReflectType.Name())
		// 				}
		// 			}
		// 		}
		// 	}
		// }),
		jsonschema.InterceptProp(func(params jsonschema.InterceptPropParams) error {
			if params.Field.Tag.Get("x-omitempty") == "true" {
				if params.PropertySchema == nil {
					return nil
				}
				if params.PropertySchema.ExtraProperties == nil {
					params.PropertySchema.ExtraProperties = map[string]any{}
				}
				params.PropertySchema.ExtraProperties["x-omitempty"] = true
			}

			if params.PropertySchema != nil && params.PropertySchema.Type != nil {
				if params.Field.Tag.Get("nullable") == "false" {
					if types := params.PropertySchema.Type.SliceOfSimpleTypeValues; len(types) > 0 {
						// fmt.Fprintf(os.Stderr, "nullable schema: %s\n", params.Name)
						params.PropertySchema.Type.SliceOfSimpleTypeValues = internalslices.Filter(types, func(item jsonschema.SimpleType, _ int) bool {
							return item != jsonschema.Null
						})
					}
				}
			}

			return nil
		}),
		jsonschema.InterceptSchema(func(params jsonschema.InterceptSchemaParams) (stop bool, err error) {
			if params.Schema.ReflectType == reflect.TypeOf(uuid.New()) {
				params.Schema.Type = &jsonschema.Type{SimpleTypes: pointers.New(jsonschema.String)}
				params.Schema.Pattern = pointers.New("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
				params.Schema.Items = &jsonschema.Items{}
				// x-go* extensions cannot be used for Models(.*) themselves,
				// but Models(.*) should not be generated at all. a ref tag is needed in structs
				params.Schema.ExtraProperties = map[string]any{
					"x-go-type": "uuid.UUID",
					"x-go-type-import": map[string]any{
						"name": "uuid",
						"path": "github.com/google/uuid",
					},
				}
			}

			return false, nil
		}),
	)

	for i, sn := range structNames {
		// we need to compile gen-schema right after PublicStructs file is updated
		// cannot import packages at runtime
		// if we have an uncompilable state then we need to update src to compile. no way around it
		// UDPATE: use https://github.com/pkujhd/goloader instead of plugin pkg which cant reload changed go file at runtime
		// or use yaegi
		st, ok := codegen.PublicStructs[sn]
		if !ok {
			log.Fatalf("struct-name %s does not exist in PublicStructs", sn)
		}
		if !hasJSONTag(st) {
			log.Fatalf("struct %s: ensure there is at least a JSON tag set", sn)
		}

		oc, err := reflector.NewOperationContext(http.MethodGet, "/dummy-op-"+strconv.Itoa(i))
		handleError(err)
		oc.AddRespStructure(st)

		handleError(reflector.AddOperation(oc))

		// IMPORTANT: ensure structs are public
		reflector.Spec.Components.Schemas.MapOfSchemaOrRefValues[sn].Schema.MapOfAnything = map[string]any{"x-postgen-struct": sn}
	}
	s, err := reflector.Spec.MarshalYAML()
	handleError(err)

	fmt.Println(string(s))
}

func hasJSONTag(input any) bool {
	t := reflect.TypeOf(input)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

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
