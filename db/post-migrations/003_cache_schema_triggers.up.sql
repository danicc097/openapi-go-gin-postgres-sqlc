create or replace function create_work_item_cache_table (project_name text)
  returns VOID
  as $$
declare
  project_table_col_and_type text;
  foreign_key_constraints_text text;
  work_items_col_and_type text;
  constraint_exists boolean;
begin
  select
    STRING_AGG(constraint_definition , ', ')
  from (
    select
      --conname as constraint_name
      --, conrelid::regclass as table_name
      PG_GET_CONSTRAINTDEF(c.oid) as constraint_definition
    from
      pg_constraint c
    where (conrelid = 'public.demo_work_items'::regclass
      or conrelid = 'public.work_items'::regclass)
    and contype = 'f') into foreign_key_constraints_text;

  select
    STRING_AGG(
      case when column_name != 'work_item_id' then
        column_name || ' ' || data_type || ' ' || case when is_nullable = 'YES' then
          ' NULL'
        else
          ' NOT NULL'
        end
      else
        'work_item_id bigint primary key'
      end , ', ')
  from
    information_schema.columns
  where
    table_name = 'work_items'
    and table_schema = 'public' into work_items_col_and_type;

  execute FORMAT('
	SELECT string_agg(column_name || '' '' || data_type || '' '' || CASE WHEN is_nullable = ''YES'' THEN
	  '' NULL'' ELSE '' NOT NULL'' END, '', '')
        FROM information_schema.columns
        WHERE table_name = ''%I'' AND table_schema = ''public'' AND column_name != ''work_item_id''' , project_name) into project_table_col_and_type;
  -- execute 'CREATE SCHEMA IF NOT EXISTS cache;';
  execute FORMAT('CREATE TABLE IF NOT EXISTS cache__%I (%s)' , project_name , project_table_col_and_type || ',' || work_items_col_and_type ||
    ',' || foreign_key_constraints_text);
  -- execute FORMAT('comment on column cache__%I.work_item_id is ''"type":WorkItemID && "properties":refs-ignore''' , project_name);
  execute FORMAT('comment on column cache__%I.work_item_id is ''"properties":refs-ignore,share-ref-constraints''' , project_name);
  -- TODO: xo will duplicate M2M and M2O constraints in constraints slice for the referenced column if share-ref-constraints set``
end;
$$
language plpgsql;

create or replace function sync_work_items ()
  returns trigger
  as $$
declare
  project_name text;
  all_columns_with_type text;
  all_columns text;
  conflict_update_columns text;
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

  with data as (
    select
      FORMAT('%I'
        , column_name) as c
      , FORMAT('%I::%s'
        , column_name
        , data_type) as cwt
    from
      information_schema.columns
    where (table_name = 'work_items'
      or table_name = project_name)
    and not (table_name = project_name
      and column_name = 'work_item_id') -- PK is FK
    and table_schema = 'public'
  order by
    ordinal_position
)
select
  ARRAY_TO_STRING(array (
      select
        c
      from data) , ', ')
  , ARRAY_TO_STRING(array (
      select
        cwt
      from data) , ', ')
  , ARRAY_TO_STRING(array (
      select
        FORMAT('%s = EXCLUDED.%s' , c , c)
      from data) , ', ') into all_columns
  , all_columns_with_type
  , conflict_update_columns;

  execute FORMAT( '
insert into cache__%I
  (%s)
  select
    %s
  from public.work_items wi
  join public.%I using (work_item_id)
  where
    wi.work_item_id = $1
  on conflict (work_item_id)
  do update set
    %s -- construct for all rows c = EXCLUDED.c (excluded is populated with all rows)
  ' , project_name , all_columns , all_columns_with_type , project_name , conflict_update_columns)
  using new.work_item_id;
  return NEW;
end;
$$
language plpgsql;

--
--
-- Sync project work item tables cache
--
--
do $BODY$
declare
  project_name text;
begin
  for project_name in
  select
    work_items_table_name
  from
    projects loop
      -- IMPORTANT: now gin indexes can cover everything we want,
      -- but kept in sync via regular migrations:
      -- CREATE UNIQUE INDEX CONCURRENTLY newidx ON tab (name, price, sku);
      -- DROP INDEX cache_demo_work_items_<...>_index;
      -- ALTER INDEX newidx RENAME TO cache_demo_work_items_<...>_index;
      perform
        create_work_item_cache_table (project_name);

      execute FORMAT('create or replace trigger work_items_sync_trigger_%1$I
        after insert or update on %1$I for each row
        execute function sync_work_items (%1$s);' , project_name);
    end loop;
end;
$BODY$
language plpgsql;
