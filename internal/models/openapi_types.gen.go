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

// Defines values for Demo2WorkItemTypes.
const (
	Demo2WorkItemTypesAnotherType Demo2WorkItemTypes = "Another type"
	Demo2WorkItemTypesType1       Demo2WorkItemTypes = "Type 1"
	Demo2WorkItemTypesType2       Demo2WorkItemTypes = "Type 2"
)

// AllDemo2WorkItemTypesValues returns all possible values for Demo2WorkItemTypes.
func AllDemo2WorkItemTypesValues() []Demo2WorkItemTypes {
	return []Demo2WorkItemTypes{
		Demo2WorkItemTypesAnotherType,
		Demo2WorkItemTypesType1,
		Demo2WorkItemTypesType2,
	}
}

// Defines values for DemoKanbanSteps.
const (
	DemoKanbanStepsDisabled       DemoKanbanSteps = "Disabled"
	DemoKanbanStepsReceived       DemoKanbanSteps = "Received"
	DemoKanbanStepsUnderReview    DemoKanbanSteps = "Under review"
	DemoKanbanStepsWorkInProgress DemoKanbanSteps = "Work in progress"
)

// AllDemoKanbanStepsValues returns all possible values for DemoKanbanSteps.
func AllDemoKanbanStepsValues() []DemoKanbanSteps {
	return []DemoKanbanSteps{
		DemoKanbanStepsDisabled,
		DemoKanbanStepsReceived,
		DemoKanbanStepsUnderReview,
		DemoKanbanStepsWorkInProgress,
	}
}

// Defines values for DemoProject2KanbanSteps.
const (
	DemoProject2KanbanStepsReceived DemoProject2KanbanSteps = "Received"
)

