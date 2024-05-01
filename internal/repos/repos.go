package repos

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

type GetPaginatedUsersParams struct {
	Limit     int
	Direction models.Direction
	Cursor    models.PaginationCursor

	Items *models.PaginationItems

	RoleRank *int
}

type DemoWorkItemUpdateParams struct {
	DemoProject *models.DemoWorkItemUpdateParams `json:"demoProject"`
	Base        *models.WorkItemUpdateParams     `json:"base"`
}

type DemoWorkItemCreateParams struct {
	DemoProject models.DemoWorkItemCreateParams `json:"demoProject" required:"true"`
	Base        models.WorkItemCreateParams     `json:"base"        required:"true"`
}

type DemoTwoWorkItemUpdateParams struct {
	DemoTwoProject *models.DemoTwoWorkItemUpdateParams `json:"demoTwoProject"`
	Base           *models.WorkItemUpdateParams        `json:"base"`
}

type DemoTwoWorkItemCreateParams struct {
	DemoTwoProject models.DemoTwoWorkItemCreateParams `json:"demoTwoProject" required:"true"`
	Base           models.WorkItemCreateParams        `json:"base"           required:"true"`
}

// WorkItem defines the datastore/repository handling retrieving WorkItem records.
type WorkItem interface {
	// ByID returns a generic WorkItem by default.
	ByID(ctx context.Context, d models.DBTX, id models.WorkItemID, opts ...models.WorkItemSelectConfigOption) (*models.WorkItem, error)
	Delete(ctx context.Context, d models.DBTX, id models.WorkItemID) (*models.WorkItem, error)
	Restore(ctx context.Context, d models.DBTX, id models.WorkItemID) (*models.WorkItem, error)
	AssignUser(ctx context.Context, d models.DBTX, params *models.WorkItemAssigneeCreateParams) error
	RemoveAssignedUser(ctx context.Context, d models.DBTX, memberID models.UserID, workItemID models.WorkItemID) error
	AssignTag(ctx context.Context, d models.DBTX, params *models.WorkItemWorkItemTagCreateParams) error
	RemoveTag(ctx context.Context, d models.DBTX, tagID models.WorkItemTagID, workItemID models.WorkItemID) error
}

// DemoWorkItem defines the datastore/repository handling persisting DemoWorkItem records.
type DemoWorkItem interface {
	// ByID returns a generic WorkItem with project-specific fields joined by default.
	ByID(ctx context.Context, d models.DBTX, id models.WorkItemID, opts ...models.WorkItemSelectConfigOption) (*models.WorkItem, error)
	Paginated(ctx context.Context, d models.DBTX, cursor models.WorkItemID, opts ...models.CacheDemoWorkItemSelectConfigOption) ([]models.CacheDemoWorkItem, error)
	// params for dedicated workItem only require workItemID (FK-as-PK)
	Create(ctx context.Context, d models.DBTX, params DemoWorkItemCreateParams) (*models.WorkItem, error)
	Update(ctx context.Context, d models.DBTX, id models.WorkItemID, params DemoWorkItemUpdateParams) (*models.WorkItem, error)
}

// DemoTwoWorkItem defines the datastore/repository handling persisting DemoTwoWorkItem records.
type DemoTwoWorkItem interface {
	// ByID returns a generic WorkItem with project-specific fields joined by default.
	ByID(ctx context.Context, d models.DBTX, id models.WorkItemID, opts ...models.WorkItemSelectConfigOption) (*models.WorkItem, error)
	Create(ctx context.Context, d models.DBTX, params DemoTwoWorkItemCreateParams) (*models.WorkItem, error)
	Update(ctx context.Context, d models.DBTX, id models.WorkItemID, params DemoTwoWorkItemUpdateParams) (*models.WorkItem, error)
}

