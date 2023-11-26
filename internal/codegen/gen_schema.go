/*
gen-schema generates OpenAPI v3 schema portions from code.
*/
package codegen

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	// kinopenapi3 "github.com/getkin/kin-openapi/openapi3".

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	internalslices "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/slices"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/structs"
	"github.com/fatih/structtag"
	"github.com/google/uuid"
	"github.com/swaggest/jsonschema-go"
	"github.com/swaggest/openapi-go/openapi3"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GenerateSpecSchemas creates OpenAPI schemas from code.
func (o *CodeGen) GenerateSpecSchemas(structNames []string) {
	reflector := newSpecReflector()

	for idx, structName := range structNames {
		// we need to compile gen-schema right after PublicStructs file is updated
		// cannot import packages at runtime
		// if we have an uncompilable state then we need to update src to compile. no way around it
		// UDPATE: use https://github.com/pkujhd/goloader instead of plugin pkg which cant reload changed go file at runtime
		// or use yaegi
		st, ok := PublicStructs[structName]
		if !ok {
			log.Fatalf("struct-name %s does not exist in PublicStructs", structName)
		}
		if !structs.HasJSONTag(st) {
			log.Fatalf("struct %s: ensure there is at least a JSON tag set", structName)
		}

		oc, err := reflector.NewOperationContext(http.MethodGet, "/dummy-op-"+strconv.Itoa(idx))
		handleError(err)
		oc.AddRespStructure(st)

		handleError(reflector.AddOperation(oc))

		// IMPORTANT: ensure structs are public
		reflector.Spec.Components.Schemas.MapOfSchemaOrRefValues[structName].Schema.MapOfAnything = map[string]any{"x-postgen-struct": structName}
	}
	s, err := reflector.Spec.MarshalYAML()
	handleError(err)

	fmt.Println(string(s))
}

// newSpecReflector returns a new SpecReflector.
func newSpecReflector() *openapi3.Reflector {
	reflector := openapi3.Reflector{Spec: &openapi3.Spec{}}

	// see https://github.com/swaggest/openapi-go/discussions/62#discussioncomment-5710581
	reflector.DefaultOptions = append(reflector.DefaultOptions,
		jsonschema.InterceptDefName(func(t reflect.Type, defaultDefName string) string {
			return defaultDefName
		}),
		jsonschema.InterceptProp(func(params jsonschema.InterceptPropParams) error {
			if params.Field.Tag.Get("openapi-go") == "ignore" {
				return jsonschema.ErrSkipProperty
			}

			if params.PropertySchema != nil {
				if params.PropertySchema.ExtraProperties == nil {
					params.PropertySchema.ExtraProperties = map[string]any{}
				}

				tags, err := structtag.Parse(string(params.Field.Tag))
				if err != nil {
					panic(fmt.Sprintf("structtag.Parse: %v", err))
				}

				for _, t := range tags.Tags() {
					if strings.HasPrefix(t.Key, "x-") {
						params.PropertySchema.ExtraProperties[t.Key] = t.Value()
					}
				}
			}

			if params.Field.Tag.Get("x-omitempty") == "true" {
				if params.PropertySchema == nil {
					return nil
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
			t := params.Schema.ReflectType
			if strings.HasSuffix(params.Schema.ReflectType.PkgPath(), "internal/models") {
				if t.Kind() == reflect.Ptr {
					t = t.Elem()
				}
				if t.Kind() == reflect.Struct {
					// will generate duplicate models otherwise, exiting only simple makes exit early it and output an empty schema
					params.Schema.ExtraProperties = map[string]any{
						"x-TO-BE-DELETED": true,
					}

					return true, nil
				}
			}

			var isCustomUUID bool
			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}
			if t.Kind() == reflect.Struct {
				if t.Field(0).Type == reflect.TypeOf(uuid.New()) {
					isCustomUUID = true
				}
			}

			// TODO: also if type script and has a field UUID
			if t == reflect.TypeOf(uuid.New()) || isCustomUUID {
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

	return &reflector
}
