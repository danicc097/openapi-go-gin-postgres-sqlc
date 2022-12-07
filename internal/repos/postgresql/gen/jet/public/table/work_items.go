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

var WorkItems = newWorkItemsTable("public", "work_items", "")

type workItemsTable struct {
	postgres.Table

	//Columns
	WorkItemID     postgres.ColumnInteger
	Title          postgres.ColumnString
	WorkItemTypeID postgres.ColumnInteger
	Metadata       postgres.ColumnString
	TeamID         postgres.ColumnInteger
	KanbanStepID   postgres.ColumnInteger
	Closed         postgres.ColumnTimestampz
	TargetDate     postgres.ColumnTimestampz
	CreatedAt      postgres.ColumnTimestampz
	UpdatedAt      postgres.ColumnTimestampz
	DeletedAt      postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type WorkItemsTable struct {
	workItemsTable

	EXCLUDED workItemsTable
}

// AS creates new WorkItemsTable with assigned alias
func (a WorkItemsTable) AS(alias string) *WorkItemsTable {
	return newWorkItemsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new WorkItemsTable with assigned schema name
func (a WorkItemsTable) FromSchema(schemaName string) *WorkItemsTable {
	return newWorkItemsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new WorkItemsTable with assigned table prefix
func (a WorkItemsTable) WithPrefix(prefix string) *WorkItemsTable {
	return newWorkItemsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new WorkItemsTable with assigned table suffix
func (a WorkItemsTable) WithSuffix(suffix string) *WorkItemsTable {
	return newWorkItemsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newWorkItemsTable(schemaName, tableName, alias string) *WorkItemsTable {
	return &WorkItemsTable{
		workItemsTable: newWorkItemsTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newWorkItemsTableImpl("", "excluded", ""),
	}
}

func newWorkItemsTableImpl(schemaName, tableName, alias string) workItemsTable {
	var (
		WorkItemIDColumn     = postgres.IntegerColumn("work_item_id")
		TitleColumn          = postgres.StringColumn("title")
		WorkItemTypeIDColumn = postgres.IntegerColumn("work_item_type_id")
		MetadataColumn       = postgres.StringColumn("metadata")
		TeamIDColumn         = postgres.IntegerColumn("team_id")
		KanbanStepIDColumn   = postgres.IntegerColumn("kanban_step_id")
		ClosedColumn         = postgres.TimestampzColumn("closed")
		TargetDateColumn     = postgres.TimestampzColumn("target_date")
		CreatedAtColumn      = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn      = postgres.TimestampzColumn("updated_at")
		DeletedAtColumn      = postgres.TimestampzColumn("deleted_at")
		allColumns           = postgres.ColumnList{WorkItemIDColumn, TitleColumn, WorkItemTypeIDColumn, MetadataColumn, TeamIDColumn, KanbanStepIDColumn, ClosedColumn, TargetDateColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn}
		mutableColumns       = postgres.ColumnList{TitleColumn, WorkItemTypeIDColumn, MetadataColumn, TeamIDColumn, KanbanStepIDColumn, ClosedColumn, TargetDateColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn}
	)

	return workItemsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		WorkItemID:     WorkItemIDColumn,
		Title:          TitleColumn,
		WorkItemTypeID: WorkItemTypeIDColumn,
		Metadata:       MetadataColumn,
		TeamID:         TeamIDColumn,
		KanbanStepID:   KanbanStepIDColumn,
		Closed:         ClosedColumn,
		TargetDate:     TargetDateColumn,
		CreatedAt:      CreatedAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		DeletedAt:      DeletedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
