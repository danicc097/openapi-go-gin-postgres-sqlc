package repos

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	repomodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/google/uuid"
)

// Boards limited to one per project. All teams in a project share the same board.
// IMPORTANT: If a board does not exist for a project, it will not be possible
// to create teams, activities, etc. The first time we must go through the project
// creation steps and create everything at once.
// Later on everything can be updated in the project settings panel, and new elements created.
type ProjectBoardCreateParams struct {
	ProjectID int `json:"projectID"`
	// TeamIDs   []int `json:"teamIDs"` // completely useless. the only check needed is to ensure at least one team
	// is associated to projectID, else prompt the user to create >=1 team before creating a board, handled at /teams/:project_name/create.
	Activities   []db.ActivityCreateParams    `json:"activities"`
	Teams        []db.TeamCreateParams        `json:"teams"`
	WorkItemTags []db.WorkItemTagCreateParams `json:"workItemTags"`
}

// ProjectBoard defines the datastore/repository handling persisting ProjectBoard records.
type ProjectBoard interface {
	/*
		Create corresponds to initial info to be filled in once a project table has been manually
		 created, before it can be used:
		 - kanban columns and their info (order, name, can log time, etc.)
		 - types of workitems (by all teams)
		 - initial teams associated (at least 1 id initially)

		  If we manually added a new project record
		 in migrations (insert into projects) and created its specific work_items_<new project> table
			at this point everything is empty and project.is_setup is False

			!project.is_setup --> dashboard shows single centered [+ Initialize project] button
			and we let the project admins create teams, work item types, etc. all at once.
			if form submitted successfully, is_setup=False and won't show again. Else, the whole thing
			is rolled back at once (should try to save this form state with zustand's persist just in case but not important)


	*/
	Create(ctx context.Context, d db.DBTX, params ProjectBoardCreateParams) (*repomodels.ProjectBoard, error)
	ByID(ctx context.Context, d db.DBTX, projectID int) (*repomodels.ProjectBoard, error)
}

type DemoWorkItemUpdateParams struct {
	DemoProject *db.DemoWorkItemUpdateParams `json:"demoProject"`
	Base        *db.WorkItemUpdateParams     `json:"base"`
}

type DemoWorkItemCreateParams struct {
	DemoProject db.DemoWorkItemCreateParams `json:"demoProject" required:"true"`
	Base        db.WorkItemCreateParams     `json:"base"        required:"true"`
}

type DemoTwoWorkItemUpdateParams struct {
	DemoTwoProject *db.DemoTwoWorkItemUpdateParams `json:"demoTwoProject"`
	Base           *db.WorkItemUpdateParams        `json:"base"`
}

type DemoTwoWorkItemCreateParams struct {
	DemoTwoProject db.DemoTwoWorkItemCreateParams `json:"demoTwoProject" required:"true"`
	Base           db.WorkItemCreateParams        `json:"base"           required:"true"`
}

/**
 *
 * TODO: all By's should accept db opts
 */

// WorkItem defines the datastore/repository handling retrieving WorkItem records.
type WorkItem interface {
	// ByID returns a generic WorkItem by default.
	ByID(ctx context.Context, d db.DBTX, id int, opts ...db.WorkItemSelectConfigOption) (*db.WorkItem, error)
}

// DemoWorkItem defines the datastore/repository handling persisting DemoWorkItem records.
type DemoWorkItem interface {
	// ByID returns a generic WorkItem with project-specific fields joined by default.
	ByID(ctx context.Context, d db.DBTX, id int, opts ...db.WorkItemSelectConfigOption) (*db.WorkItem, error)
	// params for dedicated workItem only require workItemID (FK-as-PK)
	Create(ctx context.Context, d db.DBTX, params DemoWorkItemCreateParams) (*db.WorkItem, error)
	Update(ctx context.Context, d db.DBTX, id int, params DemoWorkItemUpdateParams) (*db.WorkItem, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItem, error)
	Restore(ctx context.Context, d db.DBTX, id int) (*db.WorkItem, error)
}

// DemoTwoWorkItem defines the datastore/repository handling persisting DemoTwoWorkItem records.
type DemoTwoWorkItem interface {
	// ByID returns a generic WorkItem with project-specific fields joined by default.
	ByID(ctx context.Context, d db.DBTX, id int, opts ...db.WorkItemSelectConfigOption) (*db.WorkItem, error)
	// params for dedicated workItem only require workItemID (FK-as-PK)
	Create(ctx context.Context, d db.DBTX, params DemoTwoWorkItemCreateParams) (*db.WorkItem, error)
	Update(ctx context.Context, d db.DBTX, id int, params DemoTwoWorkItemUpdateParams) (*db.WorkItem, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItem, error)
	Restore(ctx context.Context, d db.DBTX, id int) (*db.WorkItem, error)
}

