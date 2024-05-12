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
{{- $t := $i.Table -}}
{{- $tables := .Data.Tables -}}
{{- $constraints := .Data.Constraints -}}
{{/* TODO: maybe can be init beforehand */}}
{{- $_ := initialize_constraints $t $constraints }}
// {{ func_name_context $i "" }} retrieves a row from '{{ schema $t.SQLName }}' as a {{$t.GoName}}.
//
// Generated from index '{{ $i.SQLName }}'.
{{ func_context $i "" "" $t "" }} {
	{{ initial_opts $i }}

	for _, o := range opts {
		o(c)
	}

  paramStart := {{ last_nth $i $tables }}
	nth := func ()  string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i"){
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND "+strings.Join(filterClauses, " AND ")+" "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

  orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

	{{ sqlstr_index $i $tables }}
	sqlstr += orderBy
	sqlstr += c.limit
  sqlstr = "/* {{ func_name_context $i "" }} */\n"+sqlstr

	// run
	// logf(sqlstr, {{ params $i.Fields false $t }})

{{- if $i.IsUnique }}
  rows, err := {{ db "Query" $i }}
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $t.SQLName }}/{{ $i.Func }}/db.Query: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err }))
	}
	{{ short $t }}, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[{{$t.GoName}}])
	if err != nil {
		return nil, logerror(fmt.Errorf("{{ $t.SQLName }}/{{ $i.Func }}/pgx.CollectOneRow: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err }))
	}

	{{- if $t.PrimaryKeys }}
	{{ end }}

	return &{{ short $t }}, nil
{{- else }}
	rows, err := {{ db "Query" $i }}
	if err != nil {
		return nil, logerror(fmt.Errorf("{{$t.GoName}}/{{ $i.Func }}/Query: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err }))
	}
	defer rows.Close()
	// process
  {{/* might need to use non pointer []<st> in return if we get a NumField of non-struct type*/}}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[{{$t.GoName}}])
	if err != nil {
		return nil, logerror(fmt.Errorf("{{$t.GoName}}/{{ $i.Func }}/pgx.CollectRows: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err }))
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
{{ func_context $p "" "" "" "" }} {
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
	// logf(sqlstr, {{ params $p.Params false "" }})
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
{{- $schema := .Data.Schema -}}
{{- $constraints := .Data.Constraints -}}
{{/* TODO: maybe can be init beforehand */}}
{{- $_ := initialize_constraints $t $constraints }}

{{if $t.Comment -}}
// {{ $t.Comment | eval $t.GoName }}
{{- else -}}
// {{$t.GoName}} represents a row from '{{ schema $t.SQLName }}'.
{{- end }}
type {{$t.GoName}} struct {
{{ range $t.Fields -}}
	{{ field . "Table" $t -}}
{{ end }}
{{ join_fields $t $constraints $tables }}
}
{{/* NOTE: ensure sqlc does not generate clashing names */}}

{{/* create params and helper only if theres PKs */}}
{{ if $t.PrimaryKeys -}}

// {{$t.GoName}}CreateParams represents insert params for '{{ schema $t.SQLName }}'.
type {{$t.GoName}}CreateParams struct {
{{ range sort_fields $t.Fields -}}
	{{ field . "CreateParams" $t -}}
{{ end -}}
}

// {{$t.GoName}}Params represents common params for both insert and update of '{{ schema $t.SQLName }}'.
type {{$t.GoName}}Params interface {
{{ range sort_fields $t.Fields -}}
	{{ field . "ParamsInterface" $t -}}
{{ end -}}
}

{{ range sort_fields $t.Fields -}}
	{{ field . "ParamsGetter" $t -}} {{/* will create getter for both create and update structs */}}
{{ end -}}

{{ range sort_fields $t.Fields -}}
	{{ field . "IDTypes" $t -}}
{{ end -}}

// Create{{$t.GoName}} creates a new {{$t.GoName}} in the database with the given params.
func Create{{$t.GoName}}(ctx context.Context, db DB, params *{{$t.GoName}}CreateParams) (*{{$t.GoName}}, error) {
  {{ short $t }} := &{{$t.GoName}}{
{{ range $t.Fields -}}
	{{ set_field . "CreateParams" $t -}}
{{ end -}}
  }

  return {{ short $t }}.Insert(ctx, db)
}
{{ end -}}

{{ extratypes $t.GoName $t.SQLName $constraints $t $tables }}

{{/* regular queries for a table. Ignored for mat views or views.
 */}}

{{ if $t.PrimaryKeys -}}

// {{$t.GoName}}UpdateParams represents update params for '{{ schema $t.SQLName }}'.
type {{$t.GoName}}UpdateParams struct {
{{ range sort_fields $t.Fields -}}
	{{ field . "UpdateParams" $t -}}
{{ end -}}
}

// SetUpdateParams updates {{ schema $t.SQLName }} struct fields with the specified params.
func ({{ short $t }} *{{$t.GoName}}) SetUpdateParams(params *{{$t.GoName}}UpdateParams) {
{{ range $t.Fields -}}
	{{ set_field . "UpdateParams" $t -}}
{{ end -}}
}

// {{ func_name_context "Insert" "" }} inserts the {{$t.GoName}} to the database.
{{ recv_context $t "Insert" "" }} {
{{ if and (eq (len $t.Generated) 0) (eq (len $t.Ignored) 0) -}}
	// insert (manual)
	{{ sqlstr "insert_manual" $t }}
	// run
	{{ logf $t }}
	rows, err := {{ db_prefix "Query" false false $t }}
	if err != nil {
		return nil, logerror(fmt.Errorf("{{$t.GoName}}/Insert/db.Query: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err }))
	}
	new{{ short $t }}, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[{{$t.GoName}}])
	if err != nil {
		return nil, logerror(fmt.Errorf("{{$t.GoName}}/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err }))
	}
{{- else -}}
	// insert (primary key generated and returned by database)
	{{ sqlstr "insert" $t }}
	// run
	{{ logf $t $t.Generated $t.Ignored }}

	rows, err := {{ db_prefix "Query" false false $t }}
	if err != nil {
		return nil, logerror(fmt.Errorf("{{$t.GoName}}/Insert/db.Query: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err }))
	}
	new{{ short $t }}, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[{{$t.GoName}}])
	if err != nil {
		return nil, logerror(fmt.Errorf("{{$t.GoName}}/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err }))
	}
{{ end }}
  *{{ short $t }} = new{{ short $t }}

	return {{ short $t }}, nil
}


{{ if not (table_is_updatable $t.Fields) -}}
// ------ NOTE: Update statements omitted due to lack of fields other than primary key or generated fields
{{- else -}}
// {{ func_name_context "Update" "" }} updates a {{$t.GoName}} in the database.
{{ recv_context $t "Update" "" }}  {
	// update with {{ if driver "postgres" }}composite {{ end }}primary key
	{{ sqlstr "update" $t }}
	// run
	{{ logf_update $t }}

  rows, err := {{ db_update "Query" $t }}
	if err != nil {
		return nil, logerror(fmt.Errorf("{{$t.GoName}}/Update/db.Query: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err }))
	}
	new{{ short $t }}, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[{{$t.GoName}}])
	if err != nil {
		return nil, logerror(fmt.Errorf("{{$t.GoName}}/Update/pgx.CollectOneRow: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err }))
	}
  *{{ short $t }} = new{{ short $t }}

	return {{ short $t }}, nil
}


// {{ func_name_context "Upsert" "" }} upserts a {{$t.GoName}} in the database.
// Requires appropriate PK(s) to be set beforehand.
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
			  return nil, fmt.Errorf("Upsert{{$t.GoName}}/Insert: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err })
			}
		  {{ short $t }}, err = {{ short $t }}.Update(ctx, db)
      if err != nil {
			  return nil, fmt.Errorf("Upsert{{$t.GoName}}/Update: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err })
      }
		}
	}

  return {{ short $t }}, err
}

{{- end }}

// {{ func_name_context "Delete" "" }} deletes the {{$t.GoName}} from the database.
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

{{ if $t.PrimaryKeys -}}
{{ if and (has_deleted_at $t) (table_is_updatable $t.Fields) }}

// {{ func_name_context "SoftDelete" "" }} soft deletes the {{$t.GoName}} from the database via 'deleted_at'.
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

// {{ func_name_context "Restore" "" }} restores a soft deleted {{$t.GoName}} from the database.
{{ recv_context $t "Restore" "" }} {
	{{ short $t }}.DeletedAt = nil
	new{{ short $t }}, err:= {{ short $t }}.Update(ctx,db)
	if err != nil {
		return nil, logerror(fmt.Errorf("{{$t.GoName}}/Restore/pgx.CollectRows: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err }))
	}
	return new{{ short $t }}, nil
}

{{ end }}
{{ end }}

{{ $suffix := print "Paginated" }}
// {{ func_name_context $t $suffix }} returns a cursor-paginated list of {{$t.GoName}}.
// At least one cursor is required.
{{ func_context $t $suffix "" $t "cursor PaginationCursor" }} {
	{{ initial_opts $t }}

	for _, o := range opts {
		o(c)
	}

  if cursor.Value == nil {
    {{/* last/first (desc/asc) element is meant to be queried automatically if cursor was nil, and set thereafter.
    this is necessary when Infinity cannot be used, etc. */}}
    return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
  }
  {{if not (eq $schema "public") -}}
  field, ok := {{camel_export $schema}}EntityFields[{{camel_export $schema}}TableEntity{{$t.GoName}}][cursor.Column]
  {{else -}}
  field, ok := EntityFields[TableEntity{{$t.GoName}}][cursor.Column]
  {{end -}}
  if !ok {
    return nil, logerror(fmt.Errorf("{{$t.GoName}}/Paginated/cursor: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
  }

  op := "<"
  if cursor.Direction == DirectionAsc {
    op = ">"
  }
  c.filters[fmt.Sprintf("{{$t.SQLName}}.%s %s $i", field.Db, op)] = []any{*cursor.Value}
  c.orderBy[field.Db] = cursor.Direction // no need to duplicate opts

  paramStart := 0 // all filters will come from the user
	nth := func ()  string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i"){
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters += " where " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

  orderByClause := ""
	if len(c.orderBy) > 0 {
		orderByClause += " order by "
	} else {
		return nil, logerror(fmt.Errorf("{{$t.GoName}}/Paginated/orderBy: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: fmt.Errorf("at least one sorted column is required")}))
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderByClause += " " + strings.Join(orderBys, ", ") + " "

	{{ sqlstr_paginated $t $tables }}
	sqlstr += c.limit
  sqlstr = "/* {{ func_name_context $t $suffix }} */\n"+sqlstr

	// run

	rows, err := {{ db_paginated "Query" $t }}
	if err != nil {
		return nil, logerror(fmt.Errorf("{{$t.GoName}}/Paginated/db.Query: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err }))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[{{$t.GoName}}])
	if err != nil {
		return nil, logerror(fmt.Errorf("{{$t.GoName}}/Paginated/pgx.CollectRows: %w", &XoError{Entity: "{{ sentence_case $t.SQLName }}", Err: err }))
	}
	return res, nil
}

{{ end }}
