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

type PaginatedDemoWorkItemsResponse = PaginationBaseResponse[CacheDemoWorkItem]

/**
 *
 */

type CacheDemoWorkItem struct {
	db.CacheDemoWorkItem
}

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

type GetCurrentUserQueryParameters struct {
	// if need arises somehow, exclude some joins from json via a new SQL comment annotation.
	// joins only accumulate with the ones previously set.
	Joins db.UserJoins `json:"joins"`
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
	Joins db.CacheDemoWorkItemJoins `json:"joins"`
	// TODO: Filters. easier to generate a default Filter struct via xo
	// since we know the types (see mantine-react-table filters json gen and define struct accordingly)
	// then xo filters (where or having) get automatically built. see excalidraw
	// e.g. for users filter on teams would need extra adhoc filter teamIDs
}

// type Users []User // cannot be handled by swaggest lib (only handles structs)
// panic: reflect: NumField of non-struct type rest.Users
// should use below workaround as in paginated queries (all would be paginated queries in a way...)
//
//	type Users struct {
//		Users []User `json:"users"`
//	}
type PaginatedUsersResponse = PaginationBaseResponse[User]

// NOTE: keep in sync with base workitem getSharedDBOpts.
type SharedWorkItemJoins struct {
	TimeEntries      *[]db.TimeEntry              `json:"timeEntries"`
	WorkItemComments *[]db.WorkItemComment        `json:"workItemComments"`
	Members          *[]db.WorkItemM2MAssigneeWIA `json:"members"`
	WorkItemTags     *[]db.WorkItemTag            `json:"workItemTags"`
	WorkItemType     *db.WorkItemType             `json:"workItemType"`
}

type WorkItemBase struct {
	db.WorkItem
	SharedWorkItemJoins
	ProjectName Project `json:"projectName" ref:"#/components/schemas/Project" required:"true"`
}

type DemoWorkItem struct {
	WorkItemBase

	DemoWorkItem db.DemoWorkItem `json:"demoWorkItem" required:"true"`
}
type DemoTwoWorkItem struct {
	WorkItemBase

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

type TimeEntry struct {
	db.TimeEntry
}

type CreateTimeEntryRequest struct {
	db.TimeEntryCreateParams
}

type UpdateTimeEntryRequest struct {
	db.TimeEntryUpdateParams
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
