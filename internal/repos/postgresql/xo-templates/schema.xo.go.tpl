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
// {{ func_name_context $k }} returns the {{ $k.RefTable }} associated with the {{ $k.Table.GoName }}'s ({{ names "" $k.Fields }}).
//
// Generated from foreign key '{{ $k.SQLName }}'.
{{ recv_context $k.Table $k }} {
	return {{ foreign_key_context $k }}
}
{{- if context_both }}

// {{ func_name $k }} returns the {{ $k.RefTable }} associated with the {{ $k.Table }}'s ({{ names "" $k.Fields }}).
//
// Generated from foreign key '{{ $k.SQLName }}'.
{{ recv $k.Table $k }} {
	return {{ foreign_key $k }}
}
{{- end }}
{{ end }}

{{/*
generated queries from indexes
*/}}

{{ define "index" }}
{{- $i := .Data.Index -}}
{{- $constraints := .Data.Constraints -}}
// {{ func_name_context $i }} retrieves a row from '{{ schema $i.Table.SQLName }}' as a {{ $i.Table.GoName }}.
//
// Generated from index '{{ $i.SQLName }}'.
{{ func_context $i }} {
	{{ initial_opts $i }}

	for _, o := range opts {
		o(c)
	}

	// query
	{{ sqlstr_index $i $constraints }}
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, {{ params $i.Fields false }})
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
  {{ short $i.Table }}._exists = true
	{{ end -}}

	return &{{ short $i.Table }}, nil
{{- else }}
	rows, err := {{ db "Query" $i }}
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
  {{/* might need to use non pointer []<st> in return if we get a NumField of non-struct type*/}}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[*{{$i.Table.GoName}}])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
{{- end }}
}

{{end}}

{{/* generated queries from stored procedures */}}

