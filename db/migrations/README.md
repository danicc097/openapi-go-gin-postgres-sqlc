# Schema dumping

**not implemented.**

**breaks plpgsql lsp extension, should add migration names to skip regex option,
or via flag on top of the file (preferred)**

Eventually, accumulated migrations will take too long to run compared to
creating the final schema in one go.
You can create a new migration via `project migrate.create x_schema_dump` (must be named
x_schema_dump).
These migration files will be detected before executing `up` migrations.

The latest x_schema_dump migration will be the new reference point for further
migrations.

If current revision is 0 (i.e. `error: no migration` exit 1),
it will execute `force $((latest_schema_dump_revision - 1))`
and then `up` normally. Ideal for tests, gen db, etc.

If current revision is within 1 and `$((latest_schema_dump_revision -1))` (e.g.
in prod environment)
it will skip all x_schema_dump migrations in between, executing
`goto $((schema_dump_revision - 1))` and then `force $schema_dump_revision` in ascending x_schema_dump
migrations order.

In any other case, it simply runs `up` as usual.


