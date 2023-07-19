// Code generated by 'project gen.postgen'. DO NOT EDIT.

package structs

import (
	repomodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	rest "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
)

var PublicStructs = map[string]any{
	"DbActivity":                            db.Activity{},
	"DbActivityCreateParams":                db.ActivityCreateParams{},
	"DbActivityJoins":                       db.ActivityJoins{},
	"DbActivitySelectConfig":                db.ActivitySelectConfig{},
	"DbActivityUpdateParams":                db.ActivityUpdateParams{},
	"DbDemoTwoWorkItem":                     db.DemoTwoWorkItem{},
	"DbDemoTwoWorkItemCreateParams":         db.DemoTwoWorkItemCreateParams{},
	"DbDemoTwoWorkItemJoins":                db.DemoTwoWorkItemJoins{},
	"DbDemoTwoWorkItemSelectConfig":         db.DemoTwoWorkItemSelectConfig{},
	"DbDemoTwoWorkItemUpdateParams":         db.DemoTwoWorkItemUpdateParams{},
	"DbDemoWorkItem":                        db.DemoWorkItem{},
	"DbDemoWorkItemCreateParams":            db.DemoWorkItemCreateParams{},
	"DbDemoWorkItemJoins":                   db.DemoWorkItemJoins{},
	"DbDemoWorkItemSelectConfig":            db.DemoWorkItemSelectConfig{},
	"DbDemoWorkItemUpdateParams":            db.DemoWorkItemUpdateParams{},
	"DbErrInsertFailed":                     db.ErrInsertFailed{},
	"DbErrUpdateFailed":                     db.ErrUpdateFailed{},
	"DbErrUpsertFailed":                     db.ErrUpsertFailed{},
	"DbGetUserNotificationsParams":          db.GetUserNotificationsParams{},
	"DbGetUserNotificationsRow":             db.GetUserNotificationsRow{},
	"DbGetUserParams":                       db.GetUserParams{},
	"DbGetUserRow":                          db.GetUserRow{},
	"DbKanbanStep":                          db.KanbanStep{},
	"DbKanbanStepCreateParams":              db.KanbanStepCreateParams{},
	"DbKanbanStepJoins":                     db.KanbanStepJoins{},
	"DbKanbanStepSelectConfig":              db.KanbanStepSelectConfig{},
	"DbKanbanStepUpdateParams":              db.KanbanStepUpdateParams{},
	"DbMovie":                               db.Movie{},
	"DbMovieCreateParams":                   db.MovieCreateParams{},
	"DbMovieJoins":                          db.MovieJoins{},
	"DbMovieSelectConfig":                   db.MovieSelectConfig{},
	"DbMovieUpdateParams":                   db.MovieUpdateParams{},
	"DbNotification":                        db.Notification{},
	"DbNotificationCreateParams":            db.NotificationCreateParams{},
	"DbNotificationJoins":                   db.NotificationJoins{},
	"DbNotificationSelectConfig":            db.NotificationSelectConfig{},
	"DbNotificationUpdateParams":            db.NotificationUpdateParams{},
	"DbNullNotificationType":                db.NullNotificationType{},
	"DbNullWorkItemRole":                    db.NullWorkItemRole{},
	"DbProject":                             db.Project{},
	"DbProjectCreateParams":                 db.ProjectCreateParams{},
	"DbProjectJoins":                        db.ProjectJoins{},
	"DbProjectSelectConfig":                 db.ProjectSelectConfig{},
	"DbProjectUpdateParams":                 db.ProjectUpdateParams{},
	"DbQueries":                             db.Queries{},
	"DbRegisterNewUserParams":               db.RegisterNewUserParams{},
	"DbRegisterNewUserRow":                  db.RegisterNewUserRow{},
	"DbSchemaMigration":                     db.SchemaMigration{},
	"DbSchemaMigrationCreateParams":         db.SchemaMigrationCreateParams{},
	"DbSchemaMigrationJoins":                db.SchemaMigrationJoins{},
	"DbSchemaMigrationSelectConfig":         db.SchemaMigrationSelectConfig{},
	"DbSchemaMigrationUpdateParams":         db.SchemaMigrationUpdateParams{},
	"DbTeam":                                db.Team{},
	"DbTeamCreateParams":                    db.TeamCreateParams{},
	"DbTeamJoins":                           db.TeamJoins{},
	"DbTeamSelectConfig":                    db.TeamSelectConfig{},
	"DbTeamUpdateParams":                    db.TeamUpdateParams{},
	"DbTestRow":                             db.TestRow{},
	"DbTimeEntry":                           db.TimeEntry{},
	"DbTimeEntryCreateParams":               db.TimeEntryCreateParams{},
	"DbTimeEntryJoins":                      db.TimeEntryJoins{},
	"DbTimeEntrySelectConfig":               db.TimeEntrySelectConfig{},
	"DbTimeEntryUpdateParams":               db.TimeEntryUpdateParams{},
	"DbTrigger":                             db.Trigger{},
	"DbUser":                                db.User{},
	"DbUserAPIKey":                          db.UserAPIKey{},
	"DbUserAPIKeyCreateParams":              db.UserAPIKeyCreateParams{},
	"DbUserAPIKeyJoins":                     db.UserAPIKeyJoins{},
	"DbUserAPIKeySelectConfig":              db.UserAPIKeySelectConfig{},
	"DbUserAPIKeyUpdateParams":              db.UserAPIKeyUpdateParams{},
	"DbUserCreateParams":                    db.UserCreateParams{},
	"DbUserJoins":                           db.UserJoins{},
	"DbUserNotification":                    db.UserNotification{},
	"DbUserNotificationCreateParams":        db.UserNotificationCreateParams{},
	"DbUserNotificationJoins":               db.UserNotificationJoins{},
	"DbUserNotificationSelectConfig":        db.UserNotificationSelectConfig{},
	"DbUserNotificationUpdateParams":        db.UserNotificationUpdateParams{},
	"DbUserSelectConfig":                    db.UserSelectConfig{},
	"DbUserTeam":                            db.UserTeam{},
	"DbUserTeamCreateParams":                db.UserTeamCreateParams{},
	"DbUserTeamJoins":                       db.UserTeamJoins{},
	"DbUserTeamSelectConfig":                db.UserTeamSelectConfig{},
	"DbUserTeamUpdateParams":                db.UserTeamUpdateParams{},
	"DbUserUpdateParams":                    db.UserUpdateParams{},
	"DbUser__WIAU_WorkItem":                 db.User__WIAU_WorkItem{},
	"DbUser__WIAU_WorkItemAssignedUser":     db.User__WIAU_WorkItemAssignedUser{},
	"DbWorkItem":                            db.WorkItem{},
	"DbWorkItemAssignedUser":                db.WorkItemAssignedUser{},
	"DbWorkItemAssignedUserCreateParams":    db.WorkItemAssignedUserCreateParams{},
	"DbWorkItemAssignedUserJoins":           db.WorkItemAssignedUserJoins{},
	"DbWorkItemAssignedUserSelectConfig":    db.WorkItemAssignedUserSelectConfig{},
	"DbWorkItemAssignedUserUpdateParams":    db.WorkItemAssignedUserUpdateParams{},
	"DbWorkItemComment":                     db.WorkItemComment{},
	"DbWorkItemCommentCreateParams":         db.WorkItemCommentCreateParams{},
	"DbWorkItemCommentJoins":                db.WorkItemCommentJoins{},
	"DbWorkItemCommentSelectConfig":         db.WorkItemCommentSelectConfig{},
	"DbWorkItemCommentUpdateParams":         db.WorkItemCommentUpdateParams{},
	"DbWorkItemCreateParams":                db.WorkItemCreateParams{},
	"DbWorkItemJoins":                       db.WorkItemJoins{},
	"DbWorkItemSelectConfig":                db.WorkItemSelectConfig{},
	"DbWorkItemTag":                         db.WorkItemTag{},
	"DbWorkItemTagCreateParams":             db.WorkItemTagCreateParams{},
	"DbWorkItemTagJoins":                    db.WorkItemTagJoins{},
	"DbWorkItemTagSelectConfig":             db.WorkItemTagSelectConfig{},
	"DbWorkItemTagUpdateParams":             db.WorkItemTagUpdateParams{},
	"DbWorkItemType":                        db.WorkItemType{},
	"DbWorkItemTypeCreateParams":            db.WorkItemTypeCreateParams{},
	"DbWorkItemTypeJoins":                   db.WorkItemTypeJoins{},
	"DbWorkItemTypeSelectConfig":            db.WorkItemTypeSelectConfig{},
	"DbWorkItemTypeUpdateParams":            db.WorkItemTypeUpdateParams{},
	"DbWorkItemUpdateParams":                db.WorkItemUpdateParams{},
	"DbWorkItemWorkItemTag":                 db.WorkItemWorkItemTag{},
	"DbWorkItemWorkItemTagCreateParams":     db.WorkItemWorkItemTagCreateParams{},
	"DbWorkItemWorkItemTagJoins":            db.WorkItemWorkItemTagJoins{},
	"DbWorkItemWorkItemTagSelectConfig":     db.WorkItemWorkItemTagSelectConfig{},
	"DbWorkItemWorkItemTagUpdateParams":     db.WorkItemWorkItemTagUpdateParams{},
	"DbWorkItem__WIAU_User":                 db.WorkItem__WIAU_User{},
	"DbWorkItem__WIAU_WorkItemAssignedUser": db.WorkItem__WIAU_WorkItemAssignedUser{},
	"DbXoError":                             db.XoError{},

	//

	"RestDemoTwoWorkItemsResponse":     rest.DemoTwoWorkItemsResponse{},
	"RestDemoWorkItemCreateRequest":    rest.DemoWorkItemCreateRequest{},
	"RestDemoWorkItemsResponse":        rest.DemoWorkItemsResponse{},
	"RestProjectBoardCreateRequest":    rest.ProjectBoardCreateRequest{},
	"RestProjectBoardResponse":         rest.ProjectBoardResponse{},
	"RestSharedWorkItemFields":         rest.SharedWorkItemFields{},
	"RestTeamCreateRequest":            rest.TeamCreateRequest{},
	"RestTeamUpdateRequest":            rest.TeamUpdateRequest{},
	"RestUser":                         rest.User{},
	"RestWorkItemCommentCreateRequest": rest.WorkItemCommentCreateRequest{},
	"RestWorkItemResponse":             rest.WorkItemResponse{},
	"RestWorkItemTagCreateRequest":     rest.WorkItemTagCreateRequest{},

	//

	"ModelsProjectBoard": repomodels.ProjectBoard{},
}
