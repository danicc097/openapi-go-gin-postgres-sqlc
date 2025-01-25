-- https://dba.stackexchange.com/questions/59006/what-is-a-valid-use-case-for-using-timestamp-without-time-zone
create schema if not exists v;

create schema if not exists "cache";

-- alter database postgres_test set timezone to 'America/New_York';
create or replace function jsonb_set_deep (target jsonb , path text[] , val jsonb)
  returns jsonb
  as $$
declare
  k text;
  p text[];
begin
  if (path = '{}') then
    return val;
  else
    if (target is null) then
      target = '{}'::jsonb;
    end if;
    FOREACH k in array path loop
      p := p || k;
      if (target #> p is null) then
        target := JSONB_SET(target , p , '{}'::jsonb);
      else
        target := JSONB_SET(target , p , target #> p);
      end if;
    end loop;
    -- Set the value like normal.
    return JSONB_SET(target , path , val);
  end if;
end;
$$
language plpgsql;

--
--
-- EXTRA SCHEMA showcase
--
--
create schema if not exists extra_schema;

create table extra_schema.dummy_join (
  dummy_join_id serial primary key
  , name text
);

create table extra_schema.user_api_keys (
  user_api_key_id serial primary key
  , api_key text not null unique
  , expires_on timestamp with time zone not null
);

create table extra_schema.users (
  user_id uuid default gen_random_uuid () primary key
  , name text not null unique
  , api_key_id int
  , foreign key (api_key_id) references extra_schema.user_api_keys (user_api_key_id) on delete cascade
  , created_at timestamp with time zone default current_timestamp not null unique
  , deleted_at timestamp with time zone
);

alter table extra_schema.user_api_keys
  add column user_id uuid not null unique;

alter table extra_schema.user_api_keys
  add foreign key (user_id) references extra_schema.users (user_id) on delete cascade;

comment on column extra_schema.user_api_keys.user_api_key_id is '"properties":private';

create table extra_schema.books (
  book_id serial primary key
  , name text not null
);

create table extra_schema.book_authors (
  book_id int not null
  , author_id uuid not null
  , pseudonym text
  , primary key (book_id , author_id)
  , foreign key (author_id) references extra_schema.users (user_id) on delete cascade
  , foreign key (book_id) references extra_schema.books (book_id) on delete cascade
);

comment on column extra_schema.book_authors.author_id is '"cardinality":M2M';

comment on column extra_schema.book_authors.book_id is '"cardinality":M2M';

create table extra_schema.book_authors_surrogate_key (
  book_authors_surrogate_key_id serial primary key
  , book_id int not null
  , author_id uuid not null
  , pseudonym text
  , unique (book_id , author_id)
  , foreign key (author_id) references extra_schema.users (user_id) on delete cascade
  , foreign key (book_id) references extra_schema.books (book_id) on delete cascade
);

comment on column extra_schema.book_authors_surrogate_key.author_id is '"cardinality":M2M';

comment on column extra_schema.book_authors_surrogate_key.book_id is '"cardinality":M2M';

create table extra_schema.book_sellers (
  book_id int not null
  , seller uuid not null
  , primary key (book_id , seller)
  , foreign key (seller) references extra_schema.users (user_id) on delete cascade
  , foreign key (book_id) references extra_schema.books (book_id) on delete cascade
);

create index on extra_schema.book_sellers (book_id , seller);

create index on extra_schema.book_sellers (seller , book_id);

comment on column extra_schema.book_sellers.seller is '"cardinality":M2M';

comment on column extra_schema.book_sellers.book_id is '"cardinality":M2M';

create table extra_schema.book_reviews (
  book_review_id serial primary key
  , book_id int not null
  , reviewer uuid not null
  , unique (reviewer , book_id)
  , foreign key (reviewer) references extra_schema.users (user_id) on delete cascade
  , foreign key (book_id) references extra_schema.books (book_id) on delete cascade
);

comment on column extra_schema.book_reviews.reviewer is '"cardinality":M2O';

comment on column extra_schema.book_reviews.book_id is '"cardinality":M2O';

create type extra_schema.notification_type as ENUM (
  'personal'
  , 'global'
);

create table extra_schema.notifications (
  notification_id serial primary key
  , body text not null
  , sender uuid not null
  , receiver uuid
  , notification_type extra_schema.notification_type not null
  , foreign key (sender) references extra_schema.users (user_id) on delete cascade
  , foreign key (receiver) references extra_schema.users (user_id) on delete cascade
);

comment on column extra_schema.notifications.body is '"tags":pattern:"^[A-Za-z0-9]*$" && "properties":private';

create index on extra_schema.notifications (sender);

comment on column extra_schema.notifications.sender is '"cardinality":M2O';

comment on column extra_schema.notifications.receiver is '"cardinality":M2O';

create table extra_schema.work_items (
  work_item_id bigserial primary key
  , title text
  , description text
);

create index on extra_schema.work_items using gin (title extensions.gin_trgm_ops);

create index on extra_schema.work_items using gin (description extensions.gin_trgm_ops);

create index on extra_schema.work_items using gin (title extensions.gin_trgm_ops , description extensions.gin_trgm_ops);

create index on extra_schema.work_items using gin (title , description extensions.gin_trgm_ops);

create table extra_schema.work_item_admin (
  work_item_id bigint not null
  , admin uuid not null
  , primary key (work_item_id , admin)
  , foreign key (work_item_id) references extra_schema.work_items (work_item_id) on delete cascade
  , foreign key (admin) references extra_schema.users (user_id) on delete cascade
);

create index on extra_schema.work_item_admin (admin , work_item_id);

comment on column extra_schema.work_item_admin.work_item_id is '"cardinality":M2M';

comment on column extra_schema.work_item_admin.admin is '"cardinality":M2M';

create type extra_schema.work_item_role as ENUM (
  'extra_preparer'
  , 'extra_reviewer'
);

create table extra_schema.work_item_assignee (
  work_item_id bigint
  , assignee uuid
  , role extra_schema.work_item_role
  , primary key (work_item_id , assignee)
  , foreign key (work_item_id) references extra_schema.work_items (work_item_id)
  , foreign key (assignee) references extra_schema.users (user_id)
);

create index on extra_schema.work_item_assignee (assignee , work_item_id);

comment on column extra_schema.work_item_assignee.work_item_id is '"cardinality":M2M';

comment on column extra_schema.work_item_assignee.assignee is '"cardinality":M2M';

create table extra_schema.demo_work_items (
  work_item_id bigint primary key references extra_schema.work_items (work_item_id) on delete cascade
  , checked boolean not null default false
);

create table extra_schema.pag_element (
  paginated_element_id uuid default gen_random_uuid () primary key
  , name text not null
  , created_at timestamp with time zone default current_timestamp not null unique
  , dummy int
  , foreign key (dummy) references extra_schema.dummy_join (dummy_join_id) on delete cascade
);

--
--
-- END EXTRA SCHEMA showcase
--
--
create table projects (
  project_id serial primary key
  , name text not null unique
  , description text not null
  , work_items_table_name text not null unique
  , board_config jsonb not null default '{}'
  , created_at timestamp with time zone default CLOCK_TIMESTAMP() not null
  , updated_at timestamp with time zone default CLOCK_TIMESTAMP() not null
  , check (name ~ '^[a-zA-Z0-9_\-]+$')
);

comment on table projects is 'Internal use. Update whenever a project is added (project related tables added manually via migrations).
- work_items_table_name ensures project inserts are documented properly - postmigration script checks this column';

comment on column projects.work_items_table_name is '"properties":private';

comment on column projects.board_config is '"type":ProjectConfig';

comment on column projects.name is '"type":ProjectName';

create table teams (
  team_id serial primary key
  , project_id int not null --limited to a project only
  , name text not null
  , description text not null
  , created_at timestamp with time zone default CLOCK_TIMESTAMP() not null
  , updated_at timestamp with time zone default CLOCK_TIMESTAMP() not null
  , foreign key (project_id) references projects (project_id) on delete cascade
  , unique (name , project_id)
);

comment on column teams.project_id is '"cardinality":M2O';

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
  , age int
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
  -- should have a new `id bigserial` col for pagination alongside existing uuid instead of this
  -- cannot use now()/current_timestamp if we want to use transactions reliably
  , created_at timestamp with time zone default CLOCK_TIMESTAMP() not null unique
  , updated_at timestamp with time zone default CLOCK_TIMESTAMP() not null
  , deleted_at timestamp with time zone
  , foreign key (api_key_id) references user_api_keys (user_api_key_id) on delete cascade
);

comment on column users.external_id is '"properties":private';

comment on column users.api_key_id is '"properties":private';

comment on column users.role_rank is '"properties":private';

comment on column users.scopes is '"type":Scopes';

alter table user_api_keys
  add column user_id uuid not null unique;

alter table user_api_keys
  add foreign key (user_id) references users (user_id) on delete cascade;

comment on column user_api_keys.user_api_key_id is '"properties":private';

-- create unique index on users (user_id) where deleted_at is null; -- helps if you have much more deleted rows only
-- create index on users (deleted_at);  - not worth the extra overhead.
-- does get used when filtering deleted users exclusively and there's few of them
create index on users (deleted_at)
where (deleted_at is not null);

create index on users using gin (role_rank , age , username gin_trgm_ops , email gin_trgm_ops , full_name gin_trgm_ops);

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
  , labels text[] not null
  , link text
  , created_at timestamp with time zone default CLOCK_TIMESTAMP() not null
  , sender uuid not null
  , receiver uuid -- can be null for 'global' type
  , notification_type notification_type not null
  , foreign key (sender) references users (user_id) on delete cascade
  , foreign key (receiver) references users (user_id) on delete cascade
  , check (num_nonnulls (receiver_rank , receiver) = 1)
);

