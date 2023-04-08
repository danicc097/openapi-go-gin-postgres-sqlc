// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package models

import (
	"time"
)

const (
	Api_keyScopes     = "api_key.Scopes"
	Bearer_authScopes = "bearer_auth.Scopes"
)

// Defines values for NotificationType.
const (
	NotificationTypeGlobal   NotificationType = "global"
	NotificationTypePersonal NotificationType = "personal"
)

// AllNotificationTypeValues returns all possible values for NotificationType.
func AllNotificationTypeValues() []NotificationType {
	return []NotificationType{
		NotificationTypeGlobal,
		NotificationTypePersonal,
	}
}

// Defines values for Project.
const (
	ProjectDemoProject  Project = "demoProject"
	ProjectDemoProject2 Project = "demoProject2"
)

// AllProjectValues returns all possible values for Project.
func AllProjectValues() []Project {
	return []Project{
		ProjectDemoProject,
		ProjectDemoProject2,
	}
}

// Defines values for Role.
const (
	RoleAdmin        Role = "admin"
	RoleAdvancedUser Role = "advancedUser"
	RoleGuest        Role = "guest"
	RoleManager      Role = "manager"
	RoleSuperAdmin   Role = "superAdmin"
	RoleUser         Role = "user"
)

// AllRoleValues returns all possible values for Role.
func AllRoleValues() []Role {
	return []Role{
		RoleAdmin,
		RoleAdvancedUser,
		RoleGuest,
		RoleManager,
		RoleSuperAdmin,
		RoleUser,
	}
}

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

// AllScopeValues returns all possible values for Scope.
func AllScopeValues() []Scope {
	return []Scope{
		ScopeProjectSettingsWrite,
		ScopeScopesWrite,
		ScopeTeamSettingsWrite,
		ScopeTestScope,
		ScopeUsersRead,
		ScopeUsersWrite,
		ScopeWorkItemReview,
	}
}

// Defines values for Topics.
const (
	TopicsGlobalAlerts Topics = "GlobalAlerts"
)

// AllTopicsValues returns all possible values for Topics.
func AllTopicsValues() []Topics {
	return []Topics{
		TopicsGlobalAlerts,
	}
}

// Defines values for WorkItemRole.
const (
	WorkItemRolePreparer WorkItemRole = "preparer"
	WorkItemRoleReviewer WorkItemRole = "reviewer"
)

// AllWorkItemRoleValues returns all possible values for WorkItemRole.
func AllWorkItemRoleValues() []WorkItemRole {
	return []WorkItemRole{
		WorkItemRolePreparer,
		WorkItemRoleReviewer,
	}
}

// Defines values for DemoProjectKanbanSteps.
const (
	DemoProjectKanbanStepsDisabled       DemoProjectKanbanSteps = "Disabled"
	DemoProjectKanbanStepsReceived       DemoProjectKanbanSteps = "Received"
	DemoProjectKanbanStepsUnderReview    DemoProjectKanbanSteps = "Under review"
	DemoProjectKanbanStepsWorkInProgress DemoProjectKanbanSteps = "Work in progress"
)

// AllDemoProjectKanbanStepsValues returns all possible values for DemoProjectKanbanSteps.
func AllDemoProjectKanbanStepsValues() []DemoProjectKanbanSteps {
	return []DemoProjectKanbanSteps{
		DemoProjectKanbanStepsDisabled,
		DemoProjectKanbanStepsReceived,
		DemoProjectKanbanStepsUnderReview,
		DemoProjectKanbanStepsWorkInProgress,
	}
}

// DbActivity defines the model for DbActivity.
type DbActivity struct {
	ActivityID   int            `json:"activityID"`
	Description  string         `json:"description"`
	IsProductive bool           `json:"isProductive"`
	Name         string         `json:"name"`
	ProjectID    int            `json:"projectID"`
	TimeEntries  *[]DbTimeEntry `json:"timeEntries"`
}

// DbDemoProjectWorkItem defines the model for DbDemoProjectWorkItem.
type DbDemoProjectWorkItem struct {
	LastMessageAt time.Time   `json:"lastMessageAt"`
	Line          string      `json:"line"`
	Ref           string      `json:"ref"`
	Reopened      bool        `json:"reopened"`
	WorkItem      *DbWorkItem `json:"workItem,omitempty"`
	WorkItemID    int         `json:"workItemID"`
}