// AllDemoProject2KanbanStepsValues returns all possible values for DemoProject2KanbanSteps.
func AllDemoProject2KanbanStepsValues() []DemoProject2KanbanSteps {
	return []DemoProject2KanbanSteps{
		DemoProject2KanbanStepsReceived,
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

// Defines values for DemoTwoKanbanSteps.
const (
	DemoTwoKanbanStepsReceived DemoTwoKanbanSteps = "Received"
)

// AllDemoTwoKanbanStepsValues returns all possible values for DemoTwoKanbanSteps.
func AllDemoTwoKanbanStepsValues() []DemoTwoKanbanSteps {
	return []DemoTwoKanbanSteps{
		DemoTwoKanbanStepsReceived,
	}
}

// Defines values for DemoTwoWorkItemTypes.
const (
	DemoTwoWorkItemTypesAnotherType DemoTwoWorkItemTypes = "Another type"
	DemoTwoWorkItemTypesType1       DemoTwoWorkItemTypes = "Type 1"
	DemoTwoWorkItemTypesType2       DemoTwoWorkItemTypes = "Type 2"
)

// AllDemoTwoWorkItemTypesValues returns all possible values for DemoTwoWorkItemTypes.
func AllDemoTwoWorkItemTypesValues() []DemoTwoWorkItemTypes {
	return []DemoTwoWorkItemTypes{
		DemoTwoWorkItemTypesAnotherType,
		DemoTwoWorkItemTypesType1,
		DemoTwoWorkItemTypesType2,
	}
}

// Defines values for DemoWorkItemTypes.
const (
	DemoWorkItemTypesType1 DemoWorkItemTypes = "Type 1"
)

// AllDemoWorkItemTypesValues returns all possible values for DemoWorkItemTypes.
func AllDemoWorkItemTypesValues() []DemoWorkItemTypes {
	return []DemoWorkItemTypes{
		DemoWorkItemTypesType1,
	}
}

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
	ProjectDemo    Project = "demo"
	ProjectDemoTwo Project = "demo_two"
)

// AllProjectValues returns all possible values for Project.
func AllProjectValues() []Project {
	return []Project{
		ProjectDemo,
		ProjectDemoTwo,
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
	ScopeUsersRead            Scope = "users:read"
	ScopeUsersWrite           Scope = "users:write"
	ScopeWorkItemReview       Scope = "work-item:review"
	ScopeWorkItemTagCreate    Scope = "work-item-tag:create"
	ScopeWorkItemTagDelete    Scope = "work-item-tag:delete"
	ScopeWorkItemTagEdit      Scope = "work-item-tag:edit"
)

// AllScopeValues returns all possible values for Scope.
func AllScopeValues() []Scope {
	return []Scope{
		ScopeProjectSettingsWrite,
		ScopeScopesWrite,
		ScopeTeamSettingsWrite,
		ScopeUsersRead,
		ScopeUsersWrite,
		ScopeWorkItemReview,
		ScopeWorkItemTagCreate,
		ScopeWorkItemTagDelete,
		ScopeWorkItemTagEdit,
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

// DbActivity defines the model for DbActivity.
type DbActivity struct {
	ActivityID   int    `json:"activityID"`
	Description  string `json:"description"`
	IsProductive bool   `json:"isProductive"`
	Name         string `json:"name"`
	ProjectID    int    `json:"projectID"`
}

// DbActivityCreateParams defines the model for DbActivityCreateParams.
type DbActivityCreateParams struct {
	Description  string `json:"description"`
	IsProductive bool   `json:"isProductive"`
	Name         string `json:"name"`
	ProjectID    int    `json:"projectID"`
}

// DbDemoWorkItem defines the model for DbDemoWorkItem.
type DbDemoWorkItem struct {
	LastMessageAt time.Time `json:"lastMessageAt"`
	Line          string    `json:"line"`
	Ref           string    `json:"ref"`
	Reopened      bool      `json:"reopened"`
	WorkItemID    int       `json:"workItemID"`
}

// DbDemoWorkItemCreateParams defines the model for DbDemoWorkItemCreateParams.
type DbDemoWorkItemCreateParams struct {
	LastMessageAt time.Time `json:"lastMessageAt"`
	Line          string    `json:"line"`
	Ref           string    `json:"ref"`
	Reopened      bool      `json:"reopened"`
	WorkItemID    int       `json:"workItemID"`
}

// DbKanbanStep defines the model for DbKanbanStep.
type DbKanbanStep struct {
	Color         string `json:"color"`
	Description   string `json:"description"`
	KanbanStepID  int    `json:"kanbanStepID"`
	Name          string `json:"name"`
	ProjectID     int    `json:"projectID"`
	StepOrder     int    `json:"stepOrder"`
	TimeTrackable bool   `json:"timeTrackable"`
}

// DbProject defines the model for DbProject.
type DbProject struct {
	BoardConfig ProjectConfig `json:"boardConfig"`
	CreatedAt   time.Time     `json:"createdAt"`
	Description string        `json:"description"`
	Name        Project       `json:"name"`
	ProjectID   int           `json:"projectID"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}

// DbTeam defines the model for DbTeam.
type DbTeam struct {
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	ProjectID   int       `json:"projectID"`
	TeamID      int       `json:"teamID"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// DbTeamCreateParams defines the model for DbTeamCreateParams.
type DbTeamCreateParams struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	ProjectID   int    `json:"projectID"`
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
	CreatedAt                time.Time  `json:"createdAt"`
	DeletedAt                *time.Time `json:"deletedAt"`
	Email                    string     `json:"email"`
	FirstName                *string    `json:"firstName"`
	FullName                 *string    `json:"fullName"`
	HasGlobalNotifications   bool       `json:"hasGlobalNotifications"`
	HasPersonalNotifications bool       `json:"hasPersonalNotifications"`
	LastName                 *string    `json:"lastName"`
	Scopes                   Scopes     `json:"scopes"`
	UserID                   UuidUUID   `json:"userID"`
	Username                 string     `json:"username"`
}

// DbUserAPIKey defines the model for DbUserAPIKey.
type DbUserAPIKey struct {
	ApiKey    string    `json:"apiKey"`
	ExpiresOn time.Time `json:"expiresOn"`
	UserID    UuidUUID  `json:"userID"`
}

// DbWorkItem defines the model for DbWorkItem.
type DbWorkItem struct {
	Closed         *time.Time `json:"closed"`
	CreatedAt      time.Time  `json:"createdAt"`
	DeletedAt      *time.Time `json:"deletedAt"`
	Description    string     `json:"description"`
	KanbanStepID   int        `json:"kanbanStepID"`
	Metadata       *[]int     `json:"metadata"`
	TargetDate     time.Time  `json:"targetDate"`
	TeamID         int        `json:"teamID"`
	Title          string     `json:"title"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	WorkItemID     int        `json:"workItemID"`
	WorkItemTypeID int        `json:"workItemTypeID"`
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

// DbWorkItemCreateParams defines the model for DbWorkItemCreateParams.
type DbWorkItemCreateParams struct {
	Closed         *time.Time `json:"closed"`
	Description    string     `json:"description"`
	KanbanStepID   int        `json:"kanbanStepID"`
	Metadata       *[]int     `json:"metadata"`
	TargetDate     time.Time  `json:"targetDate"`
	TeamID         int        `json:"teamID"`
	Title          string     `json:"title"`
	WorkItemTypeID int        `json:"workItemTypeID"`
}

// DbWorkItemRole defines the model for DbWorkItemRole.
type DbWorkItemRole = string

// DbWorkItemTag defines the model for DbWorkItemTag.
type DbWorkItemTag struct {
	Color         string `json:"color"`
	Description   string `json:"description"`
	Name          string `json:"name"`
	ProjectID     int    `json:"projectID"`
	WorkItemTagID int    `json:"workItemTagID"`
}

// DbWorkItemTagCreateParams defines the model for DbWorkItemTagCreateParams.
type DbWorkItemTagCreateParams struct {
	Color       string `json:"color"`
	Description string `json:"description"`
	Name        string `json:"name"`
	ProjectID   int    `json:"projectID"`
}

// DbWorkItemType defines the model for DbWorkItemType.
type DbWorkItemType struct {
	Color          string `json:"color"`
	Description    string `json:"description"`
	Name           string `json:"name"`
	ProjectID      int    `json:"projectID"`
	WorkItemTypeID int    `json:"workItemTypeID"`
}

// DbWorkItemAssignedUser defines the model for DbWorkItem_AssignedUser.
type DbWorkItemAssignedUser struct {
	// Role represents a database 'work_item_role'
	Role WorkItemRole `json:"role"`
	User *DbUser      `json:"user,omitempty"`
}

// Demo2WorkItemTypes defines the model for Demo2WorkItemTypes.
type Demo2WorkItemTypes string

// DemoKanbanSteps defines the model for DemoKanbanSteps.
type DemoKanbanSteps string

// DemoProject2KanbanSteps defines the model for DemoProject2KanbanSteps.
type DemoProject2KanbanSteps string

// DemoProjectKanbanSteps defines the model for DemoProjectKanbanSteps.
type DemoProjectKanbanSteps string

// DemoTwoKanbanSteps defines the model for DemoTwoKanbanSteps.
type DemoTwoKanbanSteps string

// DemoTwoWorkItemTypes defines the model for DemoTwoWorkItemTypes.
type DemoTwoWorkItemTypes string

// DemoWorkItemTypes defines the model for DemoWorkItemTypes.
type DemoWorkItemTypes string

// HTTPError represents an error message response.
type HTTPError struct {
	Detail          string               `json:"detail"`
	Error           string               `json:"error"`
	Status          int                  `json:"status"`
	Title           string               `json:"title"`
	Type            string               `json:"type"`
	ValidationError *HTTPValidationError `json:"validationError,omitempty"`
}

// HTTPValidationError defines the model for HTTPValidationError.
type HTTPValidationError struct {
	// Detail Additional details for validation errors
	Detail *[]ValidationError `json:"detail,omitempty"`

	// Messages Descriptive error messages to show in a callout
	Messages []string `json:"messages"`
}

// InitializeProjectRequest defines the model for InitializeProjectRequest.
type InitializeProjectRequest struct {
	Activities   *[]DbActivityCreateParams    `json:"activities"`
	ProjectID    *int                         `json:"projectID,omitempty"`
	Teams        *[]DbTeamCreateParams        `json:"teams"`
	WorkItemTags *[]DbWorkItemTagCreateParams `json:"workItemTags"`
}

// ModelsWorkItemRole defines the model for ModelsWorkItemRole.
type ModelsWorkItemRole = string

// NotificationType represents a database 'notification_type'
type NotificationType string

// PgtypeJSONB defines the model for PgtypeJSONB.
type PgtypeJSONB = map[string]interface{}

// Project defines the model for Project.
type Project string

// ProjectConfig defines the model for ProjectConfig.
type ProjectConfig struct {
	Fields []ProjectConfigField `json:"fields"`
	Header []string             `json:"header"`
}

// ProjectConfigField defines the model for ProjectConfigField.
type ProjectConfigField struct {
	IsEditable    bool   `json:"isEditable"`
	IsVisible     bool   `json:"isVisible"`
	Name          string `json:"name"`
	Path          string `json:"path"`
	ShowCollapsed bool   `json:"showCollapsed"`
}

// RestDemoWorkItemCreateRequest defines the model for RestDemoWorkItemCreateRequest.
type RestDemoWorkItemCreateRequest struct {
	Base        DbWorkItemCreateParams     `json:"base"`
	DemoProject DbDemoWorkItemCreateParams `json:"demoProject"`
	Members     *[]ServicesMember          `json:"members"`
	TagIDs      *[]int                     `json:"tagIDs"`
}

// RestDemoWorkItemsResponse defines the model for RestDemoWorkItemsResponse.
type RestDemoWorkItemsResponse struct {
	Closed           *time.Time           `json:"closed"`
	CreatedAt        time.Time            `json:"createdAt"`
	DeletedAt        *time.Time           `json:"deletedAt"`
	DemoWorkItem     DbDemoWorkItem       `json:"demoWorkItem"`
	Description      string               `json:"description"`
	KanbanStepID     int                  `json:"kanbanStepID"`
	Members          *[]DbUser            `json:"members"`
	Metadata         *[]int               `json:"metadata"`
	TargetDate       time.Time            `json:"targetDate"`
	TeamID           int                  `json:"teamID"`
	TimeEntries      *[]DbTimeEntry       `json:"timeEntries"`
	Title            string               `json:"title"`
	UpdatedAt        time.Time            `json:"updatedAt"`
	WorkItemComments *[]DbWorkItemComment `json:"workItemComments"`
	WorkItemID       int                  `json:"workItemID"`
	WorkItemTags     *[]DbWorkItemTag     `json:"workItemTags"`
	WorkItemType     *DbWorkItemType      `json:"workItemType,omitempty"`
	WorkItemTypeID   int                  `json:"workItemTypeID"`
}

// RestProjectBoardResponse defines the model for RestProjectBoardResponse.
type RestProjectBoardResponse struct {
	Activities    *[]DbActivity     `json:"activities"`
	BoardConfig   ProjectConfig     `json:"boardConfig"`
	CreatedAt     time.Time         `json:"createdAt"`
	Description   string            `json:"description"`
	KanbanSteps   *[]DbKanbanStep   `json:"kanbanSteps"`
	Name          Project           `json:"name"`
	ProjectID     int               `json:"projectID"`
	Teams         *[]DbTeam         `json:"teams"`
	UpdatedAt     time.Time         `json:"updatedAt"`
	WorkItemTags  *[]DbWorkItemTag  `json:"workItemTags"`
	WorkItemTypes *[]DbWorkItemType `json:"workItemTypes"`
}

// RestWorkItemCommentCreateRequest defines the model for RestWorkItemCommentCreateRequest.
type RestWorkItemCommentCreateRequest struct {
	Message    string   `json:"message"`
	UserID     UuidUUID `json:"userID"`
	WorkItemID int      `json:"workItemID"`
}

// RestWorkItemTagCreateRequest defines the model for RestWorkItemTagCreateRequest.
type RestWorkItemTagCreateRequest struct {
	Color       string `json:"color"`
	Description string `json:"description"`
	Name        string `json:"name"`
	ProjectID   int    `json:"projectID"`
}

// Role defines the model for Role.
type Role string

// Scope defines the model for Scope.
type Scope string

// Scopes defines the model for Scopes.
type Scopes = []Scope

// ServicesMember defines the model for ServicesMember.
type ServicesMember struct {
	Role   ModelsWorkItemRole `json:"role"`
	UserID UuidUUID           `json:"userID"`
}

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
	ApiKey                   *DbUserAPIKey `json:"apiKey,omitempty"`
	CreatedAt                time.Time     `json:"createdAt"`
	DeletedAt                *time.Time    `json:"deletedAt"`
	Email                    string        `json:"email"`
	FirstName                *string       `json:"firstName"`
	FullName                 *string       `json:"fullName"`
	HasGlobalNotifications   bool          `json:"hasGlobalNotifications"`
	HasPersonalNotifications bool          `json:"hasPersonalNotifications"`
	LastName                 *string       `json:"lastName"`
	Projects                 *[]DbProject  `json:"projects"`
	Role                     Role          `json:"role"`
	Scopes                   Scopes        `json:"scopes"`
	Teams                    *[]DbTeam     `json:"teams"`
	UserID                   UuidUUID      `json:"userID"`
	Username                 string        `json:"username"`
}

// UuidUUID defines the model for UuidUUID.
type UuidUUID = string

// ValidationError defines the model for ValidationError.
type ValidationError struct {
	Ctx *map[string]interface{} `json:"ctx,omitempty"`

	// Detail verbose details of the error
	Detail struct {
		Schema map[string]interface{} `json:"schema"`
		Value  string                 `json:"value"`
	} `json:"detail"`

	// Loc location in body path, if any
	Loc []string `json:"loc"`

	// Msg should always be shown to the user
	Msg string `json:"msg"`
}

// WorkItemRole represents a database 'work_item_role'
type WorkItemRole string

// ProjectName defines the model for ProjectName.
type ProjectName = Project

// Serial defines the model for Serial.
type Serial = int

// UUID defines the model for UUID.
type UUID = string

// EventsParams defines parameters for Events.
type EventsParams struct {
	ProjectName Project `form:"projectName" json:"projectName"`
}

// GetProjectWorkitemsParams defines parameters for GetProjectWorkitems.
type GetProjectWorkitemsParams struct {
	Open    *bool `form:"open,omitempty" json:"open,omitempty"`
	Deleted *bool `form:"deleted,omitempty" json:"deleted,omitempty"`
}

// UpdateProjectConfigJSONRequestBody defines body for UpdateProjectConfig for application/json ContentType.
type UpdateProjectConfigJSONRequestBody = ProjectConfig

// InitializeProjectJSONRequestBody defines body for InitializeProject for application/json ContentType.
type InitializeProjectJSONRequestBody = InitializeProjectRequest

// CreateWorkitemTagJSONRequestBody defines body for CreateWorkitemTag for application/json ContentType.
type CreateWorkitemTagJSONRequestBody = RestWorkItemTagCreateRequest

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = UpdateUserRequest

// UpdateUserAuthorizationJSONRequestBody defines body for UpdateUserAuthorization for application/json ContentType.
type UpdateUserAuthorizationJSONRequestBody = UpdateUserAuthRequest

// CreateWorkitemJSONRequestBody defines body for CreateWorkitem for application/json ContentType.
type CreateWorkitemJSONRequestBody = RestDemoWorkItemCreateRequest

// CreateWorkitemCommentJSONRequestBody defines body for CreateWorkitemComment for application/json ContentType.
type CreateWorkitemCommentJSONRequestBody = RestWorkItemCommentCreateRequest
