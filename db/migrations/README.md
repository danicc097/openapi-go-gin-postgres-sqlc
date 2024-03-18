> # Schema dumping (obtuse version, do not implement)
> **breaks plpgsql lsp extension, would need opt migration names to skip regex > option,
> or via flag on top of the file (preferred)**
>
> Eventually, accumulated migrations will take too long to run compared to
> creating the final schema in one go.
> You can create a new migration via `project migrate.create x_schema_dump` > (must be named
> x_schema_dump).
> These migration files will be detected before executing `up` migrations.
>
> The latest x_schema_dump migration will be the new reference point for further
> migrations.
>
> If current revision is 0 (i.e. `error: no migration` exit 1),
> it will execute `force $((latest_schema_dump_revision - 1))`
> and then `up` normally. Ideal for tests, gen db, etc.
>
> If current revision is within 1 and `$((latest_schema_dump_revision -1))` (e.g.
> in prod environment)
> it will skip all x_schema_dump migrations in between, executing
> `goto $((schema_dump_revision - 1))` and then `force $schema_dump_revision` in > ascending x_schema_dump
> migrations order.
>
> In any other case, it simply runs `up` as usual.


# Schema dumping (less obtuse version)

Dump schema from up to date gen_db:

```bash
project db.bash
# we must include data, not just schema.
/$ pg_dump gen_db --column-inserts  --exclude-table schema_migrations --exclude-table schema_post_migrations  > /var/lib/postgresql/dump.sql
```

TODO: all comments will be lost. should instead make use of comments on tables
instead of sql comments so they're saved after migrating to a schema dump
and removing old files. (not like anyone would read old migration files comments
anyway). Can use `project dev-utils.show-table-comments` to make maintenance
easier.

Once prod is up to date (i.e. revision at latest_schema_dump_revision minus
one),
we will `force` latest_schema_dump_revision on prod so that it is ignored on the
next deployment.
We must delete all migration files before latest_schema_dump_revision.
`golang-migrate` will not care that they don't exist anymore, it will start at the
first migration it encounters.

`.down.sql` for the schema dump will have to look something like this if we use
the output of pg_dump without postprocessing (cannot add `if exists` to create
statements).

```sql
drop schema if exists "extra_schema" cascade;

drop schema if exists "cache" cascade;

drop schema if exists "v" cascade;

do $$
declare
  table_rec RECORD;
  type_rec RECORD;
  function_rec RECORD;
  schema_name text;
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
  -- Drop functions
  for function_rec in
  select
    proname
  from
    pg_proc
  where
    pronamespace = 'public'::regnamespace loop
      execute 'DROP FUNCTION IF EXISTS public.' || QUOTE_IDENT(function_rec.proname) || ' CASCADE';
    end loop;
end
$$;

drop extension "supa_audit";

drop schema if exists "audit" cascade;

drop schema if exists "extensions" cascade;

```
