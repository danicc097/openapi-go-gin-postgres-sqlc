package rest

/**
 * IMPORTANT: add omitempty tag option for pointer to structs. If adding to slice of structs, include a x-omitempty:"true" tag.
 */

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
)

/**
 * Pagination
 */

type PaginationPage struct {
	NextCursor string `json:"nextCursor"`
}

type PaginationBaseResponse[T any] struct {
	Page  PaginationPage `json:"page"  required:"true"`
	Items []T            `json:"items" required:"true"`
}

type PaginatedNotificationsResponse = PaginationBaseResponse[Notification]

/**
 *
 */

type Notification struct {
	db.UserNotification
	Notification db.Notification `json:"notification" required:"true"` // notification_id clash
}

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

// DemoWorkItems represents an OpenAPI schema response for a ProjectBoard.
type DemoWorkItems struct {
	db.WorkItem
	SharedWorkItemFields
	DemoWorkItem db.DemoWorkItem `json:"demoWorkItem" required:"true"`
}

// DemoTwoWorkItems represents an OpenAPI schema response for a ProjectBoard.
type DemoTwoWorkItems struct {
	db.WorkItem
	SharedWorkItemFields
	DemoTwoWorkItem db.DemoTwoWorkItem `json:"demoTwoWorkItem" required:"true"`
}

// ProjectBoard represents an OpenAPI schema response for a ProjectBoard.
type ProjectBoard struct {
	ProjectName
}

type ProjectBoardCreateRequest struct {
	// services models not needed yet, projectId is trivial to include in every request...
	// if services use db CreateParams as is we can also have specific per-project logic
	// anyway
	Teams *[]db.TeamCreateParams        `json:"teams"`
	Tags  *[]db.WorkItemTagCreateParams `json:"tags"`
}

type WorkItemTagCreateRequest struct {
	db.WorkItemTagCreateParams
}
type WorkItemTagUpdateRequest struct {
	db.WorkItemTagUpdateParams
}
type WorkItemTag struct {
	db.WorkItemTag
	// NOTE: project join useless here, entities associated to project and do not need its own endpoint
}
type WorkItemTypeCreateRequest struct {
	db.WorkItemTypeCreateParams
}
type WorkItemTypeUpdateRequest struct {
	db.WorkItemTypeUpdateParams
}
type WorkItemType struct {
	db.WorkItemType
}

type Team struct {
	db.Team
	// NOTE: project join useless here, entities associated to project and do not need its own endpoint
}

type TeamCreateRequest struct {
	db.TeamCreateParams
}

type TeamUpdateRequest struct {
	db.TeamUpdateParams
}

type DemoWorkItemCreateRequest struct {
	ProjectName
	services.DemoWorkItemCreateParams
}

type DemoTwoWorkItemCreateRequest struct {
	ProjectName
	services.DemoTwoWorkItemCreateParams
}

type WorkItemCommentCreateRequest struct {
	db.WorkItemCommentCreateParams
}

type ProjectName struct {
	ProjectName models.Project `json:"projectName" ref:"#/components/schemas/Project" required:"true"`
}
