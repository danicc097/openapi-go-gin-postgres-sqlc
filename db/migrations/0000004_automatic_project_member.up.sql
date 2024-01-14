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

create or replace function create_dynamic_table (project_name text)
  returns VOID
  as $$
declare
  project_table_col_and_type text;
  work_items_col_and_type text;
begin
  -- Dynamically fetch column names and data types from work_items
  execute '
        SELECT string_agg(column_name || '' '' || data_type, '', '')
        FROM information_schema.columns
        WHERE table_name = ''work_items'' AND table_schema = ''public''' into work_items_col_and_type;

  project_table_col_and_type := project_table_col_and_type || ',';

  execute FORMAT('
        SELECT string_agg(column_name || '' '' || data_type, '', '')
        FROM information_schema.columns
        WHERE table_name = ''%I'' AND table_schema = ''public''' , project_name) into project_table_col_and_type;
  -- Dynamically create the cache.demo_work_items table
  execute 'CREATE SCHEMA if not exists cache;';
  -- EXECUTE 'CREATE TABLE cache.demo_work_items (' || project_table_col_and_type || work_items_col_and_type || ')';
  raise notice '% fields: %' , project_name , project_table_col_and_type;
  raise notice 'work_item fields: %' , work_items_col_and_type;

end;
$$
language plpgsql;

select
  create_dynamic_table ('demo_work_items');

create or replace function sync_work_items ()
  returns trigger
  as $$
declare
  project_name text;
  project_table_cols text[];
  work_items_cols text[];
  sync_cols text[];
  update_cols text[];
begin
  project_name := TG_ARGV[0];
  -- Make sure the project table exists
  perform
    1
  from
    pg_catalog.pg_class c
    join pg_catalog.pg_namespace n on n.oid = c.relnamespace
  where
    n.nspname = 'public'
    and c.relname = project_name
    and c.relkind = 'r';

  if not FOUND then
    raise exception 'Project table "%" does not exist' , project_name;
  end if;
  -- Dynamically fetch column names
  execute FORMAT('
        SELECT ARRAY_AGG(column_name)
        FROM information_schema.columns
        WHERE table_name = ''%I'' AND table_schema = ''public''' , project_name) into project_table_cols;
  execute '
        SELECT ARRAY_AGG(column_name)
        FROM information_schema.columns
        WHERE table_name = ''work_items'' AND table_schema = ''public''' into work_items_cols;

  raise notice 'work_items_cols: %' , work_items_cols;
  -- Construct the list of columns to synchronize
  sync_cols := array (
    select
      'wi.' || column_name || ' AS ' || column_name
    from
      UNNEST(work_items_cols) as column_name) || array (
    select
      'new.' || column_name || ' AS ' || column_name
    from
      UNNEST(project_table_cols) as column_name);
  -- Construct the list of columns for the ON CONFLICT DO UPDATE part
  update_cols := array (
    select
      column_name || ' = EXCLUDED.' || column_name
    from
      UNNEST(project_table_cols) as column_name) || array (
    select
      column_name || ' = wi.' || column_name
    from
      UNNEST(work_items_cols) as column_name);

  -- TODO: should exit early on clashing column names.
  -- need a trigger like post-migration/2-check-projects.sql so its catched in development
  -- Dynamically execute the synchronization query
  raise notice 'insert stmt: %' , FORMAT('
        INSERT INTO cache.%I (%s)
        SELECT %s
        FROM %I
        JOIN work_items wi USING (work_item_id)
        ON CONFLICT (work_item_id) DO UPDATE
        SET %s
    ' , project_name ,
    -- Concatenate all column names in order
    ARRAY_TO_STRING(array (
        select
          UNNEST(project_table_cols)
      union
      select
        UNNEST(work_items_cols)) , ',') ,
    -- Constructing SELECT clause with wi. and NEW. prefixes
    ARRAY_TO_STRING(sync_cols , ', ') , project_name ,
    -- Constructing SET clause for conflict resolution
    ARRAY_TO_STRING(update_cols , ', '));

  return NEW;
end;
$$
language plpgsql;

create trigger work_items_sync_trigger
  after insert or update on demo_work_items for each row
  execute function sync_work_items ('demo_work_items');

-- INSERT INTO cache.demo_work_items (work_item_id,ref,line,last_message_at,reopened, <here would go work_items_cols>)
-- SELECT wi.work_item_id AS work_item_id, wi.ref AS ref, wi.line AS line, wi.last_message_at AS last_message_at, wi.reopened AS reopened
-- FROM demo_work_items wi
-- JOIN work_items nw ON wi.work_item_id = nw.work_item_id
-- ON CONFLICT (work_item_id) DO UPDATE
-- SET
-- work_item_id = EXCLUDED.work_item_id,
-- ref = EXCLUDED.ref,
-- line = EXCLUDED.line,
-- last_message_at = EXCLUDED.last_message_at,
-- reopened = EXCLUDED.reopened
-- ,
-- <here would go work_items_cols doing <col> = wi.<col>>
