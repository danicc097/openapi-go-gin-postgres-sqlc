// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package models

import (
	"encoding/json"
	"errors"
	"time"

	uuid "github.com/google/uuid"
)

const (
	Api_keyScopes     = "api_key.Scopes"
	Bearer_authScopes = "bearer_auth.Scopes"
)

// DemoKanbanSteps is generated from kanban_steps table.
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

// DemoTwoKanbanSteps is generated from kanban_steps table.
const (
	DemoTwoKanbanStepsReceived DemoTwoKanbanSteps = "Received"
)

// AllDemoTwoKanbanStepsValues returns all possible values for DemoTwoKanbanSteps.
func AllDemoTwoKanbanStepsValues() []DemoTwoKanbanSteps {
	return []DemoTwoKanbanSteps{
		DemoTwoKanbanStepsReceived,
	}
}

// DemoTwoWorkItemTypes is generated from work_item_types table.
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

// DemoWorkItemTypes is generated from work_item_types table.
const (
	DemoWorkItemTypesType1 DemoWorkItemTypes = "Type 1"
)

// AllDemoWorkItemTypesValues returns all possible values for DemoWorkItemTypes.
func AllDemoWorkItemTypesValues() []DemoWorkItemTypes {
	return []DemoWorkItemTypes{
		DemoWorkItemTypesType1,
	}
}

// Defines values for Direction.
const (
	DirectionAsc  Direction = "asc"
	DirectionDesc Direction = "desc"
)

// AllDirectionValues returns all possible values for Direction.
func AllDirectionValues() []Direction {
	return []Direction{
		DirectionAsc,
		DirectionDesc,
	}
}

// ErrorCode Represents standardized HTTP error types.
// Notes:
// - 'Private' marks an error to be hidden in response.
const (
	ErrorCodeAlreadyExists      ErrorCode = "AlreadyExists"
	ErrorCodeInvalidArgument    ErrorCode = "InvalidArgument"
	ErrorCodeInvalidRole        ErrorCode = "InvalidRole"
	ErrorCodeInvalidScope       ErrorCode = "InvalidScope"
	ErrorCodeInvalidUUID        ErrorCode = "InvalidUUID"
	ErrorCodeNotFound           ErrorCode = "NotFound"
	ErrorCodeOIDC               ErrorCode = "OIDC"
	ErrorCodePrivate            ErrorCode = "Private"
	ErrorCodeRequestValidation  ErrorCode = "RequestValidation"
	ErrorCodeResponseValidation ErrorCode = "ResponseValidation"
	ErrorCodeUnauthenticated    ErrorCode = "Unauthenticated"
	ErrorCodeUnauthorized       ErrorCode = "Unauthorized"
	ErrorCodeUnknown            ErrorCode = "Unknown"
)

// AllErrorCodeValues returns all possible values for ErrorCode.
func AllErrorCodeValues() []ErrorCode {
	return []ErrorCode{
		ErrorCodeAlreadyExists,
		ErrorCodeInvalidArgument,
		ErrorCodeInvalidRole,
		ErrorCodeInvalidScope,
		ErrorCodeInvalidUUID,
		ErrorCodeNotFound,
		ErrorCodeOIDC,
		ErrorCodePrivate,
		ErrorCodeRequestValidation,
		ErrorCodeResponseValidation,
		ErrorCodeUnauthenticated,
		ErrorCodeUnauthorized,
		ErrorCodeUnknown,
	}
}

// NotificationType is generated from database enum 'notification_type'.
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

// Project is generated from projects table.
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

// Role is generated from roles.json keys.
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

