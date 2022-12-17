package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	repomodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// UserResponse represents an OpenAPI schema response for a User.
type UserResponse struct {
	Role     models.Role          `json:"role" ref:"#/components/schemas/Role" required:"true"`
	Scopes   []string             `json:"scopes" ref:"#/components/schemas/Scopes" required:"true"`
	APIKey   *db.UserAPIKeyPublic `json:"apiKey"`
	Teams    *[]db.TeamPublic     `json:"teams"`
	Projects *[]db.ProjectPublic  `json:"projects"`

	db.UserPublic
}

// DemoProjectWorkItemsResponse represents an OpenAPI schema response for a ProjectBoard.
type DemoProjectWorkItemsResponse struct {
	db.WorkItemPublic
	DemoProjectWorkItem db.DemoProjectWorkItemPublic `json:"demoProjectWorkItem" required:"true"`
	TimeEntries         *[]db.TimeEntryPublic        `json:"timeEntries"`
	WorkItemComments    *[]db.WorkItemCommentPublic  `json:"workItemComments"`
	Members             *[]db.UserPublic             `json:"members"`
	WorkItemTags        *[]db.WorkItemTagPublic      `json:"workItemTags"`
	WorkItemType        *db.WorkItemTypePublic       `json:"workItemType"`
}

// ProjectBoardResponse represents an OpenAPI schema response for a ProjectBoard.
type ProjectBoardResponse struct {
	repomodels.ProjectBoardPublic
}

// ProjectBoardConfigResponse represents an OpenAPI schema response for board configuration.
type ProjectBoardConfigResponse struct {
	CardConfig CardConfig `json:"cardConfig"`
}

// CardConfigUpdateRequest represents an OpenAPI schema request for board card configuration.
type CardConfigUpdateRequest struct {
	CardConfig
}

type CardConfig struct {
	Body         map[string]any `json:"body"`
	HeaderFields []string       `json:"headerFields"`
}

type ProjectBoardCreateRequest struct {
	repos.ProjectBoardCreateParams
}

// WorkItemResponse represents an OpenAPI schema response for a WorkItem.
type WorkItemResponse struct {
	db.WorkItemPublic
}

type TeamCreateRequest struct {
	repos.TeamCreateParams
}

type TeamUpdateRequest struct {
	repos.TeamUpdateParams
}
