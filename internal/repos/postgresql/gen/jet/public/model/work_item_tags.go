//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type WorkItemTags struct {
	WorkItemTagID int32  `sql:"primary_key" db:"work_item_tag_id"`
	ProjectID     int32  `db:"project_id"`
	Name          string `db:"name"`
	Description   string `db:"description"`
	Color         string `db:"color"`
}
