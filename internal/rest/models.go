package rest

/**
 * IMPORTANT: add omitempty tag option for pointer to structs. If adding to slice of structs, include a x-omitempty:"true" tag.
 */

import (
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
	*db.User
	// Role replaces db RoleRank
	Role Role `json:"role" ref:"#/components/schemas/Role" required:"true"`

	APIKey   *db.UserAPIKey `json:"apiKey,omitempty"`
	Teams    *[]db.Team     `json:"teams"`
	Projects *[]db.Project  `json:"projects"`
}

// type Users []User // type Users []User is ignored in is_rest_type and not in structs.gen.go since its not a struct.
// TODO: maybe worth it to include ALL types in models.go by default for swaggest instead of below workaround.
// but first: can swaggest handle array gen properly?
//
//	type Users struct {
//		Users []User `json:""`
//	}
type PaginatedUsersResponse = PaginationBaseResponse[User]

type SharedWorkItemFields struct {
	TimeEntries      *[]db.TimeEntry           `json:"timeEntries"`
	WorkItemComments *[]db.WorkItemComment     `json:"workItemComments"`
	Members          *[]db.User__WIAU_WorkItem `json:"members"`
	WorkItemTags     *[]db.WorkItemTag         `json:"workItemTags"`
	WorkItemType     *db.WorkItemType          `json:"workItemType"`
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
	ProjectName Project `json:"projectName" ref:"#/components/schemas/Project" required:"true"`
}

type CreateProjectBoardRequest struct {
	// services models not needed yet, projectId is trivial to include in every request...
	// if services use db CreateParams as is we can also have specific per-project logic
	// anyway
	Teams *[]db.TeamCreateParams        `json:"teams"`
	Tags  *[]db.WorkItemTagCreateParams `json:"tags"`
}

type CreateWorkItemTagRequest struct {
	db.WorkItemTagCreateParams
}
type UpdateWorkItemTagRequest struct {
	db.WorkItemTagUpdateParams
}
type WorkItemTag struct {
	db.WorkItemTag
	// NOTE: project join useless here, entities associated to project and do not need its own endpoint
}
type CreateWorkItemTypeRequest struct {
	db.WorkItemTypeCreateParams
}
type UpdateWorkItemTypeRequest struct {
	db.WorkItemTypeUpdateParams
}
type WorkItemType struct {
	db.WorkItemType
}

type Team struct {
	db.Team
	// NOTE: project join useless here, entities associated to project and do not need its own endpoint
}

type CreateTeamRequest struct {
	db.TeamCreateParams
}

type UpdateTeamRequest struct {
	db.TeamUpdateParams
}

type Activity struct {
	db.Activity
	// NOTE: project join useless here, entities associated to project and do not need its own endpoint
}

type CreateActivityRequest struct {
	db.ActivityCreateParams
}

type UpdateActivityRequest struct {
	db.ActivityUpdateParams
}

type CreateDemoWorkItemRequest struct {
	ProjectName Project `json:"projectName" ref:"#/components/schemas/Project" required:"true"`
	services.DemoWorkItemCreateParams
}

type CreateDemoTwoWorkItemRequest struct {
	ProjectName Project `json:"projectName" ref:"#/components/schemas/Project" required:"true"`
	services.DemoTwoWorkItemCreateParams
}

type CreateWorkItemCommentRequest struct {
	db.WorkItemCommentCreateParams
}
