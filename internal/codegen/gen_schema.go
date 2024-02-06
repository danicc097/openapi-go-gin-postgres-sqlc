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
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/slices"
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

	sns := slices.Unique(structNames)
	for idx, structName := range sns {
		// fmt.Fprintf(os.Stderr, "Generating struct %s\n", structName)
		// We need to compile gen-schema right after PublicStructs file is updated
		// cannot import packages at runtime
		// if we have an uncompilable state then we need to update src to compile. no way around it
		// to work around this would need something like yaegi, but might not support swaggest libs
		st, ok := PublicStructs[structName]
		if !ok {
			// FIXME: before generating spec via gen-schema we will find all x-gen-struct keys in
			// spec and search for them in publicstructs. if they dont exist and x-is-generated,
			// we will add x-TO-BE-DELETED: true. if they simply dont exist warn the user that the reference
			// is broken.
			// in any case, we will simply skip schema instead of failing
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
		schemaName := strings.TrimPrefix(structName, "Rest")
		x, ok := reflector.Spec.Components.Schemas.MapOfSchemaOrRefValues[schemaName]
		if !ok {
			s, err := reflector.Spec.MarshalYAML()
			handleError(err)
			fmt.Fprint(os.Stderr, string(s))
			log.Fatalf("Could not generate %s", schemaName)
		}
		x.Schema.MapOfAnything = map[string]any{"x-gen-struct": schemaName, "x-is-generated": true}
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
			isRestStruct := strings.HasPrefix(defaultDefName, "Rest")
			schemaName := strings.TrimPrefix(defaultDefName, "Rest")
			pkg := t.PkgPath()[strings.LastIndex(t.PkgPath(), "/")+1:]
			prefix := strcase.ToCamel(pkg)
			if reflectType, ok := reflectTypeNames[pkg]; ok {
				if structName, ok := reflectType[t.Name()]; ok {
					// there should be a schema already that maps to rest model without prefix, e.g.
					// User:
					//   x-gen-struct: RestUser
					// (will raise error loading openapi specification if not)
					// TODO: generate these for all rest structs:
					//  Rest<Struct>: # Generated by ...
					// 	   x-gen-struct: <Struct>

					schemaName = strings.TrimPrefix(prefix, "Rest") + structName
				}
			}

			if isRestStruct && strings.HasPrefix(schemaName, "Db") {
				log.Fatalf("Db prefix is restricted. Please rename %q in package rest\n", schemaName)
			}

			return schemaName
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
			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}

			// duplicate from gen itself
			if strings.HasSuffix(params.Schema.ReflectType.PkgPath(), "internal/models") {
				if t.Kind() == reflect.Struct {
					// will generate duplicate models otherwise
					params.Schema.ExtraProperties = map[string]any{
						"x-TO-BE-DELETED": true,
					}

					return true, nil
				}
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
					"x-is-generated": true,
				}
			}

			return false, nil
		}),
	)

	return &reflector
}
