-- NOTE: do not use `extensions.` prefix for indexes. Already in search_path.
create or replace function normalize_index_def (input_string text)
  returns text
  as $$
declare
  output_string text;
begin
  -- Remove /* */ style comments
  output_string := REGEXP_REPLACE(input_string , '/\*([^*]|\*+[^*/])*\*+/' , '' , 'g');
  -- Remove -- style comments
  output_string := REGEXP_REPLACE(output_string , '--.*?\n' , '' , 'g');
  -- Replace newlines
  output_string := REGEXP_REPLACE(output_string , E'[\n\r]+' , ' ' , 'g');
  -- No ; in stored def
  output_string := REGEXP_REPLACE(output_string , ';$' , '' , 'g');
  output_string := REGEXP_REPLACE(output_string , ' +' , ' ' , 'g');
  output_string := REGEXP_REPLACE(output_string , '\s*\(\s*' , '(' , 'g');
  output_string := REGEXP_REPLACE(output_string , '\s*\)\s*' , ')' , 'g');
  output_string := REGEXP_REPLACE(output_string , '\s*,\s*' , ',' , 'g');
  output_string := TRIM(LOWER(output_string));

  return output_string;
end;
$$
language plpgsql;

create or replace function same_index_definition (index_name text , new_index_def text)
  returns boolean
  as $$
declare
  existing_def text;
begin
  perform
    indexdef
  from
    pg_indexes
  where
    indexname = index_name;
  if not found then
    return false;
  end if;

  execute FORMAT('SELECT normalize_index_def(pg_get_indexdef(''%I''::regclass));' , index_name) into existing_def;
  new_index_def := normalize_index_def (new_index_def);
  if existing_def is null or existing_def <> new_index_def then
    return false;
  end if;

  return true;
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

create or replace function create_or_update_work_item_cache_table (project_name text)
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

-- NOTE: concurrently not supported in migration transactions
create or replace function create_or_update_index (idx_name text , project_name text , idx_def text)
  returns VOID
  as $$
begin
  if idx_def <> '' and not same_index_definition (idx_name , idx_def) then
    raise notice 'recreating differing index: %' , idx_name;
    execute FORMAT('create index newidx on cache__%I %s;' , project_name , idx_def);
    execute FORMAT('drop index if exists %I;' , idx_name);
    execute FORMAT('alter index newidx rename to %I;' , idx_name);
  else
    raise notice 'skipping identical create index statement: %' , idx_name;
  end if;
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
  idx_def text;
  idx_name text;
begin
  for project_name in
  select
    work_items_table_name
  from
    projects loop
      --
      -- sync cache table - idempotent
      --
      perform
        create_or_update_work_item_cache_table (project_name);
      --
      -- at least one gin index mandatory for all projects
      --
      idx_name := FORMAT('public.cache__%I_gin_index' , project_name);


      /*
       - gin cannot use sort index, just btrees (or using btree_gin since pg12). the only alternative would be RUM
       but its not guaranteed to be always faster. also doesnt support like op unless changing code:
       see https://github.com/postgrespro/rum/issues/34 for supporting LIKE.
       - gin is apparently not too helpful for very short string searches.
       - btree not usable for text except for equals and startsWith.


       make sure the index used in explain is the one being tested...

       drop index if exists lastmsg;
       create index lastmsg on cache__demo_work_items using btree (
       last_message_at desc);

       drop index if exists abc;
       create index abc on cache__demo_work_items using gin (
       description gin_trgm_ops
       , last_message_at
       , reopened);

       set enable_seqscan = "off";
       -- with 1000 rows
       explain analyze select * from cache__demo_work_items where description  ilike '%4%' order by last_message_at  desc limit 20;
       ->  Index Scan using lastmsg on cache__demo_work_items  (cost=0.28..58.23 rows=263 width=145) (actual time=0.018..0.160 rows=20 loops=1)
       */
      case project_name
      when 'demo_work_items' then
        idx_def := 'using gin (
        title gin_trgm_ops
        , line gin_trgm_ops
        , description gin_trgm_ops
        , ref gin_trgm_ops
        , reopened)';
      when 'demo_two_work_items' then
        idx_def := 'using gin (
        title gin_trgm_ops
        , description gin_trgm_ops
        )';
      else
        idx_def := ''; raise exception 'No index definition found for cache__%' , project_name;
      end case;

      perform
        create_or_update_index (idx_name , project_name , idx_def);
        --
        -- adhoc indexes. TODO: abstract away idx create/replace
        --
        case project_name
        when 'demo_work_items' then
          idx_name := FORMAT('public.cache__%I_last_message_at_index' , project_name);
          --
          idx_def := 'using btree (last_message_at)';
          perform
            create_or_update_index (idx_name , project_name , idx_def);
          else
        end case;
        --
        -- triggers
        --
        execute FORMAT('create or replace trigger work_items_sync_trigger_%1$I
        after insert or update on %1$I for each row
        execute function sync_work_items (%1$s);' , project_name);
        end loop;
      end;
$BODY$
language plpgsql;
