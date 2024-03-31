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

func arrayFilter(ss []string, mode models.PaginationFilterModes) *models.PaginationFilter {
	v := models.PaginationFilterArray{
		Value:      ss,
		FilterMode: mode,
	}
	j, _ := json.Marshal(v)
	p := models.PaginationFilter{}
	_ = json.Unmarshal(j, &p)

	return &p
}

func primitiveFilter(s string, mode models.PaginationFilterModes) *models.PaginationFilter {
	v := models.PaginationFilterPrimitive{
		Value:      pointers.New(s),
		FilterMode: mode,
	}
	j, _ := json.Marshal(v)
	p := models.PaginationFilter{}
	_ = json.Unmarshal(j, &p)

	return &p
}

func TestGenerateFilters(t *testing.T) {
	const testEntity db.TableEntity = "testEntity"
	db.EntityFilters[testEntity] = map[string]db.Filter{
		"createdAt": {Type: "date-time", Db: "created_at", Nullable: false},
		"fullName":  {Type: "string", Db: "full_name", Nullable: true},
		"count":     {Type: "integer", Db: "db_count", Nullable: false},
		"countF":    {Type: "float", Db: "db_countf", Nullable: false},
		"bool":      {Type: "boolean", Db: "db_bool", Nullable: false},
	}

	tests := []struct {
		name        string
		pagParams   models.PaginationItems
		expected    map[string][]interface{}
		errContains string
	}{
		{
			name: "unknown column ignored",
			pagParams: models.PaginationItems{
				"=?¿": {
					Filter: primitiveFilter("abc", models.PaginationFilterModesBetween),
				},
			},
			expected: map[string][]interface{}{},
		},
		{
			name: "empty filter ignored",
			pagParams: models.PaginationItems{
				"fullName": {
					Filter: nil,
				},
			},
			expected: map[string][]interface{}{},
		},
		{
			name: "null",
			pagParams: models.PaginationItems{
				"fullName": {
					Filter: primitiveFilter("", models.PaginationFilterModesEmpty),
				},
			},
			expected: map[string][]interface{}{
				"full_name is null": {},
			},
		},
		{
			name: "not null",
			pagParams: models.PaginationItems{
				"fullName": {
					Filter: primitiveFilter("", models.PaginationFilterModesNotEmpty),
				},
			},
			expected: map[string][]interface{}{
				"full_name is not null": {},
			},
		},
		{
			name: "string equals",
			pagParams: models.PaginationItems{
				"fullName": {
					Filter: primitiveFilter("John Doe", models.PaginationFilterModesEquals),
				},
			},
			expected: map[string][]interface{}{
				"full_name = $i": {"John Doe"},
			},
		},
		{
			name: "string startsWith",
			pagParams: models.PaginationItems{
				"fullName": {
					Filter: primitiveFilter("John Doe", models.PaginationFilterModesStartsWith),
				},
			},
			expected: map[string][]interface{}{
				"full_name ILIKE $i": {"John Doe%"},
			},
		},
		{
			name: "string endsWith",
			pagParams: models.PaginationItems{
				"fullName": {
					Filter: primitiveFilter("John Doe", models.PaginationFilterModesEndsWith),
				},
			},
			expected: map[string][]interface{}{
				"full_name ILIKE $i": {"%John Doe"},
			},
		},
		{
			name: "string contains",
			pagParams: models.PaginationItems{
				"fullName": {
					Filter: primitiveFilter("John Doe", models.PaginationFilterModesContains),
				},
			},
			expected: map[string][]interface{}{
				"full_name ILIKE $i": {"%John Doe%"},
			},
		},
		{
			name: "integer equals",
			pagParams: models.PaginationItems{
				"count": {
					Filter: primitiveFilter("30", models.PaginationFilterModesEquals),
				},
			},
			expected: map[string][]interface{}{
				"db_count = $i": {30},
			},
		},
		{
			name: "bad integer",
			pagParams: models.PaginationItems{
				"count": {
					Filter: primitiveFilter("30.123", models.PaginationFilterModesEquals),
				},
			},
			errContains: "db_count: invalid integer \"30.123\"",
		},
		{
			name: "float equals",
			pagParams: models.PaginationItems{
				"countF": {
					Filter: primitiveFilter("1.123", models.PaginationFilterModesEquals),
				},
			},
			expected: map[string][]interface{}{
				"db_countf = $i": {1.123},
			},
		},
		{
			name: "boolean equals",
			pagParams: models.PaginationItems{
				"bool": {
					Filter: primitiveFilter("true", models.PaginationFilterModesEquals),
				},
			},
			expected: map[string][]interface{}{
				"db_bool = $i": {true},
			},
		},
		{
			name: "date-time between",
			pagParams: models.PaginationItems{
				"createdAt": {
					Filter: arrayFilter(
						[]string{"2023-01-01T00:00:00Z", "2023-12-31T23:59:59Z"},
						models.PaginationFilterModesBetween,
					),
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
			name: "date-time betweenInclusive",
			pagParams: models.PaginationItems{
				"createdAt": {
					Filter: arrayFilter(
						[]string{"2023-01-01T00:00:00Z", "2023-12-31T23:59:59Z"},
						models.PaginationFilterModesBetweenInclusive,
					),
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
			got, err := postgresql.GenerateFilters(testEntity, tc.pagParams)
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
