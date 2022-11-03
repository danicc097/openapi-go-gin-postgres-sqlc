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
  , metadata json not null
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
  , metadata json not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , primary key (team_id)
  , foreign key (project_id) references projects (project_id) on delete cascade
  , unique (name)
);

-- TODO postgres 15 for nulls not distinct in external_id
create table users (
  user_id uuid default gen_random_uuid () not null
  , username text not null
  , email text not null
  , scopes text[] default '{}' not null -- defined in spec only
  , first_name text
  , last_name text
  , full_name text generated always as (((first_name) || ' ') || (last_name)) stored
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
CREATE UNIQUE INDEX ON users (user_id, external_id)
WHERE external_id IS NOT NULL;
CREATE UNIQUE INDEX ON users (user_id)
WHERE external_id IS NULL;

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
create index on user_team (team_id, user_id);
comment on column user_team.user_id is 'cardinality:M2M';

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
  , metadata json not null
  , team_id int not null
  , kanban_step_id int not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , deleted_at timestamp with time zone
  , primary key (work_item_id)
  , foreign key (team_id) references teams (team_id) on delete cascade
  , foreign key (kanban_step_id) references kanban_steps (kanban_step_id) on delete cascade
);

create type work_item_role as ENUM (
  'preparer'
  , 'reviewer'
);

-- we can aggregate members from tasks
-- but need a role for the work_item and every member
-- or can we ignore members per work_item?
-- need timelog for every task but thats saved elsewhere
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
  , unique (team_id, name)
  , foreign key (team_id) references teams (team_id) on delete cascade
);

-- customize keys per project only.
-- these keys will be used to dynamically show data in ui, regardless of current project or team
create table work_item_fields (
  project_id bigint not null
  , key text not null -- for work_items.metadata->"key" filtering (and we can dynamically create indeces on work_items.metadata when a new key is added)
  , primary key (project_id, key)
  , foreign key (project_id) references projects (project_id) on delete cascade
);

create table tasks (
  task_id bigserial not null
  , task_type_id int not null
  , title text not null
  , metadata json not null
  , target_date timestamp without time zone not null
  , target_date_timezone text not null
  , created_at timestamp with time zone default current_timestamp not null
  , updated_at timestamp with time zone default current_timestamp not null
  , deleted_at timestamp with time zone
  , primary key (task_id)
  , foreign key (task_type_id) references task_types (task_type_id) on delete cascade
);

create table task_member (
  task_id bigint not null
  , member uuid not null
  , primary key (task_id, member)
  , foreign key (task_id) references tasks (task_id) on delete cascade
  , foreign key (member) references users (user_id) on delete cascade
);
create index on task_member (member, task_id);

create table work_item_task (
  task_id bigint not null
  , work_item_id bigint not null
  , primary key (work_item_id , task_id)
  , foreign key (work_item_id) references work_items (work_item_id) on delete cascade
  , foreign key (task_id) references tasks (task_id) on delete cascade
);
create index on work_item_task (task_id, work_item_id);


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

create table user_api_key (
  api_key text not null
  , user_id uuid not null
  , expires_on timestamp without time zone not null
  , primary key (api_key) -- read bearer -> hash -> GetAPIKeyByAPIKey -> exists? -> GetUserByAPIKey
  , unique (user_id) -- already know it's O2O
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

insert into users (user_id , username , email , first_name , last_name ,
  "role" )
  values ('99270107-1b9c-4f52-a578-7390d5b31513' , 'user 1' , 'user1@email.com' , 'John' ,
    'Doe' , 'user'::user_role);

insert into users (user_id , username , email , first_name , last_name ,
  "role" )
  values ('59270107-1b9c-4f52-a578-7390d5b31513' , 'user 2' , 'user2@email.com' , 'Jane' ,
    'Doe' , 'user'::user_role);

insert into projects ("name" , description, metadata , created_at , updated_at)
  values ('org 1' , 'this is org 1', '{}' , current_timestamp , current_timestamp);

insert into projects ("name" , description, metadata , created_at , updated_at)
  values ('org 2' ,  'this is org 2','{}' , current_timestamp , current_timestamp);

insert into teams ("name" , project_id, description, metadata , created_at , updated_at)
  values ('team 1', 1 , 'this is team 1', '{}' , current_timestamp , current_timestamp);

insert into teams ("name" , project_id, description, metadata , created_at , updated_at)
  values ('team 2', 1 ,  'this is team 2','{}' , current_timestamp , current_timestamp);

insert into user_team (team_id , user_id)
  values (1 , '99270107-1b9c-4f52-a578-7390d5b31513');
insert into user_team (team_id , user_id)
  values (1 , '59270107-1b9c-4f52-a578-7390d5b31513');

insert into user_team (team_id , user_id)
  values (2 , '99270107-1b9c-4f52-a578-7390d5b31513');
