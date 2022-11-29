-- https://dba.stackexchange.com/questions/59006/what-is-a-valid-use-case-for-using-timestamp-without-time-zone
create schema if not exists v;

create schema if not exists
cache;

create table projects (
  project_id serial not null primary key
  , name text not null unique
  , description text not null
  , metadata jsonb not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
);

create table teams (
  team_id serial not null primary key
  , project_id int not null --limited to a project only
  , name text not null
  , description text not null
  , metadata jsonb not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , foreign key (project_id) references projects (project_id) on delete cascade
  , unique (name , project_id)
);

create table user_api_keys (
  user_api_key_id serial not null primary key
  , api_key text not null unique
  , expires_on timestamp with time zone not null
  -- don't use,  see https://github.com/jackc/pgx/issues/924
  --, expires_on timestamp without time zone not null
);

create table users (
  user_id uuid default gen_random_uuid () not null primary key
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
  , external_id text not null unique
  , api_key_id int
  -- so that later on we can (1) append scopes and remove duplicates:
  --    update users
  --    set    scopes = (select array_agg(distinct e) from unnest(scopes || '{"newscope-1","newscope-2"}') e)
  --    where  not scopes @> '{"newscope-1","newscope-2"}' and role_rank >= @minimum_role_rank
  -- and also (2) easily update if we add a new role:
  --    update ... set rank = rank +1 where rank >= @new_role_rank
  , scopes text[] default '{}' not null
  , role_rank smallint default 1 not null check (role_rank > 0)
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , deleted_at timestamp with time zone
  , foreign key (api_key_id) references user_api_keys (user_api_key_id) on delete cascade
);

comment on column users.external_id is 'property:private, property:something-else';

comment on column users.api_key_id is 'property:private';

comment on column users.role_rank is 'property:private';

comment on column users.scopes is 'property:private';

comment on column users.updated_at is 'property:private';

alter table user_api_keys
  add column user_id uuid not null unique;

alter table user_api_keys
  add foreign key (user_id) references users (user_id) on delete cascade;

-- circular schema ref. (users' join is ultimately useless, we will start from an apikey
-- and go from there to the user that owns it)
--  generates join in users table
-- comment on column users.api_key_id IS 'cardinality:O2O';
comment on column user_api_keys.user_id is 'cardinality:O2O';

comment on column user_api_keys.user_api_key_id is 'property:private';

-- -- pg13 alt for CONSTRAINT uq_external_id UNIQUE NULLS NOT DISTINCT (external_id)
-- create unique index on users (user_id , external_id)
-- where
--   external_id is not null;
-- create unique index on users (user_id)
-- where
--   external_id is null;
create index on users (deleted_at);

create index on users (created_at);

create index on users (updated_at);

create table user_team (
  team_id int not null
  , user_id uuid not null
  , primary key (user_id , team_id)
  , foreign key (user_id) references users (user_id) on delete cascade
  , foreign key (team_id) references teams (team_id) on delete cascade
);

create index on user_team (team_id , user_id);

comment on column user_team.user_id is 'cardinality:M2M';

comment on column user_team.team_id is 'cardinality:M2M';

create table kanban_steps (
  kanban_step_id serial not null primary key
  , team_id int not null
  , step_order smallint
  , name text not null
  , description text not null
  , color text not null
  , time_trackable bool not null default false
  , disabled bool not null default false
  , unique (team_id , step_order)
  , foreign key (team_id) references teams (team_id) on delete cascade
  , check (color ~* '^#[a-f0-9]{6}$')
  , check (step_order > 0)
);

create table work_item_types (
  work_item_type_id serial primary key
  , project_id bigint not null
  , name text not null
  , description text not null
  , color text not null
  , unique (project_id , name)
  , foreign key (project_id) references projects (project_id) on delete cascade
  , check (color ~* '^#[a-f0-9]{6}$')
);


/*
keep track of per-project overrides in shared json, indexed by project name (unique).
Can be directly used in backend (codegen alternative struct) and frontend
internally the storage is the same and doesn't affect in any way.
 */
