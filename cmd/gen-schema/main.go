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

	// see https://github.com/swaggest/openapi-go/discussions/62#discussioncomment-5710581
	reflector.DefaultOptions = append(reflector.DefaultOptions,
		jsonschema.InterceptDefName(func(t reflect.Type, defaultDefName string) string {
			if strings.HasSuffix(t.PkgPath(), "internal/models") {
				fmt.Fprintf(os.Stderr, "Generated models package type found in spec: %+v\n", t)
			}
			return defaultDefName
		}),
		jsonschema.InterceptProp(func(params jsonschema.InterceptPropParams) error {
			if params.Field.Tag.Get("openapi-go") == "ignore" {
				return jsonschema.ErrSkipProperty
			}

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
