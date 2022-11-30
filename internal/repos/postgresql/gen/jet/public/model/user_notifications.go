//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type UserNotifications struct {
	UserNotificationID int64     `sql:"primary_key" db:"user_notification_id"`
	NotificationID     int32     `db:"notification_id"`
	Read               bool      `db:"read"`
	CreatedAt          time.Time `db:"created_at"`
	UserID             uuid.UUID `db:"user_id"`
}