// DbKanbanStep defines the model for DbKanbanStep.
type DbKanbanStep struct {
	Color         string `json:"color"`
	Description   string `json:"description"`
	KanbanStepID  int    `json:"kanbanStepID"`
	Name          string `json:"name"`
	ProjectID     int    `json:"projectID"`
	StepOrder     *int   `json:"stepOrder"`
	TimeTrackable bool   `json:"timeTrackable"`
}

// DbProject defines the model for DbProject.
type DbProject struct {
	Activities    *[]DbActivity     `json:"activities"`
	CreatedAt     time.Time         `json:"createdAt"`
	Description   string            `json:"description"`
	Initialized   bool              `json:"initialized"`
	KanbanSteps   *[]DbKanbanStep   `json:"kanbanSteps"`
	Name          string            `json:"name"`
	ProjectID     int               `json:"projectID"`
	Teams         *[]DbTeam         `json:"teams"`
	UpdatedAt     time.Time         `json:"updatedAt"`
	WorkItemTags  *[]DbWorkItemTag  `json:"workItemTags"`
	WorkItemTypes *[]DbWorkItemType `json:"workItemTypes"`
}

// DbProject2WorkItem defines the model for DbProject2WorkItem.
type DbProject2WorkItem struct {
	CustomDateForProject2 *time.Time  `json:"customDateForProject2"`
	WorkItem              *DbWorkItem `json:"workItem,omitempty"`
	WorkItemID            int         `json:"workItemID"`
}

