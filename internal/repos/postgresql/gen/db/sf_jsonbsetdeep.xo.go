package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// JsonbSetDeep calls the stored function 'public.jsonb_set_deep(jsonb, text, jsonb) jsonb' on db.
func JsonbSetDeep(ctx context.Context, db DB, target map[string]any, path []string, val map[string]any) (map[string]any, error) {
	// call public.jsonb_set_deep
	sqlstr := `SELECT * FROM public.jsonb_set_deep($1, $2, $3) `
	// run
	var r0 map[string]any
	// logf(sqlstr, target, path, val)
	if err := db.QueryRow(ctx, sqlstr, target, path, val).Scan(&r0); err != nil {
		return map[string]any{}, logerror(err)
	}
	return r0, nil
}