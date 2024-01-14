do $BODY$
declare
  t text;
  project_exists boolean;
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
    end loop;
end;
$BODY$
language plpgsql;