create index on notifications (receiver_rank , notification_type , created_at);

-- will use role enums for upper layers. this way we can extend create and update params
-- instead of manually keeping in sync.
comment on column notifications.receiver_rank is '"properties":private';

comment on column notifications.sender is '"cardinality":M2O';

comment on column notifications.receiver is '"cardinality":M2O';

create table user_notifications (
  user_notification_id bigserial primary key
  , notification_id int not null
  , read boolean default false not null -- for badge. frontend simply sends a list of user_notification_id to mark as read
  , user_id uuid not null
  , unique (notification_id , user_id)
  , foreign key (user_id) references users (user_id) on delete cascade
  , foreign key (notification_id) references notifications (notification_id) on delete cascade
);

comment on column user_notifications.user_id is '"cardinality":M2O';

-- user_notif are fan out
comment on column user_notifications.notification_id is '"cardinality":M2O';

create index on user_notifications (user_id);

-- read field simply used to add 'NEW' badge but there is no filtering
create or replace function notification_fan_out ()
  returns trigger
  language plpgsql
  as $function$
declare
  receiver_id uuid;
begin
  if new.notification_type = 'personal' then
    update
      users
    set
      has_personal_notifications = true
    where
      user_id = new.receiver;
    --
    insert into user_notifications (notification_id , user_id)
      values (new.notification_id , new.receiver);
  end if;
  if new.notification_type = 'global' then
    update
      users
    set
      has_global_notifications = true
    where
      role_rank >= new.receiver_rank;
    --
    for receiver_id in (
      -- in parallel tests fan out will loop through all users with rank >= X, but one of those may be a user
      -- that will be deleted, leading to could not create notification: Key (user_id)=(**) is not present in table "users".
      select
        user_id from users
        where
          role_rank >= new.receiver_rank)
    loop
      begin
        insert into user_notifications (notification_id , user_id)
          values (new.notification_id , receiver_id);
      exception
        when others then
          -- ignore all errors since this may loop through a user that is getting deleted (tests), etc.
          raise notice 'Error inserting notification for user_id=(%): % ' , receiver_id , SQLERRM;
      end;
    end loop;
  end if;
  -- it's after trigger so wouldn't mattern anyway
  return null;
