- having gen/ (db package) and crud packages is less than ideal, need to merge
  output.
  models wont clash (sqlc plural xo singular).
  we could possibly merge them to xo's models (xo has custom struct fields, sqlc custom
  tags `db:` but is not in use).

  we can have xo output a single file
  -S, --single=<file> enable single file output
  so we can parse it more easily if necessary.

  will need to get rid of xo's generated enum types files since they clash
  (sqlc's better)

  xo queries will receive sqlc's q.db (for easier transaction use, which in sqlc
  is

  ```go
  tx, err := db.Begin()
  if err != nil {
  	return err
  }
  defer tx.Rollback()
  qtx := queries.WithTx(tx) // a new *Queries
  ```

  )

- xo does not return excluded fields (since theyre not in model), e.g.
  created_at and updated_at.

  we want to return them in selects, but ignore in the rest of queries.
  models should include the fields with a comment "Read-only and managed by
  database".

  See `ignoreName` in template to fix this
