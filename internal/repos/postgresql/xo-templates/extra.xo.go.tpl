{{ define "extra" -}}
{{- $tables := .Data.Tables -}}
{{- $schema := .Data.Schema -}}


{{ if or (eq $schema "public") }}
type Filter struct {
  // Type is one of: string, number, integer, boolean, date-time
  // Arrays and objects are ignored for default filter generation
  Type string `json:"type"`
  // Db is the corresponding db column name
  Db       string `json:"db"`
  Nullable bool   `json:"nullable"`
}

type DbField struct{}

func newPointer[T any](v T) *T {
	return &v
}

type XoError struct {
	Entity string
	Err error
}

// Error satisfies the error interface.
func (e *XoError) Error() string {
	return fmt.Sprintf("%s %v", e.Entity, e.Err)
}

// Unwrap satisfies the unwrap interface.
func (err *XoError) Unwrap() error {
	return err.Err
}
{{ end }}


{{entities $schema $tables}}

{{ generate_entity_filters $schema $tables }}

{{ generate_entity_fields $schema $tables }}

{{- end }}

