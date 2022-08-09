-- name: GetUser :one
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
where (email = LOWER(sqlc.narg('email'))::text
  or sqlc.narg('email')::text is null)
and (username = sqlc.narg('username')::text
  or sqlc.narg('username')::text is null)
and (user_id = sqlc.narg('user_id')::int
  or sqlc.narg('user_id')::int is null)
limit 1;

-- name: RegisterNewUser :one
insert into users (username, email, password, salt, is_superuser, is_verified)
  values (@username, @email, @password, @salt, @is_superuser, @is_verified)
returning
  user_id, username, email, role, is_verified, is_active, is_superuser, created_at, updated_at;

-- name: UpdateUserById :one
update
  users
set
  password = COALESCE(sqlc.narg('password'), password),
  salt = COALESCE(sqlc.narg('salt'), salt),
  username = COALESCE(sqlc.narg('username'), username),
  email = COALESCE(LOWER(sqlc.narg('email')), email)
where
  user_id = @user_id
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
  updated_at;

-- name: ListAllUsers :many
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
  is_verified = sqlc.narg('is_verified')::boolean
  or sqlc.narg('is_verified')::boolean is null;

-- name: VerifyUserByEmail :one
update
  users
set
  is_verified = 'true'
where
  email = LOWER(@user_email)
returning
  email;

-- name: ResetUserPassword :exec
update
  users
set
  password = @password,
  salt = @salt
where
  email = LOWER(@email);

-- name: UpdateUserRole :exec
update
  users
set
  role = @role
where
  user_id = @user_id;
