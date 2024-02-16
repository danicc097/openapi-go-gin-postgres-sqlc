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
      /* TODO: will have to include extra denormalized fields accordingly
       per project_name inside sync_work_items.
       Unclear if this can remain in post migration or has to be moved to regular
       migrations (unlikely that we will have data migrations for cache table that require trigger functions to be up
       to date before executing post migrations)
       */
      execute FORMAT('create or replace trigger work_items_sync_trigger_%1$I
        after insert or update on %1$I for each row
        execute function sync_work_items (%1$s);' , project_name);

    end loop;
end;
$BODY$
language plpgsql;
