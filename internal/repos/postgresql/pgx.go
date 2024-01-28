package postgresql

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/jackc/pgx/v5"
)

func rowsToJSON(rows pgx.Rows) []byte {
	var result any

	for rows.Next() {
		values, _ := rows.Values()
		for _, v := range values {
			result = v
		}
	}

	json, _ := json.Marshal(result)

	return json
}

// DynamicQuery returns an SQL query result as JSON.
func DynamicQuery(d db.DBTX, query string) (string, error) {
	rows, err := d.Query(context.Background(), fmt.Sprintf("select row_to_json(t) from (%s) as t", query))
	if err != nil {
		return "", fmt.Errorf("d.Query: %w", err)
	}

	return string(rowsToJSON(rows)), nil
}
