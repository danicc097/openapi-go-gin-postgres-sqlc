//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type EntityNotifications struct {
	EntityNotificationID int32     `sql:"primary_key" db:"entity_notification_id"`
	Entity               string    `db:"entity"`
	ID                   string    `db:"id"`
	Message              string    `db:"message"`
	Topic                string    `db:"topic"`
	CreatedAt            time.Time `db:"created_at"`
}
