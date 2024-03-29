//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Activities = newActivitiesTable("public", "activities", "")

type activitiesTable struct {
	postgres.Table

	// Columns
	ActivityID   postgres.ColumnInteger
	ProjectID    postgres.ColumnInteger
	Name         postgres.ColumnString
	Description  postgres.ColumnString
	IsProductive postgres.ColumnBool
	DeletedAt    postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ActivitiesTable struct {
	activitiesTable

	EXCLUDED activitiesTable
}

// AS creates new ActivitiesTable with assigned alias
func (a ActivitiesTable) AS(alias string) *ActivitiesTable {
	return newActivitiesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ActivitiesTable with assigned schema name
func (a ActivitiesTable) FromSchema(schemaName string) *ActivitiesTable {
	return newActivitiesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ActivitiesTable with assigned table prefix
func (a ActivitiesTable) WithPrefix(prefix string) *ActivitiesTable {
	return newActivitiesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ActivitiesTable with assigned table suffix
func (a ActivitiesTable) WithSuffix(suffix string) *ActivitiesTable {
	return newActivitiesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newActivitiesTable(schemaName, tableName, alias string) *ActivitiesTable {
	return &ActivitiesTable{
		activitiesTable: newActivitiesTableImpl(schemaName, tableName, alias),
		EXCLUDED:        newActivitiesTableImpl("", "excluded", ""),
	}
}

func newActivitiesTableImpl(schemaName, tableName, alias string) activitiesTable {
	var (
		ActivityIDColumn   = postgres.IntegerColumn("activity_id")
		ProjectIDColumn    = postgres.IntegerColumn("project_id")
		NameColumn         = postgres.StringColumn("name")
		DescriptionColumn  = postgres.StringColumn("description")
		IsProductiveColumn = postgres.BoolColumn("is_productive")
		DeletedAtColumn    = postgres.TimestampzColumn("deleted_at")
		allColumns         = postgres.ColumnList{ActivityIDColumn, ProjectIDColumn, NameColumn, DescriptionColumn, IsProductiveColumn, DeletedAtColumn}
		mutableColumns     = postgres.ColumnList{ProjectIDColumn, NameColumn, DescriptionColumn, IsProductiveColumn, DeletedAtColumn}
	)

	return activitiesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ActivityID:   ActivityIDColumn,
		ProjectID:    ProjectIDColumn,
		Name:         NameColumn,
		Description:  DescriptionColumn,
		IsProductive: IsProductiveColumn,
		DeletedAt:    DeletedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
