// Package client provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	Api_keyScopes     = "api_key.Scopes"
	Bearer_authScopes = "bearer_auth.Scopes"
)

// Defines values for NotificationType.
const (
	Global   NotificationType = "global"
	Personal NotificationType = "personal"
)

// Defines values for Role.
const (
	Admin        Role = "admin"
	AdvancedUser Role = "advancedUser"
	Guest        Role = "guest"
	Manager      Role = "manager"
	SuperAdmin   Role = "superAdmin"
	User         Role = "user"
)

// Defines values for Scope.
const (
	ProjectSettingsWrite Scope = "project-settings:write"
	ScopesWrite          Scope = "scopes:write"
	TeamSettingsWrite    Scope = "team-settings:write"
	TestScope            Scope = "test-scope"
	UsersRead            Scope = "users:read"
	UsersWrite           Scope = "users:write"
	WorkItemReview       Scope = "work-item:review"
)

// Defines values for WorkItemRole.
const (
	Preparer WorkItemRole = "preparer"
	Reviewer WorkItemRole = "reviewer"
)

// HTTPValidationError defines model for HTTPValidationError.
type HTTPValidationError struct {
	Detail *[]ValidationError `json:"detail,omitempty"`
}

// ModelsRole defines model for ModelsRole.
type ModelsRole = string

// ModelsScope defines model for ModelsScope.
type ModelsScope = string

