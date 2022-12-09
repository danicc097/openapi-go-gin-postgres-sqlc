-- https://dba.stackexchange.com/questions/59006/what-is-a-valid-use-case-for-using-timestamp-without-time-zone
create schema if not exists v;

create schema if not exists "cache";

create extension if not exists pg_stat_statements schema extensions;

create extension if not exists pg_trgm schema extensions;

create extension if not exists btree_gin schema extensions;

-- internal use. update whenever a project with its related workitems,
--  etc. tables are created in migrations
create table projects (
  project_id serial primary key
  , name text not null unique
  , description text not null
  , work_items_table_name text not null
  , initialized boolean not null default false
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
);

comment on column projects.work_items_table_name is 'property:private';

insert into projects (
  name
  , description
  , work_items_table_name
  , initialized)
values (
  'demo project'
  , 'description for demo project'
  , 'work_items_demo_project'
  , true -- just for demo since it will be programmatically initialized
);

create table teams (
  team_id serial primary key
  , project_id int not null --limited to a project only
  , name text not null
  , description text not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , foreign key (project_id) references projects (project_id) on delete cascade
  , unique (name , project_id)
);

comment on column teams.project_id is 'cardinality:O2M';

create table user_api_keys (
  user_api_key_id serial primary key
  , api_key text not null unique
  , expires_on timestamp with time zone not null
  -- don't use,  see https://github.com/jackc/pgx/issues/924
  --, expires_on timestamp without time zone not null
);

create table users (
  user_id uuid default gen_random_uuid () primary key
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
  , has_personal_notifications boolean default false not null
  , has_global_notifications boolean default false not null
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
-- composite on id, deleted_at, email, deleted_at, etc. will not improve speed
-- create unique index on users (user_id) where deleted_at is null; -- helps if you have much more deleted rows only
create index on users (created_at);

-- create index on users (deleted_at);  - not worth the extra overhead.
-- for finding all deleted users exclusively
create index on users (deleted_at)
where (deleted_at is not null);

create index on users (updated_at);

-- notification_types are append-only
create type notification_type as ENUM (
  'personal'
  , 'global'
);

create table notifications (
  notification_id serial primary key
  , receiver_rank smallint check (receiver_rank > 0) --check will not prevent null values
  , title text not null
  , body text not null
  , label text not null
  , link text null
  , created_at timestamp with time zone default current_timestamp not null
  , sender uuid not null
  , receiver uuid -- can be null for 'global' type
  , notification_type notification_type not null
  , foreign key (sender) references users (user_id) on delete cascade
  , foreign key (receiver) references users (user_id) on delete cascade
  , check (num_nonnulls (receiver_rank , receiver) = 1)
);

create index on notifications (receiver_rank , notification_type , created_at);

create table user_notifications (
  user_notification_id bigserial primary key
  , notification_id int not null
  , read boolean default false not null -- frontend simply sends a list of user_notification_id to mark as read
  , created_at timestamp with time zone default current_timestamp not null
  , user_id uuid not null
  , foreign key (user_id) references users (user_id) on delete cascade
  , foreign key (notification_id) references notifications (notification_id) on delete cascade
);

create index on user_notifications (user_id);

-- read field simply used to show 'NEW' label but there is no filtering
create or replace function notification_fan_out ()
  returns trigger
  language plpgsql
  as $function$
declare
  receiver_id uuid;
begin
  case when new.notification_type = 'personal' then
    update
      users
    set
      has_personal_notifications = true
    where
      user_id = new.receiver;
      --
      insert into user_notifications (
        notification_id
        , user_id)
      values (
        new.notification_id
        , new.receiver);
  when new.notification_type = 'global' then
    update
      users
    set
      has_global_notifications = true
    where
      role_rank >= new.receiver_rank;
      --
      for receiver_id in (
        select
          user_id
        from
          users
        where
          role_rank >= new.receiver_rank)
        loop
          insert into user_notifications (
            notification_id
            , user_id)
          values (
            new.notification_id
            , receiver_id);
          end loop;
  end case;
  -- it's after trigger so wouldn't mattern anyway
  return null;
end
$function$;

-- deletes get cascaded
create trigger notifications_fan_out
  after insert on notifications for each row
  execute function notification_fan_out ();

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
  kanban_step_id serial primary key
  , project_id int not null
  , step_order smallint -- null -> disabled
  , name text not null
  , description text not null
  , color text not null
  , time_trackable bool not null default false
  -- , disabled bool not null default false
  , unique (project_id , step_order)
  , foreign key (project_id) references projects (project_id) on delete cascade
  , check (color ~* '^#[a-f0-9]{6}$')
  , check (step_order > 0)
);

create unique index on kanban_steps (project_id , name , step_order)
where
  step_order is not null;

create unique index on kanban_steps (project_id , name)
where
  step_order is null;

comment on column kanban_steps.project_id is 'cardinality:O2M';

-- types restricted per project
create table work_item_types (
  work_item_type_id serial primary key
  , project_id int not null
  , name text not null
  , description text not null
  , color text not null
  , unique (name , project_id)
  , foreign key (project_id) references projects (project_id) on delete cascade
  , check (color ~* '^#[a-f0-9]{6}$')
);

