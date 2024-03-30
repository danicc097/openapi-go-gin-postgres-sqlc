package postgresql_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/require"
)

func TestGenerateFilters(t *testing.T) {
	const testEntity db.TableEntity = "testEntity"
	db.EntityFilters[testEntity] = map[string]db.Filter{
		"createdAt": {Type: "date-time", Db: "created_at", Nullable: false},
		"fullName":  {Type: "string", Db: "full_name", Nullable: true},
		"count":     {Type: "integer", Db: "db_count", Nullable: false},
	}

	tests := []struct {
		name        string
		queryParam  map[string]models.PaginationFilter
		expected    map[string][]interface{}
		errContains string
	}{
		{
			name: "StringEqualsFilter",
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
			name: "StringContainsFilter",
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
			name: "IntegerEqualsFilter",
			queryParam: map[string]models.PaginationFilter{
				"count": {
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
				"db_count = $i": {30},
			},
		},
		{
			name: "DateTimeBetweenFilter",
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
		{
			name: "DateTimeBetweenFilter",
			queryParam: map[string]models.PaginationFilter{
				"createdAt": {
					Value: func() models.PaginationFilterValue {
						v := models.PaginationFilterArrayValue{
							Value:      []string{"2023-01-01T00:00:00Z", "2023-12-31T23:59:59Z"},
							FilterMode: models.PaginationFilterModesBetweenInclusive,
						}
						j, _ := json.Marshal(v)
						p := models.PaginationFilterValue{}
						_ = json.Unmarshal(j, &p)
						return p
					}(),
				},
			},
			expected: map[string][]interface{}{
				"created_at >= $i AND created_at <= $i": {
					time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := postgresql.GenerateFilters(testEntity, tc.queryParam)
			if err != nil && tc.errContains == "" {
				t.Errorf("unexpected error: %v", err)

				return
			}
			if tc.errContains != "" {
				if err == nil {
					t.Errorf("expected error but got nothing")

					return
				}
				require.ErrorContains(t, err, tc.errContains)

				return
			}

			require.EqualValues(t, tc.expected, got)
		})
	}
}
