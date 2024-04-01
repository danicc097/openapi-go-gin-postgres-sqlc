package postgresql

import (
	"fmt"
	"strconv"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

const (
	like         = " LIKE $i"
	ilike        = " ILIKE $i"
	equal        = " = $i"
	greater      = " > $i"
	greaterEqual = " >= $i"
	less         = " < $i"
	lessEqual    = " <= $i"
	isNull       = " is null"
	isNotNull    = " is not null"

	// https://www.postgresql.org/docs/16/functions-array.html#ARRAY-OPERATORS-TABLE
	// https://www.postgresql.org/docs/16/functions-subquery.html
	// TODO: should allow pagination with array filters on db fields that are not arrays,
	// e.g. find documents with the given refs, which the user enters in the combobox separated by ;
	// and we generate all combobox options (separator customizable)
	// this can't be done with default text filtering.
	// not to be confused with filtering on array columns which isn't implemented at the moment
)

// GenerateDefaultFilters generates SQL where clauses for a given set of pagination params.
func GenerateDefaultFilters(entity db.TableEntity, paginationParams models.PaginationItems) (map[string][]interface{}, error) {
	filters := make(map[string][]interface{})

	if _, ok := db.EntityFilters[entity]; !ok {
		return nil, fmt.Errorf("invalid entity: %v", entity)
	}

	for id, pag := range paginationParams {
		dbfilter, ok := db.EntityFilters[entity][id]
		if !ok {
			continue
		}
		if pag.Filter == nil {
			continue
		}

		disc, err := pag.Filter.Discriminator()
		if err != nil {
			return nil, fmt.Errorf("discriminator: %w", err)
		}
		filterMode := models.PaginationFilterModes(disc)

		switch filterMode {
		case models.PaginationFilterModesEmpty:
			filters[dbfilter.Db+isNull] = []interface{}{}
			continue
		case models.PaginationFilterModesNotEmpty:
			filters[dbfilter.Db+isNotNull] = []interface{}{}
			continue
		}

		v, _ := pag.Filter.ValueByDiscriminator()
		switch t := v.(type) {
		case models.PaginationFilterArray:
			vv := t.Value
			switch dbfilter.Type {
			case "integer", "string": //...
			// should support filtering multiple exact values at once
			case "date-time": // min,max
				switch filterMode {
				case models.PaginationFilterModesBetween, models.PaginationFilterModesBetweenInclusive:
					var min, max interface{}
					min, err = time.Parse(time.RFC3339, vv[0])
					if err != nil {
						min = "null"
					}

					max, err = time.Parse(time.RFC3339, vv[1])
					if err != nil {
						max = "null"
					}
					if filterMode == models.PaginationFilterModesBetween {
						filters[fmt.Sprintf("%[1]s > $i AND %[1]s < $i", dbfilter.Db)] = []interface{}{min, max}
					} else {
						filters[fmt.Sprintf("%[1]s >= $i AND %[1]s <= $i", dbfilter.Db)] = []interface{}{min, max}
					}
				}
			}
			// we can have arrincludessome, arrincludesall, arrincludes.
			// will not share modes with single values.
		case models.PaginationFilterPrimitive:
			if t.Value == nil {
				continue
			}
			v := *t.Value

			switch dbfilter.Type {
			case "string":

				op := ilike
				if t.CaseSensitive != nil && *t.CaseSensitive {
					op = like
				}
				switch filterMode {
				case models.PaginationFilterModesEquals:
					filters[dbfilter.Db+op] = []interface{}{v}
				case models.PaginationFilterModesContains:
					filters[dbfilter.Db+op] = []interface{}{"%" + v + "%"}
				case models.PaginationFilterModesStartsWith:
					filters[dbfilter.Db+op] = []interface{}{v + "%"}
				case models.PaginationFilterModesEndsWith:
					filters[dbfilter.Db+op] = []interface{}{"%" + v}
				}
			case "float", "integer":
				var num interface{}
				switch dbfilter.Type {
				case "integer":
					num, err = strconv.Atoi(v)
				case "float":
					num, err = strconv.ParseFloat(v, 64)
				}
				if err != nil {
					return nil, fmt.Errorf("%s: invalid %s %q", dbfilter.Db, dbfilter.Type, v)
				}

				switch filterMode {
				case models.PaginationFilterModesEquals:
					filters[dbfilter.Db+equal] = []interface{}{num}
				case models.PaginationFilterModesGreaterThan:
					filters[dbfilter.Db+greater] = []interface{}{num}
				case models.PaginationFilterModesGreaterThanOrEqualTo:
					filters[dbfilter.Db+greaterEqual] = []interface{}{num}
				case models.PaginationFilterModesLessThan:
					filters[dbfilter.Db+less] = []interface{}{num}
				case models.PaginationFilterModesLessThanOrEqualTo:
					filters[dbfilter.Db+lessEqual] = []interface{}{num}
				}
			case "boolean":
				if v != "true" && v != "false" {
					return nil, fmt.Errorf("%s: invalid boolean %q", dbfilter.Db, v)
				}
				filters[dbfilter.Db+equal] = []interface{}{v == "true"}
			}
		default:
			return nil, fmt.Errorf("unsupported filter mode type: %v", t)
		}
	}

	return filters, nil
}
