-- https://dba.stackexchange.com/questions/59006/what-is-a-valid-use-case-for-using-timestamp-without-time-zone
create schema if not exists v;

create schema if not exists
cache;

create table projects (
  project_id serial not null
  , name text not null unique
  , description text not null
  , metadata jsonb not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , primary key (project_id)
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

create table user_api_keys (
  user_api_key_id serial not null
  , api_key text not null unique
  , expires_on timestamp with time zone not null
  -- don't use,  see https://github.com/jackc/pgx/issues/924
  --, expires_on timestamp without time zone not null
  , primary key (user_api_key_id)
);

create table users (
  user_id uuid default gen_random_uuid () not null
  , username text not null unique
  , email text not null unique
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
  , api_key_id int
  , scopes text[] default '{}' not null
  , role_rank smallint default 1 not null check (role_rank > 0)
  -- so that later on we can (1) append scopes and remove duplicates:
  --    update users
  --    set    scopes = (select array_agg(distinct e) from unnest(scopes || '{"newscope-1","newscope-2"}') e)
  --    where  not scopes @> '{"newscope-1","newscope-2"}' and role_rank >= @minimum_role_rank
  -- and also (2) easily update if we add a new role:
  --    update ... set rank = rank +1 where rank >= @new_role_rank
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , deleted_at timestamp with time zone
  , primary key (user_id)
  , foreign key (api_key_id) references user_api_keys (user_api_key_id) on delete cascade
);

alter table user_api_keys
  add column user_id uuid not null unique;

alter table user_api_keys
  add foreign key (user_id) references users (user_id) on delete cascade;

comment on column users.api_key_id is 'cardinality:O2O';

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
  kanban_step_id serial not null
  , team_id int not null
  , step_order smallint
  , name text not null
  , description text not null
  , color text not null
  , time_trackable bool not null default false
  , disabled bool not null default false
  , primary key (kanban_step_id)
  , unique (team_id , step_order)
  , foreign key (team_id) references teams (team_id) on delete cascade
);

create table work_items (
  work_item_id bigserial not null
  , title text not null
  , metadata jsonb not null
  , team_id int not null
  , kanban_step_id int not null
  , closed bool default false not null
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

-- we want to generate the following:
-- work_items.xo.go --> joinComments (pluralize) to return all comments associated to a workitem
-- already done by xo: WorkItemCommentsByWorkItemID (returns []*WorkItemComment)
-- work_item_comments.xo.go --> joinWorkItem (singularize) to return the workitem object every time is useless in this case.
-- work_item_comments acts as a link table
create index on work_item_comments (work_item_id);

create table work_item_tags (
  work_item_tag_id serial not null
  , name text not null unique
  , description text not null
  , color text not null
  , primary key (work_item_tag_id)
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
create type task_role as ENUM (
  'preparer'
  , 'reviewer'
);

-- only need different types for tasks. work items are all the same, just containers
create table task_types (
  task_type_id serial
  , team_id bigint not null
  , name text not null
  , description text not null
  , color text not null
  , primary key (task_type_id)
  , unique (team_id , name)
  , foreign key (team_id) references teams (team_id) on delete cascade
);

-- NOTE: will not use. Everything hardcoded.
-- customize keys per project only.
-- these keys will be used to dynamically show data in ui, regardless of current project or team
-- create table work_item_fields (
--   project_id bigint not null
--   , key text not null -- for work_items.metadata->"key" filtering (and we can dynamically create indeces on work_items.metadata when a new key is added)
--   , primary key (project_id , key)
--   , foreign key (project_id) references projects (project_id) on delete cascade
-- );
create table tasks (
  task_id bigserial not null
  , task_type_id int not null
  , work_item_id bigint not null
  , title text not null
  , metadata jsonb not null
  , target_date timestamp with time zone not null
  -- don't use,  see https://github.com/jackc/pgx/issues/924
  --, target_date timestamp without time zone not null
  , target_date_timezone text not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , deleted_at timestamp with time zone
  , primary key (task_id)
  , foreign key (task_type_id) references task_types (task_type_id) on delete cascade
  , foreign key (work_item_id) references work_items (work_item_id) on delete cascade -- not unique, many tasks for the same work_item_id
);

comment on column tasks.work_item_id is 'cardinality:O2M';

comment on column tasks.task_type_id is 'cardinality:O2O';

-- we're doing it the wrong way. O2O
create table work_item_member (
  work_item_id bigint not null
  , member uuid not null
  , primary key (work_item_id , member)
  , foreign key (work_item_id) references work_items (work_item_id) on delete cascade
  , foreign key (member) references users (user_id) on delete cascade
);

create index on work_item_member (member , work_item_id);

comment on column work_item_member.work_item_id is 'cardinality:M2M';

comment on column work_item_member.member is 'cardinality:M2M';

-- must be completely dynamic on a team basis
create table activities (
  activity_id serial not null
  , name text not null unique
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
  , team_id int
  , user_id uuid not null
  , comment text not null
  , start timestamp with time zone default current_timestamp not null
  , duration_minutes int -- NULL -> active
  , primary key (time_entry_id)
  , foreign key (user_id) references users (user_id) on delete cascade
  , foreign key (task_id) references tasks (task_id) on delete cascade
  , foreign key (activity_id) references activities (activity_id) on delete cascade -- need to know where we're allocating time
  , foreign key (team_id) references teams (team_id) on delete cascade -- need to know where we're allocating time
  , check (num_nonnulls (team_id , task_id) = 1) -- team_id null when a task id is associated and viceversa
);

-- TODO revisit all comments and fix.
-- We need cardinality comments only on FK columns, never base tables.
-- cardinality will be different for every table making use of a pk.
-- think of user api. For joins: we want to provide possibilities to:
-- user.xo.go:
-- select for time_entries that joins with:
-- a task can be associated to many time entries, and any time entry is linked back to only one task: O2M
-- another way to see it: one task shares many time_entries, and time_entries are part of only one task.
comment on column time_entries.task_id is 'cardinality:O2M';

-- a team can be associated to many time entries, and any time entry is linked back to only one team: O2M
comment on column time_entries.team_id is 'cardinality:O2M';

-- an activity can be associated to many time entries, and any time entry is linked back to only one activity: O2M
comment on column time_entries.activity_id is 'cardinality:O2M';

-- a user can be associated to many time entries, and any time entry is linked back to only one user: O2M
comment on column time_entries.user_id is 'cardinality:O2M';

-- A multicolumn B-tree index can be used with query conditions that involve any subset of the index's
-- columns, but the index is most efficient when there are constraints on the leading (leftmost) columns.
create index on time_entries (user_id , team_id);

-- show user his timelog based on what projects are selected
create index on time_entries (task_id , team_id);


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
