// Code generated by 'project postgen'. DO NOT EDIT.

package postgen

import (
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	rest "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
)

var PublicStructs = map[string]any{
	"DbWorkItemsDemoProjectPublic":       db.WorkItemsDemoProjectPublic{},
	"DbWorkItemsDemoProject":             db.WorkItemsDemoProject{},
	"DbWorkItemsDemoProjectSelectConfig": db.WorkItemsDemoProjectSelectConfig{},
	"DbWorkItemsDemoProjectJoins":        db.WorkItemsDemoProjectJoins{},
	"DbProjectPublic":                    db.ProjectPublic{},
	"DbProject":                          db.Project{},
	"DbProjectSelectConfig":              db.ProjectSelectConfig{},
	"DbProjectJoins":                     db.ProjectJoins{},
	"DbWorkItemCommentPublic":            db.WorkItemCommentPublic{},
	"DbWorkItemComment":                  db.WorkItemComment{},
	"DbWorkItemCommentSelectConfig":      db.WorkItemCommentSelectConfig{},
	"DbWorkItemCommentJoins":             db.WorkItemCommentJoins{},
	"DbTrigger":                          db.Trigger{},
	"DbErrInsertFailed":                  db.ErrInsertFailed{},
	"DbErrUpdateFailed":                  db.ErrUpdateFailed{},
	"DbErrUpsertFailed":                  db.ErrUpsertFailed{},
	"DbNullNotificationType":             db.NullNotificationType{},
	"DbUserTeamPublic":                   db.UserTeamPublic{},
	"DbUserTeam":                         db.UserTeam{},
	"DbUserTeamSelectConfig":             db.UserTeamSelectConfig{},
	"DbUserTeamJoins":                    db.UserTeamJoins{},
	"DbWorkItemMemberPublic":             db.WorkItemMemberPublic{},
	"DbWorkItemMember":                   db.WorkItemMember{},
	"DbWorkItemMemberSelectConfig":       db.WorkItemMemberSelectConfig{},
	"DbWorkItemMemberJoins":              db.WorkItemMemberJoins{},
	"DbQueries":                          db.Queries{},
	"DbTimeEntryPublic":                  db.TimeEntryPublic{},
	"DbTimeEntry":                        db.TimeEntry{},
	"DbTimeEntrySelectConfig":            db.TimeEntrySelectConfig{},
	"DbTimeEntryJoins":                   db.TimeEntryJoins{},
	"DbSchemaMigrationPublic":            db.SchemaMigrationPublic{},
	"DbSchemaMigration":                  db.SchemaMigration{},
	"DbSchemaMigrationSelectConfig":      db.SchemaMigrationSelectConfig{},
	"DbSchemaMigrationJoins":             db.SchemaMigrationJoins{},
	"DbWorkItemPublic":                   db.WorkItemPublic{},
	"DbWorkItem":                         db.WorkItem{},
	"DbWorkItemSelectConfig":             db.WorkItemSelectConfig{},
	"DbWorkItemJoins":                    db.WorkItemJoins{},
	"DbWorkItemTypePublic":               db.WorkItemTypePublic{},
	"DbWorkItemType":                     db.WorkItemType{},
	"DbWorkItemTypeSelectConfig":         db.WorkItemTypeSelectConfig{},
	"DbWorkItemTypeJoins":                db.WorkItemTypeJoins{},
	"DbWorkItemWorkItemTagPublic":        db.WorkItemWorkItemTagPublic{},
	"DbWorkItemWorkItemTag":              db.WorkItemWorkItemTag{},
	"DbWorkItemWorkItemTagSelectConfig":  db.WorkItemWorkItemTagSelectConfig{},
	"DbWorkItemWorkItemTagJoins":         db.WorkItemWorkItemTagJoins{},
	"DbKanbanStepPublic":                 db.KanbanStepPublic{},
	"DbKanbanStep":                       db.KanbanStep{},
	"DbKanbanStepSelectConfig":           db.KanbanStepSelectConfig{},
	"DbKanbanStepJoins":                  db.KanbanStepJoins{},
	"DbTeamPublic":                       db.TeamPublic{},
	"DbTeam":                             db.Team{},
	"DbTeamSelectConfig":                 db.TeamSelectConfig{},
	"DbTeamJoins":                        db.TeamJoins{},
	"DbActivityPublic":                   db.ActivityPublic{},
	"DbActivity":                         db.Activity{},
	"DbActivitySelectConfig":             db.ActivitySelectConfig{},
	"DbActivityJoins":                    db.ActivityJoins{},
	"DbGetUserParams":                    db.GetUserParams{},
	"DbGetUserRow":                       db.GetUserRow{},
	"DbListAllUsers2Row":                 db.ListAllUsers2Row{},
	"DbRegisterNewUserParams":            db.RegisterNewUserParams{},
	"DbRegisterNewUserRow":               db.RegisterNewUserRow{},
	"DbTestRow":                          db.TestRow{},
	"DbMoviePublic":                      db.MoviePublic{},
	"DbMovie":                            db.Movie{},
	"DbMovieSelectConfig":                db.MovieSelectConfig{},
	"DbMovieJoins":                       db.MovieJoins{},
	"DbWorkItemTagPublic":                db.WorkItemTagPublic{},
	"DbWorkItemTag":                      db.WorkItemTag{},
	"DbWorkItemTagSelectConfig":          db.WorkItemTagSelectConfig{},
	"DbWorkItemTagJoins":                 db.WorkItemTagJoins{},
	"DbWorkItemsProject2Public":          db.WorkItemsProject2Public{},
	"DbWorkItemsProject2":                db.WorkItemsProject2{},
	"DbWorkItemsProject2SelectConfig":    db.WorkItemsProject2SelectConfig{},
	"DbWorkItemsProject2Joins":           db.WorkItemsProject2Joins{},
	"DbNullWorkItemRole":                 db.NullWorkItemRole{},
	"DbUserNotificationPublic":           db.UserNotificationPublic{},
	"DbUserNotification":                 db.UserNotification{},
	"DbUserNotificationSelectConfig":     db.UserNotificationSelectConfig{},
	"DbUserNotificationJoins":            db.UserNotificationJoins{},
	"DbUserPublic":                       db.UserPublic{},
	"DbUser":                             db.User{},
	"DbUserSelectConfig":                 db.UserSelectConfig{},
	"DbUserJoins":                        db.UserJoins{},
	"DbNotificationPublic":               db.NotificationPublic{},
	"DbNotification":                     db.Notification{},
	"DbNotificationSelectConfig":         db.NotificationSelectConfig{},
	"DbNotificationJoins":                db.NotificationJoins{},
	"DbUserAPIKeyPublic":                 db.UserAPIKeyPublic{},
	"DbUserAPIKey":                       db.UserAPIKey{},
	"DbUserAPIKeySelectConfig":           db.UserAPIKeySelectConfig{},
	"DbUserAPIKeyJoins":                  db.UserAPIKeyJoins{},

	//

	"RestUserResponse":      rest.UserResponse{},
	"RestWorkItemResponse":  rest.WorkItemResponse{},
	"RestTeamCreateRequest": rest.TeamCreateRequest{},
	"RestTeamUpdateRequest": rest.TeamUpdateRequest{},
}