{{ define "procs" }}
{{- $ps := .Data -}}
{{- range $p := $ps -}}
// {{ func_name_context $p }} calls the stored {{ $p.Type }} '{{ $p.Signature }}' on db.
{{ func_context $p }} {
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
	logf(sqlstr, {{ params $p.Params false }})
	if err := {{ db "QueryRow" $p }}.Scan({{ names "&" $p.Returns }}); err != nil {
		return {{ zero $p.Returns }}, logerror(err)
	}
	return {{ range $p.Returns }}{{ check_name .GoName }}, {{ end }}nil
{{- else }}
	logf(sqlstr)
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
{{- $constraints := .Data.Constraints -}}

{{if $t.Comment -}}
// {{ $t.Comment | eval $t.GoName }}
{{- else -}}
// {{ $t.GoName }} represents a row from '{{ schema $t.SQLName }}'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
{{- end }}
type {{ $t.GoName }} struct {
{{ range $t.Fields -}}
	{{ field . "Table" $t -}}
{{ end }}
{{ join_fields $t.SQLName false $constraints }}
{{- if $t.PrimaryKeys -}}
	// xo fields
	_exists, _deleted bool
{{ end -}}
}
{{/* NOTE: using this notation since sqlc uses
<query_name>{{ $t.GoName }}Params which may clash if we also were to
call it Create or Update*/}}
// {{ $t.GoName }}CreateParams represents insert params for '{{ schema $t.SQLName }}'
type {{ $t.GoName }}CreateParams struct {
{{ range $t.Fields -}}
	{{ field . "CreateParams" $t -}}
{{ end -}}
}

// {{ $t.GoName }}UpdateParams represents update params for '{{ schema $t.SQLName }}'
type {{ $t.GoName }}UpdateParams struct {
{{ range $t.Fields -}}
	{{ field . "UpdateParams" $t -}}
{{ end -}}
}

{{ extratypes $t.GoName $t.SQLName $constraints $t }}

{{/* regular queries for a table. Ignored for mat views or views.
 */}}

{{ if $t.PrimaryKeys -}}
// Exists returns true when the {{ $t.GoName }} exists in the database.
func ({{ short $t }} *{{ $t.GoName }}) Exists() bool {
	return {{ short $t }}._exists
}

// Deleted returns true when the {{ $t.GoName }} has been marked for deletion from
// the database.
func ({{ short $t }} *{{ $t.GoName }}) Deleted() bool {
	return {{ short $t }}._deleted
}

// {{ func_name_context "Insert" }} inserts the {{ $t.GoName }} to the database.
{{ recv_context $t "Insert" }} {
	switch {
	case {{ short $t }}._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case {{ short $t }}._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
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
{{- end }}
	new{{ short $t }}._exists = true
  *{{ short $t }} = new{{ short $t }}

	return {{ short $t }}, nil
}

{{ if context_both -}}
// Insert inserts the {{ $t.GoName }} to the database.
{{ recv $t "Insert" }} {
	return {{ short $t }}.InsertContext(context.Background(), db)
}
{{- end }}


{{ if not_updatable $t.Fields -}}
// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------
{{- else -}}
// {{ func_name_context "Update" }} updates a {{ $t.GoName }} in the database.
{{ recv_context $t "Update" }}  {
	switch {
	case !{{ short $t }}._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case {{ short $t }}._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
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
  new{{ short $t }}._exists = true
  *{{ short $t }} = new{{ short $t }}

	return {{ short $t }}, nil
}

{{ if context_both -}}
// Update updates a {{ $t.GoName }} in the database.
{{ recv $t "Update" }} {
	return {{ short $t }}.UpdateContext(context.Background(), db)
}
{{- end }}

// {{ func_name_context "Save" }} saves the {{ $t.GoName }} to the database.
{{ recv_context $t "Save" }} {
	if {{ short $t }}.Exists() {
		return {{ short $t }}.{{ func_name_context "Update" }}({{ if context }}ctx, {{ end }}db)
	}
	return {{ short $t }}.{{ func_name_context "Insert" }}({{ if context }}ctx, {{ end }}db)
}

{{ if context_both -}}
// Save saves the {{ $t.GoName }} to the database.
{{ recv $t "Save" }} {
	if {{ short $t }}._exists {
		return {{ short $t }}.UpdateContext(context.Background(), db)
	}
	return {{ short $t }}.InsertContext(context.Background(), db)
}
{{- end }}

// {{ func_name_context "Upsert" }} performs an upsert for {{ $t.GoName }}.
{{ recv_context $t "Upsert" }} {
	switch {
	case {{ short $t }}._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	{{ sqlstr "upsert" $t }}
	// run
	{{ logf $t $t.Ignored }}{{/* upsert will require generated fields, but exclude ignored fields */}}
	if _, err := {{ db_prefix "Exec" true false $t }}; err != nil {
		return logerror(err)
	}
	// set exists
	{{ short $t }}._exists = true
	return nil
}

{{ if context_both -}}
// Upsert performs an upsert for {{ $t.GoName }}.
{{ recv $t "Upsert" }} {
	return {{ short $t }}.UpsertContext(context.Background(), db)
}
{{- end -}}
{{- end }}

// {{ func_name_context "Delete" }} deletes the {{ $t.GoName }} from the database.
{{ recv_context $t "Delete" }} {
	switch {
	case !{{ short $t }}._exists: // doesn't exist
		return nil
	case {{ short $t }}._deleted: // deleted
		return nil
	}
{{ if eq (len $t.PrimaryKeys) 1 -}}
	// delete with single primary key
	{{ sqlstr "delete" $t }}
	// run
	{{ logf_pkeys $t }}
	if _, err := {{ db "Exec" (print (short $t) "." (index $t.PrimaryKeys 0).GoName) }}; err != nil {
		return logerror(err)
	}
{{- else -}}
	// delete with composite primary key
	{{ sqlstr "delete" $t }}
	// run
	{{ logf_pkeys $t }}
	if _, err := {{ db "Exec" (names (print (short $t) ".") $t.PrimaryKeys) }}; err != nil {
		return logerror(err)
	}
{{- end }}
	// set deleted
	{{ short $t }}._deleted = true
	return nil
}

{{ if context_both -}}
// Delete deletes the {{ $t.GoName }} from the database.
{{ recv $t "Delete" }} {
	return {{ short $t }}.DeleteContext(context.Background(), db)
}
{{- end -}}
{{- end }}
{{ end }}
