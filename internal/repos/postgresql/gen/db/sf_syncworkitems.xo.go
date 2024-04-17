// Code generated by xo. DO NOT EDIT.

//lint:ignore

package db

import (
	"context"
)

// SyncWorkItems calls the stored function 'public.sync_work_items() trigger' on db.
func SyncWorkItems(ctx context.Context, db DB) (Trigger, error) {
	// call public.sync_work_items
	sqlstr := `SELECT * FROM public.sync_work_items() `
	// run
	var r0 Trigger
	// logf(sqlstr, )
	if err := db.QueryRow(ctx, sqlstr).Scan(&r0); err != nil {
		return Trigger{}, logerror(err)
	}
	return r0, nil
}
