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

create table organizations (
  organization_id serial not null
  , name text not null
  , description text not null
  , metadata json not null
  , created_at timestamp without time zone default current_timestamp not null
  , updated_at timestamp without time zone default current_timestamp not null
  , primary key (organization_id)
  , unique (name)
);

create table projects (
  project_id serial not null
  , organization_id int not null
  , name text not null
  , description text not null
  , metadata json not null
  , created_at timestamp without time zone default current_timestamp not null
  , updated_at timestamp without time zone default current_timestamp not null
  , primary key (project_id)
  , foreign key (organization_id) references organizations (organization_id) on delete cascade
  , unique (name)
);

-- TODO postgres 15 for nulls not distinct in external_id
create table users (
  user_id uuid default gen_random_uuid () not null
  , username text not null
  , email text not null
  , first_name text
  , last_name text
  , full_name text generated always as (((first_name) || ' ') || (last_name)) stored
  , external_id text
  , role user_role default 'user' not null
  , is_superuser boolean default 'False' not null
  , created_at timestamp without time zone default current_timestamp not null
  , updated_at timestamp without time zone default current_timestamp not null
  , deleted_at timestamp without time zone
  , primary key (user_id)
  , unique (email)
  , unique (username)
);
-- pg13 alt for CONSTRAINT uq_external_id UNIQUE NULLS NOT DISTINCT (external_id)
CREATE UNIQUE INDEX external_id_2col_user_id_idx ON users (user_id, external_id)
WHERE external_id IS NOT NULL;
CREATE UNIQUE INDEX external_id_user_id_idx ON users (user_id)
WHERE external_id IS NULL;

create index users_deleted_at_idx on users (deleted_at);
create index users_created_at_idx on users (created_at);
create index users_updated_at_idx on users (updated_at);

create table user_organization (
  organization_id int not null
  , user_id uuid not null
  , primary key (user_id , organization_id)
  , foreign key (user_id) references users (user_id) on delete cascade
  , foreign key (organization_id) references organizations (organization_id) on delete cascade
);
create index user_organization_organization_id_user_id_idx on user_organization (organization_id, user_id);

create table work_items (
  work_item_id bigserial not null
  , title text not null
  , metadata json not null
  , created_at timestamp without time zone default current_timestamp not null
  , updated_at timestamp without time zone default current_timestamp not null
  , deleted_at timestamp without time zone
  , primary key (work_item_id)
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


create table tasks (
  task_id bigserial not null
  , title text not null
  , metadata json not null
  , created_at timestamp without time zone default current_timestamp not null
  , updated_at timestamp without time zone default current_timestamp not null
  , deleted_at timestamp without time zone
  , primary key (task_id)
);

create table task_member (
  task_id bigint not null
  , member uuid not null
  , primary key (task_id, member)
  , foreign key (member) references users (user_id) on delete cascade
);

create table work_item_task (
  task_id bigint not null
  , work_item_id bigint not null
  , primary key (work_item_id , task_id)
  , foreign key (work_item_id) references work_items (work_item_id) on delete cascade
  , foreign key (task_id) references tasks (task_id) on delete cascade
);
create index work_item_task_task_id_work_item_id_idx on work_item_task (task_id, work_item_id);

create table kanban_steps (
  kanban_step_id int not null
  , project_id int not null
  , name text not null
  , description text not null
  , time_trackable bool not null default false
  , disabled bool not null default false
  , primary key (kanban_step_id)
  , foreign key (project_id) references projects (project_id) on delete cascade
);

create table project_kanban_steps (
  kanban_step_id int not null
  , project_id int not null
  , step_order smallint not null
  , primary key (kanban_step_id , project_id)
  , foreign key (kanban_step_id) references kanban_steps (kanban_step_id) on delete cascade
  , foreign key (project_id) references projects (project_id) on delete cascade
);
create index kanban_step_project_project_id_kanban_step_id_idx on project_kanban_steps (project_id, kanban_step_id);

/*
get org names for a given user_id, etc.
with xo we would have to make a ton of different queries to get the same result.
alternative: tell xo when to inner join using (<fk>)
user.xo.go could have a selectUserWith* query for each fk we tell it to join:
e.g. selectUserWithOrganizations, which would join everything.
we would specify an option in generation: public.users<-user_organization:name
to indicate we want to use the lookup table to get an array aggregate of organization names
per user.
we could have more than one of these:
- public.users<-user_organization:name,
- public.users<-user_organization:name
on the other hand we would have:
public.organizations<-user_organization:email would give us an array of user emails per organization
UPDATE:
or just join tables with json_agg: also supported in sqlc https://github.com/kyleconroy/sqlc/issues/1894
that will generate the struct with a nested json object that is simply the same struct from another file,
with json tags already solved.
UPDATE 2: we will inner join and select every subfield `as <prefix>_...` then scan to nested struct
Organizations Organizations `json:organizations,...`
UPDATE 3: sqlc - we get exactly the fields we want -> struct Get...Row with json tags
and our openapi spec has x-db-model: db.Get...Row so we create the schema properties automatically in the spec
and a type ***Res = db.Get...Row instead of oapi-codegen generated struct (either hack into oapi or remove with sed)

~~However organizations~~
~~could have more fks that need to be joined. If we already told xo it should join those fk~~
~~(selectOrganizationWith<fk1>, ...) it should use that same query when we selectUserWithOrganizations.~~
~~we can select fields with a prefix to avoid clashes:~~
~~select organizations.name as organizations_name~~
 */
create index user_organization_user_idx on user_organization (user_id);

create table movies (
  movie_id serial not null
  , title text not null
  , year integer not null
  , synopsis text not null
  , primary key (movie_id)
);

create table api_keys (
  api_key_id serial not null
  , api_key text not null
  , user_id uuid not null
  , expires_on timestamp without time zone not null
  , primary key (api_key_id)
  , unique (api_key)
  , foreign key (user_id) references users (user_id) on delete cascade
);

create or replace view v.users as
select
  *
from
  users
  join (
    select
      user_id
      , ARRAY_AGG(o.name)::text[] as organizations
    from
      user_organization uo
      join organizations o using (organization_id)
    where
      user_id in (
        select
          user_id
        from
          user_organization
        where
          organization_id = any (
            select
              organization_id
            from
              organizations))
        group by
          user_id) joined_organizations using (user_id);

create materialized view if not exists cache.users as
select
  *
from
  v.users with no data;

insert into users (user_id , username , email , first_name , last_name ,
  "role" , is_superuser)
  values ('99270107-1b9c-4f52-a578-7390d5b31513' , 'user 1' , 'user1@email.com' , 'John' ,
    'Doe' , 'user'::user_role , false);

insert into users (user_id , username , email , first_name , last_name ,
  "role" , is_superuser)
  values ('59270107-1b9c-4f52-a578-7390d5b31513' , 'user 2' , 'user2@email.com' , 'Jane' ,
    'Doe' , 'user'::user_role , false);

insert into organizations ("name" , description, metadata , created_at , updated_at)
  values ('org 1' , 'this is org 1', '{}' , current_timestamp , current_timestamp);

insert into organizations ("name" , description, metadata , created_at , updated_at)
  values ('org 2' ,  'this is org 2','{}' , current_timestamp , current_timestamp);

insert into user_organization (organization_id , user_id)
  values (1 , '99270107-1b9c-4f52-a578-7390d5b31513');

insert into user_organization (organization_id , user_id)
  values (2 , '99270107-1b9c-4f52-a578-7390d5b31513');
