//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Activities struct {
	ActivityID   int32  `sql:"primary_key" db:"activity_id"`
	ProjectID    int32  `db:"project_id"`
	Name         string `db:"name"`
	Description  string `db:"description"`
	IsProductive bool   `db:"is_productive"`
}
