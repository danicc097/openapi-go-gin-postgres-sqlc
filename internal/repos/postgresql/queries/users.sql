-- name: GetUser :one
select
  username,
  email,
  role,
  is_superuser,
  created_at,
  updated_at,
  user_id
  -- case when @get_db_data::boolean then
  --   (user_id)
  -- end as user_id, -- TODO sqlc.yaml overrides sql.NullInt64
from
  users
where (email = LOWER(sqlc.narg('email'))::text
  or sqlc.narg('email')::text is null)
and (username = sqlc.narg('username')::text
  or sqlc.narg('username')::text is null)
and (user_id = sqlc.narg('user_id')::int
  or sqlc.narg('user_id')::int is null)
limit 1;

-- name: UpdateUserById :exec
update
  users
set
  username = COALESCE(sqlc.narg('username'), username),
  email = COALESCE(LOWER(sqlc.narg('email')), email)
where
  user_id = @user_id;

-- name: ListAllUsers :many
select
  user_id,
  username,
  email,
  role,
  is_superuser,
  created_at,
  updated_at
from
  users;
