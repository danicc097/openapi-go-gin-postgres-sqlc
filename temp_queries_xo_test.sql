-- plpgsql-language-server:use-keyword-query-parameters

--set enable_seqscan='off';

select
  -- since we return jsons, we can unmarshal row results into xo's structs,
  -- which already account for which fields are ignored, excluded, etc.
  -- then for openapi requests we can have manual or generated adapters to convert request bodies to xo models.
  -- For responses use a struct as is: see https://github.com/swaggest/rest/ -> we could generate
  -- openapi schema refs from xo models (we just care about types, this will be used for responses only)
  -- so we could easily respond with a huge json object.
  -- specifically see https://github.com/swaggest/openapi-go (Type-based reflection of Go structures to OpenAPI 3 schema.)

  -- NOTE: ignore
  -- see https://www.alexedwards.net/blog/using-postgresql-jsonb for database sql
  -- When the fields in a JSON/JSONB column are known in advance, you can map the contents
  -- of the JSON/JSONB column to and from a struct. To do this, you'll need make sure the struct implements:
  -- The driver.Valuer interface, such that it marshals the object into a JSON byte slice that can be understood by the database.
  -- The sql.Scanner interface, such that it unmarshals a JSON byte slice from the database into the struct fields.

  -- https://github.com/jackc/pgtype/blob/e19b507b9d6f953edb4c8852a548ac4234d5e2ef/jsonb_test.go#L99
  -- see pgtypes for pgx specific (no need for extra methods) https://github.com/jackc/pgtype/blob/master/jsonb_array_test.go

  -- TLDR what we want is probably this (much easier, a bit slower): https://github.com/jackc/pgx/issues/809#issuecomment-921135843
  -- what we also want is for xo models to have `json` tags be SQLName exactly as db, so that we can unmarshal directly.
  -- (low prio if all works fine) Then we can easily generate the same models in every xo.go file but all structs get `Response` appended and json tags
  -- use {{camel GoName}}. Each struct has a single one-way convert method to Response (almost 1-1 except struct name changes, should have small overhead).
  -- also checkout https://github.com/georgysavva/scany for direct scans
  -- https://github.com/georgysavva/scany/issues/16

  -- alternative in case the above marshalling json row results doesn't work (although pgx has full support for jsonb and json
  -- ): https://github.com/jackc/pgx/issues/760
  (case when @joinwork_items = true then joined_work_items.work_items end) as work_items -- if M2M
  , (case when @joinTeams = true then joined_teams.teams end) as teams -- if M2M
  , (case when @joinUserApiKeys = true then row_to_json(user_api_keys.*) end) as user_api_key -- if O2O
  , (case when @joinTimeEntries = true then joined_time_entries.time_entries end) as time_entries -- if O2M
  , users.*
-- if all joinXXX are false all that is done is a scan of users (see explain analyze).
from
  users
------------------------------
left join (
  select
    member as work_items_user_id
    , json_agg(work_items.*) as work_items
  from
    task_member uo
    join work_items using (task_id)
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
            work_items))
      group by
        member) joined_work_items on joined_work_items.work_items_user_id = users.user_id
------------------------------
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
------------------------------
-- this below would be O2M (we return an array agg)
-- same as with work_item comments when selecting work_items
-- since work_item_id is not unique in work_item_comments
-- we assume cardinality:O2M. to distinguish O2M and M2M, cardinality:M2M comment, else O2M is assumed
left join (
  select
  user_id as time_entries_user_id
    , json_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
------------------------------
-- this below would be O2O
left join user_api_keys using (user_id); -- if O2O. This is discovered from FK on user_api_keys.user_id being also a unique constraint