end
$function$;

-- deletes get cascaded
create trigger notifications_fan_out
  after insert on notifications for each row
  execute function notification_fan_out ();

--
-- notifications table cant be adapted properly
-- e.g. WorkItem auditing. we can have `user <...> assigned member <...>`
-- We will also reuse topics (used in SSE and notifications)
create table entity_notifications (
  entity_notification_id serial primary key
  , id text not null
  , message text not null
  , topic text not null
  , created_at timestamp with time zone default CLOCK_TIMESTAMP() not null
);

comment on column entity_notifications.topic is '"type":Topics';

create table user_team (
  team_id int not null
  , member uuid not null
  , primary key (member , team_id)
  , foreign key (member) references users (user_id) on delete cascade
  , foreign key (team_id) references teams (team_id) on delete cascade
);

create index on user_team (team_id , member);

create index on user_team (member);

comment on column user_team.member is '"cardinality":M2M';

comment on column user_team.team_id is '"cardinality":M2M';

create table kanban_steps (
  kanban_step_id serial primary key
  , project_id int not null
  , step_order int not null -- 0: disabled
  , name text not null
  , description text not null
  , color text not null
  , time_trackable bool not null default false
  -- , disabled bool not null default false
  , unique (project_id , step_order)
  , foreign key (project_id) references projects (project_id) on delete cascade
  , check (step_order >= 0)
);

