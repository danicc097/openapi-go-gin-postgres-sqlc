package rest

/**
 * IMPORTANT: add omitempty tag option for pointer to structs. If adding to slice of structs, include a x-omitempty:"true" tag.
 */

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	repomodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
)

// User represents an OpenAPI schema response for a User.
type User struct {
	db.User
	Role models.Role `json:"role" ref:"#/components/schemas/Role" required:"true"`

	APIKey   *db.UserAPIKey `json:"apiKey,omitempty"`
	Teams    *[]db.Team     `json:"teams"`
	Projects *[]db.Project  `json:"projects"`
}

type SharedWorkItemFields struct {
	TimeEntries      *[]db.TimeEntry       `json:"timeEntries"`
	WorkItemComments *[]db.WorkItemComment `json:"workItemComments"`
	Members          *[]db.User            `json:"members"`
	WorkItemTags     *[]db.WorkItemTag     `json:"workItemTags"`
	WorkItemType     *db.WorkItemType      `json:"workItemType"`
}

// DemoWorkItemsResponse represents an OpenAPI schema response for a ProjectBoard.
type DemoWorkItemsResponse struct {
	db.WorkItem
	SharedWorkItemFields
	DemoWorkItem db.DemoWorkItem `json:"demoWorkItem" required:"true"`
}

// DemoTwoWorkItemsResponse represents an OpenAPI schema response for a ProjectBoard.
type DemoTwoWorkItemsResponse struct {
	db.WorkItem
	SharedWorkItemFields
	DemoTwoWorkItem db.DemoTwoWorkItem `json:"demoTwoWorkItem" required:"true"`
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

type DemoWorkItemCreateRequest struct {
	services.DemoWorkItemCreateParams
}

type WorkItemTagCreateRequest struct {
	db.WorkItemTagCreateParams
}

type WorkItemCommentCreateRequest struct {
	db.WorkItemCommentCreateParams
}
