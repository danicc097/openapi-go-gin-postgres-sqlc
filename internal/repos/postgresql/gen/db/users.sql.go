// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

const GetUser = `-- name: GetUser :one
select
  username
  , email
  , role
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
	Email    sql.NullString `db:"email" json:"email"`
	Username sql.NullString `db:"username" json:"username"`
	UserID   uuid.NullUUID  `db:"user_id" json:"user_id"`
}

type GetUserRow struct {
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	Role      UserRole  `db:"role" json:"role"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
}

// plpgsql-language-server:use-keyword-query-parameters
func (q *Queries) GetUser(ctx context.Context, db DBTX, arg GetUserParams) (GetUserRow, error) {
	row := db.QueryRow(ctx, GetUser, arg.Email, arg.Username, arg.UserID)
	var i GetUserRow
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const GetUsersWithJoins = `-- name: GetUsersWithJoins :many
select
  (case when $1::boolean = true then joined_tasks.tasks end)::jsonb as tasks -- if M2M
  , (case when $2::boolean = true then joined_teams.teams end)::jsonb as teams -- if M2M
  , (case when $3::boolean = true then row_to_json(user_api_keys.*) end)::jsonb as user_api_key -- if O2O
  , (case when $4::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries -- if O2M
  , users.user_id, users.username, users.email, users.scopes, users.first_name, users.last_name, users.full_name, users.external_id, users.role, users.created_at, users.updated_at, users.deleted_at
from
  users
left join (
  select
    member as tasks_user_id
    , json_agg(tasks.*) as tasks
  from
    task_member uo
    join tasks using (task_id)
  where
    member in (
      select
        member
      from
        task_member
      where
        task_id = any (
          select
            task_id
          from
            tasks))
      group by
        member) joined_tasks on joined_tasks.tasks_user_id = users.user_id
left join (
  select
    user_id as teams_user_id
    , json_agg(teams.*) as teams
  from
    user_team uo
    join teams using (team_id)
  where
    user_id in (
      select
        user_id
      from
        user_team
      where
        team_id = any (
          select
            team_id
          from
            teams))
      group by
        user_id) joined_teams on joined_teams.teams_user_id = users.user_id
left join (
  select
  user_id
    , json_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries using (user_id)
left join user_api_keys using (user_id)
`

type GetUsersWithJoinsParams struct {
	JoinTasks       bool `db:"join_tasks" json:"join_tasks"`
	JoinTeams       bool `db:"join_teams" json:"join_teams"`
	JoinUserApiKeys bool `db:"join_user_api_keys" json:"join_user_api_keys"`
	JoinTimeEntries bool `db:"join_time_entries" json:"join_time_entries"`
}

type GetUsersWithJoinsRow struct {
	Tasks       pgtype.JSONB   `db:"tasks" json:"tasks"`
	Teams       pgtype.JSONB   `db:"teams" json:"teams"`
	UserApiKey  pgtype.JSONB   `db:"user_api_key" json:"user_api_key"`
	TimeEntries pgtype.JSONB   `db:"time_entries" json:"time_entries"`
	UserID      uuid.UUID      `db:"user_id" json:"user_id"`
	Username    string         `db:"username" json:"username"`
	Email       string         `db:"email" json:"email"`
	Scopes      []string       `db:"scopes" json:"scopes"`
	FirstName   sql.NullString `db:"first_name" json:"first_name"`
	LastName    sql.NullString `db:"last_name" json:"last_name"`
	FullName    sql.NullString `db:"full_name" json:"full_name"`
	ExternalID  sql.NullString `db:"external_id" json:"external_id"`
	Role        UserRole       `db:"role" json:"role"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at" json:"deleted_at"`
}

// ----------------------------
// ----------------------------
// ----------------------------
// this below would be O2M (we return an array agg)
// same as with work_item comments when selecting work_items
// since work_item_id is not unique in work_item_comments
// we assume cardinality:O2M. to distinguish O2M and M2M, cardinality:M2M comment, else O2M is assumed
// ----------------------------
// this below would be O2O
func (q *Queries) GetUsersWithJoins(ctx context.Context, db DBTX, arg GetUsersWithJoinsParams) ([]GetUsersWithJoinsRow, error) {
	rows, err := db.Query(ctx, GetUsersWithJoins,
		arg.JoinTasks,
		arg.JoinTeams,
		arg.JoinUserApiKeys,
		arg.JoinTimeEntries,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUsersWithJoinsRow{}
	for rows.Next() {
		var i GetUsersWithJoinsRow
		if err := rows.Scan(
			&i.Tasks,
			&i.Teams,
			&i.UserApiKey,
			&i.TimeEntries,
			&i.UserID,
			&i.Username,
			&i.Email,
			&i.Scopes,
			&i.FirstName,
			&i.LastName,
			&i.FullName,
			&i.ExternalID,
			&i.Role,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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

const ListAllUsers = `-- name: ListAllUsers :many
select
  user_id
  , username
  , email
  , role
  , created_at
  , updated_at
from
  users
`

type ListAllUsersRow struct {
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	Role      UserRole  `db:"role" json:"role"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// -- name: Test :exec
// update
//
//	users
//
// set
//
//	username = '@test'
//	, email = COALESCE(LOWER(sqlc.narg('email')) , email)
//
// where
//
//	user_id = @user_id;
func (q *Queries) ListAllUsers(ctx context.Context, db DBTX) ([]ListAllUsersRow, error) {
	rows, err := db.Query(ctx, ListAllUsers)
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

const UpdateUserById = `-- name: UpdateUserById :exec


update
  users
set
  username = COALESCE($1 , username)
  , email = COALESCE(LOWER($2) , email)
where
  user_id = $3
`

type UpdateUserByIdParams struct {
	Username sql.NullString `db:"username" json:"username"`
	Email    sql.NullString `db:"email" json:"email"`
	UserID   uuid.UUID      `db:"user_id" json:"user_id"`
}

// if O2O. This is discovered from FK on user_api_keys.user_id being also a unique constraint
func (q *Queries) UpdateUserById(ctx context.Context, db DBTX, arg UpdateUserByIdParams) error {
	_, err := db.Exec(ctx, UpdateUserById, arg.Username, arg.Email, arg.UserID)
	return err
}