// Notification defines the datastore/repository handling persisting Notification records.
type Notification interface {
	LatestUserNotifications(ctx context.Context, d db.DBTX, params *db.GetUserNotificationsParams) ([]db.GetUserNotificationsRow, error)
	Create(ctx context.Context, d db.DBTX, params *db.NotificationCreateParams) (*db.Notification, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.Notification, error)
}

// User defines the datastore/repository handling persisting User records.
type User interface {
	ByID(ctx context.Context, d db.DBTX, id uuid.UUID, opts ...db.UserSelectConfigOption) (*db.User, error)
	ByEmail(ctx context.Context, d db.DBTX, email string, opts ...db.UserSelectConfigOption) (*db.User, error)
	ByUsername(ctx context.Context, d db.DBTX, username string, opts ...db.UserSelectConfigOption) (*db.User, error)
	ByExternalID(ctx context.Context, d db.DBTX, extID string, opts ...db.UserSelectConfigOption) (*db.User, error)
	ByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.User, error)
	Create(ctx context.Context, d db.DBTX, params *db.UserCreateParams) (*db.User, error)
	Update(ctx context.Context, d db.DBTX, id uuid.UUID, params *db.UserUpdateParams) (*db.User, error)
	Delete(ctx context.Context, d db.DBTX, id uuid.UUID) (*db.User, error)
	// CreateAPIKey requires an existing user.
	CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (*db.UserAPIKey, error)
}

// Project defines the datastore/repository handling persisting Project records.
// Projects are manually created on demand.
// NOTE: Read-only. Managed via migrations.
type Project interface {
	ByName(ctx context.Context, d db.DBTX, name models.Project, opts ...db.ProjectSelectConfigOption) (*db.Project, error)
	ByID(ctx context.Context, d db.DBTX, id int, opts ...db.ProjectSelectConfigOption) (*db.Project, error)
}

// Team defines the datastore/repository handling persisting Team records.
type Team interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.Team, error)
	ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.Team, error)
	Create(ctx context.Context, d db.DBTX, params *db.TeamCreateParams) (*db.Team, error)
	Update(ctx context.Context, d db.DBTX, id int, params *db.TeamUpdateParams) (*db.Team, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.Team, error)
}

// KanbanStep defines the datastore/repository handling persisting KanbanStep records.
// NOTE: Read-only. Managed via migrations.
type KanbanStep interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.KanbanStep, error)
	ByProject(ctx context.Context, d db.DBTX, projectID int) ([]db.KanbanStep, error)
}

// WorkItemType defines the datastore/repository handling persisting WorkItemType records.
// NOTE: Read-only. Managed via migrations.
type WorkItemType interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemType, error)
	// TODO ByProjectID(ctx context.Context, d db.DBTX, id int) ([]*db.WorkItemType, error)
	ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.WorkItemType, error)
}

// WorkItemComment defines the datastore/repository handling persisting WorkItemComment records.
type WorkItemComment interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemComment, error)
	Create(ctx context.Context, d db.DBTX, params *db.WorkItemCommentCreateParams) (*db.WorkItemComment, error)
	Update(ctx context.Context, d db.DBTX, id int, params *db.WorkItemCommentUpdateParams) (*db.WorkItemComment, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItemComment, error)
}

// WorkItemTag defines the datastore/repository handling persisting WorkItemTag records.
type WorkItemTag interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error)
	// TODO ByProjectID(ctx context.Context, d db.DBTX, id int) ([]*db.WorkItemTag, error)
	ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.WorkItemTag, error)
	Create(ctx context.Context, d db.DBTX, params *db.WorkItemTagCreateParams) (*db.WorkItemTag, error)
	Update(ctx context.Context, d db.DBTX, id int, params *db.WorkItemTagUpdateParams) (*db.WorkItemTag, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error)
}

// Activity defines the datastore/repository handling persisting Activity records.
type Activity interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.Activity, error)
	ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.Activity, error)
	ByProjectID(ctx context.Context, d db.DBTX, projectID int) ([]db.Activity, error)
	Create(ctx context.Context, d db.DBTX, params *db.ActivityCreateParams) (*db.Activity, error)
	Update(ctx context.Context, d db.DBTX, id int, params *db.ActivityUpdateParams) (*db.Activity, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.Activity, error)
}

// TimeEntry defines the datastore/repository handling persisting TimeEntry records.
type TimeEntry interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.TimeEntry, error)
	Create(ctx context.Context, d db.DBTX, params *db.TimeEntryCreateParams) (*db.TimeEntry, error)
	Update(ctx context.Context, d db.DBTX, id int, params *db.TimeEntryUpdateParams) (*db.TimeEntry, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.TimeEntry, error)
}
