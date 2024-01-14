do $BODY$
declare
  t text;
  project_exists boolean;
  work_items_columns text[];
  project_columns text[];
begin
  for t in
  select
    work_items_table_name
  from
    projects loop
      select
        exists (
          select
          from
            pg_catalog.pg_class c
            join pg_catalog.pg_namespace n on n.oid = c.relnamespace
          where
            n.nspname = 'public'
            and c.relname = t
            and c.relkind = 'r' -- only tables
) into project_exists;
      if not project_exists then
        raise exception 'Project table "%" does not exist' , t;
      end if;
      -- work_items_columns := array[]::text[];
      -- project_columns := array[]::text[];
      -- -- Extract work_items table columns to a text[]
      -- select
      --   ARRAY_AGG(column_name) into work_items_columns
      -- from
      --   information_schema.columns
      -- where
      --   table_name = 'work_items';
      -- -- Extract project table columns to a text[]
      -- select
      --   ARRAY_AGG(column_name) into project_columns
      -- from
      --   information_schema.columns
      -- where
      --   table_name = t;
      -- -- Check for column name conflicts
      -- if work_items_columns && project_columns then
      --   raise exception 'Column names overlap between work_items and project table "%".
      --   work_items_columns: %
      --   project_columns: %
      --   ' , t , work_items_columns , project_columns;
      -- end if;
    end loop;
end;
$BODY$
language plpgsql;
