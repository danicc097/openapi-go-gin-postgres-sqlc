package repos

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
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
	StepOrder     int16  `json:"stepOrder"`
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

type UserCreateParams struct {
	Username   string
	Email      string
	FirstName  *string
	LastName   *string
	ExternalID string
	Scopes     []string
	RoleRank   int16
}

type UserUpdateParams struct {
	FirstName *string
	LastName  *string
	Rank      *int16
	Scopes    *[]string
}

// ProjectBoard defines the datastore/repository handling persisting ProjectBoard records.
type ProjectBoard interface {
	/*
		Create corresponds to initial info to be filled in once a project table has been manually
		 created, before it can be used:
		 - kanban columns and their info (order, name, can log time, etc.)
		 - types of workitems (shared by all teams)
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
	ProjectBoardByID(ctx context.Context, d db.DBTX, projectID int) (*models.ProjectBoard, error)
}

// User defines the datastore/repository handling persisting User records.
type User interface {
	UserByID(ctx context.Context, d db.DBTX, id string) (*db.User, error)
	UserByEmail(ctx context.Context, d db.DBTX, email string) (*db.User, error)
	UserByUsername(ctx context.Context, d db.DBTX, username string) (*db.User, error)
	UserByExternalID(ctx context.Context, d db.DBTX, extID string) (*db.User, error)
	UserByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.User, error)
	Create(ctx context.Context, d db.DBTX, params UserCreateParams) (*db.User, error)
	Update(ctx context.Context, d db.DBTX, id string, params UserUpdateParams) (*db.User, error)
	Delete(ctx context.Context, d db.DBTX, id string) (*db.User, error)
	// CreateAPIKey requires an existing user.
	CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (*db.UserAPIKey, error)
}

// Project defines the datastore/repository handling persisting Project records.
// Projects are manually created on demand.
type Project interface {
	ProjectByName(ctx context.Context, d db.DBTX, name string) (*db.Project, error)
	ProjectByID(ctx context.Context, d db.DBTX, id int) (*db.Project, error)
}

// Team defines the datastore/repository handling persisting Team records.
type Team interface {
	TeamByID(ctx context.Context, d db.DBTX, id int) (*db.Team, error)
	TeamByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.Team, error)
	Create(ctx context.Context, d db.DBTX, params TeamCreateParams) (*db.Team, error)
	Update(ctx context.Context, d db.DBTX, id int, params TeamUpdateParams) (*db.Team, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.Team, error)
}

// WorkItemType defines the datastore/repository handling persisting WorkItemType records.
type WorkItemType interface {
	WorkItemTypeByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemType, error)
	WorkItemTypeByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.WorkItemType, error)
	Create(ctx context.Context, d db.DBTX, params WorkItemTypeCreateParams) (*db.WorkItemType, error)
	Update(ctx context.Context, d db.DBTX, id int, params WorkItemTypeUpdateParams) (*db.WorkItemType, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItemType, error)
}

// WorkItemTag defines the datastore/repository handling persisting WorkItemTag records.
type WorkItemTag interface {
	WorkItemTagByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error)
	WorkItemTagByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.WorkItemTag, error)
	Create(ctx context.Context, d db.DBTX, params WorkItemTagCreateParams) (*db.WorkItemTag, error)
	Update(ctx context.Context, d db.DBTX, id int, params WorkItemTagUpdateParams) (*db.WorkItemTag, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error)
}
