{{ define "extra" -}}

func newPointer[T any](v T) *T {
	return &v
}

{{- end }}
