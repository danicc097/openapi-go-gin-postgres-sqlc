{{/*see sqlc/internal/codegen/golang/templates/template.tmpl for enum compatibility */}}
{{ define "enum" }}
{{- $e := .Data -}}
// {{ $e.GoName }} is the '{{ $e.SQLName }}' enum type from schema '{{ schema }}'.
type {{ $e.GoName }} string

// {{ $e.GoName }} values.
const (
{{ range $e.Values -}}
	// {{ $e.GoName }}{{ .GoName }} is the '{{ .SQLName }}' {{ $e.SQLName }}.
	{{ $e.GoName }}{{ .GoName }} {{ $e.GoName }} = {{ .ConstValue }}
{{ end -}}
)

// Value satisfies the driver.Valuer interface.
func ({{ short $e.GoName }} {{ $e.GoName }}) Value() (driver.Value, error) {
	return string({{ short $e.GoName }}), nil
}

// Scan satisfies the sql.Scanner interface.
func ({{ short $e.GoName }} *{{ $e.GoName }}) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*{{ short $e.GoName }} = {{ $e.GoName }}(s)
	case string:
		*{{ short $e.GoName }} = {{ $e.GoName }}(s)
	default:
		return fmt.Errorf("unsupported scan type for {{ $e.GoName }}: %T", src)
	}
	return nil
}

{{ $nullName := (printf "%s%s" "Null" $e.GoName) -}}
{{- $nullShort := (short $nullName) -}}
// {{ $nullName }} represents a null '{{ $e.SQLName }}' enum for schema '{{ schema }}'.
type {{ $nullName }} struct {
	{{ $e.GoName }} {{ $e.GoName }}
	// Valid is true if {{ $e.GoName }} is not null.
	Valid bool
}

// Value satisfies the driver.Valuer interface.
func ({{ $nullShort }} {{ $nullName }}) Value() (driver.Value, error) {
	if !{{ $nullShort }}.Valid {
		return nil, nil
	}
	return {{ $nullShort }}.{{ $e.GoName }}.Value()
}

// Scan satisfies the sql.Scanner interface.
func ({{ $nullShort }} *{{ $nullName }}) Scan(v interface{}) error {
	if v == nil {
		{{ $nullShort }}.{{ $e.GoName }}, {{ $nullShort }}.Valid = "", false
		return nil
	}
	err := {{ $nullShort }}.{{ $e.GoName }}.Scan(v)
	{{ $nullShort }}.Valid = err == nil
	return err
}

// ErrInvalid{{ $e.GoName }} is the invalid {{ $e.GoName }} error.
type ErrInvalid{{ $e.GoName }} string

// Error satisfies the error interface.
func (err ErrInvalid{{ $e.GoName }}) Error() string {
	return fmt.Sprintf("invalid {{ $e.GoName }}(%s)", string(err))
}

func All{{ $e.GoName }}Values() []{{ $e.GoName }} {
	return []{{ $e.GoName }}{ {{ range $e.Values}}{{ "\n" }}{{ $e.GoName }}{{ .GoName }},{{ end }}
	}
}

{{ end }}


{{/* generated queries from foreign keys */}}

{{ define "foreignkey" }}
{{- $k := .Data -}}
// {{ func_name_context $k "" }} returns the {{ $k.RefTable }} associated with the {{ $k.Table.GoName }}'s ({{ names "" $k.Fields }}).
//
// Generated from foreign key '{{ $k.SQLName }}'.
{{ recv_context $k.Table $k ""}} {
	return {{ foreign_key_context $k }}
}
{{ end }}

{{/*
  generated queries from indexes
*/}}

{{ define "index" }}
{{- $i := .Data.Index -}}
{{- $tables := .Data.Tables -}}
{{- $constraints := .Data.Constraints -}}
// {{ func_name_context $i "" }} retrieves a row from '{{ schema $i.Table.SQLName }}' as a {{ $i.Table.GoName }}.
//
// Generated from index '{{ $i.SQLName }}'.
{{ func_context $i "" "" }} {
	{{ initial_opts $i }}

	for _, o := range opts {
		o(c)
	}

  filters := ""

	{{ sqlstr_index $i $constraints $tables }}
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, {{ params $i.Fields false }})

{{- if $i.IsUnique }}
  rows, err := {{ db "Query" $i }}
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $i.Table.SQLName }}/{{ $i.Func }}/db.Query: %w", err))
	}
	{{ short $i.Table }}, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[{{$i.Table.GoName}}])
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $i.Table.SQLName }}/{{ $i.Func }}/pgx.CollectOneRow: %w", err))
	}

	{{- if $i.Table.PrimaryKeys }}
	{{ end }}

	return &{{ short $i.Table }}, nil
{{- else }}
	rows, err := {{ db "Query" $i }}
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $i.Table.GoName }}/{{ $i.Func }}/Query: %w", err))
	}
	defer rows.Close()
	// process
  {{/* might need to use non pointer []<st> in return if we get a NumField of non-struct type*/}}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[{{$i.Table.GoName}}])
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $i.Table.GoName }}/{{ $i.Func }}/pgx.CollectRows: %w", err))
	}
	return res, nil
{{- end }}
}

