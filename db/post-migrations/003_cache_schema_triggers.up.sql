create or replace function create_work_item_cache_table (project_name text)
  returns VOID
  as $$
declare
  project_table_col_and_type text;
  work_items_col_and_type text;
  constraint_exists boolean;
begin
  select
    STRING_AGG(column_name || ' ' || data_type || ' ' || case when is_nullable = 'YES' then
        ' NULL'
      else
        ' NOT NULL'
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
  -- Dynamically create the cache.demo_work_items table
  execute 'CREATE SCHEMA IF NOT EXISTS cache;';
  execute FORMAT('CREATE TABLE IF NOT EXISTS cache.%I (%s)' , project_name , project_table_col_and_type || ',' || work_items_col_and_type);
  execute FORMAT('comment on column cache.%I.work_item_id is ''"type":WorkItemID && "properties":ignore-constraints''' , project_name);
  -- constraints
  select
    exists (
      select
        1
      from
        information_schema.table_constraints
      where
        constraint_name = 'fk_cache_' || project_name || '_work_item_id'
        and table_schema = 'cache'
        and table_name = project_name) into constraint_exists;
  if not constraint_exists then
    execute FORMAT('ALTER TABLE cache.%I ADD CONSTRAINT fk_cache_%s_work_item_id
    FOREIGN KEY (work_item_id) REFERENCES public.work_items (work_item_id) ON DELETE CASCADE' , project_name , project_name);
  end if;
  select
    exists (
      select
        1
      from
        information_schema.table_constraints
      where
        constraint_name = 'cache_' || project_name || '_work_item_id_unique'
        and table_schema = 'cache'
        and table_name = project_name) into constraint_exists;
  if not constraint_exists then
    execute FORMAT('ALTER TABLE cache.%I ADD CONSTRAINT cache_%s_work_item_id_unique
    UNIQUE (work_item_id)' , project_name , project_name);
  end if;
  -- IMPORTANT: we will use extra cache table columns so logic to add/modify cols will be messy.
  -- altering existing types would have to be done manually either way.
  -- better do these steps manually since its just a duplicate statement, ie when doing migrations (triggers will fail on migration if not synced so there's no risk of out of date cache schema).
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
insert into cache.%I
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
      perform
        create_work_item_cache_table (project_name);

      execute FORMAT('create or replace trigger work_items_sync_trigger_%1$I
        after insert or update on %1$I for each row
        execute function sync_work_items (%1$s);' , project_name);
    end loop;
end;
$BODY$
language plpgsql;
