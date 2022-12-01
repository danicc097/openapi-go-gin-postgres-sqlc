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

var UserTeam = newUserTeamTable("public", "user_team", "")

type userTeamTable struct {
	postgres.Table

	//Columns
	TeamID postgres.ColumnInteger
	UserID postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type UserTeamTable struct {
	userTeamTable

	EXCLUDED userTeamTable
}

// AS creates new UserTeamTable with assigned alias
func (a UserTeamTable) AS(alias string) *UserTeamTable {
	return newUserTeamTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UserTeamTable with assigned schema name
func (a UserTeamTable) FromSchema(schemaName string) *UserTeamTable {
	return newUserTeamTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UserTeamTable with assigned table prefix
func (a UserTeamTable) WithPrefix(prefix string) *UserTeamTable {
	return newUserTeamTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UserTeamTable with assigned table suffix
func (a UserTeamTable) WithSuffix(suffix string) *UserTeamTable {
	return newUserTeamTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUserTeamTable(schemaName, tableName, alias string) *UserTeamTable {
	return &UserTeamTable{
		userTeamTable: newUserTeamTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newUserTeamTableImpl("", "excluded", ""),
	}
}

func newUserTeamTableImpl(schemaName, tableName, alias string) userTeamTable {
	var (
		TeamIDColumn   = postgres.IntegerColumn("team_id")
		UserIDColumn   = postgres.StringColumn("user_id")
		allColumns     = postgres.ColumnList{TeamIDColumn, UserIDColumn}
		mutableColumns = postgres.ColumnList{}
	)

	return userTeamTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		TeamID: TeamIDColumn,
		UserID: UserIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
