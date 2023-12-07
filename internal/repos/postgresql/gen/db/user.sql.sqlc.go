// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: user.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const GetUser = `-- name: GetUser :one
select
  username
  , email
  , role_rank
  , created_at
  , updated_at
  , user_id
  -- case when @get_db_data::boolean then
  --   (user_id)
  -- end as user_id,
from
  users
where (email = LOWER($1)::text
  or $1::text is null)
and (username = $2::text
  or $2::text is null)
and (user_id = $3::uuid
  or $3::uuid is null)
limit 1
`

type GetUserParams struct {
	Email    *string    `db:"email" json:"email"`
	Username *string    `db:"username" json:"username"`
	UserID   *uuid.UUID `db:"user_id" json:"user_id"`
}

type GetUserRow struct {
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	RoleRank  int16     `db:"role_rank" json:"role_rank"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
}

// plpgsql-language-server:use-keyword-query-parameter
func (q *Queries) GetUser(ctx context.Context, db DBTX, arg GetUserParams) (GetUserRow, error) {
	row := db.QueryRow(ctx, GetUser, arg.Email, arg.Username, arg.UserID)
	var i GetUserRow
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.RoleRank,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const IsUserInProject = `-- name: IsUserInProject :one
select
  exists (
    select
      1
    from
      user_team ut
      join teams t on ut.team_id = t.team_id
    where
      ut.member = $1
      and t.project_id = $2)
`

type IsUserInProjectParams struct {
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
	ProjectID int32     `db:"project_id" json:"project_id"`
}

func (q *Queries) IsUserInProject(ctx context.Context, db DBTX, arg IsUserInProjectParams) (bool, error) {
	row := db.QueryRow(ctx, IsUserInProject, arg.UserID, arg.ProjectID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