-- if it were null as before, unique index would need 2 parts:
-- create unique index on kanban_steps (project_id , name , step_order)
-- where
--   step_order is not null;
-- create unique index on kanban_steps (project_id , name)
-- where
--   step_order is null;
create unique index on kanban_steps (project_id , name , step_order);

comment on column kanban_steps.project_id is '"cardinality":M2O';

-- types restricted per project
create table work_item_types (
  work_item_type_id serial primary key
  , project_id int not null
  , name text not null
  , description text not null
  , color text not null
  , unique (name , project_id)
  , foreign key (project_id) references projects (project_id) on delete cascade
);

comment on column work_item_types.project_id is '"cardinality":M2O';

-- create table invoice_types (
--   invoice_type_id serial primary key
--   , project_id int not null
--   , name text not null
--   , foreign key (project_id) references projects (project_id) on delete cascade
-- );
-- create table default_invoice_type (
--   team_id int not null
--   , work_item_type_id int not null
--   , invoice_type_id int not null
--   , primary key (team_id , work_item_type_id)
--   , foreign key (team_id) references teams (team_id) on delete cascade
--   , foreign key (work_item_type_id) references work_item_types (work_item_type_id) on delete cascade
--   , foreign key (invoice_type_id) references invoice_types (invoice_type_id) on delete cascade
-- );
create table work_items (
  work_item_id bigserial primary key
  , title text not null
  , description text not null
  , work_item_type_id int not null
  , metadata jsonb not null
  , team_id int not null
  , kanban_step_id int not null
  , closed_at timestamp with time zone -- NULL: active
  , target_date timestamp with time zone not null
  , created_at timestamp with time zone default CLOCK_TIMESTAMP() not null
  , updated_at timestamp with time zone default CLOCK_TIMESTAMP() not null
  , deleted_at timestamp with time zone
  , foreign key (team_id) references teams (team_id) on delete cascade
  , foreign key (work_item_type_id) references work_item_types (work_item_type_id) on delete cascade
  , foreign key (kanban_step_id) references kanban_steps (kanban_step_id) on delete cascade
);

