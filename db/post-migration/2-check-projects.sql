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

      select
        ARRAY_AGG(column_name) into work_items_columns
      from
        information_schema.columns
      where
        table_name = 'work_items'
        and table_schema = 'public';

      select
        ARRAY_AGG(column_name) into project_columns
      from
        information_schema.columns
      where
        table_name = t
        and table_schema = 'public'
        --PK is FK
        and column_name != 'work_item_id';
      -- Check for column name clashing.
      -- we will autogenerate and maintain cache schema tables so they must be unique.
      if work_items_columns && project_columns then
        raise exception '
	column names overlap between work_items and project table "%".
    work_items_columns: %
    project_columns: % ' , t , work_items_columns , project_columns;
      end if;
    end loop;
end;
$BODY$
language plpgsql;