{{end}}

{{/* generated queries from stored procedures */}}

{{ define "procs" }}
{{- $ps := .Data -}}
{{- range $p := $ps -}}
// {{ func_name_context $p "" }} calls the stored {{ $p.Type }} '{{ $p.Signature }}' on db.
{{ func_context $p "" "" }} {
{{- if and (driver "mysql") (eq $p.Type "procedure") (not $p.Void) }}
	// At the moment, the Go MySQL driver does not support stored procedures
	// with out parameters
	return {{ zero $p.Returns }}, fmt.Errorf("unsupported")
{{- else }}
	// call {{ schema $p.SQLName }}
	{{ sqlstr "proc" $p }}
	// run
{{- if not $p.Void }}
{{- range $p.Returns }}
	var {{ check_name .GoName }} {{ type .Type }}
{{- end }}
	// logf(sqlstr, {{ params $p.Params false }})
	if err := {{ db "QueryRow" $p }}.Scan({{ names "&" $p.Returns }}); err != nil {
		return {{ zero $p.Returns }}, logerror(err)
	}
	return {{ range $p.Returns }}{{ check_name .GoName }}, {{ end }}nil
{{- else }}
	// logf(sqlstr)
	if _, err := {{ db "Exec" $p }}; err != nil {
		return logerror(err)
	}
	return nil
{{- end }}
{{- end }}
}

{{ if context_both -}}
// {{ func_name $p }} calls the {{ $p.Type }} '{{ $p.Signature }}' on db.
{{ func $p }} {
	return {{ func_name_context $p }}({{ names_all "" "context.Background()" "db" $p.Params }})
}
{{- end -}}
{{- end }}
{{ end }}

{{/* generated structs */}}

{{ define "typedef" }}
{{- $t := .Data.Table -}}
{{- $tables := .Data.Tables -}}
{{- $constraints := .Data.Constraints -}}

{{if $t.Comment -}}
// {{ $t.Comment | eval $t.GoName }}
{{- else -}}
// {{ $t.GoName }} represents a row from '{{ schema $t.SQLName }}'.
// Change properties via SQL column comments, joined with ",":
//     - "property:private" to exclude a field from JSON.
//     - "type:<pkg.type>" to override the type annotation.
//     - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
{{- end }}
type {{ $t.GoName }} struct {
{{ range $t.Fields -}}
	{{ field . "Table" $t -}}
{{ end }}
{{ join_fields $t $constraints $tables }}
}
{{/* NOTE: ensure sqlc does not generate clashing names */}}
// {{ $t.GoName }}CreateParams represents insert params for '{{ schema $t.SQLName }}'.
type {{ $t.GoName }}CreateParams struct {
{{ range $t.Fields -}}
	{{ field . "CreateParams" $t -}}
{{ end -}}
}

// Create{{ $t.GoName }} creates a new {{ $t.GoName }} in the database with the given params.
func Create{{ $t.GoName }}(ctx context.Context, db DB, params *{{ $t.GoName }}CreateParams) (*{{ $t.GoName }}, error) {
  {{ short $t }} := &{{ $t.GoName }}{
{{ range $t.Fields -}}
	{{ set_field . "CreateParams" $t -}}
{{ end -}}
  }

  return {{ short $t }}.Insert(ctx, db)
}


// {{ $t.GoName }}UpdateParams represents update params for '{{ schema $t.SQLName }}'
type {{ $t.GoName }}UpdateParams struct {
{{ range $t.Fields -}}
	{{ field . "UpdateParams" $t -}}
{{ end -}}
}

// SetUpdateParams updates {{ schema $t.SQLName }} struct fields with the specified params.
func ({{ short $t }} *{{ $t.GoName }}) SetUpdateParams(params *{{ $t.GoName }}UpdateParams) {
{{ range $t.Fields -}}
	{{ set_field . "UpdateParams" $t -}}
{{ end -}}
}

{{ extratypes $t.GoName $t.SQLName $constraints $t $tables }}

{{/* regular queries for a table. Ignored for mat views or views.
 */}}

{{ if $t.PrimaryKeys -}}

// {{ func_name_context "Insert" "" }} inserts the {{ $t.GoName }} to the database.
{{ recv_context $t "Insert" "" }} {
{{ if and (eq (len $t.Generated) 0) (eq (len $t.Ignored) 0) -}}
	// insert (manual)
	{{ sqlstr "insert_manual" $t }}
	// run
	{{ logf $t }}
	rows, err := {{ db_prefix "Query" false false $t }}
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $t.GoName }}/Insert/db.Query: %w", err))
	}
	new{{ short $t }}, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[{{$t.GoName}}])
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $t.GoName }}/Insert/pgx.CollectOneRow: %w", err))
	}
{{- else -}}
	// insert (primary key generated and returned by database)
	{{ sqlstr "insert" $t }}
	// run
	{{ logf $t $t.Generated $t.Ignored }}

	rows, err := {{ db_prefix "Query" false false $t }}
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $t.GoName }}/Insert/db.Query: %w", err))
	}
	new{{ short $t }}, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[{{$t.GoName}}])
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $t.GoName }}/Insert/pgx.CollectOneRow: %w", err))
	}
{{ end }}
  *{{ short $t }} = new{{ short $t }}

	return {{ short $t }}, nil
}


