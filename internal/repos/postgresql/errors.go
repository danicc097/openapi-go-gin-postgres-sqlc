package postgresql

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kenshaw/snaker"

	"github.com/jackc/pgerrcode"
)

var errorUniqueViolationRegex = regexp.MustCompile(`\((.*)\)=\((.*)\)`)

func parseDBErrorDetail(err error) error {
	newErr := internal.WrapErrorf(err, models.ErrorCodeUnknown, err.Error())

	/**
	 * TODO: will have generic xo Error struct, which has Entity field.
	 * Then in responses.go we use errors.As for this Error struct (will stop at Error, not pgx error)
	 * So we would grab e.Entity and construct the new string based on wrapped errors in Error
	 * which we already are handling (pgErr, pgx.ErrNoRows)...
	 * the end goal is that error.Title in responses.go err.Cause() gives something like: `<.Entity> not found`, `... already exists`
	 * which can directly be shown in a callout.
	 *
	 *
	 */

	var xoErr *db.XoError

	if errors.As(err, &xoErr) {
		if errors.Is(err, pgx.ErrNoRows) {
			return internal.NewErrorf(models.ErrorCodeNotFound, xoErr.Entity+" not found")
		}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		fmt.Printf("ColumnName: %+v\n", pgErr.ColumnName)
		fmt.Printf("Hint: %+v\n", pgErr.Hint)
		fmt.Printf("TableName: %+v\n", pgErr.TableName)
		fmt.Printf("SchemaName: %+v\n", pgErr.SchemaName)
		fmt.Printf("Where: %+v\n", pgErr.Where)
		fmt.Printf("DataTypeName: %+v\n", pgErr.DataTypeName)
		fmt.Printf("pgErr err: %+v\n", pgErr.Error())

		var column, value string
		switch pgErr.Code {
		// TODO: better error detail message, e.g. trim id and if idTrimmed then build message
		// with "invalid <prefix> ID (<val>)", or "<prefix>" not found.
		// where prefix is trimmed _id and sentence case -> workItemTag=>work item tag
		case pgerrcode.UniqueViolation:
			matches := errorUniqueViolationRegex.FindStringSubmatch(pgErr.Detail)
			if len(matches) == 0 {
				break
			}
			// TODO: handle multicolumn (should be empty loc slice, which will show error on whole object)
			column, value = matches[1], matches[2]
			jsonTag := snaker.ForceLowerCamelIdentifier(column)
			newErr = internal.NewErrorWithLocf(models.ErrorCodeAlreadyExists, []string{jsonTag}, fmt.Sprintf("%s %q already exists", jsonTag, value))
		case pgerrcode.ForeignKeyViolation:
			matches := errorUniqueViolationRegex.FindStringSubmatch(pgErr.Detail)
			if len(matches) == 0 {
				break
			}
			// TODO: handle multicolumn (should be empty loc slice, which will show error on whole object)
			// TODO: in case of error in field unrelated to request params, frontend will simply attempt to show in nearest parent that does
			// exist and default to generic callout.
			// custom tag mappers done at service level or above
			column, value = matches[1], matches[2]
			jsonTag := snaker.ForceLowerCamelIdentifier(column)
			newErr = internal.NewErrorWithLocf(models.ErrorCodeInvalidArgument, []string{jsonTag}, fmt.Sprintf("%s %q is invalid", jsonTag, value))
		default:
			newErr = internal.NewErrorf(models.ErrorCodeUnknown, fmt.Sprintf("%s | %s", pgErr.Detail, pgErr.Message))
		}
	}

	return newErr
}