// Notification defines the datastore/repository handling persisting Notification records.
type Notification interface {
	// now can replace GetUserNotifications with `...WithFilters`, in lieu of sqlc
	LatestNotifications(ctx context.Context, d models.DBTX, params *models.GetUserNotificationsParams) ([]models.GetUserNotificationsRow, error)
	PaginatedUserNotifications(ctx context.Context, d models.DBTX, userID models.UserID, params models.GetPaginatedNotificationsParams) ([]models.UserNotification, error)
	Create(ctx context.Context, d models.DBTX, params *models.NotificationCreateParams) (*models.UserNotification, error)
	Delete(ctx context.Context, d models.DBTX, id models.NotificationID) (*models.Notification, error)
}

// User defines the datastore/repository handling persisting User records.
type User interface {
	ByID(ctx context.Context, d models.DBTX, id models.UserID, opts ...models.UserSelectConfigOption) (*models.User, error)
	ByTeam(ctx context.Context, d models.DBTX, teamID models.TeamID) ([]models.User, error)
	ByProject(ctx context.Context, d models.DBTX, projectID models.ProjectID) ([]models.User, error)
	ByEmail(ctx context.Context, d models.DBTX, email string, opts ...models.UserSelectConfigOption) (*models.User, error)
	ByUsername(ctx context.Context, d models.DBTX, username string, opts ...models.UserSelectConfigOption) (*models.User, error)
	ByExternalID(ctx context.Context, d models.DBTX, extID string, opts ...models.UserSelectConfigOption) (*models.User, error)
	ByAPIKey(ctx context.Context, d models.DBTX, apiKey string) (*models.User, error)
	Paginated(ctx context.Context, d models.DBTX, params GetPaginatedUsersParams) ([]models.User, error)
	Create(ctx context.Context, d models.DBTX, params *models.UserCreateParams) (*models.User, error)
	Update(ctx context.Context, d models.DBTX, id models.UserID, params *models.UserUpdateParams) (*models.User, error)
	Delete(ctx context.Context, d models.DBTX, id models.UserID) (*models.User, error)
	// CreateAPIKey requires an existing user.
	CreateAPIKey(ctx context.Context, d models.DBTX, user *models.User) (*models.UserAPIKey, error)
	DeleteAPIKey(ctx context.Context, d models.DBTX, apiKey string) (*models.UserAPIKey, error)
}

// Project defines the datastore/repository handling persisting Project records.
// Projects are manually created on demand.
// NOTE: Read-only. Managed via migrations.
type Project interface {
	ByName(ctx context.Context, d models.DBTX, name models.ProjectName, opts ...models.ProjectSelectConfigOption) (*models.Project, error)
	ByID(ctx context.Context, d models.DBTX, id models.ProjectID, opts ...models.ProjectSelectConfigOption) (*models.Project, error)
	IsTeamInProject(ctx context.Context, d models.DBTX, arg models.IsTeamInProjectParams) (bool, error)
}

// Team defines the datastore/repository handling persisting Team records.
type Team interface {
	ByID(ctx context.Context, d models.DBTX, id models.TeamID, opts ...models.TeamSelectConfigOption) (*models.Team, error)
	ByName(ctx context.Context, d models.DBTX, name string, projectID models.ProjectID, opts ...models.TeamSelectConfigOption) (*models.Team, error)
	Create(ctx context.Context, d models.DBTX, params *models.TeamCreateParams) (*models.Team, error)
	Update(ctx context.Context, d models.DBTX, id models.TeamID, params *models.TeamUpdateParams) (*models.Team, error)
	Delete(ctx context.Context, d models.DBTX, id models.TeamID) (*models.Team, error)
}

// KanbanStep defines the datastore/repository handling persisting KanbanStep records.
// NOTE: Read-only. Managed via migrations.
type KanbanStep interface {
	ByID(ctx context.Context, d models.DBTX, id models.KanbanStepID, opts ...models.KanbanStepSelectConfigOption) (*models.KanbanStep, error)
	ByProject(ctx context.Context, d models.DBTX, projectID models.ProjectID, opts ...models.KanbanStepSelectConfigOption) ([]models.KanbanStep, error)
}

