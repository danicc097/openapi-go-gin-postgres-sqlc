{{ define "extra" -}}
{{- $tables := .Data.Tables -}}

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

{{/* entities $tables */}}

{{/* TODO: initialize f.entityFilters via code for all tables. we already have $tables here... */}}
{{/* else we dont get all entities */}}
{{ generate_entity_filters $tables }}

{{- end }}

