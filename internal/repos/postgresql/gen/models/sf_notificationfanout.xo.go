// Code generated by xo. DO NOT EDIT.

//lint:ignore

package models

import (
	"context"
)

// NotificationFanOut calls the stored function 'public.notification_fan_out() trigger' on db.
func NotificationFanOut(ctx context.Context, db DB) (Trigger, error) {
	// call public.notification_fan_out
	sqlstr := `SELECT * FROM public.notification_fan_out() `
	// run
	var r0 Trigger
	// logf(sqlstr, )
	if err := db.QueryRow(ctx, sqlstr).Scan(&r0); err != nil {
		return Trigger{}, logerror(err)
	}
	return r0, nil
}
