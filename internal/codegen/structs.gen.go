// Code generated by project. DO NOT EDIT.

package codegen

import (
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	rest "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
)

var PublicStructs = map[string]any{
	"DbActivity":             new(db.Activity),
	"DbActivityCreateParams": new(db.ActivityCreateParams),
	"DbActivityJoins":        new(db.ActivityJoins),
	"DbActivitySelectConfig": new(db.ActivitySelectConfig),
	"DbActivityUpdateParams": new(db.ActivityUpdateParams),
	"DbBook__BASK_ExtraSchemaBookAuthorsSurrogateKey":  new(db.Book__BASK_ExtraSchemaBookAuthorsSurrogateKey),
	"DbBook__BASK_ExtraSchemaUser":                     new(db.Book__BASK_ExtraSchemaUser),
	"DbBook__BA_ExtraSchemaBookAuthor":                 new(db.Book__BA_ExtraSchemaBookAuthor),
	"DbBook__BA_ExtraSchemaUser":                       new(db.Book__BA_ExtraSchemaUser),
	"DbCacheDemoTwoWorkItem":                           new(db.CacheDemoTwoWorkItem),
	"DbCacheDemoTwoWorkItemJoins":                      new(db.CacheDemoTwoWorkItemJoins),
	"DbCacheDemoTwoWorkItemSelectConfig":               new(db.CacheDemoTwoWorkItemSelectConfig),
	"DbCacheDemoWorkItem":                              new(db.CacheDemoWorkItem),
	"DbCacheDemoWorkItemJoins":                         new(db.CacheDemoWorkItemJoins),
	"DbCacheDemoWorkItemSelectConfig":                  new(db.CacheDemoWorkItemSelectConfig),
	"DbDemoTwoWorkItem":                                new(db.DemoTwoWorkItem),
	"DbDemoTwoWorkItemCreateParams":                    new(db.DemoTwoWorkItemCreateParams),
	"DbDemoTwoWorkItemJoins":                           new(db.DemoTwoWorkItemJoins),
	"DbDemoTwoWorkItemSelectConfig":                    new(db.DemoTwoWorkItemSelectConfig),
	"DbDemoTwoWorkItemUpdateParams":                    new(db.DemoTwoWorkItemUpdateParams),
	"DbDemoWorkItem":                                   new(db.DemoWorkItem),
	"DbDemoWorkItemCreateParams":                       new(db.DemoWorkItemCreateParams),
	"DbDemoWorkItemJoins":                              new(db.DemoWorkItemJoins),
	"DbDemoWorkItemSelectConfig":                       new(db.DemoWorkItemSelectConfig),
	"DbDemoWorkItemUpdateParams":                       new(db.DemoWorkItemUpdateParams),
	"DbEntityNotification":                             new(db.EntityNotification),
	"DbEntityNotificationCreateParams":                 new(db.EntityNotificationCreateParams),
	"DbEntityNotificationJoins":                        new(db.EntityNotificationJoins),
	"DbEntityNotificationSelectConfig":                 new(db.EntityNotificationSelectConfig),
	"DbEntityNotificationUpdateParams":                 new(db.EntityNotificationUpdateParams),
	"DbErrInsertFailed":                                new(db.ErrInsertFailed),
	"DbErrUpdateFailed":                                new(db.ErrUpdateFailed),
	"DbErrUpsertFailed":                                new(db.ErrUpsertFailed),
	"DbExtraSchemaBook":                                new(db.ExtraSchemaBook),
	"DbExtraSchemaBookAuthor":                          new(db.ExtraSchemaBookAuthor),
	"DbExtraSchemaBookAuthorCreateParams":              new(db.ExtraSchemaBookAuthorCreateParams),
	"DbExtraSchemaBookAuthorJoins":                     new(db.ExtraSchemaBookAuthorJoins),
	"DbExtraSchemaBookAuthorSelectConfig":              new(db.ExtraSchemaBookAuthorSelectConfig),
	"DbExtraSchemaBookAuthorUpdateParams":              new(db.ExtraSchemaBookAuthorUpdateParams),
	"DbExtraSchemaBookAuthorsSurrogateKey":             new(db.ExtraSchemaBookAuthorsSurrogateKey),
	"DbExtraSchemaBookAuthorsSurrogateKeyCreateParams": new(db.ExtraSchemaBookAuthorsSurrogateKeyCreateParams),
	"DbExtraSchemaBookAuthorsSurrogateKeyJoins":        new(db.ExtraSchemaBookAuthorsSurrogateKeyJoins),
	"DbExtraSchemaBookAuthorsSurrogateKeySelectConfig": new(db.ExtraSchemaBookAuthorsSurrogateKeySelectConfig),
	"DbExtraSchemaBookAuthorsSurrogateKeyUpdateParams": new(db.ExtraSchemaBookAuthorsSurrogateKeyUpdateParams),
	"DbExtraSchemaBookCreateParams":                    new(db.ExtraSchemaBookCreateParams),
	"DbExtraSchemaBookJoins":                           new(db.ExtraSchemaBookJoins),
	"DbExtraSchemaBookReview":                          new(db.ExtraSchemaBookReview),
	"DbExtraSchemaBookReviewCreateParams":              new(db.ExtraSchemaBookReviewCreateParams),
	"DbExtraSchemaBookReviewJoins":                     new(db.ExtraSchemaBookReviewJoins),
	"DbExtraSchemaBookReviewSelectConfig":              new(db.ExtraSchemaBookReviewSelectConfig),
	"DbExtraSchemaBookReviewUpdateParams":              new(db.ExtraSchemaBookReviewUpdateParams),
	"DbExtraSchemaBookSelectConfig":                    new(db.ExtraSchemaBookSelectConfig),
	"DbExtraSchemaBookSeller":                          new(db.ExtraSchemaBookSeller),
	"DbExtraSchemaBookSellerCreateParams":              new(db.ExtraSchemaBookSellerCreateParams),
	"DbExtraSchemaBookSellerJoins":                     new(db.ExtraSchemaBookSellerJoins),
	"DbExtraSchemaBookSellerSelectConfig":              new(db.ExtraSchemaBookSellerSelectConfig),
	"DbExtraSchemaBookSellerUpdateParams":              new(db.ExtraSchemaBookSellerUpdateParams),
	"DbExtraSchemaBookUpdateParams":                    new(db.ExtraSchemaBookUpdateParams),
	"DbExtraSchemaDemoWorkItem":                        new(db.ExtraSchemaDemoWorkItem),
	"DbExtraSchemaDemoWorkItemCreateParams":            new(db.ExtraSchemaDemoWorkItemCreateParams),
	"DbExtraSchemaDemoWorkItemJoins":                   new(db.ExtraSchemaDemoWorkItemJoins),
	"DbExtraSchemaDemoWorkItemSelectConfig":            new(db.ExtraSchemaDemoWorkItemSelectConfig),
	"DbExtraSchemaDemoWorkItemUpdateParams":            new(db.ExtraSchemaDemoWorkItemUpdateParams),
	"DbExtraSchemaDummyJoin":                           new(db.ExtraSchemaDummyJoin),
	"DbExtraSchemaDummyJoinCreateParams":               new(db.ExtraSchemaDummyJoinCreateParams),
	"DbExtraSchemaDummyJoinJoins":                      new(db.ExtraSchemaDummyJoinJoins),
	"DbExtraSchemaDummyJoinSelectConfig":               new(db.ExtraSchemaDummyJoinSelectConfig),
	"DbExtraSchemaDummyJoinUpdateParams":               new(db.ExtraSchemaDummyJoinUpdateParams),
	"DbExtraSchemaNotification":                        new(db.ExtraSchemaNotification),
	"DbExtraSchemaNotificationCreateParams":            new(db.ExtraSchemaNotificationCreateParams),
	"DbExtraSchemaNotificationJoins":                   new(db.ExtraSchemaNotificationJoins),
	"DbExtraSchemaNotificationSelectConfig":            new(db.ExtraSchemaNotificationSelectConfig),
	"DbExtraSchemaNotificationUpdateParams":            new(db.ExtraSchemaNotificationUpdateParams),
	"DbExtraSchemaPagElement":                          new(db.ExtraSchemaPagElement),
	"DbExtraSchemaPagElementCreateParams":              new(db.ExtraSchemaPagElementCreateParams),
	"DbExtraSchemaPagElementID":                        new(db.ExtraSchemaPagElementID),
	"DbExtraSchemaPagElementJoins":                     new(db.ExtraSchemaPagElementJoins),
	"DbExtraSchemaPagElementSelectConfig":              new(db.ExtraSchemaPagElementSelectConfig),
	"DbExtraSchemaPagElementUpdateParams":              new(db.ExtraSchemaPagElementUpdateParams),
	"DbExtraSchemaUser":                                new(db.ExtraSchemaUser),
	"DbExtraSchemaUserAPIKey":                          new(db.ExtraSchemaUserAPIKey),
	"DbExtraSchemaUserAPIKeyCreateParams":              new(db.ExtraSchemaUserAPIKeyCreateParams),
	"DbExtraSchemaUserAPIKeyJoins":                     new(db.ExtraSchemaUserAPIKeyJoins),
	"DbExtraSchemaUserAPIKeySelectConfig":              new(db.ExtraSchemaUserAPIKeySelectConfig),
	"DbExtraSchemaUserAPIKeyUpdateParams":              new(db.ExtraSchemaUserAPIKeyUpdateParams),
	"DbExtraSchemaUserCreateParams":                    new(db.ExtraSchemaUserCreateParams),
	"DbExtraSchemaUserID":                              new(db.ExtraSchemaUserID),
	"DbExtraSchemaUserJoins":                           new(db.ExtraSchemaUserJoins),
	"DbExtraSchemaUserSelectConfig":                    new(db.ExtraSchemaUserSelectConfig),
	"DbExtraSchemaUserUpdateParams":                    new(db.ExtraSchemaUserUpdateParams),
	"DbExtraSchemaWorkItem":                            new(db.ExtraSchemaWorkItem),
	"DbExtraSchemaWorkItemAssignedUser":                new(db.ExtraSchemaWorkItemAssignedUser),
	"DbExtraSchemaWorkItemAssignedUserCreateParams":    new(db.ExtraSchemaWorkItemAssignedUserCreateParams),
	"DbExtraSchemaWorkItemAssignedUserJoins":           new(db.ExtraSchemaWorkItemAssignedUserJoins),
	"DbExtraSchemaWorkItemAssignedUserSelectConfig":    new(db.ExtraSchemaWorkItemAssignedUserSelectConfig),
	"DbExtraSchemaWorkItemAssignedUserUpdateParams":    new(db.ExtraSchemaWorkItemAssignedUserUpdateParams),
	"DbExtraSchemaWorkItemCreateParams":                new(db.ExtraSchemaWorkItemCreateParams),
	"DbExtraSchemaWorkItemJoins":                       new(db.ExtraSchemaWorkItemJoins),
	"DbExtraSchemaWorkItemSelectConfig":                new(db.ExtraSchemaWorkItemSelectConfig),
	"DbExtraSchemaWorkItemUpdateParams":                new(db.ExtraSchemaWorkItemUpdateParams),
	"DbGetExtraSchemaNotificationsParams":              new(db.GetExtraSchemaNotificationsParams),
	"DbGetExtraSchemaNotificationsRow":                 new(db.GetExtraSchemaNotificationsRow),
	"DbGetUserNotificationsParams":                     new(db.GetUserNotificationsParams),
	"DbGetUserNotificationsRow":                        new(db.GetUserNotificationsRow),
	"DbGetUserParams":                                  new(db.GetUserParams),
	"DbGetUserRow":                                     new(db.GetUserRow),
	"DbIsTeamInProjectParams":                          new(db.IsTeamInProjectParams),
	"DbKanbanStep":                                     new(db.KanbanStep),
	"DbKanbanStepCreateParams":                         new(db.KanbanStepCreateParams),
	"DbKanbanStepJoins":                                new(db.KanbanStepJoins),
	"DbKanbanStepSelectConfig":                         new(db.KanbanStepSelectConfig),
	"DbKanbanStepUpdateParams":                         new(db.KanbanStepUpdateParams),
	"DbMovie":                                          new(db.Movie),
	"DbMovieCreateParams":                              new(db.MovieCreateParams),
	"DbMovieJoins":                                     new(db.MovieJoins),
	"DbMovieSelectConfig":                              new(db.MovieSelectConfig),
	"DbMovieUpdateParams":                              new(db.MovieUpdateParams),
	"DbNotification":                                   new(db.Notification),
	"DbNotificationCreateParams":                       new(db.NotificationCreateParams),
	"DbNotificationJoins":                              new(db.NotificationJoins),
	"DbNotificationSelectConfig":                       new(db.NotificationSelectConfig),
	"DbNotificationUpdateParams":                       new(db.NotificationUpdateParams),
	"DbProject":                                        new(db.Project),
	"DbProjectCreateParams":                            new(db.ProjectCreateParams),
	"DbProjectJoins":                                   new(db.ProjectJoins),
	"DbProjectSelectConfig":                            new(db.ProjectSelectConfig),
	"DbProjectUpdateParams":                            new(db.ProjectUpdateParams),
	"DbQueries":                                        new(db.Queries),
	"DbSchemaMigration":                                new(db.SchemaMigration),
	"DbSchemaMigrationCreateParams":                    new(db.SchemaMigrationCreateParams),
	"DbSchemaMigrationJoins":                           new(db.SchemaMigrationJoins),
	"DbSchemaMigrationSelectConfig":                    new(db.SchemaMigrationSelectConfig),
	"DbSchemaMigrationUpdateParams":                    new(db.SchemaMigrationUpdateParams),
	"DbTeam":                                           new(db.Team),
	"DbTeamCreateParams":                               new(db.TeamCreateParams),
	"DbTeamJoins":                                      new(db.TeamJoins),
	"DbTeamSelectConfig":                               new(db.TeamSelectConfig),
	"DbTeamUpdateParams":                               new(db.TeamUpdateParams),
	"DbTimeEntry":                                      new(db.TimeEntry),
	"DbTimeEntryCreateParams":                          new(db.TimeEntryCreateParams),
	"DbTimeEntryJoins":                                 new(db.TimeEntryJoins),
	"DbTimeEntrySelectConfig":                          new(db.TimeEntrySelectConfig),
	"DbTimeEntryUpdateParams":                          new(db.TimeEntryUpdateParams),
	"DbTrigger":                                        new(db.Trigger),
	"DbUser":                                           new(db.User),
	"DbUserAPIKey":                                     new(db.UserAPIKey),
	"DbUserAPIKeyCreateParams":                         new(db.UserAPIKeyCreateParams),
	"DbUserAPIKeyJoins":                                new(db.UserAPIKeyJoins),
	"DbUserAPIKeySelectConfig":                         new(db.UserAPIKeySelectConfig),
	"DbUserAPIKeyUpdateParams":                         new(db.UserAPIKeyUpdateParams),
	"DbUserCreateParams":                               new(db.UserCreateParams),
	"DbUserID":                                         new(db.UserID),
	"DbUserJoins":                                      new(db.UserJoins),
	"DbUserNotification":                               new(db.UserNotification),
	"DbUserNotificationCreateParams":                   new(db.UserNotificationCreateParams),
	"DbUserNotificationJoins":                          new(db.UserNotificationJoins),
	"DbUserNotificationSelectConfig":                   new(db.UserNotificationSelectConfig),
	"DbUserNotificationUpdateParams":                   new(db.UserNotificationUpdateParams),
	"DbUserProject":                                    new(db.UserProject),
	"DbUserProjectCreateParams":                        new(db.UserProjectCreateParams),
	"DbUserProjectJoins":                               new(db.UserProjectJoins),
	"DbUserProjectSelectConfig":                        new(db.UserProjectSelectConfig),
	"DbUserProjectUpdateParams":                        new(db.UserProjectUpdateParams),
	"DbUserSelectConfig":                               new(db.UserSelectConfig),
	"DbUserTeam":                                       new(db.UserTeam),
	"DbUserTeamCreateParams":                           new(db.UserTeamCreateParams),
	"DbUserTeamJoins":                                  new(db.UserTeamJoins),
	"DbUserTeamSelectConfig":                           new(db.UserTeamSelectConfig),
	"DbUserTeamUpdateParams":                           new(db.UserTeamUpdateParams),
	"DbUserUpdateParams":                               new(db.UserUpdateParams),
	"DbUser__BASK_ExtraSchemaBook":                     new(db.User__BASK_ExtraSchemaBook),
	"DbUser__BASK_ExtraSchemaBookAuthorsSurrogateKey":  new(db.User__BASK_ExtraSchemaBookAuthorsSurrogateKey),
	"DbUser__BA_ExtraSchemaBook":                       new(db.User__BA_ExtraSchemaBook),
	"DbUser__BA_ExtraSchemaBookAuthor":                 new(db.User__BA_ExtraSchemaBookAuthor),
	"DbUser__WIAU_ExtraSchemaWorkItem":                 new(db.User__WIAU_ExtraSchemaWorkItem),
	"DbUser__WIAU_ExtraSchemaWorkItemAssignedUser":     new(db.User__WIAU_ExtraSchemaWorkItemAssignedUser),
	"DbUser__WIAU_WorkItem":                            new(db.User__WIAU_WorkItem),
	"DbUser__WIAU_WorkItemAssignedUser":                new(db.User__WIAU_WorkItemAssignedUser),
	"DbWorkItem":                                       new(db.WorkItem),
	"DbWorkItemAssignedUser":                           new(db.WorkItemAssignedUser),
	"DbWorkItemAssignedUserCreateParams":               new(db.WorkItemAssignedUserCreateParams),
	"DbWorkItemAssignedUserJoins":                      new(db.WorkItemAssignedUserJoins),
	"DbWorkItemAssignedUserSelectConfig":               new(db.WorkItemAssignedUserSelectConfig),
	"DbWorkItemAssignedUserUpdateParams":               new(db.WorkItemAssignedUserUpdateParams),
	"DbWorkItemComment":                                new(db.WorkItemComment),
	"DbWorkItemCommentCreateParams":                    new(db.WorkItemCommentCreateParams),
	"DbWorkItemCommentJoins":                           new(db.WorkItemCommentJoins),
	"DbWorkItemCommentSelectConfig":                    new(db.WorkItemCommentSelectConfig),
	"DbWorkItemCommentUpdateParams":                    new(db.WorkItemCommentUpdateParams),
	"DbWorkItemCreateParams":                           new(db.WorkItemCreateParams),
	"DbWorkItemJoins":                                  new(db.WorkItemJoins),
	"DbWorkItemSelectConfig":                           new(db.WorkItemSelectConfig),
	"DbWorkItemTag":                                    new(db.WorkItemTag),
	"DbWorkItemTagCreateParams":                        new(db.WorkItemTagCreateParams),
	"DbWorkItemTagJoins":                               new(db.WorkItemTagJoins),
	"DbWorkItemTagSelectConfig":                        new(db.WorkItemTagSelectConfig),
	"DbWorkItemTagUpdateParams":                        new(db.WorkItemTagUpdateParams),
	"DbWorkItemType":                                   new(db.WorkItemType),
	"DbWorkItemTypeCreateParams":                       new(db.WorkItemTypeCreateParams),
	"DbWorkItemTypeJoins":                              new(db.WorkItemTypeJoins),
	"DbWorkItemTypeSelectConfig":                       new(db.WorkItemTypeSelectConfig),
	"DbWorkItemTypeUpdateParams":                       new(db.WorkItemTypeUpdateParams),
	"DbWorkItemUpdateParams":                           new(db.WorkItemUpdateParams),
	"DbWorkItemWorkItemTag":                            new(db.WorkItemWorkItemTag),
	"DbWorkItemWorkItemTagCreateParams":                new(db.WorkItemWorkItemTagCreateParams),
	"DbWorkItemWorkItemTagJoins":                       new(db.WorkItemWorkItemTagJoins),
	"DbWorkItemWorkItemTagSelectConfig":                new(db.WorkItemWorkItemTagSelectConfig),
	"DbWorkItemWorkItemTagUpdateParams":                new(db.WorkItemWorkItemTagUpdateParams),
	"DbWorkItem__WIAU_ExtraSchemaUser":                 new(db.WorkItem__WIAU_ExtraSchemaUser),
	"DbWorkItem__WIAU_ExtraSchemaWorkItemAssignedUser": new(db.WorkItem__WIAU_ExtraSchemaWorkItemAssignedUser),
	"DbWorkItem__WIAU_User":                            new(db.WorkItem__WIAU_User),
	"DbWorkItem__WIAU_WorkItemAssignedUser":            new(db.WorkItem__WIAU_WorkItemAssignedUser),
	"DbXoError":                                        new(db.XoError),

	//

	"RestActivity":                       new(rest.Activity),
	"RestCreateActivityRequest":          new(rest.CreateActivityRequest),
	"RestCreateDemoTwoWorkItemRequest":   new(rest.CreateDemoTwoWorkItemRequest),
	"RestCreateDemoWorkItemRequest":      new(rest.CreateDemoWorkItemRequest),
	"RestCreateProjectBoardRequest":      new(rest.CreateProjectBoardRequest),
	"RestCreateTeamRequest":              new(rest.CreateTeamRequest),
	"RestCreateWorkItemCommentRequest":   new(rest.CreateWorkItemCommentRequest),
	"RestCreateWorkItemTagRequest":       new(rest.CreateWorkItemTagRequest),
	"RestCreateWorkItemTypeRequest":      new(rest.CreateWorkItemTypeRequest),
	"RestDemoTwoWorkItems":               new(rest.DemoTwoWorkItems),
	"RestDemoWorkItems":                  new(rest.DemoWorkItems),
	"RestNotification":                   new(rest.Notification),
	"RestPaginatedNotificationsResponse": new(rest.PaginatedNotificationsResponse),
	"RestPaginatedUsersResponse":         new(rest.PaginatedUsersResponse),
	"RestPaginationPage":                 new(rest.PaginationPage),
	"RestProjectBoard":                   new(rest.ProjectBoard),
	"RestSharedWorkItemFields":           new(rest.SharedWorkItemFields),
	"RestTeam":                           new(rest.Team),
	"RestUpdateActivityRequest":          new(rest.UpdateActivityRequest),
	"RestUpdateTeamRequest":              new(rest.UpdateTeamRequest),
	"RestUpdateWorkItemCommentRequest":   new(rest.UpdateWorkItemCommentRequest),
	"RestUpdateWorkItemTagRequest":       new(rest.UpdateWorkItemTagRequest),
	"RestUpdateWorkItemTypeRequest":      new(rest.UpdateWorkItemTypeRequest),
	"RestUser":                           new(rest.User),
	"RestWorkItemComment":                new(rest.WorkItemComment),
	"RestWorkItemTag":                    new(rest.WorkItemTag),
	"RestWorkItemType":                   new(rest.WorkItemType),

	//

}
