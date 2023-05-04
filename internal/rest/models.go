package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	repomodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
)

// UserResponse represents an OpenAPI schema response for a User.
type UserResponse struct {
	db.User
	Role models.Role `json:"role" ref:"#/components/schemas/Role" required:"true"`

	APIKey   *db.UserAPIKey `json:"apiKey,omitempty"`
	Teams    *[]db.Team     `json:"teams,omitempty"`
	Projects *[]db.Project  `json:"projects,omitempty"`
}

// DemoWorkItemsResponse represents an OpenAPI schema response for a ProjectBoard.
type DemoWorkItemsResponse struct {
	db.WorkItem
	DemoWorkItem     db.DemoWorkItem       `json:"demoWorkItem" required:"true"`
	TimeEntries      *[]db.TimeEntry       `json:"timeEntries"`
	WorkItemComments *[]db.WorkItemComment `json:"workItemComments"`
	Members          *[]db.User            `json:"members"`
	WorkItemTags     *[]db.WorkItemTag     `json:"workItemTags"`
	WorkItemType     *db.WorkItemType      `json:"workItemType"`
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
	db.TeamCreateParams
}

type TeamUpdateRequest struct {
	db.TeamUpdateParams
}

type UserCreateRequest struct {
	services.UserRegisterParams
}

type DemoWorkItemCreateRequest struct {
	services.DemoWorkItemCreateParams
}
