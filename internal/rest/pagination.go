package rest

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

func formatCursorValue(value interface{}) (string, error) {
	switch v := value.(type) {
	case time.Time:
		return v.Format(time.RFC3339Nano), nil
	case string:
		return v, nil
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v), nil
	default:
		return "", fmt.Errorf("unhandled cursor type: %v", v)
	}
}

// getNextCursor returns the next cursor from an entity struct by a given json tag.
// This allows for dynamic pagination parameters.
func getNextCursor(entity interface{}, jsonFieldName string, tableEntity models.TableEntity) (string, error) {
	if entity == nil {
		return "", errors.New("no entity given")
	}

	if _, ok := models.EntityFields[tableEntity]; !ok {
		return "", errors.New("no entity found")
	}

	entityType := reflect.TypeOf(entity)
	entityValue := reflect.ValueOf(entity)

	for i := range entityType.NumField() {
		structField := entityType.Field(i)
		jsonTag := structField.Tag.Get("json")
		if jsonTag == jsonFieldName {
			fieldValue := entityValue.Field(i).Interface()
			return formatCursorValue(fieldValue)
		}
	}

	return "", fmt.Errorf("no json tag with value: %v", jsonFieldName)
}
