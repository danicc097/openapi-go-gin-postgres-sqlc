/**
 * IMPORTANT: all structs in models.go will always be generated as OpenAPI schemas.
 * Generated "Rest*" schemas must be ignored.
 * TODO: struct type grouping not supported in gen.
 */
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

// type Users []User // cannot be handled by swaggest lib (only handles structs)
// panic: reflect: NumField of non-struct type rest.Users
// should use below workaround as in paginated queries (all would be paginated queries in a way...)
//
//	type Users struct {
//		Users []User `json:"users"`
//	}
type PaginatedUsersResponse = PaginationBaseResponse[User]

type SharedWorkItemJoins struct {
	TimeEntries      *[]db.TimeEntry           `json:"timeEntries"`
	WorkItemComments *[]db.WorkItemComment     `json:"workItemComments"`
	Members          *[]db.User__WIAU_WorkItem `json:"members"`
	WorkItemTags     *[]db.WorkItemTag         `json:"workItemTags"`
	WorkItemType     *db.WorkItemType          `json:"workItemType"`
	KanbanStep       *db.KanbanStep            `json:"kanbanStep"`
	Team             *db.Team                  `json:"team"`
}

// 	// TODO: losing all joins.
// type DemoWorkItems struct {
// 	SharedWorkItemJoins
// 	// e.g. we have comment on column work_item_work_item_tag.work_item_id is '"cardinality":M2M';
// 	// therefore this join (internally handled as constraint) should be propagated if
// 	// we set a flag propagate-fk-constraints  (don't want to always propagate).
// 	// if so we duplicate constraints that reference work_item_id original table so that its generated
// 	// in cache
// 	WorkItem db.CacheDemoWorkItem `json:"workItem" required:"true"`
// }

type DemoWorkItems struct {
	db.WorkItem
	SharedWorkItemJoins
	DemoWorkItem db.DemoWorkItem `json:"demoWorkItem" required:"true"`
}

type DemoTwoWorkItems struct {
	db.WorkItem
	SharedWorkItemJoins
	DemoTwoWorkItem db.DemoTwoWorkItem `json:"demoTwoWorkItem" required:"true"`
}

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

type WorkItemComment struct {
	db.WorkItemComment
}

type CreateWorkItemCommentRequest struct {
	db.WorkItemCommentCreateParams
}

type UpdateWorkItemCommentRequest struct {
	db.WorkItemCommentUpdateParams
}
