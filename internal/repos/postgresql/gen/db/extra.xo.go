package db

// Code generated by xo. DO NOT EDIT.

import (
	"fmt"
)

type Direction string

const (
	DirectionAsc  Direction = "asc"
	DirectionDesc Direction = "desc"
)

func newPointer[T any](v T) *T {
	return &v
}

type XoError struct {
	Entity string
	Err    error
}

// Error satisfies the error interface.
func (e *XoError) Error() string {
	return fmt.Sprintf("%s %v", e.Entity, e.Err)
}

// Unwrap satisfies the unwrap interface.
func (err *XoError) Unwrap() error {
	return err.Err
}

type Entity string

const (
	ActivityEntity             Entity = "Activity"
	DemoTwoWorkItemEntity      Entity = "DemoTwoWorkItem"
	DemoWorkItemEntity         Entity = "DemoWorkItem"
	EntityNotificationEntity   Entity = "EntityNotification"
	KanbanStepEntity           Entity = "KanbanStep"
	MovieEntity                Entity = "Movie"
	NotificationEntity         Entity = "Notification"
	ProjectEntity              Entity = "Project"
	SchemaMigrationEntity      Entity = "SchemaMigration"
	TeamEntity                 Entity = "Team"
	TimeEntryEntity            Entity = "TimeEntry"
	UserEntity                 Entity = "User"
	UserAPIKeyEntity           Entity = "UserAPIKey"
	UserNotificationEntity     Entity = "UserNotification"
	UserTeamEntity             Entity = "UserTeam"
	WorkItemEntity             Entity = "WorkItem"
	WorkItemAssignedUserEntity Entity = "WorkItemAssignedUser"
	WorkItemCommentEntity      Entity = "WorkItemComment"
	WorkItemTagEntity          Entity = "WorkItemTag"
	WorkItemTypeEntity         Entity = "WorkItemType"
	WorkItemWorkItemTagEntity  Entity = "WorkItemWorkItemTag"
)
