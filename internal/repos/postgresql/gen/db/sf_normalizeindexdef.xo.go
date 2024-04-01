package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// NormalizeIndexDef calls the stored function 'public.normalize_index_def(text) text' on db.
func NormalizeIndexDef(ctx context.Context, db DB, inputString string) (string, error) {
	// call public.normalize_index_def
	sqlstr := `SELECT * FROM public.normalize_index_def($1) `
	// run
	var r0 string
	// logf(sqlstr, inputString)
	if err := db.QueryRow(ctx, sqlstr, inputString).Scan(&r0); err != nil {
		return "", logerror(err)
	}
	return r0, nil
}
