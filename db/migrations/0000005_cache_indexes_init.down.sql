do $do$
declare
  _tbl text;
begin
  for _tbl in
  select
    QUOTE_IDENT(table_schema) || '.' || QUOTE_IDENT(table_name)
  from
    information_schema.tables
  where
    table_name like 'cache__' || '%'
    and table_schema not like 'pg\_%' loop
      execute FORMAT('DROP TABLE ' || _tbl);
    end loop;
end
$do$;
