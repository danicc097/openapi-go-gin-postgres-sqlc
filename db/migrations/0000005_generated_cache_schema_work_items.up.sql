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
  -- Dynamically fetch column names and data types from the project_table
  execute FORMAT('
        SELECT string_agg(column_name || '' '' || data_type, '', '')
        FROM information_schema.columns
        WHERE table_name = ''%I'' AND table_schema = ''public'' AND column_name != ''work_item_id''' , project_name) into project_table_col_and_type;
  -- Dynamically create the cache.demo_work_items table
  execute 'CREATE SCHEMA IF NOT EXISTS cache;';
  execute FORMAT('CREATE TABLE IF NOT EXISTS cache.%I (%s)' , project_name , project_table_col_and_type || ',' || work_items_col_and_type);
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
  project_table_cols text[];
  work_items_cols text[];
  sync_cols text[];
  update_cols text[];
  project_table_col_values text;
  all_cols_names text;
  all_values record;
  all_values_columns text[];
  wid bigint;
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

  select
    new.work_item_id into wid;

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
  -- Construct the list of columns to synchronize
  sync_cols := array (
    select
      'wi.' || column_name || ' AS ' || column_name
    from
      UNNEST(work_items_cols) as column_name) || array (
    select
      'NEW.' || column_name || ' AS ' || column_name
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
  all_cols_names := ARRAY_TO_STRING(array (
      select
        UNNEST(project_table_cols)
    union
    select
      UNNEST(work_items_cols)) , ', ');

  execute FORMAT('
        SELECT *
        FROM work_items wi
        JOIN %I USING (work_item_id)
        WHERE wi.work_item_id = $1' , project_name) into all_values
  using new.work_item_id;
  -- FIXME: get column names of all_values and use that instead of all_cols_names...
  -- this is overly complicated for no reason.
  -- all_values_columns := ARRAY_TO_STRING(array (
  --     select
  --       *
  --     from PG_TYPEOF(all_values)) , ', ');
  raise notice 'all_values_columns: %' , all_values_columns;
  raise notice 'all_values: %' , all_values;
  -- TODO: these come from all columns in the NEW KEYWORD
  -- Get values for work_items_cols once and use them in the subsequent INSERT INTO ... VALUES ...
  raise notice 'sssss %' , FORMAT(' insert into cache.%I (%s)
      values (%s)
    on conflict (work_item_id)
      do update set
        %s' , project_name , all_cols_names, all_values , ARRAY_TO_STRING(update_cols , ' , '));
  -- using NEW;
  return NEW;
end;
$$
language plpgsql;

-- TODO: anon func loop for project in projects table and drop and replace triggers always
select
  create_dynamic_table ('demo_work_items');

create trigger work_items_sync_trigger_demo_work_items
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
