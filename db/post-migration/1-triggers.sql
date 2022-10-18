--
-- Update updated_at timestamp if values changed on update
--
create or replace function before_update_updated_at ()
  returns trigger
  as $BODY$
begin
  if row (new.*::text) is distinct from row (old.*::text) then
    new.updated_at = NOW();
  end if;
  return NEW;
end;
$BODY$
language plpgsql;

--
-- Apply before_update_updated_at function to all schemas and tables as trigger
--
do $BODY$
declare
  t text;
  declare s text;
begin
  for t,
  s in
  select
    table_name,
    table_schema
  from
    information_schema.columns
  where (column_name = 'updated_at'
    and (
      select
        1
      from
        information_schema.triggers
      where
        trigger_name = 'before_update_updated_at_' || table_schema || '_' || table_name) is null)
    loop
      execute FORMAT('
            CREATE TRIGGER before_update_updated_at_%s_%s
            BEFORE UPDATE ON %I
            FOR EACH ROW EXECUTE PROCEDURE before_update_updated_at();
        ', s, t, t);
    end loop;
end;
$BODY$
language plpgsql;