create index on work_items (team_id);

-- project for tour. when starting it user joins the only demo team. when exiting it user is removed.
-- we can reset it every X hours
create table demo_work_items (
  work_item_id bigint primary key references work_items (work_item_id) on delete cascade
  , ref text not null
  , line text not null
  , last_message_at timestamp with time zone not null
  , reopened boolean not null default false
);

create index on demo_work_items (ref , line);

create table demo_two_work_items (
  work_item_id bigint primary key references work_items (work_item_id) on delete cascade
  , custom_date_for_project_2 timestamp with time zone
);

-- FIXME xo cannot properly infer edge case when PK is FK
comment on column work_items.work_item_id is '"cardinality":O2O';

-- for finding all deleted work items exclusively
create index on work_items (deleted_at)
where (deleted_at is not null);

create table work_item_comments (
  work_item_comment_id bigserial primary key
  , work_item_id bigint not null
  , user_id uuid not null
  , message text not null
  , created_at timestamp with time zone default CLOCK_TIMESTAMP() not null
  , updated_at timestamp with time zone default CLOCK_TIMESTAMP() not null
  , foreign key (user_id) references users (user_id) on delete cascade
  , foreign key (work_item_id) references work_items (work_item_id) on delete cascade
);

comment on column work_item_comments.work_item_id is '"cardinality":M2O';

comment on column work_item_comments.user_id is '"cardinality":M2O';

create index on work_item_comments (work_item_id);

create table work_item_tags (
  work_item_tag_id serial primary key
  , project_id int not null
  , name text not null
  , description text not null
  , color text not null
  , deleted_at timestamp with time zone
  , unique (name , project_id)
  , foreign key (project_id) references projects (project_id) on delete cascade
);

create table work_item_work_item_tag (
  work_item_tag_id int not null
  , work_item_id bigint not null
  , primary key (work_item_id , work_item_tag_id) -- M2M, work_item can have multple tags. tags can be in multiple work_items (same as book authors example)
  , foreign key (work_item_id) references work_items (work_item_id) on delete cascade
  , foreign key (work_item_tag_id) references work_item_tags (work_item_tag_id) on delete cascade
);

create index on work_item_work_item_tag (work_item_tag_id , work_item_id);

comment on column work_item_work_item_tag.work_item_tag_id is '"cardinality":M2M';

comment on column work_item_work_item_tag.work_item_id is '"cardinality":M2M';

-- roles are append-only
create type work_item_role as ENUM (
  'preparer'
  , 'reviewer'
);

create table work_item_assignee (
  work_item_id bigint not null
  , assignee uuid not null
  , role work_item_role not null
  , primary key (work_item_id , assignee)
  , foreign key (work_item_id) references work_items (work_item_id) on delete cascade
  , foreign key (assignee) references users (user_id) on delete cascade
);

create index on work_item_assignee (assignee , work_item_id);

comment on column work_item_assignee.role is '"type":WorkItemRole';

comment on column work_item_assignee.work_item_id is '"cardinality":M2M';

comment on column work_item_assignee.assignee is '"cardinality":M2M';

-- must be completely dynamic on a project basis
create table activities (
  activity_id serial primary key
  , project_id int not null
  , name text not null
  , description text not null
  , is_productive boolean default false not null
  , deleted_at timestamp with time zone
  -- can't have multiple unrelated projects see each other's activities
  , unique (name , project_id)
  , foreign key (project_id) references projects (project_id) on delete cascade
);

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

comment on column time_entries.work_item_id is '"cardinality":M2O';

comment on column time_entries.team_id is '"cardinality":M2O';

comment on column time_entries.activity_id is '"cardinality":M2O';

comment on column time_entries.user_id is '"cardinality":M2O';

comment on column demo_work_items.ref is '"tags":pattern:"^[0-9]{8}$"';

