//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
)

type WorkItemAssignee struct {
	WorkItemID int64        `sql:"primary_key" db:"work_item_id"`
	Assignee   uuid.UUID    `sql:"primary_key" db:"assignee"`
	Role       WorkItemRole `db:"role"`
}
