package postgresql_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
)

func TestGenerateFilters(t *testing.T) {
	const testEntity db.TableEntity = "testEntity"
	db.EntityFilters[testEntity] = map[string]db.Filter{
		"createdAt": {Type: "date-time", Db: "created_at", Nullable: false},
		"fullName":  {Type: "string", Db: "full_name", Nullable: true},
		"age":       {Type: "integer", Db: "age", Nullable: false},
	}

	tests := []struct {
		name       string
		entity     db.TableEntity
		queryParam map[string]models.PaginationFilter
		expected   map[string][]interface{}
	}{
		{
			name:   "StringEqualsFilter",
			entity: testEntity,
			queryParam: map[string]models.PaginationFilter{
				"fullName": {
					Value: func() models.PaginationFilterValue {
						v := models.PaginationFilterSingleValue{
							Value:      pointers.New("John Doe"),
							FilterMode: models.PaginationFilterModesEquals,
						}
						j, _ := json.Marshal(v)
						p := models.PaginationFilterValue{}
						_ = json.Unmarshal(j, &p)
						return p
					}(),
				},
			},
			expected: map[string][]interface{}{
				"full_name = $i": {"John Doe"},
			},
		},
		{
			name:   "StringContainsFilter",
			entity: testEntity,
			queryParam: map[string]models.PaginationFilter{
				"fullName": {
					Value: func() models.PaginationFilterValue {
						v := models.PaginationFilterSingleValue{
							Value:      pointers.New("John"),
							FilterMode: models.PaginationFilterModesContains,
						}
						j, _ := json.Marshal(v)
						p := models.PaginationFilterValue{}
						_ = json.Unmarshal(j, &p)
						return p
					}(),
				},
			},
			expected: map[string][]interface{}{
				"full_name ILIKE $i": {"%John%"},
			},
		},
		{
			name:   "IntegerEqualsFilter",
			entity: testEntity,
			queryParam: map[string]models.PaginationFilter{
				"age": {
					Value: func() models.PaginationFilterValue {
						v := models.PaginationFilterSingleValue{
							Value:      pointers.New("30"),
							FilterMode: models.PaginationFilterModesEquals,
						}
						j, _ := json.Marshal(v)
						p := models.PaginationFilterValue{}
						_ = json.Unmarshal(j, &p)
						return p
					}(),
				},
			},
			expected: map[string][]interface{}{
				"age = $i": {30},
			},
		},
		{
			name:   "DateTimeBetweenFilter",
			entity: testEntity,
			queryParam: map[string]models.PaginationFilter{
				"createdAt": {
					Value: func() models.PaginationFilterValue {
						v := models.PaginationFilterArrayValue{
							Value:      []string{"2023-01-01T00:00:00Z", "2023-12-31T23:59:59Z"},
							FilterMode: models.PaginationFilterModesBetween,
						}
						j, _ := json.Marshal(v)
						p := models.PaginationFilterValue{}
						_ = json.Unmarshal(j, &p)
						return p
					}(),
				},
			},
			expected: map[string][]interface{}{
				"created_at > $i AND created_at < $i": {
					time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := postgresql.GenerateFilters(tt.entity, tt.queryParam)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, actual)
			}
		})
	}
}
