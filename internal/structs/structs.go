package structs

import (
	"reflect"
	"strings"
)

// GetStructKeys returns a slice of json keys extracted from an initialized struct's tags.
func GetStructKeys(s any, parent string) []string {
	keys := []string{}

	if s == nil {
		return keys
	}

	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return keys
		}
		val = val.Elem()
	}

	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		for j := 0; j < val.Len(); j++ {
			elem := val.Index(j).Interface()
			subkeys := GetStructKeys(elem, "")
			for _, subkey := range subkeys {
				keys = append(keys, parent+"."+subkey)
			}
		}
	}

	if val.Kind() == reflect.Struct {
		for idx := 0; idx < val.NumField(); idx++ {
			typeField := val.Type().Field(idx)
			jsonTag := typeField.Tag.Get("json")
			if jsonTag == "" {
				continue
			}
			key := strings.Split(jsonTag, ",")[0]

			switch typeField.Type.Kind() {
			case reflect.Array, reflect.Slice:
				if val.Field(idx).Len() > 0 {
					keys = append(keys, key)
				}
				for j := 0; j < val.Field(idx).Len(); j++ {
					elem := val.Field(idx).Index(j).Interface()
					subkeys := GetStructKeys(elem, key)
					for _, subkey := range subkeys {
						keys = append(keys, key+"."+subkey)
					}
				}
			case reflect.Struct:
				keys = append(keys, key)
				subkeys := GetStructKeys(val.Field(idx).Interface(), key)
				for _, subkey := range subkeys {
					keys = append(keys, key+"."+subkey)
				}
			case reflect.Pointer:
				if val.Field(idx).IsNil() {
					continue
				}
				keys = append(keys, key)
				subkeys := GetStructKeys(val.Field(idx).Interface(), key)
				for _, subkey := range subkeys {
					keys = append(keys, key+"."+subkey)
				}
			default:
				keys = append(keys, key)
			}
		}
	}

	return keys
}
