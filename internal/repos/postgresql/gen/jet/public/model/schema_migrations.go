//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type SchemaMigrations struct {
	Version int64 `sql:"primary_key" db:"version"`
	Dirty   bool  `db:"dirty"`
}
