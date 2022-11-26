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

// ModelsRole defines model for ModelsRole.
type ModelsRole = string

// ModelsScope defines model for ModelsScope.
type ModelsScope = string

// PgtypeJSONB defines model for PgtypeJSONB.
type PgtypeJSONB = map[string]interface{}

// Role defines model for Role.
type Role string

// Scope defines model for Scope.
type Scope string

// Scopes defines model for Scopes.
type Scopes = []Scope

// TaskPublic defines model for TaskPublic.
type TaskPublic struct {
	CreatedAt  *time.Time      `json:"createdAt,omitempty"`
	DeletedAt  *time.Time      `json:"deletedAt"`
	Finished   *bool           `json:"finished"`
	Metadata   *PgtypeJSONB    `json:"metadata,omitempty"`
	TaskID     *int            `json:"taskID,omitempty"`
	TaskType   *TaskTypePublic `json:"taskType"`
	TaskTypeID *int            `json:"taskTypeID,omitempty"`
	Title      *string         `json:"title,omitempty"`
	UpdatedAt  *time.Time      `json:"updatedAt,omitempty"`
	WorkItemID *int            `json:"workItemID,omitempty"`
}

// TaskTypePublic defines model for TaskTypePublic.
type TaskTypePublic struct {
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	TaskTypeID  *int    `json:"taskTypeID,omitempty"`
	TeamID      *int    `json:"teamID,omitempty"`
}

// TeamPublic defines model for TeamPublic.
type TeamPublic struct {
	CreatedAt   time.Time   `json:"createdAt"`
	Description string      `json:"description"`
	Metadata    PgtypeJSONB `json:"metadata"`
	Name        string      `json:"name"`
	ProjectID   int         `json:"projectID"`
	TeamID      int         `json:"teamID"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

// TimeEntryPublic defines model for TimeEntryPublic.
type TimeEntryPublic struct {
	ActivityID      *int       `json:"activityID,omitempty"`
	Comment         *string    `json:"comment,omitempty"`
	DurationMinutes *int       `json:"durationMinutes"`
	Start           *time.Time `json:"start,omitempty"`
	TeamID          *int       `json:"teamID"`
	TimeEntryID     *int       `json:"timeEntryID,omitempty"`
	UserID          *UuidUUID  `json:"userID,omitempty"`
	WorkItemID      *int       `json:"workItemID"`
}

// UpdateUserAuthRequest represents User authorization data to update
type UpdateUserAuthRequest struct {
	Role   *Role   `json:"role,omitempty"`
	Scopes *Scopes `json:"scopes,omitempty"`
}

// UpdateUserRequest represents User data to update
type UpdateUserRequest struct {
	// FirstName originally from auth server but updatable
	FirstName *string `json:"first_name,omitempty"`

	// LastName originally from auth server but updatable
	LastName *string `json:"last_name,omitempty"`
}

// UserAPIKeyPublic defines model for UserAPIKeyPublic.
type UserAPIKeyPublic struct {
	ApiKey    string    `json:"apiKey"`
	ExpiresOn time.Time `json:"expiresOn"`
	UserID    UuidUUID  `json:"userID"`
}

// UserPublic defines model for UserPublic.
type UserPublic struct {
	ApiKeyID    *int               `json:"apiKeyID"`
	CreatedAt   *time.Time         `json:"createdAt,omitempty"`
	DeletedAt   *time.Time         `json:"deletedAt"`
	Email       *string            `json:"email,omitempty"`
	FirstName   *string            `json:"firstName"`
	FullName    *string            `json:"fullName"`
	LastName    *string            `json:"lastName"`
	Teams       *[]TeamPublic      `json:"teams"`
	TimeEntries *[]TimeEntryPublic `json:"timeEntries"`
	UserID      *UuidUUID          `json:"userID,omitempty"`
	Username    *string            `json:"username,omitempty"`
	WorkItems   *[]WorkItemPublic  `json:"workItems"`
}

// UserResponse defines model for UserResponse.
type UserResponse struct {
	ApiKey    *UserAPIKeyPublic `json:"apiKey"`
	CreatedAt time.Time         `json:"createdAt"`
	DeletedAt *time.Time        `json:"deletedAt"`
	Email     string            `json:"email"`
	FirstName *string           `json:"firstName"`
	FullName  *string           `json:"fullName"`
	LastName  *string           `json:"lastName"`
	Role      *Role             `json:"role,omitempty"`
	Scopes    *Scopes           `json:"scopes,omitempty"`
	Teams     *[]TeamPublic     `json:"teams"`
	UserID    UuidUUID          `json:"userID"`
	Username  string            `json:"username"`
}

// UuidUUID defines model for UuidUUID.
type UuidUUID = string

// ValidationError defines model for ValidationError.
type ValidationError struct {
	Loc  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
}

// WorkItemCommentPublic defines model for WorkItemCommentPublic.
type WorkItemCommentPublic struct {
	CreatedAt         *time.Time `json:"createdAt,omitempty"`
	Message           *string    `json:"message,omitempty"`
	UpdatedAt         *time.Time `json:"updatedAt,omitempty"`
	UserID            *UuidUUID  `json:"userID,omitempty"`
	WorkItemCommentID *int       `json:"workItemCommentID,omitempty"`
	WorkItemID        *int       `json:"workItemID,omitempty"`
}

// WorkItemPublic defines model for WorkItemPublic.
type WorkItemPublic struct {
	Closed           *bool                    `json:"closed,omitempty"`
	CreatedAt        *time.Time               `json:"createdAt,omitempty"`
	DeletedAt        *time.Time               `json:"deletedAt"`
	KanbanStepID     *int                     `json:"kanbanStepID,omitempty"`
	Metadata         *PgtypeJSONB             `json:"metadata,omitempty"`
	Tasks            *[]TaskPublic            `json:"tasks"`
	TeamID           *int                     `json:"teamID,omitempty"`
	TimeEntries      *[]TimeEntryPublic       `json:"timeEntries"`
	Title            *string                  `json:"title,omitempty"`
	UpdatedAt        *time.Time               `json:"updatedAt,omitempty"`
	Users            *[]UserPublic            `json:"users"`
	WorkItemComments *[]WorkItemCommentPublic `json:"workItemComments"`
	WorkItemID       *int                     `json:"workItemID,omitempty"`
	WorkItemTypeID   *int                     `json:"workItemTypeID,omitempty"`
}

// WorkItemRole Role in work item for a member.
type WorkItemRole string

// UserID defines model for UserID.
type UserID = string

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = UpdateUserRequest

// UpdateUserAuthorizationJSONRequestBody defines body for UpdateUserAuthorization for application/json ContentType.
type UpdateUserAuthorizationJSONRequestBody = UpdateUserAuthRequest
