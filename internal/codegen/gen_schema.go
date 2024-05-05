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
	"golang.org/x/exp/slices"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GenerateSpecSchemas creates OpenAPI schemas from code.
func (o *CodeGen) GenerateSpecSchemas(structNames, existingStructs, dbIDs []string) {
	reflector := newSpecReflector(dbIDs)

	sns := internalslices.Unique(structNames)
	for idx, structName := range sns {
		// fmt.Fprintf(os.Stderr, "Generating struct %s\n", structName)
		// We need to compile gen-schema right after PublicStructs file is updated
		// cannot import packages at runtime
		// if we have an uncompilable state then we need to update src to compile. no way around it
		// to work around this would need something like yaegi, but might not support swaggest libs
		st, ok := PublicStructs[structName]
		if !ok {
			if !slices.Contains(existingStructs, structName) {
				fmt.Fprintf(os.Stderr, "generated struct %q no longer exists and will be deleted", structName)
				continue
			}

			log.Fatalf("struct %s does not exist in rest or models packages but is referenced in x-gen-struct", structName)
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

func newSpecReflector(dbIDs []string) *openapi3.Reflector {
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
			isRestStruct := strings.HasPrefix(defaultDefName, "Rest")
			schemaName := strings.TrimPrefix(defaultDefName, "Rest")
			pkg := t.PkgPath()[strings.LastIndex(t.PkgPath(), "/")+1:]
			prefix := strcase.ToCamel(pkg)
			if reflectType, ok := reflectTypeNames[pkg]; ok {
				if structName, ok := reflectType[t.Name()]; ok {
					schemaName = strings.TrimPrefix(prefix, "Rest") + structName
				}
			}

			if isRestStruct && strings.HasPrefix(schemaName, "Db") {
				log.Fatalf("Db prefix is restricted. Please rename %q in package rest\n", schemaName)
			}

			return schemaName
		}),
		jsonschema.InterceptProperty(func(name string, field reflect.StructField, propertySchema *jsonschema.Schema) error {
			if slices.Contains(dbIDs, field.Name) {
				propertySchema.ExtraProperties = map[string]any{
					"x-go-type": field.Name,
				}
			}

			// intercept arrays of ids
			if ii := propertySchema.Items; ii != nil && ii.SchemaOrBool != nil && ii.SchemaOrBool.TypeObject != nil {
				obj := ii.SchemaOrBool.TypeObject
				goName := obj.ReflectType.Name()
				if slices.Contains(dbIDs, goName) {
					obj.ExtraProperties = map[string]any{
						"x-go-type": goName,
					}
				}
			}

			return nil
		}),
		jsonschema.InterceptProp(func(params jsonschema.InterceptPropParams) error {
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

			// PkgPath is empty for top level struct instance via struct{} (see docs)
			// therefore we cannot append vendor extensions via gen_schema.
			// for top level schemas we must check for Db[A-Z]* and append
			// NOTE: now with models merged this is not needed
			// isDbType := strings.HasSuffix(t.PkgPath(), "/models")
			// if n := t.Name(); isDbType {
			// 	params.Schema.ExtraProperties = map[string]any{
			// 		"x-go-type": strings.TrimPrefix(n, "Db"),
			// 		// in case this is needed later, import-mappings config affects this
			// 		// "x-go-name": strings.TrimPrefix(n, "Db"),
			// 		// "x-go-type-import": map[string]any{
			// 		// 	"name": "db",
			// 		// 	"path": "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models",
			// 		// },
			// 		"x-is-generated": true,
			// 	}
			// }

			var isCustomUUID bool
			if t.Kind() == reflect.Struct && t.Field(0).Type == reflect.TypeOf(uuid.New()) {
				isCustomUUID = true
			}

			if t == reflect.TypeOf(uuid.New()) || isCustomUUID {
				params.Schema.Type = &jsonschema.Type{SimpleTypes: pointers.New(jsonschema.String)}
				params.Schema.Pattern = pointers.New("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
				params.Schema.Items = &jsonschema.Items{}
				params.Schema.Examples = []interface{}{"cdb15f83-1c5d-4727-98d1-8924ccd1fc01"}
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
