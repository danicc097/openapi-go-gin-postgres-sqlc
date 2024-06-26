// Code generated by xo. DO NOT EDIT.

//lint:ignore

package models

import (
	"context"
)

// CreateOrUpdateWorkItemCacheTable calls the stored function 'public.create_or_update_work_item_cache_table(text)' on db.
func CreateOrUpdateWorkItemCacheTable(ctx context.Context, db DB, projectName string) error {
	// call public.create_or_update_work_item_cache_table
	sqlstr := `SELECT * FROM public.create_or_update_work_item_cache_table($1) `
	// run
	// logf(sqlstr)
	if _, err := db.Exec(ctx, sqlstr, projectName); err != nil {
		return logerror(err)
	}
	return nil
}