// Scope is generated from scopes.json keys.
const (
	ScopeActivityCreate       Scope = "activity:create"
	ScopeActivityDelete       Scope = "activity:delete"
	ScopeActivityEdit         Scope = "activity:edit"
	ScopeProjectMember        Scope = "project-member"
	ScopeProjectSettingsWrite Scope = "project-settings:write"
	ScopeScopesWrite          Scope = "scopes:write"
	ScopeTeamSettingsWrite    Scope = "team-settings:write"
	ScopeUsersDelete          Scope = "users:delete"
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
		ScopeActivityCreate,
		ScopeActivityDelete,
		ScopeActivityEdit,
		ScopeProjectMember,
		ScopeProjectSettingsWrite,
		ScopeScopesWrite,
		ScopeTeamSettingsWrite,
		ScopeUsersDelete,
		ScopeUsersRead,
		ScopeUsersWrite,
		ScopeWorkItemReview,
		ScopeWorkItemTagCreate,
		ScopeWorkItemTagDelete,
		ScopeWorkItemTagEdit,
	}
}

// Topics string identifiers for SSE event listeners.
const (
	TopicsGlobalAlerts Topics = "GlobalAlerts"
)

// AllTopicsValues returns all possible values for Topics.
func AllTopicsValues() []Topics {
	return []Topics{
		TopicsGlobalAlerts,
	}
}

// WorkItemRole is generated from database enum 'work_item_role'.
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

// Activity defines the model for Activity.
type Activity struct {
	ActivityID   int        `json:"activityID"`
	DeletedAt    *time.Time `json:"deletedAt"`
	Description  string     `json:"description"`
	IsProductive bool       `json:"isProductive"`
	Name         string     `json:"name"`
	ProjectID    int        `json:"projectID"`
}

// CreateActivityRequest defines the model for CreateActivityRequest.
type CreateActivityRequest struct {
	Description  string `json:"description"`
	IsProductive bool   `json:"isProductive"`
	Name         string `json:"name"`
}

// CreateDemoTwoWorkItemRequest defines the model for CreateDemoTwoWorkItemRequest.
type CreateDemoTwoWorkItemRequest struct {
	Base           DbWorkItemCreateParams        `json:"base"`
	DemoTwoProject DbDemoTwoWorkItemCreateParams `json:"demoTwoProject"`
	Members        []ServicesMember              `json:"members"`

	// ProjectName is generated from projects table.
	ProjectName Project `json:"projectName"`
	TagIDs      []int   `json:"tagIDs"`
}

// CreateDemoWorkItemRequest defines the model for CreateDemoWorkItemRequest.
type CreateDemoWorkItemRequest struct {
	Base        DbWorkItemCreateParams     `json:"base"`
	DemoProject DbDemoWorkItemCreateParams `json:"demoProject"`
	Members     []ServicesMember           `json:"members"`

	// ProjectName is generated from projects table.
	ProjectName Project `json:"projectName"`
	TagIDs      []int   `json:"tagIDs"`
}

// CreateTeamRequest defines the model for CreateTeamRequest.
type CreateTeamRequest struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// CreateWorkItemCommentRequest defines the model for CreateWorkItemCommentRequest.
type CreateWorkItemCommentRequest struct {
	Message    string   `json:"message"`
	UserID     DbUserID `json:"userID"`
	WorkItemID int      `json:"workItemID"`
}

// CreateWorkItemRequest defines the model for CreateWorkItemRequest.
type CreateWorkItemRequest struct {
	union json.RawMessage
}

