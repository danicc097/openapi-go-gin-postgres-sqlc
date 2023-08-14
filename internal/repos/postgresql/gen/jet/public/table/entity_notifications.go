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

var EntityNotifications = newEntityNotificationsTable("public", "entity_notifications", "")

type entityNotificationsTable struct {
	postgres.Table

	// Columns
	EntityNotificationID postgres.ColumnInteger
	Entity               postgres.ColumnString
	ID                   postgres.ColumnString
	Message              postgres.ColumnString
	Topic                postgres.ColumnString
	CreatedAt            postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type EntityNotificationsTable struct {
	entityNotificationsTable

	EXCLUDED entityNotificationsTable
}

// AS creates new EntityNotificationsTable with assigned alias
func (a EntityNotificationsTable) AS(alias string) *EntityNotificationsTable {
	return newEntityNotificationsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EntityNotificationsTable with assigned schema name
func (a EntityNotificationsTable) FromSchema(schemaName string) *EntityNotificationsTable {
	return newEntityNotificationsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EntityNotificationsTable with assigned table prefix
func (a EntityNotificationsTable) WithPrefix(prefix string) *EntityNotificationsTable {
	return newEntityNotificationsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EntityNotificationsTable with assigned table suffix
func (a EntityNotificationsTable) WithSuffix(suffix string) *EntityNotificationsTable {
	return newEntityNotificationsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEntityNotificationsTable(schemaName, tableName, alias string) *EntityNotificationsTable {
	return &EntityNotificationsTable{
		entityNotificationsTable: newEntityNotificationsTableImpl(schemaName, tableName, alias),
		EXCLUDED:                 newEntityNotificationsTableImpl("", "excluded", ""),
	}
}

func newEntityNotificationsTableImpl(schemaName, tableName, alias string) entityNotificationsTable {
	var (
		EntityNotificationIDColumn = postgres.IntegerColumn("entity_notification_id")
		EntityColumn               = postgres.StringColumn("entity")
		IDColumn                   = postgres.StringColumn("id")
		MessageColumn              = postgres.StringColumn("message")
		TopicColumn                = postgres.StringColumn("topic")
		CreatedAtColumn            = postgres.TimestampzColumn("created_at")
		allColumns                 = postgres.ColumnList{EntityNotificationIDColumn, EntityColumn, IDColumn, MessageColumn, TopicColumn, CreatedAtColumn}
		mutableColumns             = postgres.ColumnList{EntityColumn, IDColumn, MessageColumn, TopicColumn, CreatedAtColumn}
	)

	return entityNotificationsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		EntityNotificationID: EntityNotificationIDColumn,
		Entity:               EntityColumn,
		ID:                   IDColumn,
		Message:              MessageColumn,
		Topic:                TopicColumn,
		CreatedAt:            CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