comment on column work_item_types.project_id is 'cardinality:O2M';


/*
keep track of per-project overrides in shared json, indexed by project name (unique).
Can be directly used in backend (codegen alternative struct) and frontend
internally the storage is the same and doesn't affect in any way.
 */
create table work_items (
  work_item_id bigserial primary key
  , title text not null
  , description text not null
  , work_item_type_id int not null
  , metadata jsonb not null
  , team_id int not null
  , kanban_step_id int not null
  , closed timestamp with time zone -- NULL: active
  , target_date timestamp with time zone not null
  /*
   -- IMPORTANT: implement this:
   alternative to sharing all keys for different projects in the same table:
   https://stackoverflow.com/questions/10068033/postgresql-foreign-key-referencing-primary-keys-of-two-different-tables
   for every custom project we reference the common columns with its own PK being a FK
   We would then query this custom table explicitly and join with the common columns via PI which is really fast
   we would need to join every single record but models are much much cleaner.
   every default workitem still keeps a team_id reference.
   tables with extra fields are for a given project. so we could have another team_id column
   in this new table with check (team_id in select team_id ... join teams where project_id ... )
   the same concept is also seen in https://dba.stackexchange.com/questions/232262/solving-supertype-subtype-relationship-without-sacrificing-data-consistency-in-a
   */
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , deleted_at timestamp with time zone
  , foreign key (team_id) references teams (team_id) on delete cascade
  , foreign key (work_item_type_id) references work_item_types (work_item_type_id) on delete cascade
  , foreign key (kanban_step_id) references kanban_steps (kanban_step_id) on delete cascade
);

-- to get join directly instead of having to call xo's generated FK for every single one
comment on column work_items.work_item_type_id is 'cardinality:O2O';


/*
when a new project is required -> manual table creation with empty new fields, just
 work_item_id bigint primary key.
 When a new field is added, possibilities are:
 - not nullable -> must set default value for the existing rows
 - nullable and custom business logic when it's required or not. previous rows remain null or with default as required
 */
-- project for tour. when starting it user joins the only demo project team. when exiting it user is removed.
-- we can reset it every X hours
create table work_items_demo_project (
  work_item_id bigint primary key references work_items (work_item_id) on delete cascade
  , ref text not null
  , line text not null
  , last_message_at timestamp with time zone not null
  , reopened boolean not null default false
);

create index on work_items_demo_project (ref , line);

create table work_items_project_2 (
  work_item_id bigint primary key references work_items (work_item_id) on delete cascade
  , custom_date_for_project_2 timestamp with time zone
);

comment on column work_items.work_item_id is 'cardinality:O2O';

comment on column work_items_demo_project.work_item_id is 'cardinality:O2O';

comment on column work_items_project_2.work_item_id is 'cardinality:O2O';

-- for finding all deleted work items exclusively
create index on work_items (deleted_at)
where (deleted_at is not null);

create table work_item_comments (
  work_item_comment_id bigserial primary key
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
  work_item_tag_id serial primary key
  , project_id int not null
  , name text not null
  , description text not null
  , color text not null
  , unique (name , project_id)
  , check (color ~* '^#[a-f0-9]{6}$')
  , foreign key (project_id) references projects (project_id) on delete cascade
);

comment on column work_item_tags.project_id is 'cardinality:O2M';

create table work_item_work_item_tag (
  work_item_tag_id int not null
  , work_item_id bigint not null
  , primary key (work_item_id , work_item_tag_id) -- M2M, work_item can have multple tags. tags can be in multiple work_items (same as book authors example)
  , foreign key (work_item_id) references work_items (work_item_id) on delete cascade
  , foreign key (work_item_tag_id) references work_item_tags (work_item_tag_id) on delete cascade
);

create index on work_item_work_item_tag (work_item_tag_id , work_item_id);

comment on column work_item_work_item_tag.work_item_tag_id is 'cardinality:M2M';

comment on column work_item_work_item_tag.work_item_id is 'cardinality:M2M';

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

-- must be completely dynamic on a project basis
create table activities (
  activity_id serial primary key
  , project_id int not null
  , name text not null
  , description text not null
  , is_productive boolean default false not null
  , unique (name , project_id)
  -- can't have multiple unrelated projects see each other's activities
  , foreign key (project_id) references projects (project_id) on delete cascade
);

comment on column activities.project_id is 'cardinality:O2M';

-- will restrict available activities on a per-project basis
-- where project_id is null (shared) or project_id = @project_id
-- table will be tiny, don't even index
-- create index on activities (project_id);
create table time_entries (
  time_entry_id bigserial primary key
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

-- grpc demo
create table movies (
  movie_id serial primary key
  , title text not null
  , year integer not null
  , synopsis text not null
);

--
-- audit
--
select
  audit.enable_tracking ('public.kanban_steps');

select
  audit.enable_tracking ('public.projects');

select
  audit.enable_tracking ('public.teams');

select
  audit.enable_tracking ('public.work_items');