// CreateWorkItemTagRequest defines the model for CreateWorkItemTagRequest.
type CreateWorkItemTagRequest struct {
	Color       string `json:"color"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// CreateWorkItemTypeRequest defines the model for CreateWorkItemTypeRequest.
type CreateWorkItemTypeRequest struct {
	Color       string `json:"color"`
	Description string `json:"description"`
	Name        string `json:"name"`
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
	ProjectID    *int   `json:"projectID,omitempty"`
}

// DbDemoTwoWorkItem defines the model for DbDemoTwoWorkItem.
type DbDemoTwoWorkItem struct {
	CustomDateForProject2 *time.Time `json:"customDateForProject2"`
	WorkItemID            int        `json:"workItemID"`
}

// DbDemoTwoWorkItemCreateParams defines the model for DbDemoTwoWorkItemCreateParams.
type DbDemoTwoWorkItemCreateParams struct {
	CustomDateForProject2 *time.Time `json:"customDateForProject2"`
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

// DbNotification defines the model for DbNotification.
type DbNotification struct {
	Body           string    `json:"body"`
	CreatedAt      time.Time `json:"createdAt"`
	Labels         []string  `json:"labels"`
	Link           *string   `json:"link"`
	NotificationID int       `json:"notificationID"`

	// NotificationType is generated from database enum 'notification_type'.
	NotificationType NotificationType `json:"notificationType"`
	Receiver         *DbUserID        `json:"receiver,omitempty"`
	Sender           DbUserID         `json:"sender"`
	Title            string           `json:"title"`
}

// DbNotificationID defines the model for DbNotificationID.
type DbNotificationID = interface{}

// DbProject defines the model for DbProject.
type DbProject struct {
	BoardConfig ProjectConfig `json:"boardConfig"`
	CreatedAt   time.Time     `json:"createdAt"`
	Description string        `json:"description"`

	// Name is generated from projects table.
	Name      Project   `json:"name"`
	ProjectID int       `json:"projectID"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// DbProjectID defines the model for DbProjectID.
type DbProjectID = interface{}

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
}

// DbTimeEntry defines the model for DbTimeEntry.
type DbTimeEntry struct {
	ActivityID      int       `json:"activityID"`
	Comment         string    `json:"comment"`
	DurationMinutes *int      `json:"durationMinutes"`
	Start           time.Time `json:"start"`
	TeamID          *int      `json:"teamID"`
	TimeEntryID     int       `json:"timeEntryID"`
	UserID          DbUserID  `json:"userID"`
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
	UserID                   DbUserID   `json:"userID"`
	Username                 string     `json:"username"`
}

// DbUserAPIKey defines the model for DbUserAPIKey.
type DbUserAPIKey struct {
	ApiKey    string    `json:"apiKey"`
	ExpiresOn time.Time `json:"expiresOn"`
	UserID    DbUserID  `json:"userID"`
}

// DbUserID defines the model for DbUserID.
type DbUserID = uuid.UUID

// DbUserNotification defines the model for DbUserNotification.
type DbUserNotification struct {
	NotificationID     int      `json:"notificationID"`
	Read               bool     `json:"read"`
	UserID             DbUserID `json:"userID"`
	UserNotificationID int      `json:"userNotificationID"`
}

// DbUserWIAUWorkItem defines the model for DbUserWIAUWorkItem.
type DbUserWIAUWorkItem struct {
	// Role is generated from database enum 'work_item_role'.
	Role WorkItemRole `json:"role"`
	User DbUser       `json:"user"`
}

// DbWorkItem defines the model for DbWorkItem.
type DbWorkItem struct {
	ClosedAt       *time.Time             `json:"closedAt"`
	CreatedAt      time.Time              `json:"createdAt"`
	DeletedAt      *time.Time             `json:"deletedAt"`
	Description    string                 `json:"description"`
	KanbanStepID   int                    `json:"kanbanStepID"`
	Metadata       map[string]interface{} `json:"metadata"`
	TargetDate     time.Time              `json:"targetDate"`
	TeamID         int                    `json:"teamID"`
	Title          string                 `json:"title"`
	UpdatedAt      time.Time              `json:"updatedAt"`
	WorkItemID     int                    `json:"workItemID"`
	WorkItemTypeID int                    `json:"workItemTypeID"`
}

// DbWorkItemComment defines the model for DbWorkItemComment.
type DbWorkItemComment struct {
	CreatedAt         time.Time `json:"createdAt"`
	Message           string    `json:"message"`
	UpdatedAt         time.Time `json:"updatedAt"`
	UserID            DbUserID  `json:"userID"`
	WorkItemCommentID int       `json:"workItemCommentID"`
	WorkItemID        int       `json:"workItemID"`
}

