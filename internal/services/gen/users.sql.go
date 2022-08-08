// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const GetRoles = `-- name: GetRoles :many
select
  ENUM_RANGE(null::users.role)::text[]
`

func (q *Queries) GetRoles(ctx context.Context) ([][]string, error) {
	rows, err := q.db.Query(ctx, GetRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := [][]string{}
	for rows.Next() {
		var column_1 []string
		if err := rows.Scan(&column_1); err != nil {
			return nil, err
		}
		items = append(items, column_1)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetUser = `-- name: GetUser :one
select
  username,
  email,
  role,
  is_verified,
  is_active,
  is_superuser,
  created_at,
  updated_at,
  user_id,
  salt,
  password
  -- case when @get_db_data::boolean then
  --   (user_id)
  -- end as user_id, -- TODO sqlc.yaml overrides sql.NullInt64
  -- case when @get_db_data::boolean then
  --   (salt)
  -- end as salt, -- TODO sqlc.yaml overrides sql.NullString
  -- case when @get_db_data::boolean then
  --   (password)
  -- end as password -- TODO sqlc.yaml overrides sql.NullString
from
  users
where (email = LOWER($1)::text
  or $1::text is null)
and (username = $2::text
  or $2::text is null)
and (user_id = $3::int
  or $3::int is null)
limit 1
`

type GetUserParams struct {
	Email    sql.NullString `db:"email" json:"email"`
	Username sql.NullString `db:"username" json:"username"`
	UserID   sql.NullInt32  `db:"user_id" json:"user_id"`
}

type GetUserRow struct {
	Username    string    `db:"username" json:"username"`
	Email       string    `db:"email" json:"email"`
	Role        Role      `db:"role" json:"role"`
	IsVerified  bool      `db:"is_verified" json:"is_verified"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	IsSuperuser bool      `db:"is_superuser" json:"is_superuser"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	UserID      int64     `db:"user_id" json:"user_id"`
	Salt        string    `db:"salt" json:"salt"`
	Password    string    `db:"password" json:"password"`
}

func (q *Queries) GetUser(ctx context.Context, arg GetUserParams) (GetUserRow, error) {
	row := q.db.QueryRow(ctx, GetUser, arg.Email, arg.Username, arg.UserID)
	var i GetUserRow
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.Role,
		&i.IsVerified,
		&i.IsActive,
		&i.IsSuperuser,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.Salt,
		&i.Password,
	)
	return i, err
}

const ListAllUsers = `-- name: ListAllUsers :many
select
  user_id,
  username,
  email,
  role,
  is_verified,
  salt,
  password,
  is_active,
  is_superuser,
  created_at,
  updated_at
from
  users
where
  is_verified = $1::boolean
  or $1::boolean is null
`

type ListAllUsersRow struct {
	UserID      int64     `db:"user_id" json:"user_id"`
	Username    string    `db:"username" json:"username"`
	Email       string    `db:"email" json:"email"`
	Role        Role      `db:"role" json:"role"`
	IsVerified  bool      `db:"is_verified" json:"is_verified"`
	Salt        string    `db:"salt" json:"salt"`
	Password    string    `db:"password" json:"password"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	IsSuperuser bool      `db:"is_superuser" json:"is_superuser"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

func (q *Queries) ListAllUsers(ctx context.Context, isVerified sql.NullBool) ([]ListAllUsersRow, error) {
	rows, err := q.db.Query(ctx, ListAllUsers, isVerified)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAllUsersRow{}
	for rows.Next() {
		var i ListAllUsersRow
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Email,
			&i.Role,
			&i.IsVerified,
			&i.Salt,
			&i.Password,
			&i.IsActive,
			&i.IsSuperuser,
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
insert into users (username, email, password, salt, is_superuser, is_verified)
  values ($1, $2, $3, $4, $5, $6)
returning
  user_id, username, email, role, is_verified, is_active, is_superuser, created_at, updated_at
`

type RegisterNewUserParams struct {
	Username    string `db:"username" json:"username"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"password"`
	Salt        string `db:"salt" json:"salt"`
	IsSuperuser bool   `db:"is_superuser" json:"is_superuser"`
	IsVerified  bool   `db:"is_verified" json:"is_verified"`
}