// DbTeam defines the model for DbTeam.
type DbTeam struct {
	CreatedAt   time.Time      `json:"createdAt"`
	Description string         `json:"description"`
	Name        string         `json:"name"`
	ProjectID   int            `json:"projectID"`
	TeamID      int            `json:"teamID"`
	TimeEntries *[]DbTimeEntry `json:"timeEntries"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	Users       *[]DbUser      `json:"users"`
}

// DbTimeEntry defines the model for DbTimeEntry.
type DbTimeEntry struct {
	ActivityID      int       `json:"activityID"`
	Comment         string    `json:"comment"`
	DurationMinutes *int      `json:"durationMinutes"`
	Start           time.Time `json:"start"`
	TeamID          *int      `json:"teamID"`
	TimeEntryID     int       `json:"timeEntryID"`
	UserID          UuidUUID  `json:"userID"`
	WorkItemID      *int      `json:"workItemID"`
}

// DbUser defines the model for DbUser.
type DbUser struct {
	CreatedAt                time.Time      `json:"createdAt"`
	DeletedAt                *time.Time     `json:"deletedAt"`
	Email                    string         `json:"email"`
	FirstName                *string        `json:"firstName"`
	FullName                 *string        `json:"fullName"`
	HasGlobalNotifications   bool           `json:"hasGlobalNotifications"`
	HasPersonalNotifications bool           `json:"hasPersonalNotifications"`
	LastName                 *string        `json:"lastName"`
	Teams                    *[]DbTeam      `json:"teams"`
	TimeEntries              *[]DbTimeEntry `json:"timeEntries"`
	UserAPIKey               *DbUserAPIKey  `json:"userAPIKey"`
	UserID                   UuidUUID       `json:"userID"`
	Username                 string         `json:"username"`
	WorkItems                *[]DbWorkItem  `json:"workItems"`
}

// DbUserAPIKey defines the model for DbUserAPIKey.
type DbUserAPIKey struct {
	ApiKey    string    `json:"apiKey"`
	ExpiresOn time.Time `json:"expiresOn"`
	User      *DbUser   `json:"user,omitempty"`
	UserID    UuidUUID  `json:"userID"`
}

// DbWorkItem defines the model for DbWorkItem.
type DbWorkItem struct {
	Closed              *time.Time             `json:"closed"`
	CreatedAt           time.Time              `json:"createdAt"`
	DeletedAt           *time.Time             `json:"deletedAt"`
	DemoProjectWorkItem *DbDemoProjectWorkItem `json:"demoProjectWorkItem"`
	Description         string                 `json:"description"`
	KanbanStepID        int                    `json:"kanbanStepID"`
	Members             *[]DbUser              `json:"members"`
	Metadata            PgtypeJSONB            `json:"metadata"`
	Project2workItem    *DbProject2WorkItem    `json:"project2workItem"`
	TargetDate          time.Time              `json:"targetDate"`
	TeamID              int                    `json:"teamID"`
	TimeEntries         *[]DbTimeEntry         `json:"timeEntries"`
	Title               string                 `json:"title"`
	UpdatedAt           time.Time              `json:"updatedAt"`
	WorkItemComments    *[]DbWorkItemComment   `json:"workItemComments"`
	WorkItemID          int                    `json:"workItemID"`
	WorkItemTags        *[]DbWorkItemTag       `json:"workItemTags"`
	WorkItemType        *DbWorkItemType        `json:"workItemType"`
	WorkItemTypeID      int                    `json:"workItemTypeID"`
}

// DbWorkItemComment defines the model for DbWorkItemComment.
type DbWorkItemComment struct {
	CreatedAt         time.Time `json:"createdAt"`
	Message           string    `json:"message"`
	UpdatedAt         time.Time `json:"updatedAt"`
	UserID            UuidUUID  `json:"userID"`
	WorkItemCommentID int       `json:"workItemCommentID"`
	WorkItemID        int       `json:"workItemID"`
}

// DbWorkItemTag defines the model for DbWorkItemTag.
type DbWorkItemTag struct {
	Color         string        `json:"color"`
	Description   string        `json:"description"`
	Name          string        `json:"name"`
	ProjectID     int           `json:"projectID"`
	WorkItemTagID int           `json:"workItemTagID"`
	WorkItems     *[]DbWorkItem `json:"workItems"`
}

// DbWorkItemType defines the model for DbWorkItemType.
type DbWorkItemType struct {
	Color          string      `json:"color"`
	Description    string      `json:"description"`
	Name           string      `json:"name"`
	ProjectID      int         `json:"projectID"`
	WorkItem       *DbWorkItem `json:"workItem,omitempty"`
	WorkItemTypeID int         `json:"workItemTypeID"`
}

// HTTPValidationError defines the model for HTTPValidationError.
type HTTPValidationError struct {
	Detail *[]ValidationError `json:"detail,omitempty"`
}

// InitializeProjectRequest defines the model for InitializeProjectRequest.
type InitializeProjectRequest struct {
	Activities    *[]ReposActivityCreateParams     `json:"activities"`
	KanbanSteps   *[]ReposKanbanStepCreateParams   `json:"kanbanSteps"`
	ProjectID     *int                             `json:"projectID,omitempty"`
	Teams         *[]ReposTeamCreateParams         `json:"teams"`
	WorkItemTags  *[]ReposWorkItemTagCreateParams  `json:"workItemTags"`
	WorkItemTypes *[]ReposWorkItemTypeCreateParams `json:"workItemTypes"`
}

// ModelsProjectConfigField defines the model for ModelsProjectConfigField.
type ModelsProjectConfigField struct {
	IsEditable    bool   `json:"isEditable"`
	IsVisible     bool   `json:"isVisible"`
	Name          string `json:"name"`
	Path          string `json:"path"`
	ShowCollapsed bool   `json:"showCollapsed"`
}

// ModelsRole defines the model for ModelsRole.
type ModelsRole = string

// NotificationType User notification type.
type NotificationType string

// PgtypeJSONB defines the model for PgtypeJSONB.
type PgtypeJSONB = map[string]interface{}

// Project Existing projects
type Project string

// ProjectConfig defines the model for ProjectConfig.
type ProjectConfig struct {
	Fields *[]ModelsProjectConfigField `json:"fields"`
	Header *[]string                   `json:"header"`
}

// ReposActivityCreateParams defines the model for ReposActivityCreateParams.
type ReposActivityCreateParams struct {
	Description  *string `json:"description,omitempty"`
	IsProductive *bool   `json:"isProductive,omitempty"`
	Name         *string `json:"name,omitempty"`
	ProjectID    *int    `json:"projectID,omitempty"`
}

// ReposKanbanStepCreateParams defines the model for ReposKanbanStepCreateParams.
type ReposKanbanStepCreateParams struct {
	Color         *string `json:"color,omitempty"`
	Description   *string `json:"description,omitempty"`
	Name          *string `json:"name,omitempty"`
	ProjectID     *int    `json:"projectID,omitempty"`
	StepOrder     *int    `json:"stepOrder,omitempty"`
	TimeTrackable *bool   `json:"timeTrackable,omitempty"`
}

// ReposTeamCreateParams defines the model for ReposTeamCreateParams.
type ReposTeamCreateParams struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	ProjectID   *int    `json:"projectID,omitempty"`
}

// ReposWorkItemTagCreateParams defines the model for ReposWorkItemTagCreateParams.
type ReposWorkItemTagCreateParams struct {
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	ProjectID   *int    `json:"projectID,omitempty"`
}

// ReposWorkItemTypeCreateParams defines the model for ReposWorkItemTypeCreateParams.
type ReposWorkItemTypeCreateParams struct {
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	ProjectID   *int    `json:"projectID,omitempty"`
}

// RestDemoProjectWorkItemsResponse defines the model for RestDemoProjectWorkItemsResponse.
type RestDemoProjectWorkItemsResponse struct {
	Closed              *time.Time             `json:"closed"`
	CreatedAt           time.Time              `json:"createdAt"`
	DeletedAt           *time.Time             `json:"deletedAt"`
	DemoProjectWorkItem *DbDemoProjectWorkItem `json:"demoProjectWorkItem"`
	Description         string                 `json:"description"`
	KanbanStepID        int                    `json:"kanbanStepID"`
	Members             *[]DbUser              `json:"members"`
	Metadata            PgtypeJSONB            `json:"metadata"`
	Project2workItem    *DbProject2WorkItem    `json:"project2workItem"`
	TargetDate          time.Time              `json:"targetDate"`
	TeamID              int                    `json:"teamID"`
	TimeEntries         *[]DbTimeEntry         `json:"timeEntries"`
	Title               string                 `json:"title"`
	UpdatedAt           time.Time              `json:"updatedAt"`
	WorkItemComments    *[]DbWorkItemComment   `json:"workItemComments"`
	WorkItemID          int                    `json:"workItemID"`
	WorkItemTags        *[]DbWorkItemTag       `json:"workItemTags"`
	WorkItemType        *DbWorkItemType        `json:"workItemType"`
	WorkItemTypeID      int                    `json:"workItemTypeID"`
}

// RestProjectBoardResponse defines the model for RestProjectBoardResponse.
type RestProjectBoardResponse struct {
	Project *DbProject `json:"project,omitempty"`
}

// Role defines the model for Role.
type Role string

// Scope defines the model for Scope.
type Scope string

// Scopes defines the model for Scopes.
type Scopes = []Scope

// Topics string identifiers for SSE event listeners.
type Topics string

// UpdateUserAuthRequest represents User authorization data to update
type UpdateUserAuthRequest struct {
	Role   *Role   `json:"role,omitempty"`
	Scopes *Scopes `json:"scopes,omitempty"`
}

// UpdateUserRequest represents User data to update
type UpdateUserRequest struct {
	// FirstName originally from auth server but updatable
	FirstName *string `json:"firstName,omitempty"`

	// LastName originally from auth server but updatable
	LastName *string `json:"lastName,omitempty"`
}

// UserResponse defines the model for UserResponse.
type UserResponse struct {
	ApiKey   *DbUserAPIKey `json:"apiKey"`
	Projects *[]DbProject  `json:"projects"`
	Role     Role          `json:"role"`
	Scopes   Scopes        `json:"scopes"`
	Teams    *[]DbTeam     `json:"teams"`
}

// UuidUUID defines the model for UuidUUID.
type UuidUUID = string

// ValidationError defines the model for ValidationError.
type ValidationError struct {
	Loc  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
}

// WorkItemRole Role in work item for a member.
type WorkItemRole string

// DemoProjectKanbanSteps Kanban columns for project demoProject
type DemoProjectKanbanSteps string

// PathSerial defines the model for PathSerial.
type PathSerial = int

// Uuid defines the model for uuid.
type Uuid = string

// GetProjectWorkitemsParams defines parameters for GetProjectWorkitems.
type GetProjectWorkitemsParams struct {
	Open    *bool `form:"open,omitempty" json:"open,omitempty"`
	Deleted *bool `form:"deleted,omitempty" json:"deleted,omitempty"`
}

// UpdateProjectConfigJSONRequestBody defines body for UpdateProjectConfig for application/json ContentType.
type UpdateProjectConfigJSONRequestBody = ProjectConfig

// InitializeProjectJSONRequestBody defines body for InitializeProject for application/json ContentType.
type InitializeProjectJSONRequestBody = InitializeProjectRequest

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = UpdateUserRequest

// UpdateUserAuthorizationJSONRequestBody defines body for UpdateUserAuthorization for application/json ContentType.
type UpdateUserAuthorizationJSONRequestBody = UpdateUserAuthRequest
