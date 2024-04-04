package openapi

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

// ReconstructQueryParamsValues converts query params that may include arrays represented as maps
// to the correct go types.
// This utility is necessary for generated methods when the schemas are being used as query params.
// It also supports anyOf, oneOf and allOf keywords.
func ReconstructQueryParamsValues(schema *openapi3.Schema, data interface{}, schemaName string) (interface{}, error) {
	mdata, ok := data.(map[string]interface{})
	if !ok {
		// its not map, just return the primitive (array and maps both represented as maps in query params)
		return data, nil
	}

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

		return ReconstructQueryParamsValues(matchingSchema, mdata, schemaName)
	}

	props := schema.Properties
	additPropsSchema := schema.AdditionalProperties.Schema
	if props == nil {
		if additPropsSchema == nil {
			return nil, fmt.Errorf("invalid schema")
		}
		for k, v := range mdata {
			obj, err := ReconstructQueryParamsValues(additPropsSchema.Value, v, schemaName)
			if err != nil {
				return nil, err
			}
			result[k] = obj
		}
	}

	for propName, prop := range props {
		propSchema := prop.Value
		if propSchema == nil {
			return nil, fmt.Errorf("invalid schema for property %s", propName)
		}

		propData, ok := mdata[propName]
		if !ok {
			fmt.Printf("propname %q not found in data: %v\n", propName, mdata)
			continue
		}

		switch {
		case propSchema.Type.Permits("object"):
			objData, ok := propData.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("invalid object data type for property %s: %T", propName, propData)
			}
			obj, err := ReconstructQueryParamsValues(propSchema, objData, schemaName)
			if err != nil {
				return nil, err
			}
			result[propName] = obj
		case propSchema.Type.Permits("array"):
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
