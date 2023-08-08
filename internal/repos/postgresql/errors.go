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

func parseDbErrorDetail(err error) error {
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
		var column, value string
		switch pgErr.Code {
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
			// in case of error in field unrelated to request params, frontend will simply attempt to show in nearest parent that does
			// exist and default to generic callout.
			column, value = matches[1], matches[2]
			jsonTag := snaker.ForceLowerCamelIdentifier(column)
			newErr = internal.NewErrorWithLocf(models.ErrorCodeInvalidArgument, []string{jsonTag}, fmt.Sprintf("%s %q is invalid", jsonTag, value))
		default:
			newErr = internal.NewErrorf(models.ErrorCodeUnknown, fmt.Sprintf("%s | %s", pgErr.Detail, pgErr.Message))
		}
	}

	return newErr
}
