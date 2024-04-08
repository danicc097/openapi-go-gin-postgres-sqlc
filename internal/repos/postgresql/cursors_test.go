package postgresql

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func TestSetDefaultCursors(t *testing.T) {
	tests := []struct {
		name          string
		wantQuery     string
		queryResult   string
		cursor        models.PaginationCursor
		wantCursor    interface{}
		entity        db.TableEntity
		errorContains string
	}{
		{
			name: "infinity",
			cursor: models.PaginationCursor{
				Column:    "createdAt",
				Direction: models.DirectionAsc,
				Value:     nil,
			},
			wantCursor: "-Infinity",
			entity:     db.TableEntityUser,
		},
		{
			name:       "cursor ignored if set",
			wantCursor: "something",
			cursor: models.PaginationCursor{
				Column:    "fullName",
				Direction: models.DirectionAsc,
				Value:     pointers.New[interface{}]("something"),
			},
			entity: db.TableEntityUser,
		},
		{
			name:        "find in query",
			wantQuery:   "select full_name from users order by full_name asc limit 1",
			queryResult: "my name",
			wantCursor:  "my name",
			cursor: models.PaginationCursor{
				Column:    "fullName",
				Direction: models.DirectionAsc,
				Value:     nil,
			},
			entity: db.TableEntityUser,
		},
		{
			// should collate via CREATE COLLATION numeric (provider = icu, locale = 'en@colNumeric=yes')
			// at some point, but most likely requires indexing on collated row for it to work
			name:          "no rows",
			wantQuery:     "select full_name from users order by full_name desc limit 1",
			errorContains: "no items exist yet",
			cursor: models.PaginationCursor{
				Column:    "fullName",
				Direction: models.DirectionDesc,
				Value:     nil,
			},
			entity: db.TableEntityUser,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock, err := pgxmock.NewConn()
			require.NoError(t, err)

			if tt.wantQuery != "" {
				rr := pgxmock.NewRows([]string{tt.queryResult}) // existing rows in mock db
				if tt.queryResult != "" {
					rr.AddRow(tt.queryResult) // result rows
				}
				mock.ExpectQuery(tt.wantQuery).WillReturnRows(rr)
			}

			newCursors, err := setDefaultCursors(mock, tt.entity, models.PaginationCursors{tt.cursor})

			if tt.errorContains != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errorContains)

				return
			} else {
				require.NoError(t, err)
				require.Len(t, newCursors, 1)
				require.EqualValues(t, tt.wantCursor, *newCursors[0].Value)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err, "unfulfilled expectations")
		})
	}
}
