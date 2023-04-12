package repos

import (
	"context"

	internalmodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/google/uuid"
)

// Boards limited to one per project. All teams in a project share the same board.
// IMPORTANT: If a board does not exist for a project, it will not be possible
// to create teams, activities, etc. The first time we must go through the project
// creation steps and create everything at once.
// Later on everything can be updated in the project settings panel, and new elements created.
// Once a board is created it cannot be deleted.
type ProjectBoardCreateParams struct {
	ProjectID int `json:"projectID"`
	// TeamIDs   []int `json:"teamIDs"` // completely useless. the only check needed is to ensure at least one team
	// exacts associated to projectID, else prompt the user to create at least 1 team before creating a board.
	Activities    []ActivityCreateParams     `json:"activities"`
	KanbanSteps   []KanbanStepCreateParams   `json:"kanbanSteps"`
	Teams         []TeamCreateParams         `json:"teams"`
	WorkItemTypes []WorkItemTypeCreateParams `json:"workItemTypes"`
	WorkItemTags  []WorkItemTagCreateParams  `json:"workItemTags"`
}

type WorkItemTagCreateParams struct {
	ProjectID   int    `json:"projectID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type WorkItemTagUpdateParams struct {
	ProjectID   *int    `json:"projectID"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Color       *string `json:"color"`
}

type WorkItemTypeCreateParams struct {
	ProjectID   int    `json:"projectID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type WorkItemTypeUpdateParams struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Color       *string `json:"color"`
}

type ActivityCreateParams struct {
	ProjectID    int    `json:"projectID"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsProductive bool   `json:"isProductive"`
}

type ActivityUpdateParams struct {
	Name         *string `json:"name"`
	Description  *string `json:"description"`
	IsProductive *bool   `json:"isProductive"`
}

type KanbanStepCreateParams struct {
	ProjectID     int    `json:"projectID"`
	StepOrder     *int16 `json:"stepOrder"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Color         string `json:"color"`
	TimeTrackable bool   `json:"timeTrackable"`
}

type KanbanStepUpdateParams struct {
	StepOrder     *int16  `json:"stepOrder"` // if StepOrder is changed, all that happens is current workitems are swapped visually
	Name          *string `json:"name"`
	Description   *string `json:"description"`
	Color         *string `json:"color"`
	TimeTrackable *bool   `json:"timeTrackable"` // if TimeTrackable is changed, already existing items won't change
}

type TeamCreateParams struct {
	ProjectID   int    `json:"projectID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TeamUpdateParams struct {
	Name        *string
	Description *string
}

type (
	GetUserNotificationsParams = db.GetUserNotificationsParams
	NotificationCreateParams   = db.CreateNotificationParams
)

type UserCreateParams struct {
	Username   string
	Email      string
	FirstName  *string
	LastName   *string
	ExternalID string
	Scopes     []string
	RoleRank   int16

	// what insert generates
	// Username
	// Email
	// FirstName
	// LastName
	// ExternalID
	// APIKeyID
	// Scopes
	// RoleRank
	// HasPersonalNotifications
	// HasGlobalNotifications
	// DeletedAt

	// Username                 string     `json:"username" db:"username" required:"true"`                                   // username
	// Email                    string     `json:"email" db:"email" required:"true"`                                         // email
	// FirstName                *string    `json:"firstName" db:"first_name" required:"true"`                                // first_name
	// LastName                 *string    `json:"lastName" db:"last_name" required:"true"`                                  // last_name
	// FullName                 *string    `json:"fullName" db:"full_name" required:"true"`                                  // full_name
	// ExternalID               string     `json:"-" db:"external_id" `                                                      // external_id
	// APIKeyID                 *int       `json:"-" db:"api_key_id" `                                                       // api_key_id
	// Scopes                   []string   `json:"-" db:"scopes" `                                                           // scopes
	// RoleRank                 int16      `json:"-" db:"role_rank" `                                                        // role_rank
	// HasPersonalNotifications bool       `json:"hasPersonalNotifications" db:"has_personal_notifications" required:"true"` // has_personal_notifications
	// HasGlobalNotifications   bool       `json:"hasGlobalNotifications" db:"has_global_notifications" required:"true"`     // has_global_notifications
	// CreatedAt                time.Time  `json:"createdAt" db:"created_at" required:"true"`                                // created_at
	// UpdatedAt                time.Time  `json:"-" db:"updated_at" `                                                       // updated_at
	// DeletedAt                *time.Time `json:"deletedAt" db:"deleted_at" required:"true"`                                // deleted_at

	//
	/**
		 *  TODO custom createParams changes to what currently gets generated by Insert():
		 *
		 * - skip deleted_at column always
		 * - skip fields from foreign keys, e.g. api_key_id:
	  			foreign key (api_key_id) references user_api_keys (user_api_key_id) on delete cascade
		* they will be manually updated
	*/

	/**
	 *  TODO custom updateParams:
	 * - same as insert but all pointers
	 */

	// TODO xo MarkAsDeleted() method
}

