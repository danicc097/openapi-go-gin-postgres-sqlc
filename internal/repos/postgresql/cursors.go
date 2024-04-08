package postgresql

import (
	"context"
	"fmt"
	"slices"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/jackc/pgx/v5"
)

var infinityTypes = []db.ColumnSimpleType{
	db.ColumnSimpleTypeDateTime,
	db.ColumnSimpleTypeNumber,
	db.ColumnSimpleTypeInteger,
}

func setDefaultCursors(d db.DBTX, entity db.TableEntity, cursors models.PaginationCursors) (models.PaginationCursors, error) {
	for i, v := range cursors {
		if v.Value != nil {
			continue
		}
		f, ok := db.EntityFields[entity][v.Column]
		if !ok {
			continue
		}

		if slices.Contains(infinityTypes, f.Type) {
			var defaultCursor interface{}
			if v.Direction == models.DirectionAsc {
				defaultCursor = "-Infinity"
			} else {
				defaultCursor = "Infinity"
			}
			v.Value = &defaultCursor
			cursors[i] = v

			continue
		}

		query := fmt.Sprintf("select %s from %s order by %s %s limit 1", f.Db, entity, f.Db, v.Direction)
		var defaultCursor interface{}
		err := d.QueryRow(context.Background(), query).Scan(&defaultCursor)
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(models.ErrorCodeNotFound, "no items exist yet")
		} else if err != nil {
			return nil, err
		}

		v.Value = &defaultCursor
		cursors[i] = v
	}

	return cursors, nil
}
