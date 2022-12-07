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

var Projects = newProjectsTable("public", "projects", "")

type projectsTable struct {
	postgres.Table

	//Columns
	ProjectID          postgres.ColumnInteger
	Name               postgres.ColumnString
	Description        postgres.ColumnString
	WorkItemsTableName postgres.ColumnString
	CreatedAt          postgres.ColumnTimestampz
	UpdatedAt          postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ProjectsTable struct {
	projectsTable

	EXCLUDED projectsTable
}

// AS creates new ProjectsTable with assigned alias
func (a ProjectsTable) AS(alias string) *ProjectsTable {
	return newProjectsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ProjectsTable with assigned schema name
func (a ProjectsTable) FromSchema(schemaName string) *ProjectsTable {
	return newProjectsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ProjectsTable with assigned table prefix
func (a ProjectsTable) WithPrefix(prefix string) *ProjectsTable {
	return newProjectsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ProjectsTable with assigned table suffix
func (a ProjectsTable) WithSuffix(suffix string) *ProjectsTable {
	return newProjectsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newProjectsTable(schemaName, tableName, alias string) *ProjectsTable {
	return &ProjectsTable{
		projectsTable: newProjectsTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newProjectsTableImpl("", "excluded", ""),
	}
}

func newProjectsTableImpl(schemaName, tableName, alias string) projectsTable {
	var (
		ProjectIDColumn          = postgres.IntegerColumn("project_id")
		NameColumn               = postgres.StringColumn("name")
		DescriptionColumn        = postgres.StringColumn("description")
		WorkItemsTableNameColumn = postgres.StringColumn("work_items_table_name")
		CreatedAtColumn          = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn          = postgres.TimestampzColumn("updated_at")
		allColumns               = postgres.ColumnList{ProjectIDColumn, NameColumn, DescriptionColumn, WorkItemsTableNameColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns           = postgres.ColumnList{NameColumn, DescriptionColumn, WorkItemsTableNameColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return projectsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ProjectID:          ProjectIDColumn,
		Name:               NameColumn,
		Description:        DescriptionColumn,
		WorkItemsTableName: WorkItemsTableNameColumn,
		CreatedAt:          CreatedAtColumn,
		UpdatedAt:          UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
