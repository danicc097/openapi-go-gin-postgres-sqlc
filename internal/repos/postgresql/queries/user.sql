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

-- name: IsUserInProject :one
select
  exists (
    select
      1
    from
      user_team ut
      join teams t on ut.team_id = t.team_id
    where
      ut.member = @user_id
      and t.project_id = @project_id);
