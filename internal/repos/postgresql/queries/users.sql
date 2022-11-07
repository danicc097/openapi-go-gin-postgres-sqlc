-- plpgsql-language-server:use-keyword-query-parameters
-- name: GetUser :one
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
where (email = LOWER(sqlc.narg('email'))::text
  or sqlc.narg('email')::text is null)
and (username = sqlc.narg('username')::text
  or sqlc.narg('username')::text is null)
and (user_id = sqlc.narg('user_id')::uuid
  or sqlc.narg('user_id')::uuid is null)
limit 1;

-- name: GetUsersWithJoins :many
select
  (
    case when @join_work_items::boolean = true then
      joined_work_items.work_items
    end)::jsonb as work_items -- if M2M
  , (
    case when @join_teams::boolean = true then
      joined_teams.teams
    end)::jsonb as teams -- if M2M
  , (
    case when @join_user_api_keys::boolean = true then
      ROW_TO_JSON(user_api_keys.*)
    end)::jsonb as user_api_key -- if O2O
  , (
    case when @join_time_entries::boolean = true then
      joined_time_entries.time_entries
    end)::jsonb as time_entries -- if O2M
  , users.*
from
  users
  ------------------------------
  left join (
    select
      member as work_items_user_id
      , JSON_AGG(work_items.*) as work_items
    from
      work_item_member uo
      join work_items using (work_item_id)
    where
      member in (
        select
          member
        from
          work_item_member
        where
          work_item_id = any (
            select
              work_item_id
            from
              work_items))
        group by
          member) joined_work_items on joined_work_items.work_items_user_id = users.user_id
  ------------------------------
  left join (
    select
      user_id as teams_user_id
      , JSON_AGG(teams.*) as teams
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
  ------------------------------
  -- this below would be O2M (we return an array agg)
  -- same as with work_item comments when selecting work_items
  -- since work_item_id is not unique in work_item_comments
  -- we assume cardinality:O2M. to distinguish O2M and M2M, cardinality:M2M comment, else O2M is assumed
  left join (
    select
      user_id
      , JSON_AGG(time_entries.*) as time_entries
    from
      time_entries
    group by
      user_id) joined_time_entries using (user_id)
------------------------------
-- this below would be O2O
  left join user_api_keys using (user_id);

-- if O2O. This is discovered from FK on user_api_keys.user_id being also a unique constraint
-- name: UpdateUserById :exec
update
  users
set
  username = COALESCE(sqlc.narg('username') , username)
  , email = COALESCE(LOWER(sqlc.narg('email')) , email)
where
  user_id = @user_id;

-- -- name: Test :exec
-- update
--   users
-- set
--   username = '@test'
--   , email = COALESCE(LOWER(sqlc.narg('email')) , email)
-- where
--   user_id = @user_id;
-- name: ListAllUsers :many
select
  user_id
  , username
  , email
  , role
  , created_at
  , updated_at
from
  users;
