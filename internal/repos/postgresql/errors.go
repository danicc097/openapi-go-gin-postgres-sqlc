package postgresql

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kenshaw/snaker"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/slices"
)

var errorUniqueViolationRegex = regexp.MustCompile(`\((.*)\)=\((.*)\)`)

func ParseDBErrorDetail(err error) error {
	newErr := internal.WrapErrorf(err, models.ErrorCodeUnknown, "database error")

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

	var xoErr *models.XoError

	if errors.As(err, &xoErr) {
		if errors.Is(err, pgx.ErrNoRows) {
			return internal.NewErrorf(models.ErrorCodeNotFound, "%s", xoErr.Entity+" not found")
		}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// fmt.Printf("ColumnName: %+v\n", pgErr.ColumnName)
		// fmt.Printf("Hint: %+v\n", pgErr.Hint)
		// fmt.Printf("TableName: %+v\n", pgErr.TableName)
		// fmt.Printf("SchemaName: %+v\n", pgErr.SchemaName)
		// fmt.Printf("Where: %+v\n", pgErr.Where)
		// fmt.Printf("DataTypeName: %+v\n", pgErr.DataTypeName)
		// fmt.Printf("pgErr err: %+v\n", pgErr.Error())

		var column, value string
		switch pgErr.Code {
		// TODO: better error detail message, e.g. trim id and if idTrimmed then build message
		// with "invalid <prefix> ID (<val>)", or "<prefix>" not found.
		// where prefix is trimmed _id and sentence case -> workItemTag=>work item tag
		case pgerrcode.UniqueViolation, pgerrcode.ForeignKeyViolation:
			matches := errorUniqueViolationRegex.FindStringSubmatch(pgErr.Detail)
			if len(matches) == 0 {
				break
			}
			// TODO: in case of error in field unrelated to request params, frontend will simply attempt to show in nearest parent that does
			// exist and default to generic callout.
			// custom tag mappers done at service level or above
			column, value = matches[1], matches[2]
			loc := []string{snaker.ForceLowerCamelIdentifier(column)}
			msg := fmt.Sprintf("%s %q", loc[0], value)
			if strings.Contains(column, ", ") {
				loc = []string{} // ignore loc for multicolumn constraint error
				columns := strings.Split(column, ", ")
				values := strings.Split(value, ", ")
				multierr := []string{}
				for i := range len(columns) {
					multierr = append(multierr, fmt.Sprintf("%s=%s", snaker.ForceLowerCamelIdentifier(columns[i]), values[i]))
				}
				msg = "combination of " + slices.JoinWithAnd(multierr)
			}

			msgSuffix := ""
			var code models.ErrorCode
			if pgErr.Code == pgerrcode.ForeignKeyViolation {
				msgSuffix = " is invalid"
				code = models.ErrorCodeInvalidArgument
			} else if pgErr.Code == pgerrcode.UniqueViolation {
				msgSuffix = " already exists"
				code = models.ErrorCodeAlreadyExists
			}

			newErr = internal.NewErrorWithLocf(code, loc, "%s", msg+msgSuffix)
		default:
			msg := pgErr.Message // always set
			if pgErr.Detail != "" {
				msg += fmt.Sprintf(" (%s)", pgErr.Detail)
			}
			newErr = internal.NewErrorf(models.ErrorCodeUnknown, "%s", msg)
		}
	}

	return newErr
}
