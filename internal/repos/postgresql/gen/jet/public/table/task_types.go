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

var TaskTypes = newTaskTypesTable("public", "task_types", "")

type taskTypesTable struct {
	postgres.Table

	//Columns
	TaskTypeID  postgres.ColumnInteger
	TeamID      postgres.ColumnInteger
	Name        postgres.ColumnString
	Description postgres.ColumnString
	Color       postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TaskTypesTable struct {
	taskTypesTable

	EXCLUDED taskTypesTable
}

// AS creates new TaskTypesTable with assigned alias
func (a TaskTypesTable) AS(alias string) *TaskTypesTable {
	return newTaskTypesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TaskTypesTable with assigned schema name
func (a TaskTypesTable) FromSchema(schemaName string) *TaskTypesTable {
	return newTaskTypesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TaskTypesTable with assigned table prefix
func (a TaskTypesTable) WithPrefix(prefix string) *TaskTypesTable {
	return newTaskTypesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TaskTypesTable with assigned table suffix
func (a TaskTypesTable) WithSuffix(suffix string) *TaskTypesTable {
	return newTaskTypesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTaskTypesTable(schemaName, tableName, alias string) *TaskTypesTable {
	return &TaskTypesTable{
		taskTypesTable: newTaskTypesTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newTaskTypesTableImpl("", "excluded", ""),
	}
}

func newTaskTypesTableImpl(schemaName, tableName, alias string) taskTypesTable {
	var (
		TaskTypeIDColumn  = postgres.IntegerColumn("task_type_id")
		TeamIDColumn      = postgres.IntegerColumn("team_id")
		NameColumn        = postgres.StringColumn("name")
		DescriptionColumn = postgres.StringColumn("description")
		ColorColumn       = postgres.StringColumn("color")
		allColumns        = postgres.ColumnList{TaskTypeIDColumn, TeamIDColumn, NameColumn, DescriptionColumn, ColorColumn}
		mutableColumns    = postgres.ColumnList{TeamIDColumn, NameColumn, DescriptionColumn, ColorColumn}
	)

	return taskTypesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		TaskTypeID:  TaskTypeIDColumn,
		TeamID:      TeamIDColumn,
		Name:        NameColumn,
		Description: DescriptionColumn,
		Color:       ColorColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
