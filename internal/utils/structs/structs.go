package structs

import (
	"reflect"
	"strings"
)

// GetKeys returns a slice of tag values extracted from an initialized struct.
func GetKeys(tag string, s any, parent string) []string {
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
			subkeys := GetKeys(tag, elem, "")
			for _, subkey := range subkeys {
				keys = append(keys, parent+"."+subkey)
			}
		}
	}

	if val.Kind() == reflect.Struct {
		for idx := 0; idx < val.NumField(); idx++ {
			typeField := val.Type().Field(idx)
			tagValue := typeField.Tag.Get(tag)
			if tagValue == "" {
				continue
			}
			key := strings.Split(tagValue, ",")[0]

			switch typeField.Type.Kind() {
			case reflect.Array, reflect.Slice:
				if val.Field(idx).Len() > 0 {
					keys = append(keys, key)
				}
				for j := 0; j < val.Field(idx).Len(); j++ {
					elem := val.Field(idx).Index(j).Interface()
					subkeys := GetKeys(tag, elem, key)
					for _, subkey := range subkeys {
						keys = append(keys, key+"."+subkey)
					}
				}
			case reflect.Struct:
				keys = append(keys, key)
				subkeys := GetKeys(tag, val.Field(idx).Interface(), key)
				for _, subkey := range subkeys {
					keys = append(keys, key+"."+subkey)
				}
			case reflect.Pointer:
				if val.Field(idx).IsNil() {
					continue
				}
				keys = append(keys, key)
				subkeys := GetKeys(tag, val.Field(idx).Interface(), key)
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

// InitializeFields sets struct fields up to maxDepth.
func InitializeFields(v reflect.Value, maxDepth int) reflect.Value {
	if maxDepth == 0 {
		return v
	}
	maxDepth--
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		return InitializeFields(v.Elem(), maxDepth)
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if field.CanSet() {
				zeroValue := reflect.Zero(field.Type())
				if field.Kind() == reflect.Ptr {
					if field.IsNil() {
						field.Set(reflect.New(field.Type().Elem()))
					}
					InitializeFields(field.Elem(), maxDepth)
				} else {
					InitializeFields(field.Addr(), maxDepth)
				}
				if field.IsZero() {
					field.Set(zeroValue)
				}
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			InitializeFields(v.Index(i), maxDepth)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			InitializeFields(v.MapIndex(key), maxDepth)
		}
	}

	return v
}

// HasJSONTag ensures a struct has at least a JSON tag.
func HasJSONTag(st any) bool {
	t := reflect.TypeOf(st)

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
			if HasJSONTag(reflect.New(field.Type).Elem().Interface()) {
				return true
			}
		}
	}

	return false
}
