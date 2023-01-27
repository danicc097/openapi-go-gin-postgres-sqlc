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

var WorkItemMember = newWorkItemMemberTable("public", "work_item_member", "")

type workItemMemberTable struct {
	postgres.Table

	//Columns
	WorkItemID postgres.ColumnInteger
	Member     postgres.ColumnString
	Role       postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type WorkItemMemberTable struct {
	workItemMemberTable

	EXCLUDED workItemMemberTable
}

// AS creates new WorkItemMemberTable with assigned alias
func (a WorkItemMemberTable) AS(alias string) *WorkItemMemberTable {
	return newWorkItemMemberTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new WorkItemMemberTable with assigned schema name
func (a WorkItemMemberTable) FromSchema(schemaName string) *WorkItemMemberTable {
	return newWorkItemMemberTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new WorkItemMemberTable with assigned table prefix
func (a WorkItemMemberTable) WithPrefix(prefix string) *WorkItemMemberTable {
	return newWorkItemMemberTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new WorkItemMemberTable with assigned table suffix
func (a WorkItemMemberTable) WithSuffix(suffix string) *WorkItemMemberTable {
	return newWorkItemMemberTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newWorkItemMemberTable(schemaName, tableName, alias string) *WorkItemMemberTable {
	return &WorkItemMemberTable{
		workItemMemberTable: newWorkItemMemberTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newWorkItemMemberTableImpl("", "excluded", ""),
	}
}

func newWorkItemMemberTableImpl(schemaName, tableName, alias string) workItemMemberTable {
	var (
		WorkItemIDColumn = postgres.IntegerColumn("work_item_id")
		MemberColumn     = postgres.StringColumn("member")
		RoleColumn       = postgres.StringColumn("role")
		allColumns       = postgres.ColumnList{WorkItemIDColumn, MemberColumn, RoleColumn}
		mutableColumns   = postgres.ColumnList{RoleColumn}
	)

	return workItemMemberTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		WorkItemID: WorkItemIDColumn,
		Member:     MemberColumn,
		Role:       RoleColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
