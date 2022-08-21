// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: animals.sql

package db

import (
	"context"
	"database/sql"

	"github.com/jackc/pgtype"
)

const GetPetById = `-- name: GetPetById :one
select
  pet_id,
  color,
  metadata
from
  pets
  left join animals using (animal_id)
where
  pet_id = $1
limit 1
`

type GetPetByIdRow struct {
	PetID    int64          `db:"pet_id" json:"pet_id"`
	Color    sql.NullString `db:"color" json:"color"`
	Metadata pgtype.JSONB   `db:"metadata" json:"metadata"`
}

func (q *Queries) GetPetById(ctx context.Context, petID int64) (GetPetByIdRow, error) {
	row := q.db.QueryRow(ctx, GetPetById, petID)
	var i GetPetByIdRow
	err := row.Scan(&i.PetID, &i.Color, &i.Metadata)
	return i, err
}
