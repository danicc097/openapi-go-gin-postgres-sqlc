-- https://dba.stackexchange.com/questions/59006/what-is-a-valid-use-case-for-using-timestamp-without-time-zone
create schema if not exists v;

create schema if not exists
cache;

create type user_role as ENUM (
  'guest'
  , 'user'
  , 'advanced user'
  , 'manager'
  , 'admin'
  , 'superadmin'
);

create table projects (
  project_id serial not null
  , name text not null
  , description text not null
  , metadata jsonb not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , primary key (project_id)
  , unique (name)
);

create table teams (
  team_id serial not null
  , project_id int not null
  , name text not null
  , description text not null
  , metadata jsonb not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , primary key (team_id)
  , foreign key (project_id) references projects (project_id) on delete cascade
  , unique (name , project_id)
);

create table users (
  user_id uuid default gen_random_uuid () not null
  , username text not null
  , email text not null
  , scopes text[] default '{}' not null -- defined in spec only
  , first_name text
  , last_name text
  , full_name text generated always as ( case when first_name is null then
    last_name
  when last_name is null then
    first_name
  else
    first_name || ' ' || last_name
  end) stored
  , external_id text
  , role user_role default 'user' not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , deleted_at timestamp with time zone
  , primary key (user_id)
  , unique (email)
  , unique (username)
);

-- pg13 alt for CONSTRAINT uq_external_id UNIQUE NULLS NOT DISTINCT (external_id)
create unique index on users (user_id , external_id)
where
  external_id is not null;

create unique index on users (user_id)
where
  external_id is null;

create index on users (deleted_at);

create index on users (created_at);

create index on users (updated_at);

create table user_team (
  team_id int not null
  , user_id uuid not null
  , primary key (user_id , team_id) -- M2M, user can be in multple teams. teams can have multiple users (same as book authors example)
  , foreign key (user_id) references users (user_id) on delete cascade
  , foreign key (team_id) references teams (team_id) on delete cascade
);

create index on user_team (team_id , user_id);

comment on column user_team.user_id is 'cardinality:M2M';

comment on column user_team.team_id is 'cardinality:M2M';

create table kanban_steps (
  kanban_step_id int not null
  , team_id int not null
  , step_order smallint not null
  , name text not null
  , description text not null
  , time_trackable bool not null default false
  , disabled bool not null default false
  , primary key (kanban_step_id)
  , foreign key (team_id) references teams (team_id) on delete cascade
);

create table work_items (
  work_item_id bigserial not null
  , title text not null
  , metadata jsonb not null
  , team_id int not null
  , kanban_step_id int not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , deleted_at timestamp with time zone
  , primary key (work_item_id)
  , foreign key (team_id) references teams (team_id) on delete cascade
  , foreign key (kanban_step_id) references kanban_steps (kanban_step_id) on delete cascade
);

create table work_item_comments (
  work_item_comment_id bigserial not null
  , work_item_id bigint not null
  , user_id uuid not null
  , message text not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , primary key (work_item_comment_id) -- work_item can have multiple comments. a comment is for a single work_item
  , foreign key (user_id) references users (user_id) on delete cascade
  , foreign key (work_item_id) references work_items (work_item_id) on delete cascade
);

comment on column work_item_comments.work_item_id is 'cardinality:O2M';

-- no unique index on it -> O2M
create index on work_item_comments (work_item_id);

create table work_item_tags (
  work_item_tag_id serial not null
  , name text not null
  , description text not null
  , primary key (work_item_tag_id)
  , unique (name)
);

create table work_item_work_item_tag (
  work_item_tag_id int not null
  , work_item_id bigint not null
  , primary key (work_item_id , work_item_tag_id) -- M2M, work_item can have multple tags. tags can be in multiple work_items (same as book authors example)
  , foreign key (work_item_id) references work_items (work_item_id) on delete cascade
  , foreign key (work_item_tag_id) references work_item_tags (work_item_tag_id) on delete cascade
);

create index on work_item_work_item_tag (work_item_tag_id , work_item_id);

-- dont need to index by user_id, there's no use case to filter by user_id
-- create type work_item_role as ENUM (
--   'preparer'
--   , 'reviewer'
-- );
-- we can aggregate members from tasks
-- but do we need a role for the work_item and every member
-- or can we ignore members per work_item?
-- create table work_item_member (
--   work_item_id bigint not null
--   , member uuid not null
--   , role work_item_role not null
--   , primary key (work_item_id, member)
--   , foreign key (member) references users (user_id) on delete cascade
-- );
-- only need different types for tasks. work items are all the same, just containers
create table task_types (
  task_type_id serial
  , team_id bigint not null
  , name text not null
  , primary key (task_type_id)
  , unique (team_id , name)
  , foreign key (team_id) references teams (team_id) on delete cascade
);

-- customize keys per project only.
-- these keys will be used to dynamically show data in ui, regardless of current project or team
create table work_item_fields (
  project_id bigint not null
  , key text not null -- for work_items.metadata->"key" filtering (and we can dynamically create indeces on work_items.metadata when a new key is added)
  , primary key (project_id , key)
  , foreign key (project_id) references projects (project_id) on delete cascade
);

create table tasks (
  task_id bigserial not null
  , task_type_id int not null
  , work_item_id bigint not null
  , title text not null
  , metadata jsonb not null
  , target_date timestamp without time zone not null
  , target_date_timezone text not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , deleted_at timestamp with time zone
  , primary key (task_id)
  , foreign key (task_type_id) references task_types (task_type_id) on delete cascade
  , foreign key (work_item_id) references work_items (work_item_id) on delete
    cascade -- not unique, many tasks for the same work_item_id
);