type RegisterNewUserRow struct {
	UserID      int64     `db:"user_id" json:"user_id"`
	Username    string    `db:"username" json:"username"`
	Email       string    `db:"email" json:"email"`
	Role        Role      `db:"role" json:"role"`
	IsVerified  bool      `db:"is_verified" json:"is_verified"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	IsSuperuser bool      `db:"is_superuser" json:"is_superuser"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

func (q *Queries) RegisterNewUser(ctx context.Context, arg RegisterNewUserParams) (RegisterNewUserRow, error) {
	row := q.db.QueryRow(ctx, RegisterNewUser,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Salt,
		arg.IsSuperuser,
		arg.IsVerified,
	)
	var i RegisterNewUserRow
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.Role,
		&i.IsVerified,
		&i.IsActive,
		&i.IsSuperuser,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const ResetUserPassword = `-- name: ResetUserPassword :exec
update
  users
set
  password = $1,
  salt = $2
where
  email = LOWER($3)
`

type ResetUserPasswordParams struct {
	Password string `db:"password" json:"password"`
	Salt     string `db:"salt" json:"salt"`
	Email    string `db:"email" json:"email"`
}

func (q *Queries) ResetUserPassword(ctx context.Context, arg ResetUserPasswordParams) error {
	_, err := q.db.Exec(ctx, ResetUserPassword, arg.Password, arg.Salt, arg.Email)
	return err
}

const UpdateUserById = `-- name: UpdateUserById :one
update
  users
set
  password = COALESCE($1, password),
  salt = COALESCE($2, salt),
  username = COALESCE($3, username),
  email = COALESCE(LOWER($4), email)
where
  user_id = $5
returning
  user_id,
  username,
  email,
  role,
  is_verified,
  salt,
  password,
  is_active,
  is_superuser,
  created_at,
  updated_at
`

type UpdateUserByIdParams struct {
	Password sql.NullString `db:"password" json:"password"`
	Salt     sql.NullString `db:"salt" json:"salt"`
	Username sql.NullString `db:"username" json:"username"`
	Email    sql.NullString `db:"email" json:"email"`
	UserID   int64          `db:"user_id" json:"user_id"`
}

type UpdateUserByIdRow struct {
	UserID      int64     `db:"user_id" json:"user_id"`
	Username    string    `db:"username" json:"username"`
	Email       string    `db:"email" json:"email"`
	Role        Role      `db:"role" json:"role"`
	IsVerified  bool      `db:"is_verified" json:"is_verified"`
	Salt        string    `db:"salt" json:"salt"`
	Password    string    `db:"password" json:"password"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	IsSuperuser bool      `db:"is_superuser" json:"is_superuser"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

func (q *Queries) UpdateUserById(ctx context.Context, arg UpdateUserByIdParams) (UpdateUserByIdRow, error) {
	row := q.db.QueryRow(ctx, UpdateUserById,
		arg.Password,
		arg.Salt,
		arg.Username,
		arg.Email,
		arg.UserID,
	)
	var i UpdateUserByIdRow
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.Role,
		&i.IsVerified,
		&i.Salt,
		&i.Password,
		&i.IsActive,
		&i.IsSuperuser,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const UpdateUserRole = `-- name: UpdateUserRole :exec
update
  users
set
  role = $1
where
  user_id = $2
`

type UpdateUserRoleParams struct {
	Role   Role  `db:"role" json:"role"`
	UserID int64 `db:"user_id" json:"user_id"`
}

func (q *Queries) UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) error {
	_, err := q.db.Exec(ctx, UpdateUserRole, arg.Role, arg.UserID)
	return err
}

const VerifyUserByEmail = `-- name: VerifyUserByEmail :one
update
  users
set
  is_verified = 'true'
where
  email = LOWER($1)
returning
  email
`

func (q *Queries) VerifyUserByEmail(ctx context.Context, userEmail string) (string, error) {
	row := q.db.QueryRow(ctx, VerifyUserByEmail, userEmail)
	var email string
	err := row.Scan(&email)
	return email, err
}
