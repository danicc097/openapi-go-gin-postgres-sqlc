package postgresql

import (
	"fmt"
	"strconv"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

const (
	ilike = " ILIKE $i"
	equal = " = $i"

	// https://www.postgresql.org/docs/16/functions-array.html#ARRAY-OPERATORS-TABLE
	// https://www.postgresql.org/docs/16/functions-subquery.html
	// TODO: should allow pagination with array filters on db fields that are not arrays,
	// e.g. find documents with the given refs, which the user enters in the combobox separated by ;
	// and we generate all combobox options (separator customizable)
	// this can't be done with default text filtering.
	// not to be confused with filtering on array columns which isn't implemented at the moment
)

func GenerateFilters(entity db.TableEntity, queryParams map[string]models.PaginationFilter) (map[string][]interface{}, error) {
	filters := make(map[string][]interface{})

	for id, filter := range queryParams {
		if _, ok := db.EntityFilters[entity]; !ok {
			return nil, fmt.Errorf("invalid entity: %v", entity)
		}
		dbfilter, ok := db.EntityFilters[entity][id]
		if !ok {
			continue
		}

		disc, err := filter.Value.Discriminator()
		if err != nil {
			return nil, fmt.Errorf("discriminator: %w", err)
		}
		filterMode := models.PaginationFilterModes(disc)
		v, _ := filter.Value.ValueByDiscriminator()
		switch t := v.(type) {
		case models.PaginationFilterArrayValue:
			// we can have arrincludessome, arrincludesall, arrincludes.
			// will not share modes with single values.
		case models.PaginationFilterSingleValue:
			if t.Value == nil {
				continue
			}
			v := *t.Value
			switch dbfilter.Type {
			case "string":
				switch filterMode {
				case models.PaginationFilterModesEquals:
					filters[dbfilter.Db+equal] = []interface{}{v}
				case models.PaginationFilterModesContains:
					filters[dbfilter.Db+ilike] = []interface{}{"%" + v + "%"}
				case models.PaginationFilterModesStartsWith:
					filters[dbfilter.Db+ilike] = []interface{}{v + "%"}
				case models.PaginationFilterModesEndsWith:
					filters[dbfilter.Db+ilike] = []interface{}{"%" + v}
				}
			case "integer":
				if intValue, err := strconv.Atoi(v); err == nil {
					filters[dbfilter.Db+equal] = []interface{}{intValue}
				}
			case "float":
				if floatValue, err := strconv.ParseFloat(v, 64); err == nil {
					filters[dbfilter.Db+equal] = []interface{}{floatValue}
				}
			case "date-time":
				if dateTimeValue, err := time.Parse(time.RFC3339, v); err == nil {
					filters[dbfilter.Db+equal] = []interface{}{dateTimeValue}
				}
			case "boolean":
				// we will receive actual types (boolean, time.Time) via runtime package
				if v == "true" || v == "false" {
					filters[dbfilter.Db+equal] = []interface{}{v == "true"}
				}
			}
		default:
			return nil, fmt.Errorf("unsupported filter mode type: %v", t)
		}
	}

	return filters, nil
}
