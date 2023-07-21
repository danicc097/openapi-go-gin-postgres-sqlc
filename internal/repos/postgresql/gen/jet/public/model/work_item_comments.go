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

type WorkItemComments struct {
	WorkItemCommentID int64     `sql:"primary_key" db:"work_item_comment_id"`
	WorkItemID        int64     `db:"work_item_id"`
	UserID            uuid.UUID `db:"user_id"`
	Message           string    `db:"message"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}
