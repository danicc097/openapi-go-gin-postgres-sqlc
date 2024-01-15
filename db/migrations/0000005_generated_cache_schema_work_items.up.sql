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
  execute FORMAT('ALTER TABLE cache.%I
  DROP CONSTRAINT IF EXISTS fk_cache_%s_work_item_id,
  DROP CONSTRAINT IF EXISTS cache_%s_work_item_id_unique
  ' , project_name , project_name , project_name);
  execute FORMAT('ALTER TABLE cache.%I
  ADD CONSTRAINT fk_cache_%s_work_item_id FOREIGN KEY (work_item_id)
  REFERENCES public.work_items (work_item_id) ON DELETE CASCADE,
  ADD CONSTRAINT cache_%s_work_item_id_unique UNIQUE (work_item_id)' , project_name , project_name , project_name);
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
  res record;
  all_columns_with_type text;
  all_columns text;
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

  select
    ARRAY_TO_STRING(array (
        select
          FORMAT('%I::%s' , column_name , data_type)
        from information_schema.columns
        where (table_name = 'work_items'
          or table_name = 'demo_work_items')
        and not (table_name = 'demo_work_items'
          and column_name = 'work_item_id') -- PK is FK
        and table_schema = 'public' order by ordinal_position) , ', ') into all_columns_with_type;
  select
    ARRAY_TO_STRING(array (
        select
          FORMAT('%I' , column_name)
        from information_schema.columns
        where (table_name = 'work_items'
          or table_name = 'demo_work_items')
        and not (table_name = 'demo_work_items'
          and column_name = 'work_item_id') -- PK is FK
        and table_schema = 'public' order by ordinal_position) , ', ') into all_columns;

  execute FORMAT( '
with del as (
  delete from cache.demo_work_items as t
  where t.work_item_id = $1)
insert into cache.demo_work_items
  (%s)
  select
    %s
  from work_items wi
  join demo_work_items using (work_item_id)
  where
    wi.work_item_id = $1
  on conflict (work_item_id)
  do nothing -- another tx won insert after delete'
    , all_columns , all_columns_with_type)
  using new.work_item_id;
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
-- with data as (
--   select
--     work_item_id::bigint
--     , ref::text
--     , title::text
--     , description::text
--     , line::text
--     , last_message_at::timestamp with time zone
--     , work_item_type_id::integer
--     , metadata::jsonb
--     , reopened::boolean
--     , team_id::integer
--     , kanban_step_id::integer
--     , closed_at::timestamp with time zone
--     , target_date::timestamp with time zone
--     , created_at::timestamp with time zone
--     , updated_at::timestamp with time zone
--     , deleted_at::timestamp with time zone
--   from
--     work_items wi
--     join demo_work_items using (work_item_id)
--   where
--     wi.work_item_id = 1
-- )
-- , del as (
--   delete from cache.demo_work_items as t using data d
-- -- cache t has  {ref,line,last_message_at,reopened,work_item_id,title,description,work_item_type_id,metadata,team_id,kanban_step_id,closed_at,target_date,created_at,updated_at,deleted_at}
-- -- d has a single work_item_id...
-- where t.work_item_id = d.work_item_id)
-- insert into cache.demo_work_items as t table data
-- on conflict (work_item_id)
--   do nothing
-- returning
--   t.work_item_id
