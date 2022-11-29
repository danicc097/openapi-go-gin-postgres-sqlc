package postgres

import "github.com/go-jet/jet/v2/internal/jet"

// RawStatement creates new sql statements from raw query and optional map of named arguments
func RawStatement(rawQuery string, namedArguments ...RawArgs) Statement {
	return jet.RawStatement(Dialect, rawQuery, namedArguments...)
}
