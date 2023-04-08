package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	repomodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// UserResponse represents an OpenAPI schema response for a User.
type UserResponse struct {
	Role     models.Role    `json:"role" ref:"#/components/schemas/Role" required:"true"`
	Scopes   []string       `json:"scopes" ref:"#/components/schemas/Scopes" required:"true"`
	APIKey   *db.UserAPIKey `json:"apiKey"`
	Teams    *[]db.Team     `json:"teams"`
	Projects *[]db.Project  `json:"projects"`

	*db.User
}

// DemoProjectWorkItemsResponse represents an OpenAPI schema response for a ProjectBoard.
type DemoProjectWorkItemsResponse struct {
	db.WorkItem         `required:"true"`
	DemoProjectWorkItem db.DemoProjectWorkItem `json:"demoProjectWorkItem" required:"true"`
	TimeEntries         *[]db.TimeEntry        `json:"timeEntries"`
	WorkItemComments    *[]db.WorkItemComment  `json:"workItemComments"`
	Members             *[]db.User             `json:"members"`
	WorkItemTags        *[]db.WorkItemTag      `json:"workItemTags"`
	WorkItemType        *db.WorkItemType       `json:"workItemType"`
}

// ProjectBoardResponse represents an OpenAPI schema response for a ProjectBoard.
type ProjectBoardResponse struct {
	repomodels.ProjectBoard
}

type ProjectBoardCreateRequest struct {
	repos.ProjectBoardCreateParams
}

// WorkItemResponse represents an OpenAPI schema response for a WorkItem.
type WorkItemResponse struct {
	db.WorkItem
}

type TeamCreateRequest struct {
	repos.TeamCreateParams
}

type TeamUpdateRequest struct {
	repos.TeamUpdateParams
}
