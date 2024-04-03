package rest

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// example: map[0:map[key:true] 1:map[key:false]] -> [map[key:true] map[key:false]]
func sliceMapToSlice(m map[string]interface{}) ([]interface{}, error) {
	var result []interface{}

	keys := make([]int, 0, len(m))
	for k := range m {
		key, err := strconv.Atoi(k)
		if err != nil {
			return nil, fmt.Errorf("array indexes must be integers: %w", err)
		}
		keys = append(keys, key)
	}
	for i := 0; i <= slices.Max(keys); i++ {
		val, ok := m[strconv.Itoa(i)]
		if !ok {
			result = append(result, nil)
			continue
		}
		result = append(result, val)
	}
	return result, nil
}

func reconstructMapFromSchema(schema *openapi3.Schema, data map[string]interface{}, schemaName string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	if schema == nil {
		return result, nil
	}

	if len(schema.AnyOf) > 0 || len(schema.OneOf) > 0 || len(schema.AllOf) > 0 {
		var schemas []*openapi3.SchemaRef
		if len(schema.AnyOf) > 0 {
			schemas = schema.AnyOf
		} else if len(schema.OneOf) > 0 {
			schemas = schema.OneOf
		} else {
			schemas = schema.AllOf
		}

		// will be the top level one
		var matchingSchema *openapi3.Schema
		for _, s := range schemas {
			if s == nil {
				continue
			}
			if strings.HasSuffix(s.Ref, "/"+schemaName) {
				matchingSchema = s.Value
				break
			}
		}

		if matchingSchema == nil {
			return nil, fmt.Errorf("property schema %s not found in anyOf, oneOf, or allOf", schemaName)
		}

		return reconstructMapFromSchema(matchingSchema, data, schemaName)
	}

	props := schema.Properties
	if props == nil {
		return nil, fmt.Errorf("invalid schema")
	}

	for propName, prop := range props {
		propSchema := prop.Value
		if propSchema == nil {
			return nil, fmt.Errorf("invalid schema for property %s", propName)
		}

		propData, ok := data[propName]
		if !ok {
			fmt.Printf("propname %q not found in data: %v\n", propName, data)
			continue
		}

		switch {
		case propSchema.Type.Is("object"):
			objData, ok := propData.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("invalid object data type for property %s: %T", propName, propData)
			}
			obj, err := reconstructMapFromSchema(propSchema, objData, schemaName)
			if err != nil {
				return nil, err
			}
			result[propName] = obj
		case propSchema.Type.Is("array"):
			arrData, ok := propData.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("invalid array data type for property %s: %T", propName, propData)
			}
			arrayResult, err := sliceMapToSlice(arrData)
			if err != nil {
				return nil, err
			}
			result[propName] = arrayResult
		default:
			result[propName] = propData
		}
	}

	return result, nil
}