// NotificationType User notification type.
type NotificationType string

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
	ApiKey                   *UserAPIKeyPublic `json:"apiKey"`
	CreatedAt                time.Time         `json:"createdAt"`
	DeletedAt                *time.Time        `json:"deletedAt"`
	Email                    string            `json:"email"`
	FirstName                *string           `json:"firstName"`
	FullName                 *string           `json:"fullName"`
	HasGlobalNotifications   bool              `json:"hasGlobalNotifications"`
	HasPersonalNotifications bool              `json:"hasPersonalNotifications"`
	LastName                 *string           `json:"lastName"`
	Role                     Role              `json:"role"`
	Scopes                   Scopes            `json:"scopes"`
	Teams                    *[]TeamPublic     `json:"teams"`
	UserID                   UuidUUID          `json:"userID"`
	Username                 string            `json:"username"`
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

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// AdminPing request
	AdminPing(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Events request
	Events(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpenapiYamlGet request
	OpenapiYamlGet(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Ping request
	Ping(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetCurrentUser request
	GetCurrentUser(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteUser request
	DeleteUser(ctx context.Context, id UserID, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateUser request with any body
	UpdateUserWithBody(ctx context.Context, id UserID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateUser(ctx context.Context, id UserID, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateUserAuthorization request with any body
	UpdateUserAuthorizationWithBody(ctx context.Context, id UserID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateUserAuthorization(ctx context.Context, id UserID, body UpdateUserAuthorizationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) AdminPing(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAdminPingRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Events(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewEventsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) OpenapiYamlGet(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewOpenapiYamlGetRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Ping(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPingRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetCurrentUser(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetCurrentUserRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteUser(ctx context.Context, id UserID, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteUserRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateUserWithBody(ctx context.Context, id UserID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateUserRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateUser(ctx context.Context, id UserID, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateUserRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateUserAuthorizationWithBody(ctx context.Context, id UserID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateUserAuthorizationRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateUserAuthorization(ctx context.Context, id UserID, body UpdateUserAuthorizationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateUserAuthorizationRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewAdminPingRequest generates requests for AdminPing
func NewAdminPingRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/admin/ping")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewEventsRequest generates requests for Events
func NewEventsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/events")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewOpenapiYamlGetRequest generates requests for OpenapiYamlGet
func NewOpenapiYamlGetRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/openapi.yaml")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPingRequest generates requests for Ping
func NewPingRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/ping")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetCurrentUserRequest generates requests for GetCurrentUser
func NewGetCurrentUserRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user/me")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewDeleteUserRequest generates requests for DeleteUser
func NewDeleteUserRequest(server string, id UserID) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateUserRequest calls the generic UpdateUser builder with application/json body
func NewUpdateUserRequest(server string, id UserID, body UpdateUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateUserRequestWithBody(server, id, "application/json", bodyReader)
}

// NewUpdateUserRequestWithBody generates requests for UpdateUser with any type of body
func NewUpdateUserRequestWithBody(server string, id UserID, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewUpdateUserAuthorizationRequest calls the generic UpdateUserAuthorization builder with application/json body
func NewUpdateUserAuthorizationRequest(server string, id UserID, body UpdateUserAuthorizationJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateUserAuthorizationRequestWithBody(server, id, "application/json", bodyReader)
}

// NewUpdateUserAuthorizationRequestWithBody generates requests for UpdateUserAuthorization with any type of body
func NewUpdateUserAuthorizationRequestWithBody(server string, id UserID, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user/%s/authorization", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// AdminPing request
	AdminPingWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*AdminPingResponse, error)

	// Events request
	EventsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*EventsResponse, error)

	// OpenapiYamlGet request
	OpenapiYamlGetWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*OpenapiYamlGetResponse, error)

	// Ping request
	PingWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PingResponse, error)

	// GetCurrentUser request
	GetCurrentUserWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetCurrentUserResponse, error)

	// DeleteUser request
	DeleteUserWithResponse(ctx context.Context, id UserID, reqEditors ...RequestEditorFn) (*DeleteUserResponse, error)

	// UpdateUser request with any body
	UpdateUserWithBodyWithResponse(ctx context.Context, id UserID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error)

	UpdateUserWithResponse(ctx context.Context, id UserID, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error)

	// UpdateUserAuthorization request with any body
	UpdateUserAuthorizationWithBodyWithResponse(ctx context.Context, id UserID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateUserAuthorizationResponse, error)

	UpdateUserAuthorizationWithResponse(ctx context.Context, id UserID, body UpdateUserAuthorizationJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateUserAuthorizationResponse, error)
}

type AdminPingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r AdminPingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AdminPingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type EventsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r EventsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r EventsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type OpenapiYamlGetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	YAML200      *openapi_types.File
}

// Status returns HTTPResponse.Status
func (r OpenapiYamlGetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r OpenapiYamlGetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r PingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetCurrentUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UserResponse
}

// Status returns HTTPResponse.Status
func (r GetCurrentUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetCurrentUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UserPublic
}

// Status returns HTTPResponse.Status
func (r UpdateUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateUserAuthorizationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UserResponse
}

// Status returns HTTPResponse.Status
func (r UpdateUserAuthorizationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateUserAuthorizationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// AdminPingWithResponse request returning *AdminPingResponse
func (c *ClientWithResponses) AdminPingWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*AdminPingResponse, error) {
	rsp, err := c.AdminPing(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAdminPingResponse(rsp)
}

// EventsWithResponse request returning *EventsResponse
func (c *ClientWithResponses) EventsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*EventsResponse, error) {
	rsp, err := c.Events(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseEventsResponse(rsp)
}

// OpenapiYamlGetWithResponse request returning *OpenapiYamlGetResponse
func (c *ClientWithResponses) OpenapiYamlGetWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*OpenapiYamlGetResponse, error) {
	rsp, err := c.OpenapiYamlGet(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseOpenapiYamlGetResponse(rsp)
}

// PingWithResponse request returning *PingResponse
func (c *ClientWithResponses) PingWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PingResponse, error) {
	rsp, err := c.Ping(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePingResponse(rsp)
}

// GetCurrentUserWithResponse request returning *GetCurrentUserResponse
func (c *ClientWithResponses) GetCurrentUserWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetCurrentUserResponse, error) {
	rsp, err := c.GetCurrentUser(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetCurrentUserResponse(rsp)
}

// DeleteUserWithResponse request returning *DeleteUserResponse
func (c *ClientWithResponses) DeleteUserWithResponse(ctx context.Context, id UserID, reqEditors ...RequestEditorFn) (*DeleteUserResponse, error) {
	rsp, err := c.DeleteUser(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteUserResponse(rsp)
}

// UpdateUserWithBodyWithResponse request with arbitrary body returning *UpdateUserResponse
func (c *ClientWithResponses) UpdateUserWithBodyWithResponse(ctx context.Context, id UserID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error) {
	rsp, err := c.UpdateUserWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateUserResponse(rsp)
}

func (c *ClientWithResponses) UpdateUserWithResponse(ctx context.Context, id UserID, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error) {
	rsp, err := c.UpdateUser(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateUserResponse(rsp)
}

// UpdateUserAuthorizationWithBodyWithResponse request with arbitrary body returning *UpdateUserAuthorizationResponse
func (c *ClientWithResponses) UpdateUserAuthorizationWithBodyWithResponse(ctx context.Context, id UserID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateUserAuthorizationResponse, error) {
	rsp, err := c.UpdateUserAuthorizationWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateUserAuthorizationResponse(rsp)
}

func (c *ClientWithResponses) UpdateUserAuthorizationWithResponse(ctx context.Context, id UserID, body UpdateUserAuthorizationJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateUserAuthorizationResponse, error) {
	rsp, err := c.UpdateUserAuthorization(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateUserAuthorizationResponse(rsp)
}

// ParseAdminPingResponse parses an HTTP response from a AdminPingWithResponse call
func ParseAdminPingResponse(rsp *http.Response) (*AdminPingResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AdminPingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseEventsResponse parses an HTTP response from a EventsWithResponse call
func ParseEventsResponse(rsp *http.Response) (*EventsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &EventsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseOpenapiYamlGetResponse parses an HTTP response from a OpenapiYamlGetWithResponse call
func ParseOpenapiYamlGetResponse(rsp *http.Response) (*OpenapiYamlGetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &OpenapiYamlGetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "yaml") && rsp.StatusCode == 200:
		var dest openapi_types.File
		if err := yaml.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.YAML200 = &dest

	}

	return response, nil
}

// ParsePingResponse parses an HTTP response from a PingWithResponse call
func ParsePingResponse(rsp *http.Response) (*PingResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseGetCurrentUserResponse parses an HTTP response from a GetCurrentUserWithResponse call
func ParseGetCurrentUserResponse(rsp *http.Response) (*GetCurrentUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetCurrentUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UserResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseDeleteUserResponse parses an HTTP response from a DeleteUserWithResponse call
func ParseDeleteUserResponse(rsp *http.Response) (*DeleteUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseUpdateUserResponse parses an HTTP response from a UpdateUserWithResponse call
func ParseUpdateUserResponse(rsp *http.Response) (*UpdateUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UserPublic
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseUpdateUserAuthorizationResponse parses an HTTP response from a UpdateUserAuthorizationWithResponse call
func ParseUpdateUserAuthorizationResponse(rsp *http.Response) (*UpdateUserAuthorizationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateUserAuthorizationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UserResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
