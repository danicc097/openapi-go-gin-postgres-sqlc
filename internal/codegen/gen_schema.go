/*
gen-schema generates OpenAPI v3 schema portions from code.
*/
package codegen

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	// kinopenapi3 "github.com/getkin/kin-openapi/openapi3".

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	internalslices "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/slices"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/structs"
	"github.com/fatih/structtag"
	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
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
		// fmt.Fprintf(os.Stderr, "Generating struct %s\n", structName)
		// We need to compile gen-schema right after PublicStructs file is updated
		// cannot import packages at runtime
		// if we have an uncompilable state then we need to update src to compile. no way around it
		// to work around this would need something like yaegi, but might not support swaggest libs
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

		// IMPORTANT: ensure structs are public.
		x, ok := reflector.Spec.Components.Schemas.MapOfSchemaOrRefValues[structName]
		if !ok {
			s, err := reflector.Spec.MarshalYAML()
			handleError(err)
			fmt.Fprint(os.Stderr, string(s))
			log.Fatalf("Could not generate %s", structName)
		}
		x.Schema.MapOfAnything = map[string]any{"x-postgen-struct": structName}
	}
	s, err := reflector.Spec.MarshalYAML()
	handleError(err)

	fmt.Println(string(s))
}

// newSpecReflector returns a new SpecReflector.
func newSpecReflector() *openapi3.Reflector {
	reflector := openapi3.Reflector{Spec: &openapi3.Spec{}}

	reflectTypeNames := map[string]map[string]string{}
	jsonBlob, err := os.ReadFile("internal/codegen/reflectTypeMap.gen.json")
	if err != nil {
		log.Fatalf("Error reading reflect types: %v\n", err)
	}
	err = json.Unmarshal(jsonBlob, &reflectTypeNames)
	if err != nil {
		log.Fatalf("Error unmarshaling reflect types: %v\n", err)
	}
	// see https://github.com/swaggest/openapi-go/discussions/62#discussioncomment-5710581
	reflector.DefaultOptions = append(reflector.DefaultOptions,
		jsonschema.InterceptDefName(func(t reflect.Type, defaultDefName string) string {
			// see https://stackoverflow.com/questions/74838506/get-the-type-name-of-a-generic-struct-without-type-parameters
			// both t.Name() and t.String() return a composed name
			// RestGetPaginatedNotificationsResponse has
			// defaultDefName: RestPaginationBaseResponse[GithubComDanicc097OpenapiGoGinPostgresSqlcInternalRestNotification]
			// TODO: if we use ast-parser create-generics-map we can generate a JSON mapping real names to composed names
			// right before gen-schema
			// PaginationBaseResponse[github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest.Notification] may be mapped to
			// github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest.PaginationBaseResponse[github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest.Notification]
			// NOTE: returning structName in intercept will not work since we intercept all indirect types here as well
			// and would need a way to identify generic structs regardless
			// fmt.Printf("t.Name(): %v\n", t.Name())

			// c := render.AsCode(g)
			// fmt.Printf("c: %v\n", c)
			pkg := t.PkgPath()[strings.LastIndex(t.PkgPath(), "/")+1:]
			if reflectType, ok := reflectTypeNames[pkg]; ok {
				if structName, ok := reflectType[t.Name()]; ok {
					prefix := strcase.ToCamel(pkg)

					return prefix + structName
				}
			}

			return defaultDefName
		}),
		jsonschema.InterceptProp(func(params jsonschema.InterceptPropParams) error {
			if params.Field.Tag.Get("openapi-go") == "ignore" { // does not ignore completely, it still sets as required. Probably a bug
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
					// will generate duplicate models otherwise
					params.Schema.ExtraProperties = map[string]any{
						"x-TO-BE-DELETED": true,
					}

					return true, nil
				}
			}

			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}

			if t.Kind() == reflect.Struct {
				for i := 0; i < t.NumField(); i++ {
					field := t.Field(i)
					openapiGoTag := field.Tag.Get("openapi-go")
					if openapiGoTag == "ignore" {
						tag := reflect.StructTag(`openapi-go:"ignore"`)
						field.Tag = tag
					}
				}
			}

			var isCustomUUID bool
			if t.Kind() == reflect.Struct && t.Field(0).Type == reflect.TypeOf(uuid.New()) {
				isCustomUUID = true
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
