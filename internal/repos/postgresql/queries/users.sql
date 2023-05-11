-- plpgsql-language-server:use-keyword-query-parameter
-- name: GetUser :one
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
where (email = LOWER(sqlc.narg('email'))::text
  or sqlc.narg('email')::text is null)
and (username = sqlc.narg('username')::text
  or sqlc.narg('username')::text is null)
and (user_id = sqlc.narg('user_id')::uuid
  or sqlc.narg('user_id')::uuid is null)
limit 1;

-- name: RegisterNewUser :one
-- plpgsql-language-server:disable
insert into users (username , email , role_rank)
  values (@username , @email , @role_rank)
returning
  user_id , username , email , role_rank , created_at , updated_at;

-- name: Test :exec
-- update
--   users
-- set
--   username = null
--   , email = COALESCE(LOWER(sqlc.narg('email')) , email)
-- where
--   user_id = @user_id;
-- name: ListAllUsers :many
select
  user_id
  , username
  , email
  , role_rank
  , created_at
  , updated_at
from
  users;
