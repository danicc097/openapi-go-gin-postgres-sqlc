// Code generated by project. DO NOT EDIT.

package codegen

import (
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	rest "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
)

var PublicStructs = map[string]any{
	"DbActivity":                            new(db.Activity),
	"DbActivityCreateParams":                new(db.ActivityCreateParams),
	"DbActivityJoins":                       new(db.ActivityJoins),
	"DbActivitySelectConfig":                new(db.ActivitySelectConfig),
	"DbActivityUpdateParams":                new(db.ActivityUpdateParams),
	"DbDemoTwoWorkItem":                     new(db.DemoTwoWorkItem),
	"DbDemoTwoWorkItemCreateParams":         new(db.DemoTwoWorkItemCreateParams),
	"DbDemoTwoWorkItemJoins":                new(db.DemoTwoWorkItemJoins),
	"DbDemoTwoWorkItemSelectConfig":         new(db.DemoTwoWorkItemSelectConfig),
	"DbDemoTwoWorkItemUpdateParams":         new(db.DemoTwoWorkItemUpdateParams),
	"DbDemoWorkItem":                        new(db.DemoWorkItem),
	"DbDemoWorkItemCreateParams":            new(db.DemoWorkItemCreateParams),
	"DbDemoWorkItemJoins":                   new(db.DemoWorkItemJoins),
	"DbDemoWorkItemSelectConfig":            new(db.DemoWorkItemSelectConfig),
	"DbDemoWorkItemUpdateParams":            new(db.DemoWorkItemUpdateParams),
	"DbEntityNotification":                  new(db.EntityNotification),
	"DbEntityNotificationCreateParams":      new(db.EntityNotificationCreateParams),
	"DbEntityNotificationJoins":             new(db.EntityNotificationJoins),
	"DbEntityNotificationSelectConfig":      new(db.EntityNotificationSelectConfig),
	"DbEntityNotificationUpdateParams":      new(db.EntityNotificationUpdateParams),
	"DbErrInsertFailed":                     new(db.ErrInsertFailed),
	"DbErrUpdateFailed":                     new(db.ErrUpdateFailed),
	"DbErrUpsertFailed":                     new(db.ErrUpsertFailed),
	"DbGetUserNotificationsParams":          new(db.GetUserNotificationsParams),
	"DbGetUserNotificationsRow":             new(db.GetUserNotificationsRow),
	"DbGetUserParams":                       new(db.GetUserParams),
	"DbGetUserRow":                          new(db.GetUserRow),
	"DbKanbanStep":                          new(db.KanbanStep),
	"DbKanbanStepCreateParams":              new(db.KanbanStepCreateParams),
	"DbKanbanStepJoins":                     new(db.KanbanStepJoins),
	"DbKanbanStepSelectConfig":              new(db.KanbanStepSelectConfig),
	"DbKanbanStepUpdateParams":              new(db.KanbanStepUpdateParams),
	"DbMovie":                               new(db.Movie),
	"DbMovieCreateParams":                   new(db.MovieCreateParams),
	"DbMovieJoins":                          new(db.MovieJoins),
	"DbMovieSelectConfig":                   new(db.MovieSelectConfig),
	"DbMovieUpdateParams":                   new(db.MovieUpdateParams),
	"DbNotification":                        new(db.Notification),
	"DbNotificationCreateParams":            new(db.NotificationCreateParams),
	"DbNotificationJoins":                   new(db.NotificationJoins),
	"DbNotificationSelectConfig":            new(db.NotificationSelectConfig),
	"DbNotificationUpdateParams":            new(db.NotificationUpdateParams),
	"DbNullNotificationType":                new(db.NullNotificationType),
	"DbNullWorkItemRole":                    new(db.NullWorkItemRole),
	"DbProject":                             new(db.Project),
	"DbProjectCreateParams":                 new(db.ProjectCreateParams),
	"DbProjectJoins":                        new(db.ProjectJoins),
	"DbProjectSelectConfig":                 new(db.ProjectSelectConfig),
	"DbProjectUpdateParams":                 new(db.ProjectUpdateParams),
	"DbQueries":                             new(db.Queries),
	"DbRegisterNewUserParams":               new(db.RegisterNewUserParams),
	"DbRegisterNewUserRow":                  new(db.RegisterNewUserRow),
	"DbSchemaMigration":                     new(db.SchemaMigration),
	"DbSchemaMigrationCreateParams":         new(db.SchemaMigrationCreateParams),
	"DbSchemaMigrationJoins":                new(db.SchemaMigrationJoins),
	"DbSchemaMigrationSelectConfig":         new(db.SchemaMigrationSelectConfig),
	"DbSchemaMigrationUpdateParams":         new(db.SchemaMigrationUpdateParams),
	"DbTeam":                                new(db.Team),
	"DbTeamCreateParams":                    new(db.TeamCreateParams),
	"DbTeamJoins":                           new(db.TeamJoins),
	"DbTeamSelectConfig":                    new(db.TeamSelectConfig),
	"DbTeamUpdateParams":                    new(db.TeamUpdateParams),
	"DbTestRow":                             new(db.TestRow),
	"DbTimeEntry":                           new(db.TimeEntry),
	"DbTimeEntryCreateParams":               new(db.TimeEntryCreateParams),
	"DbTimeEntryJoins":                      new(db.TimeEntryJoins),
	"DbTimeEntrySelectConfig":               new(db.TimeEntrySelectConfig),
	"DbTimeEntryUpdateParams":               new(db.TimeEntryUpdateParams),
	"DbTrigger":                             new(db.Trigger),
	"DbUser":                                new(db.User),
	"DbUserAPIKey":                          new(db.UserAPIKey),
	"DbUserAPIKeyCreateParams":              new(db.UserAPIKeyCreateParams),
	"DbUserAPIKeyJoins":                     new(db.UserAPIKeyJoins),
	"DbUserAPIKeySelectConfig":              new(db.UserAPIKeySelectConfig),
	"DbUserAPIKeyUpdateParams":              new(db.UserAPIKeyUpdateParams),
	"DbUserCreateParams":                    new(db.UserCreateParams),
	"DbUserID":                              new(db.UserID),
	"DbUserJoins":                           new(db.UserJoins),
	"DbUserNotification":                    new(db.UserNotification),
	"DbUserNotificationCreateParams":        new(db.UserNotificationCreateParams),
	"DbUserNotificationJoins":               new(db.UserNotificationJoins),
	"DbUserNotificationSelectConfig":        new(db.UserNotificationSelectConfig),
	"DbUserNotificationUpdateParams":        new(db.UserNotificationUpdateParams),
	"DbUserSelectConfig":                    new(db.UserSelectConfig),
	"DbUserTeam":                            new(db.UserTeam),
	"DbUserTeamCreateParams":                new(db.UserTeamCreateParams),
	"DbUserTeamJoins":                       new(db.UserTeamJoins),
	"DbUserTeamSelectConfig":                new(db.UserTeamSelectConfig),
	"DbUserTeamUpdateParams":                new(db.UserTeamUpdateParams),
	"DbUserUpdateParams":                    new(db.UserUpdateParams),
	"DbUser__WIAU_WorkItem":                 new(db.User__WIAU_WorkItem),
	"DbUser__WIAU_WorkItemAssignedUser":     new(db.User__WIAU_WorkItemAssignedUser),
	"DbWorkItem":                            new(db.WorkItem),
	"DbWorkItemAssignedUser":                new(db.WorkItemAssignedUser),
	"DbWorkItemAssignedUserCreateParams":    new(db.WorkItemAssignedUserCreateParams),
	"DbWorkItemAssignedUserJoins":           new(db.WorkItemAssignedUserJoins),
	"DbWorkItemAssignedUserSelectConfig":    new(db.WorkItemAssignedUserSelectConfig),
	"DbWorkItemAssignedUserUpdateParams":    new(db.WorkItemAssignedUserUpdateParams),
	"DbWorkItemComment":                     new(db.WorkItemComment),
	"DbWorkItemCommentCreateParams":         new(db.WorkItemCommentCreateParams),
	"DbWorkItemCommentJoins":                new(db.WorkItemCommentJoins),
	"DbWorkItemCommentSelectConfig":         new(db.WorkItemCommentSelectConfig),
	"DbWorkItemCommentUpdateParams":         new(db.WorkItemCommentUpdateParams),
	"DbWorkItemCreateParams":                new(db.WorkItemCreateParams),
	"DbWorkItemJoins":                       new(db.WorkItemJoins),
	"DbWorkItemSelectConfig":                new(db.WorkItemSelectConfig),
	"DbWorkItemTag":                         new(db.WorkItemTag),
	"DbWorkItemTagCreateParams":             new(db.WorkItemTagCreateParams),
	"DbWorkItemTagJoins":                    new(db.WorkItemTagJoins),
	"DbWorkItemTagSelectConfig":             new(db.WorkItemTagSelectConfig),
	"DbWorkItemTagUpdateParams":             new(db.WorkItemTagUpdateParams),
	"DbWorkItemType":                        new(db.WorkItemType),
	"DbWorkItemTypeCreateParams":            new(db.WorkItemTypeCreateParams),
	"DbWorkItemTypeJoins":                   new(db.WorkItemTypeJoins),
	"DbWorkItemTypeSelectConfig":            new(db.WorkItemTypeSelectConfig),
	"DbWorkItemTypeUpdateParams":            new(db.WorkItemTypeUpdateParams),
	"DbWorkItemUpdateParams":                new(db.WorkItemUpdateParams),
	"DbWorkItemWorkItemTag":                 new(db.WorkItemWorkItemTag),
	"DbWorkItemWorkItemTagCreateParams":     new(db.WorkItemWorkItemTagCreateParams),
	"DbWorkItemWorkItemTagJoins":            new(db.WorkItemWorkItemTagJoins),
	"DbWorkItemWorkItemTagSelectConfig":     new(db.WorkItemWorkItemTagSelectConfig),
	"DbWorkItemWorkItemTagUpdateParams":     new(db.WorkItemWorkItemTagUpdateParams),
	"DbWorkItem__WIAU_User":                 new(db.WorkItem__WIAU_User),
	"DbWorkItem__WIAU_WorkItemAssignedUser": new(db.WorkItem__WIAU_WorkItemAssignedUser),
	"DbXoError":                             new(db.XoError),

	//

	"RestDemoTwoWorkItemCreateRequest":      new(rest.DemoTwoWorkItemCreateRequest),
	"RestDemoTwoWorkItemsResponse":          new(rest.DemoTwoWorkItemsResponse),
	"RestDemoWorkItemCreateRequest":         new(rest.DemoWorkItemCreateRequest),
	"RestDemoWorkItemsResponse":             new(rest.DemoWorkItemsResponse),
	"RestGetPaginatedNotificationsResponse": new(rest.GetPaginatedNotificationsResponse),
	"RestNotification":                      new(rest.Notification),
	"RestProjectBoardCreateRequest":         new(rest.ProjectBoardCreateRequest),
	"RestProjectBoardResponse":              new(rest.ProjectBoardResponse),
	"RestProjectName":                       new(rest.ProjectName),
	"RestSharedWorkItemFields":              new(rest.SharedWorkItemFields),
	"RestTeamCreateRequest":                 new(rest.TeamCreateRequest),
	"RestTeamUpdateRequest":                 new(rest.TeamUpdateRequest),
	"RestUser":                              new(rest.User),
	"RestWorkItemCommentCreateRequest":      new(rest.WorkItemCommentCreateRequest),
	"RestWorkItemResponse":                  new(rest.WorkItemResponse),
	"RestWorkItemTagCreateRequest":          new(rest.WorkItemTagCreateRequest),

	//

}
