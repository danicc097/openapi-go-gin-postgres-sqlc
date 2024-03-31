drop schema if exists v;

drop schema if exists "cache";

drop schema if exists extra_schema cascade;

do $$
declare
  table_rec RECORD;
  type_rec RECORD;
begin
  -- Drop tables
  for table_rec in
  select
    table_name
  from
    information_schema.tables
  where
    table_schema = 'public'
    and table_type = 'BASE TABLE' loop
      execute 'DROP TABLE IF EXISTS public.' || QUOTE_IDENT(table_rec.table_name) || ' CASCADE';
    end loop;
  -- Drop enums
  for type_rec in
  select
    typname
  from
    pg_type
  where
    typnamespace = 'public'::regnamespace
    and typtype = 'e' loop
      execute 'DROP TYPE IF EXISTS public.' || QUOTE_IDENT(type_rec.typname) || ' CASCADE';
    end loop;
end
$$;