// WorkItemType defines the datastore/repository handling persisting WorkItemType records.
// NOTE: Read-only. Managed via migrations.
type WorkItemType interface {
	ByID(ctx context.Context, d models.DBTX, id models.WorkItemTypeID, opts ...models.WorkItemTypeSelectConfigOption) (*models.WorkItemType, error)
	// TODO ByProjectID(ctx context.Context, d db.DBTX, id int) ([]*db.WorkItemType, error)
	ByName(ctx context.Context, d models.DBTX, name string, projectID models.ProjectID, opts ...models.WorkItemTypeSelectConfigOption) (*models.WorkItemType, error)
}

// WorkItemTag defines the datastore/repository handling persisting WorkItemTag records.
type WorkItemTag interface {
	ByID(ctx context.Context, d models.DBTX, id models.WorkItemTagID, opts ...models.WorkItemTagSelectConfigOption) (*models.WorkItemTag, error)
	// TODO ByProjectID(ctx context.Context, d db.DBTX, id db.WorkItemTagID) ([]*db.WorkItemTag, error)
	ByName(ctx context.Context, d models.DBTX, name string, projectID models.ProjectID, opts ...models.WorkItemTagSelectConfigOption) (*models.WorkItemTag, error)
	Create(ctx context.Context, d models.DBTX, params *models.WorkItemTagCreateParams) (*models.WorkItemTag, error)
	Update(ctx context.Context, d models.DBTX, id models.WorkItemTagID, params *models.WorkItemTagUpdateParams) (*models.WorkItemTag, error)
	Delete(ctx context.Context, d models.DBTX, id models.WorkItemTagID) (*models.WorkItemTag, error)
}

// Activity defines the datastore/repository handling persisting Activity records.
type Activity interface {
	ByID(ctx context.Context, d models.DBTX, id models.ActivityID, opts ...models.ActivitySelectConfigOption) (*models.Activity, error)
	ByName(ctx context.Context, d models.DBTX, name string, projectID models.ProjectID, opts ...models.ActivitySelectConfigOption) (*models.Activity, error)
	ByProjectID(ctx context.Context, d models.DBTX, projectID models.ProjectID, opts ...models.ActivitySelectConfigOption) ([]models.Activity, error)
	Create(ctx context.Context, d models.DBTX, params *models.ActivityCreateParams) (*models.Activity, error)
	Update(ctx context.Context, d models.DBTX, id models.ActivityID, params *models.ActivityUpdateParams) (*models.Activity, error)
	Delete(ctx context.Context, d models.DBTX, id models.ActivityID) (*models.Activity, error)
	Restore(ctx context.Context, d models.DBTX, id models.ActivityID) error
}

// TimeEntry defines the datastore/repository handling persisting TimeEntry records.
type TimeEntry interface {
	ByID(ctx context.Context, d models.DBTX, id models.TimeEntryID, opts ...models.TimeEntrySelectConfigOption) (*models.TimeEntry, error)
	Create(ctx context.Context, d models.DBTX, params *models.TimeEntryCreateParams) (*models.TimeEntry, error)
	Update(ctx context.Context, d models.DBTX, id models.TimeEntryID, params *models.TimeEntryUpdateParams) (*models.TimeEntry, error)
	Delete(ctx context.Context, d models.DBTX, id models.TimeEntryID) (*models.TimeEntry, error)
}

// WorkItemComment defines the datastore/repository handling persisting work item comment records.
type WorkItemComment interface {
	ByID(ctx context.Context, d models.DBTX, id models.WorkItemCommentID, opts ...models.WorkItemCommentSelectConfigOption) (*models.WorkItemComment, error)
	Create(ctx context.Context, d models.DBTX, params *models.WorkItemCommentCreateParams) (*models.WorkItemComment, error)
	Update(ctx context.Context, d models.DBTX, id models.WorkItemCommentID, params *models.WorkItemCommentUpdateParams) (*models.WorkItemComment, error)
	Delete(ctx context.Context, d models.DBTX, id models.WorkItemCommentID) (*models.WorkItemComment, error)
}