// DbWorkItemCreateParams defines the model for DbWorkItemCreateParams.
type DbWorkItemCreateParams struct {
	ClosedAt       *time.Time             `json:"closedAt"`
	Description    string                 `json:"description"`
	KanbanStepID   int                    `json:"kanbanStepID"`
	Metadata       map[string]interface{} `json:"metadata"`
	TargetDate     time.Time              `json:"targetDate"`
	TeamID         int                    `json:"teamID"`
	Title          string                 `json:"title"`
	WorkItemTypeID int                    `json:"workItemTypeID"`
}

// DbWorkItemID defines the model for DbWorkItemID.
type DbWorkItemID = interface{}

// DbWorkItemRole defines the model for DbWorkItemRole.
type DbWorkItemRole = string

// DbWorkItemTag defines the model for DbWorkItemTag.
type DbWorkItemTag struct {
	Color         string     `json:"color"`
	DeletedAt     *time.Time `json:"deletedAt"`
	Description   string     `json:"description"`
	Name          string     `json:"name"`
	ProjectID     int        `json:"projectID"`
	WorkItemTagID int        `json:"workItemTagID"`
}

// DbWorkItemTagCreateParams defines the model for DbWorkItemTagCreateParams.
type DbWorkItemTagCreateParams struct {
	Color       string `json:"color"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// DbWorkItemType defines the model for DbWorkItemType.
type DbWorkItemType struct {
	Color          string `json:"color"`
	Description    string `json:"description"`
	Name           string `json:"name"`
	ProjectID      int    `json:"projectID"`
	WorkItemTypeID int    `json:"workItemTypeID"`
}

// DbWorkItemTypeID defines the model for DbWorkItemTypeID.
type DbWorkItemTypeID = interface{}

// DemoKanbanSteps is generated from kanban_steps table.
type DemoKanbanSteps string

// DemoTwoKanbanSteps is generated from kanban_steps table.
type DemoTwoKanbanSteps string

// DemoTwoWorkItemTypes is generated from work_item_types table.
type DemoTwoWorkItemTypes string

// DemoTwoWorkItems defines the model for DemoTwoWorkItems.
type DemoTwoWorkItems struct {
	ClosedAt         *time.Time             `json:"closedAt"`
	CreatedAt        time.Time              `json:"createdAt"`
	DeletedAt        *time.Time             `json:"deletedAt"`
	DemoTwoWorkItem  DbDemoTwoWorkItem      `json:"demoTwoWorkItem"`
	Description      string                 `json:"description"`
	KanbanStepID     int                    `json:"kanbanStepID"`
	Members          *[]DbUserWIAUWorkItem  `json:"members"`
	Metadata         map[string]interface{} `json:"metadata"`
	TargetDate       time.Time              `json:"targetDate"`
	TeamID           *int                   `json:"teamID"`
	TimeEntries      *[]DbTimeEntry         `json:"timeEntries"`
	Title            string                 `json:"title"`
	UpdatedAt        time.Time              `json:"updatedAt"`
	WorkItemComments *[]DbWorkItemComment   `json:"workItemComments"`
	WorkItemID       int                    `json:"workItemID"`
	WorkItemTags     *[]DbWorkItemTag       `json:"workItemTags"`
	WorkItemType     *DbWorkItemType        `json:"workItemType,omitempty"`
	WorkItemTypeID   int                    `json:"workItemTypeID"`
}

// DemoWorkItemTypes is generated from work_item_types table.
type DemoWorkItemTypes string

// DemoWorkItems defines the model for DemoWorkItems.
type DemoWorkItems struct {
	ClosedAt         *time.Time             `json:"closedAt"`
	CreatedAt        time.Time              `json:"createdAt"`
	DeletedAt        *time.Time             `json:"deletedAt"`
	DemoWorkItem     DbDemoWorkItem         `json:"demoWorkItem"`
	Description      string                 `json:"description"`
	KanbanStepID     int                    `json:"kanbanStepID"`
	Members          *[]DbUserWIAUWorkItem  `json:"members"`
	Metadata         map[string]interface{} `json:"metadata"`
	TargetDate       time.Time              `json:"targetDate"`
	TeamID           *int                   `json:"teamID"`
	TimeEntries      *[]DbTimeEntry         `json:"timeEntries"`
	Title            string                 `json:"title"`
	UpdatedAt        time.Time              `json:"updatedAt"`
	WorkItemComments *[]DbWorkItemComment   `json:"workItemComments"`
	WorkItemID       int                    `json:"workItemID"`
	WorkItemTags     *[]DbWorkItemTag       `json:"workItemTags"`
	WorkItemType     *DbWorkItemType        `json:"workItemType,omitempty"`
	WorkItemTypeID   int                    `json:"workItemTypeID"`
}

// Direction defines the model for Direction.
type Direction string

// ErrorCode Represents standardized HTTP error types.
// Notes:
// - 'Private' marks an error to be hidden in response.
type ErrorCode string

// HTTPError represents an error message response.
type HTTPError struct {
	Detail string `json:"detail"`
	Error  string `json:"error"`

	// Loc location in body path, if any
	Loc    *[]string `json:"loc,omitempty"`
	Status int       `json:"status"`
	Title  string    `json:"title"`

	// Type Represents standardized HTTP error types.
	// Notes:
	// - 'Private' marks an error to be hidden in response.
	Type            ErrorCode            `json:"type"`
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
	Tags  *[]DbWorkItemTagCreateParams `json:"tags"`
	Teams *[]DbTeamCreateParams        `json:"teams"`
}

// Notification defines the model for Notification.
type Notification struct {
	Notification       DbNotification `json:"notification"`
	NotificationID     int            `json:"notificationID"`
	Read               bool           `json:"read"`
	UserID             DbUserID       `json:"userID"`
	UserNotificationID int            `json:"userNotificationID"`
}

// NotificationType is generated from database enum 'notification_type'.
type NotificationType string

// PaginatedNotificationsResponse defines the model for PaginatedNotificationsResponse.
type PaginatedNotificationsResponse struct {
	Items *[]RestNotification `json:"items"`
	Page  RestPaginationPage  `json:"page"`
}

// PaginatedUsersResponse defines the model for PaginatedUsersResponse.
type PaginatedUsersResponse struct {
	Items *[]RestUser        `json:"items"`
	Page  RestPaginationPage `json:"page"`
}

// Project is generated from projects table.
type Project string

// ProjectBoard defines the model for ProjectBoard.
type ProjectBoard struct {
	// ProjectName is generated from projects table.
	ProjectName Project `json:"projectName"`
}

// ProjectConfig defines the model for ProjectConfig.
type ProjectConfig struct {
	Fields        []ProjectConfigField    `json:"fields"`
	Header        []string                `json:"header"`
	Visualization *map[string]interface{} `json:"visualization,omitempty"`
}

// ProjectConfigField defines the model for ProjectConfigField.
type ProjectConfigField struct {
	IsEditable    bool   `json:"isEditable"`
	IsVisible     bool   `json:"isVisible"`
	Name          string `json:"name"`
	Path          string `json:"path"`
	ShowCollapsed bool   `json:"showCollapsed"`
}

// RestNotification defines the model for RestNotification.
type RestNotification struct {
	Notification       DbNotification `json:"notification"`
	NotificationID     int            `json:"notificationID"`
	Read               bool           `json:"read"`
	UserID             DbUserID       `json:"userID"`
	UserNotificationID int            `json:"userNotificationID"`
}

// RestPaginationPage defines the model for RestPaginationPage.
type RestPaginationPage struct {
	NextCursor *string `json:"nextCursor,omitempty"`
}

// RestUser defines the model for RestUser.
type RestUser = User

// Role is generated from roles.json keys.
type Role string

// Scope is generated from scopes.json keys.
type Scope string

// Scopes defines the model for Scopes.
type Scopes = []Scope

// ServicesMember defines the model for ServicesMember.
type ServicesMember struct {
	// Role is generated from database enum 'work_item_role'.
	Role   WorkItemRole `json:"role"`
	UserID DbUserID     `json:"userID"`
}

// Team defines the model for Team.
type Team struct {
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	ProjectID   int       `json:"projectID"`
	TeamID      int       `json:"teamID"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Topics string identifiers for SSE event listeners.
type Topics string

// UpdateActivityRequest defines the model for UpdateActivityRequest.
type UpdateActivityRequest struct {
	Description  *string `json:"description,omitempty"`
	IsProductive *bool   `json:"isProductive,omitempty"`
	Name         *string `json:"name,omitempty"`
}

// UpdateTeamRequest defines the model for UpdateTeamRequest.
type UpdateTeamRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// UpdateUserAuthRequest represents User authorization data to update
type UpdateUserAuthRequest struct {
	// Role is generated from roles.json keys.
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

// UpdateWorkItemTagRequest defines the model for UpdateWorkItemTagRequest.
type UpdateWorkItemTagRequest struct {
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// UpdateWorkItemTypeRequest defines the model for UpdateWorkItemTypeRequest.
type UpdateWorkItemTypeRequest struct {
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// User defines the model for User.
type User struct {
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

	// Role is generated from roles.json keys.
	Role     Role      `json:"role"`
	Scopes   Scopes    `json:"scopes"`
	Teams    *[]DbTeam `json:"teams"`
	UserID   DbUserID  `json:"userID"`
	Username string    `json:"username"`
}

// UuidUUID defines the model for UuidUUID.
type UuidUUID = uuid.UUID

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

// WorkItemRole is generated from database enum 'work_item_role'.
type WorkItemRole string

// WorkItemTag defines the model for WorkItemTag.
type WorkItemTag struct {
	Color         string     `json:"color"`
	DeletedAt     *time.Time `json:"deletedAt"`
	Description   string     `json:"description"`
	Name          string     `json:"name"`
	ProjectID     int        `json:"projectID"`
	WorkItemTagID int        `json:"workItemTagID"`
}

// WorkItemType defines the model for WorkItemType.
type WorkItemType struct {
	Color          string `json:"color"`
	Description    string `json:"description"`
	Name           string `json:"name"`
	ProjectID      int    `json:"projectID"`
	WorkItemTypeID int    `json:"workItemTypeID"`
}

// ProjectName is generated from projects table.
type ProjectName = Project

// SerialID defines the model for SerialID.
type SerialID = int

// UUID defines the model for UUID.
type UUID = uuid.UUID

// MyProviderLoginParams defines parameters for MyProviderLogin.
type MyProviderLoginParams struct {
	AuthRedirectUri string `form:"auth-redirect-uri" json:"auth-redirect-uri"`
}

// EventsParams defines parameters for Events.
type EventsParams struct {
	ProjectName Project `form:"projectName" json:"projectName"`
}

// GetPaginatedNotificationsParams defines parameters for GetPaginatedNotifications.
type GetPaginatedNotificationsParams struct {
	Limit     int       `form:"limit" json:"limit"`
	Direction Direction `form:"direction" json:"direction"`
	Cursor    string    `form:"cursor" json:"cursor"`
}

// GetProjectWorkitemsParams defines parameters for GetProjectWorkitems.
type GetProjectWorkitemsParams struct {
	Open    *bool `form:"open,omitempty" json:"open,omitempty"`
	Deleted *bool `form:"deleted,omitempty" json:"deleted,omitempty"`
}

// GetPaginatedUsersParams defines parameters for GetPaginatedUsers.
type GetPaginatedUsersParams struct {
	Limit     int       `form:"limit" json:"limit"`
	Direction Direction `form:"direction" json:"direction"`
	Cursor    string    `form:"cursor" json:"cursor"`
}

// UpdateActivityJSONRequestBody defines body for UpdateActivity for application/json ContentType.

type UpdateActivityJSONRequestBody = UpdateActivityRequest

// CreateActivityJSONRequestBody defines body for CreateActivity for application/json ContentType.

type CreateActivityJSONRequestBody = CreateActivityRequest

// UpdateProjectConfigJSONRequestBody defines body for UpdateProjectConfig for application/json ContentType.

type UpdateProjectConfigJSONRequestBody = ProjectConfig

// InitializeProjectJSONRequestBody defines body for InitializeProject for application/json ContentType.

type InitializeProjectJSONRequestBody = InitializeProjectRequest

// CreateTeamJSONRequestBody defines body for CreateTeam for application/json ContentType.

type CreateTeamJSONRequestBody = CreateTeamRequest

// CreateWorkItemTagJSONRequestBody defines body for CreateWorkItemTag for application/json ContentType.

type CreateWorkItemTagJSONRequestBody = CreateWorkItemTagRequest

// CreateWorkItemTypeJSONRequestBody defines body for CreateWorkItemType for application/json ContentType.

type CreateWorkItemTypeJSONRequestBody = CreateWorkItemTypeRequest

// UpdateTeamJSONRequestBody defines body for UpdateTeam for application/json ContentType.

type UpdateTeamJSONRequestBody = UpdateTeamRequest

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.

type UpdateUserJSONRequestBody = UpdateUserRequest

// UpdateUserAuthorizationJSONRequestBody defines body for UpdateUserAuthorization for application/json ContentType.

type UpdateUserAuthorizationJSONRequestBody = UpdateUserAuthRequest

// UpdateWorkItemTagJSONRequestBody defines body for UpdateWorkItemTag for application/json ContentType.

type UpdateWorkItemTagJSONRequestBody = UpdateWorkItemTagRequest

// UpdateWorkItemTypeJSONRequestBody defines body for UpdateWorkItemType for application/json ContentType.

type UpdateWorkItemTypeJSONRequestBody = UpdateWorkItemTypeRequest

// CreateWorkitemJSONRequestBody defines body for CreateWorkitem for application/json ContentType.

type CreateWorkitemJSONRequestBody = CreateWorkItemRequest

// AsCreateDemoWorkItemRequest returns the union data inside the CreateWorkItemRequest as a CreateDemoWorkItemRequest
func (t CreateWorkItemRequest) AsCreateDemoWorkItemRequest() (CreateDemoWorkItemRequest, error) {
	var body CreateDemoWorkItemRequest
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// AsCreateDemoTwoWorkItemRequest returns the union data inside the CreateWorkItemRequest as a CreateDemoTwoWorkItemRequest
func (t CreateWorkItemRequest) AsCreateDemoTwoWorkItemRequest() (CreateDemoTwoWorkItemRequest, error) {
	var body CreateDemoTwoWorkItemRequest
	err := json.Unmarshal(t.union, &body)
	return body, err
}

func (t CreateWorkItemRequest) Discriminator() (string, error) {
	var discriminator struct {
		Discriminator string `json:"projectName"`
	}
	err := json.Unmarshal(t.union, &discriminator)
	return discriminator.Discriminator, err
}

func (t CreateWorkItemRequest) ValueByDiscriminator() (interface{}, error) {
	discriminator, err := t.Discriminator()
	if err != nil {
		return nil, err
	}
	switch discriminator {
	case "demo":
		return t.AsCreateDemoWorkItemRequest()
	case "demo_two":
		return t.AsCreateDemoTwoWorkItemRequest()
	default:
		return nil, errors.New("unknown discriminator value: " + discriminator)
	}
}

func (t CreateWorkItemRequest) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CreateWorkItemRequest) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
