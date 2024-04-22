// Code generated by project. DO NOT EDIT.

package codegen

import (
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	rest "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
)

var PublicStructs = map[string]any{
	"DbActivity":                                        new(db.Activity),
	"DbActivityCreateParams":                            new(db.ActivityCreateParams),
	"DbActivityJoins":                                   new(db.ActivityJoins),
	"DbActivitySelectConfig":                            new(db.ActivitySelectConfig),
	"DbActivityUpdateParams":                            new(db.ActivityUpdateParams),
	"DbCacheDemoTwoWorkItem":                            new(db.CacheDemoTwoWorkItem),
	"DbCacheDemoTwoWorkItemCreateParams":                new(db.CacheDemoTwoWorkItemCreateParams),
	"DbCacheDemoTwoWorkItemJoins":                       new(db.CacheDemoTwoWorkItemJoins),
	"DbCacheDemoTwoWorkItemM2MAssigneeWIA":              new(db.CacheDemoTwoWorkItemM2MAssigneeWIA),
	"DbCacheDemoTwoWorkItemSelectConfig":                new(db.CacheDemoTwoWorkItemSelectConfig),
	"DbCacheDemoTwoWorkItemUpdateParams":                new(db.CacheDemoTwoWorkItemUpdateParams),
	"DbCacheDemoWorkItem":                               new(db.CacheDemoWorkItem),
	"DbCacheDemoWorkItemCreateParams":                   new(db.CacheDemoWorkItemCreateParams),
	"DbCacheDemoWorkItemJoins":                          new(db.CacheDemoWorkItemJoins),
	"DbCacheDemoWorkItemM2MAssigneeWIA":                 new(db.CacheDemoWorkItemM2MAssigneeWIA),
	"DbCacheDemoWorkItemSelectConfig":                   new(db.CacheDemoWorkItemSelectConfig),
	"DbCacheDemoWorkItemUpdateParams":                   new(db.CacheDemoWorkItemUpdateParams),
	"DbCursor":                                          new(db.Cursor),
	"DbDbField":                                         new(db.DbField),
	"DbDemoTwoWorkItem":                                 new(db.DemoTwoWorkItem),
	"DbDemoTwoWorkItemCreateParams":                     new(db.DemoTwoWorkItemCreateParams),
	"DbDemoTwoWorkItemJoins":                            new(db.DemoTwoWorkItemJoins),
	"DbDemoTwoWorkItemSelectConfig":                     new(db.DemoTwoWorkItemSelectConfig),
	"DbDemoTwoWorkItemUpdateParams":                     new(db.DemoTwoWorkItemUpdateParams),
	"DbDemoWorkItem":                                    new(db.DemoWorkItem),
	"DbDemoWorkItemCreateParams":                        new(db.DemoWorkItemCreateParams),
	"DbDemoWorkItemJoins":                               new(db.DemoWorkItemJoins),
	"DbDemoWorkItemSelectConfig":                        new(db.DemoWorkItemSelectConfig),
	"DbDemoWorkItemUpdateParams":                        new(db.DemoWorkItemUpdateParams),
	"DbEntityNotification":                              new(db.EntityNotification),
	"DbEntityNotificationCreateParams":                  new(db.EntityNotificationCreateParams),
	"DbEntityNotificationJoins":                         new(db.EntityNotificationJoins),
	"DbEntityNotificationSelectConfig":                  new(db.EntityNotificationSelectConfig),
	"DbEntityNotificationUpdateParams":                  new(db.EntityNotificationUpdateParams),
	"DbErrInsertFailed":                                 new(db.ErrInsertFailed),
	"DbErrUpdateFailed":                                 new(db.ErrUpdateFailed),
	"DbErrUpsertFailed":                                 new(db.ErrUpsertFailed),
	"DbExtraSchemaBook":                                 new(db.ExtraSchemaBook),
	"DbExtraSchemaBookAuthor":                           new(db.ExtraSchemaBookAuthor),
	"DbExtraSchemaBookAuthorCreateParams":               new(db.ExtraSchemaBookAuthorCreateParams),
	"DbExtraSchemaBookAuthorJoins":                      new(db.ExtraSchemaBookAuthorJoins),
	"DbExtraSchemaBookAuthorM2MAuthorBA":                new(db.ExtraSchemaBookAuthorM2MAuthorBA),
	"DbExtraSchemaBookAuthorM2MBookBA":                  new(db.ExtraSchemaBookAuthorM2MBookBA),
	"DbExtraSchemaBookAuthorSelectConfig":               new(db.ExtraSchemaBookAuthorSelectConfig),
	"DbExtraSchemaBookAuthorUpdateParams":               new(db.ExtraSchemaBookAuthorUpdateParams),
	"DbExtraSchemaBookAuthorsSurrogateKey":              new(db.ExtraSchemaBookAuthorsSurrogateKey),
	"DbExtraSchemaBookAuthorsSurrogateKeyCreateParams":  new(db.ExtraSchemaBookAuthorsSurrogateKeyCreateParams),
	"DbExtraSchemaBookAuthorsSurrogateKeyJoins":         new(db.ExtraSchemaBookAuthorsSurrogateKeyJoins),
	"DbExtraSchemaBookAuthorsSurrogateKeyM2MAuthorBASK": new(db.ExtraSchemaBookAuthorsSurrogateKeyM2MAuthorBASK),
	"DbExtraSchemaBookAuthorsSurrogateKeyM2MBookBASK":   new(db.ExtraSchemaBookAuthorsSurrogateKeyM2MBookBASK),
	"DbExtraSchemaBookAuthorsSurrogateKeySelectConfig":  new(db.ExtraSchemaBookAuthorsSurrogateKeySelectConfig),
	"DbExtraSchemaBookAuthorsSurrogateKeyUpdateParams":  new(db.ExtraSchemaBookAuthorsSurrogateKeyUpdateParams),
	"DbExtraSchemaBookCreateParams":                     new(db.ExtraSchemaBookCreateParams),
	"DbExtraSchemaBookJoins":                            new(db.ExtraSchemaBookJoins),
	"DbExtraSchemaBookM2MAuthorBA":                      new(db.ExtraSchemaBookM2MAuthorBA),
	"DbExtraSchemaBookM2MAuthorBASK":                    new(db.ExtraSchemaBookM2MAuthorBASK),
	"DbExtraSchemaBookReview":                           new(db.ExtraSchemaBookReview),
	"DbExtraSchemaBookReviewCreateParams":               new(db.ExtraSchemaBookReviewCreateParams),
	"DbExtraSchemaBookReviewJoins":                      new(db.ExtraSchemaBookReviewJoins),
	"DbExtraSchemaBookReviewSelectConfig":               new(db.ExtraSchemaBookReviewSelectConfig),
	"DbExtraSchemaBookReviewUpdateParams":               new(db.ExtraSchemaBookReviewUpdateParams),
	"DbExtraSchemaBookSelectConfig":                     new(db.ExtraSchemaBookSelectConfig),
	"DbExtraSchemaBookSeller":                           new(db.ExtraSchemaBookSeller),
	"DbExtraSchemaBookSellerCreateParams":               new(db.ExtraSchemaBookSellerCreateParams),
	"DbExtraSchemaBookSellerJoins":                      new(db.ExtraSchemaBookSellerJoins),
	"DbExtraSchemaBookSellerSelectConfig":               new(db.ExtraSchemaBookSellerSelectConfig),
	"DbExtraSchemaBookSellerUpdateParams":               new(db.ExtraSchemaBookSellerUpdateParams),
	"DbExtraSchemaBookUpdateParams":                     new(db.ExtraSchemaBookUpdateParams),
	"DbExtraSchemaDemoWorkItem":                         new(db.ExtraSchemaDemoWorkItem),
	"DbExtraSchemaDemoWorkItemCreateParams":             new(db.ExtraSchemaDemoWorkItemCreateParams),
	"DbExtraSchemaDemoWorkItemJoins":                    new(db.ExtraSchemaDemoWorkItemJoins),
	"DbExtraSchemaDemoWorkItemSelectConfig":             new(db.ExtraSchemaDemoWorkItemSelectConfig),
	"DbExtraSchemaDemoWorkItemUpdateParams":             new(db.ExtraSchemaDemoWorkItemUpdateParams),
	"DbExtraSchemaDummyJoin":                            new(db.ExtraSchemaDummyJoin),
	"DbExtraSchemaDummyJoinCreateParams":                new(db.ExtraSchemaDummyJoinCreateParams),
	"DbExtraSchemaDummyJoinJoins":                       new(db.ExtraSchemaDummyJoinJoins),
	"DbExtraSchemaDummyJoinSelectConfig":                new(db.ExtraSchemaDummyJoinSelectConfig),
	"DbExtraSchemaDummyJoinUpdateParams":                new(db.ExtraSchemaDummyJoinUpdateParams),
	"DbExtraSchemaNotification":                         new(db.ExtraSchemaNotification),
	"DbExtraSchemaNotificationCreateParams":             new(db.ExtraSchemaNotificationCreateParams),
	"DbExtraSchemaNotificationJoins":                    new(db.ExtraSchemaNotificationJoins),
	"DbExtraSchemaNotificationSelectConfig":             new(db.ExtraSchemaNotificationSelectConfig),
	"DbExtraSchemaNotificationUpdateParams":             new(db.ExtraSchemaNotificationUpdateParams),
	"DbExtraSchemaPagElement":                           new(db.ExtraSchemaPagElement),
	"DbExtraSchemaPagElementCreateParams":               new(db.ExtraSchemaPagElementCreateParams),
	"DbExtraSchemaPagElementID":                         new(db.ExtraSchemaPagElementID),
	"DbExtraSchemaPagElementJoins":                      new(db.ExtraSchemaPagElementJoins),
	"DbExtraSchemaPagElementSelectConfig":               new(db.ExtraSchemaPagElementSelectConfig),
	"DbExtraSchemaPagElementUpdateParams":               new(db.ExtraSchemaPagElementUpdateParams),
	"DbExtraSchemaUser":                                 new(db.ExtraSchemaUser),
	"DbExtraSchemaUserAPIKey":                           new(db.ExtraSchemaUserAPIKey),
	"DbExtraSchemaUserAPIKeyCreateParams":               new(db.ExtraSchemaUserAPIKeyCreateParams),
	"DbExtraSchemaUserAPIKeyJoins":                      new(db.ExtraSchemaUserAPIKeyJoins),
	"DbExtraSchemaUserAPIKeySelectConfig":               new(db.ExtraSchemaUserAPIKeySelectConfig),
	"DbExtraSchemaUserAPIKeyUpdateParams":               new(db.ExtraSchemaUserAPIKeyUpdateParams),
	"DbExtraSchemaUserCreateParams":                     new(db.ExtraSchemaUserCreateParams),
	"DbExtraSchemaUserID":                               new(db.ExtraSchemaUserID),
	"DbExtraSchemaUserJoins":                            new(db.ExtraSchemaUserJoins),
	"DbExtraSchemaUserM2MBookBA":                        new(db.ExtraSchemaUserM2MBookBA),
	"DbExtraSchemaUserM2MBookBASK":                      new(db.ExtraSchemaUserM2MBookBASK),
	"DbExtraSchemaUserM2MWorkItemWIA":                   new(db.ExtraSchemaUserM2MWorkItemWIA),
	"DbExtraSchemaUserSelectConfig":                     new(db.ExtraSchemaUserSelectConfig),
	"DbExtraSchemaUserUpdateParams":                     new(db.ExtraSchemaUserUpdateParams),
	"DbExtraSchemaWorkItem":                             new(db.ExtraSchemaWorkItem),
	"DbExtraSchemaWorkItemAdmin":                        new(db.ExtraSchemaWorkItemAdmin),
	"DbExtraSchemaWorkItemAdminCreateParams":            new(db.ExtraSchemaWorkItemAdminCreateParams),
	"DbExtraSchemaWorkItemAdminJoins":                   new(db.ExtraSchemaWorkItemAdminJoins),
	"DbExtraSchemaWorkItemAdminSelectConfig":            new(db.ExtraSchemaWorkItemAdminSelectConfig),
	"DbExtraSchemaWorkItemAdminUpdateParams":            new(db.ExtraSchemaWorkItemAdminUpdateParams),
	"DbExtraSchemaWorkItemAssignee":                     new(db.ExtraSchemaWorkItemAssignee),
	"DbExtraSchemaWorkItemAssigneeCreateParams":         new(db.ExtraSchemaWorkItemAssigneeCreateParams),
	"DbExtraSchemaWorkItemAssigneeJoins":                new(db.ExtraSchemaWorkItemAssigneeJoins),
	"DbExtraSchemaWorkItemAssigneeM2MAssigneeWIA":       new(db.ExtraSchemaWorkItemAssigneeM2MAssigneeWIA),
	"DbExtraSchemaWorkItemAssigneeM2MWorkItemWIA":       new(db.ExtraSchemaWorkItemAssigneeM2MWorkItemWIA),
	"DbExtraSchemaWorkItemAssigneeSelectConfig":         new(db.ExtraSchemaWorkItemAssigneeSelectConfig),
	"DbExtraSchemaWorkItemAssigneeUpdateParams":         new(db.ExtraSchemaWorkItemAssigneeUpdateParams),
	"DbExtraSchemaWorkItemCreateParams":                 new(db.ExtraSchemaWorkItemCreateParams),
	"DbExtraSchemaWorkItemJoins":                        new(db.ExtraSchemaWorkItemJoins),
	"DbExtraSchemaWorkItemM2MAssigneeWIA":               new(db.ExtraSchemaWorkItemM2MAssigneeWIA),
	"DbExtraSchemaWorkItemSelectConfig":                 new(db.ExtraSchemaWorkItemSelectConfig),
	"DbExtraSchemaWorkItemUpdateParams":                 new(db.ExtraSchemaWorkItemUpdateParams),
	"DbGetExtraSchemaNotificationsParams":               new(db.GetExtraSchemaNotificationsParams),
	"DbGetExtraSchemaNotificationsRow":                  new(db.GetExtraSchemaNotificationsRow),
	"DbGetUserNotificationsParams":                      new(db.GetUserNotificationsParams),
	"DbGetUserNotificationsRow":                         new(db.GetUserNotificationsRow),
	"DbGetUserParams":                                   new(db.GetUserParams),
	"DbGetUserRow":                                      new(db.GetUserRow),
	"DbIsTeamInProjectParams":                           new(db.IsTeamInProjectParams),
	"DbKanbanStep":                                      new(db.KanbanStep),
	"DbKanbanStepCreateParams":                          new(db.KanbanStepCreateParams),
	"DbKanbanStepJoins":                                 new(db.KanbanStepJoins),
	"DbKanbanStepSelectConfig":                          new(db.KanbanStepSelectConfig),
	"DbKanbanStepUpdateParams":                          new(db.KanbanStepUpdateParams),
	"DbMovie":                                           new(db.Movie),
	"DbMovieCreateParams":                               new(db.MovieCreateParams),
	"DbMovieJoins":                                      new(db.MovieJoins),
	"DbMovieSelectConfig":                               new(db.MovieSelectConfig),
	"DbMovieUpdateParams":                               new(db.MovieUpdateParams),
	"DbNotification":                                    new(db.Notification),
	"DbNotificationCreateParams":                        new(db.NotificationCreateParams),
	"DbNotificationJoins":                               new(db.NotificationJoins),
	"DbNotificationSelectConfig":                        new(db.NotificationSelectConfig),
	"DbNotificationUpdateParams":                        new(db.NotificationUpdateParams),
	"DbProject":                                         new(db.Project),
	"DbProjectCreateParams":                             new(db.ProjectCreateParams),
	"DbProjectJoins":                                    new(db.ProjectJoins),
	"DbProjectSelectConfig":                             new(db.ProjectSelectConfig),
	"DbProjectUpdateParams":                             new(db.ProjectUpdateParams),
	"DbQueries":                                         new(db.Queries),
	"DbTeam":                                            new(db.Team),
	"DbTeamCreateParams":                                new(db.TeamCreateParams),
	"DbTeamJoins":                                       new(db.TeamJoins),
	"DbTeamSelectConfig":                                new(db.TeamSelectConfig),
	"DbTeamUpdateParams":                                new(db.TeamUpdateParams),
	"DbTimeEntry":                                       new(db.TimeEntry),
	"DbTimeEntryCreateParams":                           new(db.TimeEntryCreateParams),
	"DbTimeEntryJoins":                                  new(db.TimeEntryJoins),
	"DbTimeEntrySelectConfig":                           new(db.TimeEntrySelectConfig),
	"DbTimeEntryUpdateParams":                           new(db.TimeEntryUpdateParams),
	"DbTrigger":                                         new(db.Trigger),
	"DbUser":                                            new(db.User),
	"DbUserAPIKey":                                      new(db.UserAPIKey),
	"DbUserAPIKeyCreateParams":                          new(db.UserAPIKeyCreateParams),
	"DbUserAPIKeyJoins":                                 new(db.UserAPIKeyJoins),
	"DbUserAPIKeySelectConfig":                          new(db.UserAPIKeySelectConfig),
	"DbUserAPIKeyUpdateParams":                          new(db.UserAPIKeyUpdateParams),
	"DbUserCreateParams":                                new(db.UserCreateParams),
	"DbUserID":                                          new(db.UserID),
	"DbUserJoins":                                       new(db.UserJoins),
	"DbUserM2MWorkItemWIA":                              new(db.UserM2MWorkItemWIA),
	"DbUserNotification":                                new(db.UserNotification),
	"DbUserNotificationCreateParams":                    new(db.UserNotificationCreateParams),
	"DbUserNotificationJoins":                           new(db.UserNotificationJoins),
	"DbUserNotificationSelectConfig":                    new(db.UserNotificationSelectConfig),
	"DbUserNotificationUpdateParams":                    new(db.UserNotificationUpdateParams),
	"DbUserProject":                                     new(db.UserProject),
	"DbUserProjectCreateParams":                         new(db.UserProjectCreateParams),
	"DbUserProjectJoins":                                new(db.UserProjectJoins),
	"DbUserProjectSelectConfig":                         new(db.UserProjectSelectConfig),
	"DbUserProjectUpdateParams":                         new(db.UserProjectUpdateParams),
	"DbUserSelectConfig":                                new(db.UserSelectConfig),
	"DbUserTeam":                                        new(db.UserTeam),
	"DbUserTeamCreateParams":                            new(db.UserTeamCreateParams),
	"DbUserTeamJoins":                                   new(db.UserTeamJoins),
	"DbUserTeamSelectConfig":                            new(db.UserTeamSelectConfig),
	"DbUserTeamUpdateParams":                            new(db.UserTeamUpdateParams),
	"DbUserUpdateParams":                                new(db.UserUpdateParams),
	"DbWorkItem":                                        new(db.WorkItem),
	"DbWorkItemAssignee":                                new(db.WorkItemAssignee),
	"DbWorkItemAssigneeCreateParams":                    new(db.WorkItemAssigneeCreateParams),
	"DbWorkItemAssigneeJoins":                           new(db.WorkItemAssigneeJoins),
	"DbWorkItemAssigneeM2MAssigneeWIA":                  new(db.WorkItemAssigneeM2MAssigneeWIA),
	"DbWorkItemAssigneeM2MWorkItemWIA":                  new(db.WorkItemAssigneeM2MWorkItemWIA),
	"DbWorkItemAssigneeSelectConfig":                    new(db.WorkItemAssigneeSelectConfig),
	"DbWorkItemAssigneeUpdateParams":                    new(db.WorkItemAssigneeUpdateParams),
	"DbWorkItemComment":                                 new(db.WorkItemComment),
	"DbWorkItemCommentCreateParams":                     new(db.WorkItemCommentCreateParams),
	"DbWorkItemCommentJoins":                            new(db.WorkItemCommentJoins),
	"DbWorkItemCommentSelectConfig":                     new(db.WorkItemCommentSelectConfig),
	"DbWorkItemCommentUpdateParams":                     new(db.WorkItemCommentUpdateParams),
	"DbWorkItemCreateParams":                            new(db.WorkItemCreateParams),
	"DbWorkItemJoins":                                   new(db.WorkItemJoins),
	"DbWorkItemM2MAssigneeWIA":                          new(db.WorkItemM2MAssigneeWIA),
	"DbWorkItemSelectConfig":                            new(db.WorkItemSelectConfig),
	"DbWorkItemTag":                                     new(db.WorkItemTag),
	"DbWorkItemTagCreateParams":                         new(db.WorkItemTagCreateParams),
	"DbWorkItemTagJoins":                                new(db.WorkItemTagJoins),
	"DbWorkItemTagSelectConfig":                         new(db.WorkItemTagSelectConfig),
	"DbWorkItemTagUpdateParams":                         new(db.WorkItemTagUpdateParams),
	"DbWorkItemType":                                    new(db.WorkItemType),
	"DbWorkItemTypeCreateParams":                        new(db.WorkItemTypeCreateParams),
	"DbWorkItemTypeJoins":                               new(db.WorkItemTypeJoins),
	"DbWorkItemTypeSelectConfig":                        new(db.WorkItemTypeSelectConfig),
	"DbWorkItemTypeUpdateParams":                        new(db.WorkItemTypeUpdateParams),
	"DbWorkItemUpdateParams":                            new(db.WorkItemUpdateParams),
	"DbWorkItemWorkItemTag":                             new(db.WorkItemWorkItemTag),
	"DbWorkItemWorkItemTagCreateParams":                 new(db.WorkItemWorkItemTagCreateParams),
	"DbWorkItemWorkItemTagJoins":                        new(db.WorkItemWorkItemTagJoins),
	"DbWorkItemWorkItemTagSelectConfig":                 new(db.WorkItemWorkItemTagSelectConfig),
	"DbWorkItemWorkItemTagUpdateParams":                 new(db.WorkItemWorkItemTagUpdateParams),
	"DbXoError":                                         new(db.XoError),

	//

	"Activity":                            new(rest.Activity),
	"CacheDemoWorkItem":                   new(rest.CacheDemoWorkItem),
	"CreateActivityRequest":               new(rest.CreateActivityRequest),
	"CreateDemoTwoWorkItemRequest":        new(rest.CreateDemoTwoWorkItemRequest),
	"CreateDemoWorkItemRequest":           new(rest.CreateDemoWorkItemRequest),
	"CreateProjectBoardRequest":           new(rest.CreateProjectBoardRequest),
	"CreateTeamRequest":                   new(rest.CreateTeamRequest),
	"CreateTimeEntryRequest":              new(rest.CreateTimeEntryRequest),
	"CreateWorkItemCommentRequest":        new(rest.CreateWorkItemCommentRequest),
	"CreateWorkItemTagRequest":            new(rest.CreateWorkItemTagRequest),
	"CreateWorkItemTypeRequest":           new(rest.CreateWorkItemTypeRequest),
	"DemoTwoWorkItem":                     new(rest.DemoTwoWorkItem),
	"DemoWorkItem":                        new(rest.DemoWorkItem),
	"GetCacheDemoWorkItemQueryParameters": new(rest.GetCacheDemoWorkItemQueryParameters),
	"GetCurrentUserQueryParameters":       new(rest.GetCurrentUserQueryParameters),
	"Notification":                        new(rest.Notification),
	"PaginatedDemoWorkItemsResponse":      new(rest.PaginatedDemoWorkItemsResponse),
	"PaginatedNotificationsResponse":      new(rest.PaginatedNotificationsResponse),
	"PaginatedUsersResponse":              new(rest.PaginatedUsersResponse),
	"PaginationPage":                      new(rest.PaginationPage),
	"ProjectBoard":                        new(rest.ProjectBoard),
	"SharedWorkItemJoins":                 new(rest.SharedWorkItemJoins),
	"Team":                                new(rest.Team),
	"TimeEntry":                           new(rest.TimeEntry),
	"UpdateActivityRequest":               new(rest.UpdateActivityRequest),
	"UpdateTeamRequest":                   new(rest.UpdateTeamRequest),
	"UpdateTimeEntryRequest":              new(rest.UpdateTimeEntryRequest),
	"UpdateWorkItemCommentRequest":        new(rest.UpdateWorkItemCommentRequest),
	"UpdateWorkItemTagRequest":            new(rest.UpdateWorkItemTagRequest),
	"UpdateWorkItemTypeRequest":           new(rest.UpdateWorkItemTypeRequest),
	"User":                                new(rest.User),
	"WorkItemBase":                        new(rest.WorkItemBase),
	"WorkItemComment":                     new(rest.WorkItemComment),
	"WorkItemTag":                         new(rest.WorkItemTag),
	"WorkItemType":                        new(rest.WorkItemType),

	//

}