{{ if not_updatable $t.Fields -}}
// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------
{{- else -}}
// {{ func_name_context "Update" "" }} updates a {{ $t.GoName }} in the database.
{{ recv_context $t "Update" "" }}  {
	// update with {{ if driver "postgres" }}composite {{ end }}primary key
	{{ sqlstr "update" $t }}
	// run
	{{ logf_update $t }}

  rows, err := {{ db_update "Query" $t }}
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $t.GoName }}/Update/db.Query: %w", err))
	}
	new{{ short $t }}, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[{{$t.GoName}}])
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $t.GoName }}/Update/pgx.CollectOneRow: %w", err))
	}
  *{{ short $t }} = new{{ short $t }}

	return {{ short $t }}, nil
}


// {{ func_name_context "Upsert" "" }} upserts a {{ $t.GoName }} in the database.
// Requires appropiate PK(s) to be set beforehand.
{{ recv_context $t "Upsert" "" }}  {
	var err error

  {{ range $t.Fields -}}
    {{ set_field . "UpsertParams" $t -}}
  {{ end }}

  {{ short $t }}, err = {{ short $t }}.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
			  return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
		  {{ short $t }}, err = {{ short $t }}.Update(ctx, db)
      if err != nil {
			  return nil, fmt.Errorf("UpsertUser/Update: %w", err)
      }
		}
	}

  return {{ short $t }}, err
}

{{- end }}

// {{ func_name_context "Delete" "" }} deletes the {{ $t.GoName }} from the database.
{{ recv_context $t "Delete" "" }} {
{{ if eq (len $t.PrimaryKeys) 1 -}}
	// delete with single primary key
	{{ sqlstr "delete" $t }}
	// run
	if _, err := {{ db "Exec" (print (short $t) "." (index $t.PrimaryKeys 0).GoName) }}; err != nil {
		return logerror(err)
	}
{{- else -}}
	// delete with composite primary key
	{{ sqlstr "delete" $t }}
	// run
	if _, err := {{ db "Exec" (names (print (short $t) ".") $t.PrimaryKeys) }}; err != nil {
		return logerror(err)
	}
{{- end }}
	return nil
}

{{- end }}

{{ if (has_deleted_at $t) }}
// {{ func_name_context "SoftDelete" "" }} soft deletes the {{ $t.GoName }} from the database via 'deleted_at'.
{{ recv_context $t "SoftDelete" "" }} {
	{{ if eq (len $t.PrimaryKeys) 1 -}}
	// delete with single primary key
	{{ sqlstr "soft_delete" $t }}
	// run
	if _, err := {{ db "Exec" (print (short $t) "." (index $t.PrimaryKeys 0).GoName) }}; err != nil {
		return logerror(err)
	}
  {{- else -}}
	// delete with composite primary key
	{{ sqlstr "soft_delete" $t }}
	// run
	if _, err := {{ db "Exec" (names (print (short $t) ".") $t.PrimaryKeys) }}; err != nil {
		return logerror(err)
	}
  {{ end }}
	// set deleted
  {{ short $t }}.DeletedAt = newPointer(time.Now())

	return nil
}

// {{ func_name_context "Restore" "" }} restores a soft deleted {{ $t.GoName }} from the database.
{{ recv_context $t "Restore" "" }} {
	{{ short $t }}.DeletedAt = nil
	new{{ short $t }}, err:= {{ short $t }}.Update(ctx,db)
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $t.GoName }}/Restore/pgx.CollectRows: %w", err))
	}
	return new{{ short $t }}, nil
}

{{ end }}

{{ range $order := combine_values "Asc" "Desc" }}
{{ range $cursor := cursor_columns $t $constraints $tables }}
{{ $suffix := print "PaginatedBy" (fields_to_goname $cursor "") $order }}
// {{ func_name_context $t $suffix }} returns a cursor-paginated list of {{ $t.GoName }} in {{ $order }} order.
{{ func_context $t $suffix $cursor }} {
	{{ initial_opts $t }}

	for _, o := range opts {
		o(c)
	}

  filters := ""

	{{ sqlstr_paginated $t $constraints $tables $cursor $order }}
	sqlstr += c.limit

	// run

	rows, err := {{ db_paginated "Query" $t $cursor }}
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $t.GoName }}/Paginated/{{ $order }}/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[{{$t.GoName}}])
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $t.GoName }}/Paginated/{{ $order }}/pgx.CollectRows: %w", err))
	}
	return res, nil
}
{{ end }}
{{ end }}

{{ end }}
