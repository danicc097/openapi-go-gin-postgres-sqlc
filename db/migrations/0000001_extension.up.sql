create schema if not exists extensions;

-- make sure everybody can use everything in the extensions schema
grant usage on schema extensions to public;

grant execute on all functions in schema extensions to public;

-- include future extensions
alter default privileges in schema extensions grant execute on functions to public;

alter default privileges in schema extensions grant usage on types to public;

create extension if not exists plpgsql_check schema extensions;

create extension if not exists "uuid-ossp" schema extensions;

-- https://github.com/supabase/supa_audit
-- not available in supabase itself, rather pgaudit
do $$
begin
  create extension if not exists supa_audit schema extensions;
exception
  when others then
    raise notice 'supa_audit extension not available';
end
$$;

create extension if not exists pg_stat_statements schema extensions;

create extension if not exists hypopg schema extensions;

create extension if not exists index_advisor schema extensions;

create extension if not exists pg_trgm schema extensions;

create extension if not exists btree_gin schema extensions;

create extension if not exists rum schema extensions;
