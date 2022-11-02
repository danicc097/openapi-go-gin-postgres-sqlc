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

{{/* generated queries from indexes */}}

{{ define "index" }}
{{- $i := .Data -}}
// {{ func_name_context $i }} retrieves a row from '{{ schema $i.Table.SQLName }}' as a {{ $i.Table.GoName }}.
//
// Generated from index '{{ $i.SQLName }}'.
{{ func_context $i }} {
	// query
	{{ sqlstr "index" $i }}
	// run
	logf(sqlstr, {{ params $i.Fields false }})
{{- if $i.IsUnique }}
	{{ short $i.Table }} := {{ $i.Table.GoName }}{
	{{- if $i.Table.PrimaryKeys }}
		_exists: true,
	{{ end -}}
	}
	if err := {{ db "QueryRow"  $i }}.Scan({{ names (print "&" (short $i.Table) ".") $i.Table }}); err != nil {
		return nil, logerror(err)
	}
	return &{{ short $i.Table }}, nil
{{- else }}
	rows, err := {{ db "Query" $i }}
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*{{ $i.Table.GoName }}
	for rows.Next() {
		{{ short $i.Table }} := {{ $i.Table.GoName }}{
		{{- if $i.Table.PrimaryKeys }}
			_exists: true,
		{{ end -}}
		}
		// scan
		if err := rows.Scan({{ names_ignore (print "&" (short $i.Table) ".")  $i.Table }}); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &{{ short $i.Table }})
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
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
{{- $t := .Data -}}

type {{ $t.GoName }}SelectConfig struct {
	limit       *int
	orderBy     []{{ $t.GoName }}OrderBy
  joinWith    []{{ $t.GoName }}JoinBy
}

type {{ $t.GoName }}SelectConfigOption func(*{{ $t.GoName }}SelectConfig)


// {{ $t.GoName }}WithLimit limits row selection.
func {{ $t.GoName }}WithLimit(limit int) {{ $t.GoName }}SelectConfigOption {
	return func(s *{{ $t.GoName }}SelectConfig) {
		s.limit = &limit
	}
}

// {{ $t.GoName }}WithOrderBy orders results by the given columns.
func {{ $t.GoName }}WithOrderBy(rows ...{{ $t.GoName }}OrderBy) {{ $t.GoName }}SelectConfigOption {
	return func(s *{{ $t.GoName }}SelectConfig) {
		s.orderBy = rows
	}
}

type {{ $t.GoName }}JoinBy = string
type {{ $t.GoName }}OrderBy = string

{{ functype $t.GoName $t }}

{{/* TODO orderbys func to generate e.g. UserCreatedAtDesc (camelcased dyn.) = "created_at desc" */}}
{{if $t.Comment -}}
// {{ $t.Comment | eval $t.GoName }}
{{- else -}}
// {{ $t.GoName }} represents a row from '{{ schema $t.SQLName }}'.
{{- end }}
type {{ $t.GoName }} struct {
{{ range $t.Fields -}}
	{{ field . }}
{{ end }}
{{- if $t.PrimaryKeys -}}
	// xo fields
	_exists, _deleted bool
{{ end -}}
}

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
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case {{ short $t }}._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
{{ if eq (len $t.Generated) 0 -}}
	// insert (manual)
	{{ sqlstr "insert_manual" $t }}
	// run
	{{ logf $t }}
	if _, err := {{ db_prefix "Exec" false false $t }}; err != nil {
		return logerror(err)
	}
{{- else -}}
	// insert (primary key generated and returned by database)
	{{ sqlstr "insert" $t }}
	// run
	{{ logf $t $t.Generated $t.Ignored }}
{{ if (driver "postgres") -}}
	if err := {{ db_prefix "QueryRow" false false $t }}.Scan({{ names (print "&" (short $t) ".") $t.Generated }}); err != nil {
		return logerror(err)
	}
{{- else -}}
	res, err := {{ db_prefix "Exec" false false $t }}
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	}
{{- end -}}
{{ if not (driver "postgres") -}}
	// set primary key
	{{ short $t }}.{{ (index $t.PrimaryKeys 0).GoName }} = {{ (index $t.PrimaryKeys 0).Type }}(id)
{{- end }}
{{- end }}
	// set exists
	{{ short $t }}._exists = true
	return nil
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
{{ recv_context $t "Update" }} {
	switch {
	case !{{ short $t }}._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case {{ short $t }}._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with {{ if driver "postgres" }}composite {{ end }}primary key
	{{ sqlstr "update" $t }}
	// run
	{{ logf_update $t }}
	if _, err := {{ db_update "Exec" $t }}; err != nil {
		return logerror(err)
	}
	return nil
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
