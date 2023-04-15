package postgresql

import (
	"reflect"
)

// updateEntityWithParams updates repo entity with params.
// Since params are already generated from the entity, many assumptions are made.
// This is not to be used for other kinds of structs.
// Example:
//
//	updateEntityWithParams(&User{}, &Params{Name: "Jane"})
func updateEntityWithParams[T any, U any](entity *T, params *U) {
	entityValue := reflect.ValueOf(entity).Elem()
	paramsType := reflect.TypeOf(params).Elem()
	paramsValue := reflect.ValueOf(params).Elem()

	for i := 0; i < paramsType.NumField(); i++ {
		paramName := paramsType.Field(i).Name
		paramValue := paramsValue.Field(i)

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
			fieldValue.Set(paramValue.Elem())
		}
	}
}
