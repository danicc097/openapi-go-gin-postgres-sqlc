// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/danicc097/openapi-go-gin-postgres-sqlc version (devel) DO NOT EDIT.
package models

import (
	"time"
)

const (
	Api_keyScopes     = "api_key.Scopes"
	Bearer_authScopes = "bearer_auth.Scopes"
)

// Defines values for Role.
const (
	RoleAdmin        Role = "admin"
	RoleAdvancedUser Role = "advancedUser"
	RoleGuest        Role = "guest"
	RoleManager      Role = "manager"
	RoleSuperAdmin   Role = "superAdmin"
	RoleUser         Role = "user"
)

// Defines values for Scope.
const (
	ScopeProjectSettingsWrite Scope = "project-settings:write"
	ScopeScopesWrite          Scope = "scopes:write"
	ScopeTeamSettingsWrite    Scope = "team-settings:write"
	ScopeTestScope            Scope = "test-scope"
	ScopeUsersRead            Scope = "users:read"
	ScopeUsersWrite           Scope = "users:write"
	ScopeWorkItemReview       Scope = "work-item:review"
)

// Defines values for WorkItemRole.
const (
	WorkItemRolePreparer WorkItemRole = "preparer"
	WorkItemRoleReviewer WorkItemRole = "reviewer"
)

// HTTPValidationError defines model for HTTPValidationError.
type HTTPValidationError struct {
	Detail *[]ValidationError `json:"detail,omitempty"`
}

// Organization Organization a user belongs to.
type Organization = string

// PgtypeJSONB defines model for PgtypeJSONB.
type PgtypeJSONB = map[string]interface{}

// Role Role automatically generated from roles.json keys
type Role string

// Scope Scope automatically generated from scopes.json keys
type Scope string

// Task defines model for Task.
type Task struct {
	CreatedAt   *time.Time   `json:"created_at,omitempty"`
	DeletedAt   *time.Time   `json:"deleted_at"`
	Finished    *bool        `json:"finished"`
	Metadata    *PgtypeJSONB `json:"metadata,omitempty"`
	TaskId      *int         `json:"task_id,omitempty"`
	TaskType    *TaskType    `json:"task_type"`
	TaskTypeId  *int         `json:"task_type_id,omitempty"`
	TimeEntries *[]TimeEntry `json:"time_entries"`
	Title       *string      `json:"title,omitempty"`
	UpdatedAt   *time.Time   `json:"updated_at,omitempty"`
	WorkItemId  *int         `json:"work_item_id,omitempty"`
}

// TaskType defines model for TaskType.
type TaskType struct {
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	TaskTypeId  *int    `json:"task_type_id,omitempty"`
	TeamId      *int    `json:"team_id,omitempty"`
}

// Team defines model for Team.
type Team struct {
	CreatedAt   *time.Time   `json:"created_at,omitempty"`
	Description *string      `json:"description,omitempty"`
	Metadata    *PgtypeJSONB `json:"metadata,omitempty"`
	Name        *string      `json:"name,omitempty"`
	ProjectId   *int         `json:"project_id,omitempty"`
	TeamId      *int         `json:"team_id,omitempty"`
	TimeEntries *[]TimeEntry `json:"time_entries"`
	UpdatedAt   *time.Time   `json:"updated_at,omitempty"`
	Users       *[]User      `json:"users"`
}

// TimeEntry defines model for TimeEntry.
type TimeEntry struct {
	ActivityId      *int       `json:"activity_id,omitempty"`
	Comment         *string    `json:"comment,omitempty"`
	DurationMinutes *int       `json:"duration_minutes"`
	Start           *time.Time `json:"start,omitempty"`
	TaskId          *int       `json:"task_id"`
	TeamId          *int       `json:"team_id"`
	TimeEntryId     *int       `json:"time_entry_id,omitempty"`
	UserId          *UuidUUID  `json:"user_id,omitempty"`
}

// UpdateUserRequest represents User data to update
type UpdateUserRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`

	// Role Role automatically generated from roles.json keys
	Role *Role `json:"role,omitempty"`
}

// User defines model for User.
type User struct {
	ApiKeyId    *int         `json:"api_key_id"`
	CreatedAt   *time.Time   `json:"created_at,omitempty"`
	DeletedAt   *time.Time   `json:"deleted_at"`
	Email       *string      `json:"email,omitempty"`
	ExternalId  *string      `json:"external_id,omitempty"`
	FirstName   *string      `json:"first_name"`
	FullName    *string      `json:"full_name"`
	LastName    *string      `json:"last_name"`
	RoleRank    *int         `json:"role_rank,omitempty"`
	Scopes      *[]string    `json:"scopes"`
	Teams       *[]Team      `json:"teams"`
	TimeEntries *[]TimeEntry `json:"time_entries"`
	UpdatedAt   *time.Time   `json:"updated_at,omitempty"`
	UserId      *UuidUUID    `json:"user_id,omitempty"`
	Username    *string      `json:"username,omitempty"`
	WorkItems   *[]WorkItem  `json:"work_items"`
}

// UserAPIKey defines model for UserAPIKey.
type UserAPIKey struct {
	ApiKey       *string    `json:"api_key,omitempty"`
	ExpiresOn    *time.Time `json:"expires_on,omitempty"`
	UserApiKeyId *int       `json:"user_api_key_id,omitempty"`
	UserId       *UuidUUID  `json:"user_id,omitempty"`
}

// UuidUUID defines model for UuidUUID.
type UuidUUID = string

// ValidationError defines model for ValidationError.
type ValidationError struct {
	Loc  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
}

// WorkItem defines model for WorkItem.
type WorkItem struct {
	Closed           *bool              `json:"closed,omitempty"`
	CreatedAt        *time.Time         `json:"created_at,omitempty"`
	DeletedAt        *time.Time         `json:"deleted_at"`
	KanbanStepId     *int               `json:"kanban_step_id,omitempty"`
	Metadata         *PgtypeJSONB       `json:"metadata,omitempty"`
	Tasks            *[]Task            `json:"tasks"`
	TeamId           *int               `json:"team_id,omitempty"`
	Title            *string            `json:"title,omitempty"`
	UpdatedAt        *time.Time         `json:"updated_at,omitempty"`
	Users            *[]User            `json:"users"`
	WorkItemComments *[]WorkItemComment `json:"work_item_comments"`
	WorkItemId       *int               `json:"work_item_id,omitempty"`
	WorkItemTypeId   *int               `json:"work_item_type_id,omitempty"`
}

// WorkItemComment defines model for WorkItemComment.
type WorkItemComment struct {
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	Message           *string    `json:"message,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	UserId            *UuidUUID  `json:"user_id,omitempty"`
	WorkItemCommentId *int       `json:"work_item_comment_id,omitempty"`
	WorkItemId        *int       `json:"work_item_id,omitempty"`
}

// WorkItemRole Role in work item for a member.
type WorkItemRole string

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = UpdateUserRequest
