create or replace function create_work_item_cache_table (project_name text)
  returns VOID
  as $$
declare
  project_table_col_and_type text;
  foreign_key_constraints_text text;
  work_items_col_and_type text;
  constraint_exists boolean;
  tags_comment text;
  col text;
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
  -- we lose "tags" annotation from column comments in ref
  for col
  , tags_comment in
  select
    a.attname as col
    , case when d.description ~ '"tags":\s*([^,]+)' then
      REGEXP_REPLACE(d.description , '.*"tags":\s*([^,]+).*' , '"tags":\1')
    else
      null
    end as tags_comment
  from
    pg_catalog.pg_description d
    join pg_catalog.pg_attribute a on d.objoid = a.attrelid
  where
    a.attrelid = 'public.work_items'::regclass
    or a.attrelid = ('public.' || project_name)::regclass
    and a.attnum = d.objsubid loop
      begin
        continue
        when tags_comment = null;

        execute FORMAT('comment on column cache__%I.%s is ''%s''' , project_name , col , tags_comment);
      end;
    end loop;
  -- override
  execute FORMAT('comment on column cache__%I.work_item_id is ''"properties":refs-ignore,share-ref-constraints''' , project_name);

end;
$$
language plpgsql;

select
  create_work_item_cache_table ('demo_work_items');

select
  create_work_item_cache_table ('demo_two_work_items');

create index newidx on cache__demo_work_items using gin ( --
title extensions.gin_trgm_ops --
, line extensions.gin_trgm_ops --
, ref extensions.gin_trgm_ops --
, reopened);

drop index if exists cache__demo_work_items_gin_index;

alter index newidx rename to cache__demo_work_items_gin_index;

create index newidx on cache__demo_two_work_items using gin ( --
title extensions.gin_trgm_ops --
);

drop index if exists cache__demo_two_work_items_gin_index;

alter index newidx rename to cache__demo_two_work_items_gin_index;