comment on column work_item_types.color is '"tags":pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"';

comment on column work_item_tags.color is '"tags":pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"';

comment on column kanban_steps.color is '"tags":pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"';

-- will use project name as path param. Simplifies frontend without the need for model mappings as well.
-- although it's probably easier to have projectID just be a body parameter as it was
comment on column activities.project_id is '"cardinality":M2O && "properties":hidden';

comment on column teams.project_id is '"cardinality":M2O && "properties":hidden';

comment on column work_item_tags.project_id is '"cardinality":M2O && "properties":hidden';

comment on column work_item_types.project_id is '"cardinality":M2O && "properties":hidden';

comment on column kanban_steps.project_id is '"cardinality":M2O && "properties":hidden';

-- A multicolumn B-tree index can be used with query conditions that involve any subset of the index's
-- columns, but the index is most efficient when there are constraints on the leading (leftmost) columns.
create index on time_entries (user_id , team_id);

-- show user his timelog based on what projects are selected
create index on time_entries (work_item_id , team_id);

-- grpc demo
create table movies (
  movie_id serial primary key
  , title text not null
  , year integer not null
  , synopsis text not null
);

create table user_project (
  project_id int not null
  , member uuid not null
  , primary key (member , project_id)
  , foreign key (member) references users (user_id) on delete cascade
  , foreign key (project_id) references projects (project_id) on delete cascade
);

create index on user_project (project_id , member);

create index on user_project (member);

comment on column user_project.member is '"cardinality":M2M';

comment on column user_project.project_id is '"cardinality":M2M';


/*
-- sync existing users
do $BODY$
declare
 u_id uuid;
 proj_id int;
 t_id int;
 teams_in_project int[];
 user_scopes text[];
begin
 for u_id
 , proj_id
 , teams_in_project
 , user_scopes in
 select
 user_id
 , assigned_teams.project_id
 , ARRAY_AGG(assigned_teams.tip)
 , ARRAY_AGG(users.scopes)
 from
 users
 left join (
 select
 ut.member as uid
 , teams.project_id
 , ARRAY_AGG(ut.team_id) as tip
 from
 teams
 join user_team ut using (team_id)
 join users on ut.member = users.user_id
 where
 ut.member = users.user_id
 group by
 users.user_id , ut.member , teams.project_id) as assigned_teams on assigned_teams.uid = users.user_id
 left join projects on assigned_teams.uid = users.user_id
group by
 users.user_id
 , assigned_teams.project_id
 , users.scopes loop
 raise notice 'user_project project-member sync for u_id: % proj_id % ' , u_id , proj_id;
 execute FORMAT('
 INSERT INTO user_project (member, project_id)
 VALUES(%L,%L)
 ON CONFLICT DO NOTHING;
 ' , u_id , proj_id);
 -- assign to all teams in project
 if '{"project-member"}' = any (user_scopes) then
 FOREACH t_id in array teams_in_project loop
 raise notice 'user_team project-member sync for u_id: % t_id % ' , u_id , t_id;
 execute FORMAT('
 INSERT INTO user_team (member, team_id)
 VALUES(%L,%L)
 ON CONFLICT DO NOTHING;
 ' , u_id , t_id);
 end loop;
 end if;

 end loop;
end;
$BODY$
language plpgsql;
 */
create or replace function sync_user_teams ()
  returns trigger
  as $BODY$
declare
  users_to_include uuid[];
  uid uuid;