create table work_items (
  work_item_id bigserial not null primary key
  /* generic must-have fields. store naming overrides in business logic, if any
  (json with project name (unique) as key should suffice to be used by both back and frontend)
  as requested by clients to prevent useless joins.
  yq will ensure keys do exist as db column */
  , title text not null
  , work_item_type_id int not null
  , metadata jsonb not null
  , team_id int not null
  , kanban_step_id int not null
  , closed timestamp with time zone -- NULL: active
  , target_date timestamp with time zone not null
  /* if a project requests a new field that needs to be indexed (either manual or automated)
  add it as nullable. in business logic that project_id will have this field marked as required .
  If indexability is not required, dump it to metadata
  it can use the same json as above.
  since its not indexed (maybe just GIN) we dont care about schema changes over time
  (no keys before existence) */
  , some_custom_date timestamp with time zone
  --
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , deleted_at timestamp with time zone
  , foreign key (team_id) references teams (team_id) on delete cascade
  , foreign key (work_item_type_id) references work_item_types (work_item_type_id) on delete cascade
  , foreign key (kanban_step_id) references kanban_steps (kanban_step_id) on delete cascade
);

create table work_item_comments (
  work_item_comment_id bigserial not null primary key
  , work_item_id bigint not null
  , user_id uuid not null
  , message text not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , foreign key (user_id) references users (user_id) on delete cascade
  , foreign key (work_item_id) references work_items (work_item_id) on delete cascade
);

comment on column work_item_comments.work_item_id is 'cardinality:O2M';

create index on work_item_comments (work_item_id);

create table work_item_tags (
  work_item_tag_id serial not null primary key
  , name text not null unique
  , description text not null
  , color text not null
  , check (color ~* '^#[a-f0-9]{6}$')
);

create table work_item_work_item_tag (
  work_item_tag_id int not null
  , work_item_id bigint not null
  , primary key (work_item_id , work_item_tag_id) -- M2M, work_item can have multple tags. tags can be in multiple work_items (same as book authors example)
  , foreign key (work_item_id) references work_items (work_item_id) on delete cascade
  , foreign key (work_item_tag_id) references work_item_tags (work_item_tag_id) on delete cascade
);

create index on work_item_work_item_tag (work_item_tag_id , work_item_id);

-- roles are append-only
create type work_item_role as ENUM (
  'preparer'
  , 'reviewer'
);

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
  activity_id serial not null primary key
  , project_id int
  , name text not null unique
  , description text not null
  , is_productive boolean not null
  -- can't have multiple unrelated projects see each other's activities
  , foreign key (project_id) references projects (project_id) on delete cascade
);

-- will restrict available activities on a per-project basis
-- where project_id is null (shared) and project_id = @project_id
create index on activities (project_id);

create table time_entries (
  time_entry_id bigserial not null primary key
  , work_item_id bigint
  , activity_id int not null
  , team_id int
  , user_id uuid not null
  , comment text not null
  , start timestamp with time zone default current_timestamp not null
  , duration_minutes int -- NULL -> active
  , foreign key (user_id) references users (user_id) on delete cascade
  , foreign key (work_item_id) references work_items (work_item_id) on delete cascade
  , foreign key (activity_id) references activities (activity_id) on delete cascade -- need to know where we're allocating time
  , foreign key (team_id) references teams (team_id) on delete cascade -- need to know where we're allocating time
  , check (num_nonnulls (team_id , work_item_id) = 1) -- team_id null when a work_item id is associated and viceversa
);

comment on column time_entries.work_item_id is 'cardinality:O2M';

comment on column time_entries.team_id is 'cardinality:O2M';

comment on column time_entries.activity_id is 'cardinality:O2M';

comment on column time_entries.user_id is 'cardinality:O2M';

-- A multicolumn B-tree index can be used with query conditions that involve any subset of the index's
-- columns, but the index is most efficient when there are constraints on the leading (leftmost) columns.
create index on time_entries (user_id , team_id);

-- show user his timelog based on what projects are selected
create index on time_entries (work_item_id , team_id);

create index on user_team (user_id);

create table movies (
  movie_id serial not null primary key
  , title text not null
  , year integer not null
  , synopsis text not null
);
