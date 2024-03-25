package postgresql

import (
	"fmt"
	"strconv"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

func GenerateFilters(entity db.TableEntity, queryParams map[string][]string) (map[string][]interface{}, error) {
	filters := make(map[string][]interface{})

	for key, values := range queryParams {
		if _, ok := db.EntityFilters[entity]; !ok {
			return nil, fmt.Errorf("invalid entity: %v", entity)
		}
		filter, ok := db.EntityFilters[entity][key]
		if !ok {
			continue
		}

		for _, value := range values {
			switch filter.Type {
			case "string":
				switch models.PaginationFilterModes(value) {
				case models.PaginationFilterModesEquals:
					filters[fmt.Sprintf("%s = ?", filter.Db)] = []interface{}{value}
				case models.PaginationFilterModesContains:
					filters[fmt.Sprintf("%s ILIKE ?", filter.Db)] = []interface{}{"%" + value + "%"}
				case models.PaginationFilterModesStartsWith:
					filters[fmt.Sprintf("%s ILIKE ?", filter.Db)] = []interface{}{value + "%"}
				case models.PaginationFilterModesEndsWith:
					filters[fmt.Sprintf("%s ILIKE ?", filter.Db)] = []interface{}{"%" + value}
				}
			case "integer":
				if intValue, err := strconv.Atoi(value); err == nil {
					filters[fmt.Sprintf("%s = ?", filter.Db)] = []interface{}{intValue}
				}
			case "float":
				if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
					filters[fmt.Sprintf("%s = ?", filter.Db)] = []interface{}{floatValue}
				}
			case "date-time":
				if dateTimeValue, err := time.Parse(time.RFC3339, value); err == nil {
					filters[fmt.Sprintf("%s = ?", filter.Db)] = []interface{}{dateTimeValue}
				}
			case "boolean":
				// we will receive actual types (boolean, time.Time) via runtime package
				if value == "true" || value == "false" {
					boolValue := value == "true"
					filters[fmt.Sprintf("%s = ?", filter.Db)] = []interface{}{boolValue}
				}
			}
		}
	}

	return filters, nil
}
