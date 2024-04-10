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

func setDefaultCursor(d db.DBTX, entity db.TableEntity, cursor *models.PaginationCursor) error {
	if cursor.Value != nil && *cursor.Value != nil {
		return nil
	}
	f, ok := db.EntityFields[entity][cursor.Column]
	if !ok {
		return nil
	}

	if slices.Contains(infinityTypes, f.Type) {
		var defaultCursor interface{}
		if cursor.Direction == models.DirectionAsc {
			defaultCursor = "-Infinity"
		} else {
			defaultCursor = "Infinity"
		}
		cursor.Value = &defaultCursor

		return nil
	}

	query := fmt.Sprintf("select %s from %s order by %s %s limit 1", f.Db, entity, f.Db, cursor.Direction)
	var defaultCursor interface{}
	err := d.QueryRow(context.Background(), query).Scan(&defaultCursor)
	if err == pgx.ErrNoRows {
		return internal.NewErrorf(models.ErrorCodeNotFound, "no items exist yet")
	} else if err != nil {
		return err
	}

	cursor.Value = &defaultCursor

	return nil
}