type UserUpdateParams struct {
	FirstName                *string
	LastName                 *string
	Rank                     *int16
	Scopes                   *[]string
	HasPersonalNotifications *bool
	HasGlobalNotifications   *bool
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
	Create(ctx context.Context, d db.DBTX, params ProjectBoardCreateParams) (*models.ProjectBoard, error)
	ByID(ctx context.Context, d db.DBTX, projectID int) (*models.ProjectBoard, error)
}

// DemoProjectWorkItem defines the datastore/repository handling persisting DemoProjectWorkItem records.
type DemoProjectWorkItem interface {
	ByID(ctx context.Context, d db.DBTX, id int64, opts ...db.DemoProjectWorkItemSelectConfigOption) (*db.DemoProjectWorkItem, error)
	// Create,
	// Delete,
	// ByTeam(closed bool, deleted bool)
	// Update (service has Close (Update with closed=True), Move(Update with kanban step change), ...)
	// TBD if useful: ByTag, ByType (for closed workitem searches. open ones simply return everything and filter in client)
}

// Notification defines the datastore/repository handling persisting Notification records.
type Notification interface {
	LatestUserNotifications(ctx context.Context, d db.DBTX, params GetUserNotificationsParams) ([]db.GetUserNotificationsRow, error)
	Create(ctx context.Context, d db.DBTX, params NotificationCreateParams) error
	Delete(ctx context.Context, d db.DBTX, notificationID int32) error
}

// User defines the datastore/repository handling persisting User records.
type User interface {
	ByID(ctx context.Context, d db.DBTX, id uuid.UUID) (*db.User, error)
	ByEmail(ctx context.Context, d db.DBTX, email string) (*db.User, error)
	ByUsername(ctx context.Context, d db.DBTX, username string) (*db.User, error)
	ByExternalID(ctx context.Context, d db.DBTX, extID string) (*db.User, error)
	ByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.User, error)
	Create(ctx context.Context, d db.DBTX, params UserCreateParams) (*db.User, error)
	Update(ctx context.Context, d db.DBTX, id uuid.UUID, params UserUpdateParams) (*db.User, error)
	Delete(ctx context.Context, d db.DBTX, id uuid.UUID) (*db.User, error)
	// CreateAPIKey requires an existing user.
	CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (*db.UserAPIKey, error)
}

// Project defines the datastore/repository handling persisting Project records.
// Projects are manually created on demand.
type Project interface {
	ByName(ctx context.Context, d db.DBTX, name internalmodels.Project) (*db.Project, error)
	ByID(ctx context.Context, d db.DBTX, id int) (*db.Project, error)
}

// Team defines the datastore/repository handling persisting Team records.
type Team interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.Team, error)
	ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.Team, error)
	Create(ctx context.Context, d db.DBTX, params TeamCreateParams) (*db.Team, error)
	Update(ctx context.Context, d db.DBTX, id int, params TeamUpdateParams) (*db.Team, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.Team, error)
}

// WorkItemType defines the datastore/repository handling persisting WorkItemType records.
type WorkItemType interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemType, error)
	// TODO ByProjectID(ctx context.Context, d db.DBTX, id int) ([]*db.WorkItemType, error)
	ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.WorkItemType, error)
	Create(ctx context.Context, d db.DBTX, params WorkItemTypeCreateParams) (*db.WorkItemType, error)
	Update(ctx context.Context, d db.DBTX, id int, params WorkItemTypeUpdateParams) (*db.WorkItemType, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItemType, error)
}

// WorkItemTag defines the datastore/repository handling persisting WorkItemTag records.
type WorkItemTag interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error)
	// TODO ByProjectID(ctx context.Context, d db.DBTX, id int) ([]*db.WorkItemTag, error)
	ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.WorkItemTag, error)
	Create(ctx context.Context, d db.DBTX, params WorkItemTagCreateParams) (*db.WorkItemTag, error)
	Update(ctx context.Context, d db.DBTX, id int, params WorkItemTagUpdateParams) (*db.WorkItemTag, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error)
}

// Activity defines the datastore/repository handling persisting Activity records.
type Activity interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.Activity, error)
	// TODO ByProjectID(ctx context.Context, d db.DBTX, id int) ([]*db.Activity, error)
	ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.Activity, error)
	Create(ctx context.Context, d db.DBTX, params ActivityCreateParams) (*db.Activity, error)
	Update(ctx context.Context, d db.DBTX, id int, params ActivityUpdateParams) (*db.Activity, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.Activity, error)
}
