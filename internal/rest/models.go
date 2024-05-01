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
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
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

type PaginatedNotificationsResponse = PaginationBaseResponse[NotificationResponse]

type PaginatedDemoWorkItemsResponse = PaginationBaseResponse[CacheDemoWorkItemResponse]

/**
 *
 */

type CacheDemoWorkItemResponse struct {
	db.CacheDemoWorkItem
}

type NotificationResponse struct {
	db.UserNotification
	Notification db.Notification `json:"notification" required:"true"` // notification_id clash
}

// User represents an OpenAPI schema response for a User.
type UserResponse struct {
	*db.User
	// Role replaces db RoleRank
	Role Role `json:"role" ref:"#/components/schemas/Role" required:"true"`

	APIKey   *models.UserAPIKey `json:"apiKey,omitempty"`
	Teams    *[]models.Team     `json:"teams"`
	Projects *[]models.Project  `json:"projects"`
}

type GetCurrentUserQueryParameters struct {
	// if need arises somehow, exclude some joins from json via a new SQL comment annotation.
	// joins only accumulate with the ones previously set.
	Joins models.UserJoins `json:"joins"`
}

// type GetPaginatedUsersQueryParameters struct {
// 	// proper are joins needed when filtering with having clause, however we will not generate
// 	// code for these joins, so we will set them dynamically in repos based on filters
// 	// Joins db.UserJoins `json:"joins"`
// 	// custom filters:
// 	Role Role `json:"role" ref:"#/components/schemas/Role" required:"true"`
// 	// generic user filters (via openapi we know the types -> generate xo + postgresql accordingly)
// 	// NOTE: do not generate joins, just create filters with havingclause manually, its not trivial to generate.
// 	// Teams ArrayFilter[db.TeamID] or anyof in spec
// 	// frontend doesnt care,
// 	// based on generated m2m, o2m or o2o,
// 	// const userTableMemberProjectsSelectSQL = `COALESCE(
// 	// 	ARRAY_AGG( DISTINCT (
// 	// 	xo_join_user_project_projects.__projects
// 	// 	)) filter (where xo_join_user_project_projects.__projects_project_id is not null), '{}') as user_project_projects`
// }

type GetCacheDemoWorkItemQueryParameters struct {
	// if need arises, exclude some joins from json via a new SQL comment annotation.
	// joins only accumulate with the ones previously set.
	Joins models.CacheDemoWorkItemJoins `json:"joins"`
	// TODO: Filters. easier to generate a default Filter struct via xo
	// since we know the types (see mantine-react-table filters json gen and define struct accordingly)
	// then xo filters (where or having) get automatically built. see excalidraw
	// e.g. for users filter on teams would need extra adhoc filter teamIDs
}

// type Users []User // cannot be handled by swaggest lib (only handles structs)
// panic: reflect: NumField of non-struct type rest.Users
// should use below workaround as in paginated queries (all would be paginated queries in a way...)
//
//	type UsersResponse struct {
//		Users []User `json:"users"`
//	}
type PaginatedUsersResponse = PaginationBaseResponse[UserResponse]

// NOTE: keep in sync with base workitem getSharedDBOpts.
type SharedWorkItemJoins struct {
	TimeEntries      *[]models.TimeEntry              `json:"timeEntries"`
	WorkItemComments *[]models.WorkItemComment        `json:"workItemComments"`
	Members          *[]models.WorkItemM2MAssigneeWIA `json:"members"`
	WorkItemTags     *[]models.WorkItemTag            `json:"workItemTags"`
	WorkItemType     *models.WorkItemType             `json:"workItemType"`
}

type WorkItemBase struct {
	models.WorkItem
	SharedWorkItemJoins
	ProjectName ProjectName `json:"projectName" ref:"#/components/schemas/ProjectName" required:"true"`
}

type DemoWorkItemResponse struct {
	WorkItemBase

	DemoWorkItem models.DemoWorkItem `json:"demoWorkItem" required:"true"`
}
type DemoTwoWorkItemResponse struct {
	WorkItemBase

	DemoTwoWorkItem models.DemoTwoWorkItem `json:"demoTwoWorkItem" required:"true"`
}

type ProjectBoardResponse struct {
	ProjectName Project `json:"projectName" ref:"#/components/schemas/Project" required:"true"`
}

type CreateProjectBoardRequest struct {
	// services models not needed yet, projectId is trivial to include in every request...
	// if services use db CreateParams as is we can also have specific per-project logic
	// anyway
	Teams *[]models.TeamCreateParams        `json:"teams"`
	Tags  *[]models.WorkItemTagCreateParams `json:"tags"`
}

type CreateWorkItemTagRequest struct {
	models.WorkItemTagCreateParams
}
type UpdateWorkItemTagRequest struct {
	models.WorkItemTagUpdateParams
}
type WorkItemTagResponse struct {
	db.WorkItemTag
	// NOTE: project join useless here, entities associated to project and do not need its own endpoint
}
type CreateWorkItemTypeRequest struct {
	models.WorkItemTypeCreateParams
}
type UpdateWorkItemTypeRequest struct {
	models.WorkItemTypeUpdateParams
}
type WorkItemTypeResponse struct {
	db.WorkItemType
}

type TeamResponse struct {
	db.Team
	// NOTE: project join useless here, entities associated to project and do not need its own endpoint
}

type CreateTeamRequest struct {
	models.TeamCreateParams
}

type UpdateTeamRequest struct {
	models.TeamUpdateParams
}

type ActivityResponse struct {
	db.Activity
	// NOTE: project join useless here, entities associated to project and do not need its own endpoint
}

type CreateActivityRequest struct {
	models.ActivityCreateParams
}

type UpdateActivityRequest struct {
	models.ActivityUpdateParams
}

type TimeEntryResponse struct {
	db.TimeEntry
}

type CreateTimeEntryRequest struct {
	models.TimeEntryCreateParams
}

type UpdateTimeEntryRequest struct {
	models.TimeEntryUpdateParams
}

type CreateDemoWorkItemRequest struct {
	ProjectName ProjectName `json:"projectName" ref:"#/components/schemas/ProjectName" required:"true"`
	services.DemoWorkItemCreateParams
}

type CreateDemoTwoWorkItemRequest struct {
	ProjectName ProjectName `json:"projectName" ref:"#/components/schemas/ProjectName" required:"true"`
	services.DemoTwoWorkItemCreateParams
}

type WorkItemCommentResponse struct {
	db.WorkItemComment
}

type CreateWorkItemCommentRequest struct {
	models.WorkItemCommentCreateParams
}

type UpdateWorkItemCommentRequest struct {
	models.WorkItemCommentUpdateParams
}
