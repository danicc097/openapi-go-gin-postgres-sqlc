package postgresql

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/jackc/pgerrcode"
)

var errorDetailRegex = regexp.MustCompile(`\((.*)\)=\((.*)\)`)

func parseErrorDetail(err error) error {
	newErr := internal.WrapErrorf(err, internal.ErrorCodeUnknown, err.Error())

	var column, value string
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			matches := errorDetailRegex.FindStringSubmatch(pgErr.Detail)
			if len(matches) == 0 {
				break
			}
			column, value = matches[1], matches[2]
			newErr = internal.WrapErrorf(err, internal.ErrorCodeAlreadyExists, fmt.Sprintf("%s %q already exists", column, value))
		default:
			newErr = internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("%s | %s", pgErr.Detail, pgErr.Message))
		}
	}

	return newErr
}
