package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// .nolint: gochecknoglobals.
var (
	explode   = openapi3.BoolPtr(true)
	noExplode = openapi3.BoolPtr(false)
	arrayOf   = func(items *openapi3.SchemaRef) *openapi3.SchemaRef {
		return &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"array"}, Items: items}, Ref: "#/components/schemas/arrayOf"}
	}
	objectOf = func(args ...interface{}) *openapi3.SchemaRef {
		s := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"object"}, Properties: make(map[string]*openapi3.SchemaRef)}}
		if len(args)%2 != 0 {
			panic("invalid arguments. must be an even number of arguments")
		}
		for i := range len(args) / 2 {
			propName, _ := args[i*2].(string)
			propSchema, _ := args[i*2+1].(*openapi3.SchemaRef)
			s.Value.Properties[propName] = propSchema
		}

		return s
	}

	additionalPropertiesObjectOf = func(schema *openapi3.SchemaRef) *openapi3.SchemaRef {
		return &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type:                 &openapi3.Types{"object"},
				AdditionalProperties: openapi3.AdditionalProperties{Schema: schema},
			},
			Ref: "#/components/schemas/additionalPropertiesObjectOf",
		}
	}

	integerSchema                          = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}}
	numberSchema                           = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"number"}}}
	booleanSchema                          = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"boolean"}}}
	stringSchema                           = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	additionalPropertiesObjectStringSchema = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"object"}, AdditionalProperties: openapi3.AdditionalProperties{Schema: stringSchema}}}
	additionalPropertiesObjectBoolSchema   = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"object"}, AdditionalProperties: openapi3.AdditionalProperties{Schema: booleanSchema}}}
	allofSchema                            = &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			AllOf: []*openapi3.SchemaRef{
				integerSchema,
				numberSchema,
			},
		},
	}
	anyofSchema = &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			AnyOf: []*openapi3.SchemaRef{
				integerSchema,
				stringSchema,
			},
		},
	}
	oneofSchema = &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			OneOf: []*openapi3.SchemaRef{
				booleanSchema,
				integerSchema,
			},
		},
	}
	oneofSchemaObject = &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			OneOf: []*openapi3.SchemaRef{
				objectOneRSchema,
				objectTwoRSchema,
			},
		},
	}
	anyofSchemaObject = &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			AnyOf: []*openapi3.SchemaRef{
				objectOneRSchema,
				objectTwoRSchema,
			},
		},
	}
	stringArraySchema  = arrayOf(stringSchema)
	integerArraySchema = arrayOf(integerSchema)
	objectSchema       = objectOf("id", stringSchema, "name", stringSchema)
	objectTwoRSchema   = func() *openapi3.SchemaRef {
		s := objectOf("id2", stringSchema, "name2", stringSchema)
		s.Ref = "#/components/schemas/objectTwoRSchema"
		s.Value.Required = []string{"id2"}

		return s
	}()

	objectOneRSchema = func() *openapi3.SchemaRef {
		s := objectOf("id", stringSchema, "name", stringSchema)
		s.Ref = "#/components/schemas/objectOneRSchema"
		s.Value.Required = []string{"id"}

		return s
	}()

	oneofSchemaArrayObject = &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			OneOf: []*openapi3.SchemaRef{
				stringArraySchema,
				objectTwoRSchema,
			},
		},
	}
)

func TestReconstructQueryParamsValues(t *testing.T) {
	type args struct {
		schema     *openapi3.Schema
		data       interface{}
		schemaName string
	}
	tests := []struct {
		name        string
		args        args
		want        interface{}
		errContains string
	}{
		{
			name: "empty correct",
			args: args{
				schema:     oneofSchemaObject.Value,
				data:       map[string]interface{}{},
				schemaName: "objectOneRSchema",
			},
			want: map[string]interface{}{},
		},
		{
			name: "correct data",
			args: args{
				schema: oneofSchemaObject.Value,
				data: map[string]interface{}{
					"id":   "a",
					"name": "John",
				},
				schemaName: "objectOneRSchema",
			},
			want: map[string]interface{}{
				"id":   "a",
				"name": "John",
			},
		},
		{
			name: "invalid schema",
			args: args{
				schema: oneofSchemaArrayObject.Value,
				data: map[string]interface{}{
					"id": "a",
				},
				schemaName: "arrayOf",
			},
			errContains: "strconv.Atoi: parsing \"id\": invalid syntax",
		},
		{
			name: "correct array data",
			args: args{
				schema: oneofSchemaArrayObject.Value,
				data: []interface{}{
					"1", "1",
				},
				schemaName: "arrayOf",
			},
			want: []interface{}{
				"1", "1",
			},
		},
		{
			name: "correct array data in query param format",
			args: args{
				schema: oneofSchemaArrayObject.Value,
				data: map[string]interface{}{
					"0": "0",
					"2": "1",
				},
				schemaName: "arrayOf",
			},
			want: []interface{}{
				"0", nil, "1",
			},
		},
		{
			name: "correct additionalProperties data",
			args: args{
				schema: additionalPropertiesObjectOf(oneofSchemaArrayObject).Value,
				data: map[string]interface{}{
					"entry-1": []interface{}{
						"1", "2",
					},
				},
				schemaName: "arrayOf",
			},
			want: map[string]interface{}{
				"entry-1": []interface{}{
					"1", "2",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ReconstructQueryParamsValues(tc.args.schema, tc.args.data, tc.args.schemaName)
			if err != nil && tc.errContains == "" {
				t.Errorf("unexpected error: %v", err)

				return
			}
			if tc.errContains != "" {
				if err == nil {
					t.Errorf("expected error but got nothing")

					return
				}
				require.ErrorContains(t, err, tc.errContains)

				return
			}

			assert.Equal(t, tc.want, got)
		})
	}
}
