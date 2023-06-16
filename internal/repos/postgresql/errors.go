package postgresql

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/jackc/pgerrcode"
)

var errorDetailRegex = regexp.MustCompile(`\((.*)\)=\((.*)\)`)

func parseErrorDetail(err error) error {
	newErr := internal.WrapErrorf(err, internal.ErrorCodeUnknown, err.Error())

	/**
	 * TODO: will have generic xo Error struct, which has Entity field.
	 * so we build error with NewErrorf("User", err).
	 * Then in responses.go we use errors.As for this Error struct (will stop at Error, not pgx error)
	 * So we would grab e.Entity and construct the new string based on wrapped errors in Error
	 * which we already are handling (pgErr, pgx.ErrNoRows)...
	 * the end goal is that error.Title in responses.go err.Cause() gives something like: `<.Entity> not found`, `... already exists`
	 * which can directly be shown in a callout.
	 *
	 *
	 */

	var column, value string
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		if errors.Is(err, pgx.ErrNoRows) {
			newErr = internal.NewErrorf(internal.ErrorCodeNotFound, "not found")
		}
	} else {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			matches := errorDetailRegex.FindStringSubmatch(pgErr.Detail)
			if len(matches) == 0 {
				break
			}
			column, value = matches[1], matches[2]
			newErr = internal.NewErrorf(internal.ErrorCodeAlreadyExists, fmt.Sprintf("%s %q already exists", column, value))
		default:
			newErr = internal.NewErrorf(internal.ErrorCodeUnknown, fmt.Sprintf("%s | %s", pgErr.Detail, pgErr.Message))
		}
	}

	return newErr
}
