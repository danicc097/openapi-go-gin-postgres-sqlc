package postgresql

import (
	"log"
	"reflect"
)

// updateEntityWithParams updates repo entity with params.
// Since params are already generated from the entity, many assumptions are made.
// This is not to be used for other kinds of structs.
// For performance critical stuff refrain from this.
// Example:
//
//	updateEntityWithParams(&User{}, &Params{Name: "Jane"})
func updateEntityWithParams(entity any, params any) {
	entityValue := reflect.ValueOf(entity).Elem()
	paramsType := reflect.TypeOf(params).Elem()
	paramsValue := reflect.ValueOf(params).Elem()

	for i := 0; i < paramsType.NumField(); i++ {
		paramName := paramsType.Field(i).Name
		paramValue := paramsValue.Field(i)

		// "<Entity>UpdateParams" struct are all pointers to values or other pointers
		if paramValue.Kind() == reflect.Ptr && paramValue.IsNil() {
			continue
		}

		// corresponding field. Generated with same names.
		fieldName, ok := paramsType.FieldByName(paramName)
		if !ok {
			continue
		}

		fieldValue := entityValue.FieldByName(fieldName.Name)

		// params first interface will always be a pointer.
		if fieldValue.CanSet() {
			if paramValue.Kind() == reflect.Struct {
				log.Default().Printf("WARNING: struct is not supported: %s", paramName)
				continue
			}
			fieldValue.Set(paramValue.Elem())
		}
	}
}