create table task_member (
  task_id bigint not null
  , member uuid not null
  , primary key (task_id , member)
  , foreign key (task_id) references tasks (task_id) on delete cascade
  , foreign key (member) references users (user_id) on delete cascade
);

create index on task_member (member , task_id);

comment on column task_member.task_id is 'cardinality:M2M';

comment on column task_member.member is 'cardinality:M2M';

-- must be completely dynamic on a team basis
create table activities (
  activity_id int not null
  , name text not null
  , description text not null
  , is_productive boolean not null
  , primary key (activity_id)
  -- , foreign key (team_id) references teams (team_id) on delete cascade --not needed, shared for all teams and projects
  -- and managed by admin
);

-- no unique indexes at all
create table time_entries (
  time_entry_id bigserial not null
  , task_id bigint
  , activity_id int not null
  , team_id int not null
  , user_id uuid not null
  , message text not null
  , start timestamp with time zone default current_timestamp not null
  , duration_minutes int -- NULL -> active
  , primary key (time_entry_id)
  , foreign key (user_id) references users (user_id) on delete cascade
  , foreign key (task_id) references tasks (task_id) on delete cascade
  , foreign key (activity_id) references activities (activity_id) on delete cascade -- need to know where we're allocating time
  , foreign key (team_id) references teams (team_id) on delete cascade -- need to know where we're allocating time
);

-- TODO revisit all comments and fix.
-- We need cardinality comments only on FK columns, never base tables.
-- cardinality will be different for every table making use of a pk.
-- it would be O2O for these below if we had unique indexes for every single one of them.
-- in that case we would simply apply a join using (fk_name).
-- if not, FK by definition allows duplicates --> O2M -> need 2 nested joins to get an agg
comment on column time_entries.task_id is 'cardinality:O2M';

comment on column time_entries.team_id is 'cardinality:O2M';

comment on column time_entries.activity_id is 'cardinality:O2M';

comment on column time_entries.user_id is 'cardinality:O2M';

-- A multicolumn B-tree index can be used with query conditions that involve any subset of the index's
-- columns, but the index is most efficient when there are constraints on the leading (leftmost) columns.
create index on time_entries (user_id , team_id);

-- show user his timelog based on what projects are selected
create index on time_entries (task_id , team_id);

-- for joins aggregating time spent on any task by any user
-- FIXME this makes no sense. its not M2M its o2m. a given task_id can only be in one work_item
-- create table work_item_task (
--   task_id bigint not null
--   , work_item_id bigint not null
--   , primary key (work_item_id , task_id)
--   , foreign key (work_item_id) references work_items (work_item_id) on delete cascade
--   , foreign key (task_id) references tasks (task_id) on delete cascade
-- );
-- create index on work_item_task (task_id , work_item_id);
/*
get org names for a given user_id, etc.
with xo we would have to make a ton of different queries to get the same result.
alternative: tell xo when to inner join using (<fk>)
user.xo.go could have a selectUserWith* query for each fk we tell it to join:
e.g. selectUserWithprojects, which would join everything.
we would specify an option in generation: public.users<user_team:name
to indicate we want to use the lookup table to get an array aggregate of project names
per user.
we could have more than one of these:
- public.users<user_team:name,
- public.users<user_team:name
on the other hand we would have:
public.projects<user_team:email would give us an array of user emails per project
UPDATE:
or just join tables with json_agg: also supported in sqlc https://github.com/kyleconroy/sqlc/issues/1894
that will generate the struct with a nested json object that is simply the same struct from another file,
with json tags already solved.
UPDATE 2: we will inner join and select every subfield `as <prefix>_...` then scan to nested struct
projects projects `json:projects,...`
UPDATE 3: sqlc - we get exactly the fields we want -> struct Get...Row with json tags
and our openapi spec has x-db-model: db.Get...Row so we create the schema properties automatically in the spec
and a type ***Res = db.Get...Row instead of oapi-codegen generated struct (either hack into oapi or remove with sed)

~~However projects~~
~~could have more fks that need to be joined. If we already told xo it should join those fk~~
~~(selectprojectWith<fk1>, ...) it should use that same query when we selectUserWithprojects.~~
~~we can select fields with a prefix to avoid clashes:~~
~~select projects.name as projects_name~~
 */
create index user_team_user_idx on user_team (user_id);

create table movies (
  movie_id serial not null
  , title text not null
  , year integer not null
  , synopsis text not null
  , primary key (movie_id)
);

-- we can infer O2O for the user_id fk since PK and FK in a given table are the same
-- comment on column user_api_keys.user_id IS 'cardinality:O2O';
create table user_api_keys (
  user_id uuid not null
  , api_key text not null
  , expires_on timestamp without time zone not null
  , primary key (user_id)
  , unique (api_key)
  , foreign key (user_id) references users (user_id) on delete cascade -- generates GetUserByAPIKey
);

create or replace view v.users as
select
  *
from
  users
  left join (
    select
      user_id
      , ARRAY_AGG(teams.*) as teams
    from
      user_team uo
      left join teams using (team_id)
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
          user_id) joined_teams using (user_id);

create materialized view if not exists cache.users as
select
  *
from
  v.users with no data;

create index on cache.users (external_id);