begin
  select
    ARRAY_AGG(user_id)
  from
    users
    join user_project up on up.member = users.user_id
  where
    up.project_id = new.project_id
    -- automatically include user with these scopes in all new teams
    and users.scopes @> '{"project-member"}' into users_to_include;
  if (users_to_include is null) then
    return new;
  end if;
  FOREACH uid in array users_to_include loop
    execute FORMAT('
            INSERT INTO user_team (member, team_id)
            VALUES(%L,%L)
            ON CONFLICT (member, team_id)
            DO NOTHING;
        ' , uid , new.team_id);
  end loop;
  raise notice 'team id % initialized with user ids: % ' , new.team_id , users_to_include;
  return NEW;
  end;
$BODY$
language plpgsql;

create trigger sync_user_teams
  after insert on teams for each row
  execute function sync_user_teams ();

-- assign user to team's project automatically.
-- we won't assign to projects individually, it's implicit.
create or replace function sync_user_projects ()
  returns trigger
  as $BODY$
begin
  insert into user_project (project_id , member)
  select
    teams.project_id
    , new.member
  from
    teams
    join user_team ut on ut.team_id = new.team_id
  where
    ut.member = new.member
  on conflict
    do nothing;
  raise notice 'user_project for  new.member and teamid % % ' , new.member , new.team_id;
  return NEW;
end;
$BODY$
language plpgsql;

create trigger sync_user_projects
  after insert or update on user_team for each row
  execute function sync_user_projects ();

--
-- audit
--
/*
select
 audit.enable_tracking ('public.kanban_steps');

select
 audit.enable_tracking ('public.projects');

select
 audit.enable_tracking ('public.teams');

select
 audit.enable_tracking ('public.work_items');
 */
----
create or replace function project_exists (project_name text)
  returns boolean
  as $$
declare
  project_exists_boolean boolean;
begin
  select
    exists (
      select
        1
      from
        pg_catalog.pg_class c
        join pg_catalog.pg_namespace n on n.oid = c.relnamespace
      where
        n.nspname = 'public'
        and c.relname = project_name
        and c.relkind = 'r' -- only tables
) into project_exists_boolean;

  return project_exists_boolean;
end;
$$
language plpgsql;


/*

 INIT
 */
insert into projects (name , description , work_items_table_name)
  values ('demo' , 'description for demo' , 'demo_work_items');

insert into projects (name , description , work_items_table_name)
  values ('demo_two' , 'description for demo_two' , 'demo_two_work_items');

insert into kanban_steps (name , description , project_id , color , step_order)
  values ('Disabled' , 'This column is disabled' , (
      select
        project_id
      from
        projects
      where
        name = 'demo') , '#aaaaaa' , 0);

insert into kanban_steps (name , description , project_id , color , step_order)
  values ('Received' , 'description for Received column' , (
      select
        project_id
      from
        projects
      where
        name = 'demo') , '#aaaaaa' , 1);

insert into kanban_steps (name , description , project_id , color , step_order)
  values ('Under review' , 'description for Under review column' , (
      select
        project_id
      from
        projects
      where
        name = 'demo') , '#f6f343' , 2);

insert into kanban_steps (name , description , project_id , color , step_order)
  values ('Work in progress' , 'description for Work in progress column' , (
      select
        project_id
      from
        projects
      where
        name = 'demo') , '#2b2444' , 3);

insert into work_item_types (name , description , project_id , color)
  values ('Type 1' , 'description for Type 1 work item type' , (
      select
        project_id
      from
        projects
      where
        name = 'demo') , '#282828');

insert into kanban_steps (name , description , project_id , color , step_order)
  values ('Received' , 'description for Received column' , (
      select
        project_id
      from
        projects
      where
        name = 'demo_two') , '#bbbbbb' , 1);

insert into work_item_types (name , description , project_id , color)
  values ('Type 1' , 'description for Type 1 work item type' , (
      select
        project_id
      from
        projects
      where
        name = 'demo_two') , '#282828');

insert into work_item_types (name , description , project_id , color)
  values ('Type 2' , 'description for Type 2 work item type' , (
      select
        project_id
      from
        projects
      where
        name = 'demo_two') , '#d0f810');

insert into work_item_types (name , description , project_id , color)
  values ('Another type' , 'description for Another type work item type' , (
      select
        project_id
      from
        projects
      where
        name = 'demo_two') , '#d0f810');
