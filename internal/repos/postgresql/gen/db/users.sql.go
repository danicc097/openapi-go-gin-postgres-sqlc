// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
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
  -- end as user_id, -- TODO sqlc.yaml overrides sql.NullInt64
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
	Email    pgtype.Text `db:"email" json:"email"`
	Username pgtype.Text `db:"username" json:"username"`
	UserID   pgtype.UUID `db:"user_id" json:"user_id"`
}

type GetUserRow struct {
	Username  string             `db:"username" json:"username"`
	Email     string             `db:"email" json:"email"`
	RoleRank  int16              `db:"role_rank" json:"role_rank"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
	UserID    pgtype.UUID        `db:"user_id" json:"user_id"`
}

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

const ListAllUsers2 = `-- name: ListAllUsers2 :many
select
  user_id
  , username
  , email
  , role_rank
  , created_at
  , updated_at
from
  users
`

type ListAllUsers2Row struct {
	UserID    pgtype.UUID        `db:"user_id" json:"user_id"`
	Username  string             `db:"username" json:"username"`
	Email     string             `db:"email" json:"email"`
	RoleRank  int16              `db:"role_rank" json:"role_rank"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

func (q *Queries) ListAllUsers2(ctx context.Context, db DBTX) ([]ListAllUsers2Row, error) {
	rows, err := db.Query(ctx, ListAllUsers2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAllUsers2Row{}
	for rows.Next() {
		var i ListAllUsers2Row
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Email,
			&i.RoleRank,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const RegisterNewUser = `-- name: RegisterNewUser :one
insert into users (
  username
  , email
  , role_rank)
values (
  $1
  , $2
  , $3)
returning
  user_id
  , username
  , email
  , role_rank
  , created_at
  , updated_at
`

type RegisterNewUserParams struct {
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	RoleRank int16  `db:"role_rank" json:"role_rank"`
}

type RegisterNewUserRow struct {
	UserID    pgtype.UUID        `db:"user_id" json:"user_id"`
	Username  string             `db:"username" json:"username"`
	Email     string             `db:"email" json:"email"`
	RoleRank  int16              `db:"role_rank" json:"role_rank"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

// plpgsql-language-server:disable
func (q *Queries) RegisterNewUser(ctx context.Context, db DBTX, arg RegisterNewUserParams) (RegisterNewUserRow, error) {
	row := db.QueryRow(ctx, RegisterNewUser, arg.Username, arg.Email, arg.RoleRank)
	var i RegisterNewUserRow
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.RoleRank,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const Test = `-- name: Test :exec
select
  user_id
  , username
  , email
  , role_rank
  , created_at
  , updated_at
from
  users
`

type TestRow struct {
	UserID    pgtype.UUID        `db:"user_id" json:"user_id"`
	Username  string             `db:"username" json:"username"`
	Email     string             `db:"email" json:"email"`
	RoleRank  int16              `db:"role_rank" json:"role_rank"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

// update
//
//	users
//
// set
//
//	username = null
//	, email = COALESCE(LOWER(sqlc.narg('email')) , email)
//
// where
//
//	user_id = @user_id;
func (q *Queries) Test(ctx context.Context, db DBTX) error {
	_, err := db.Exec(ctx, Test)
	return err
}
