// Package client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/danicc097/openapi-go-gin-postgres-sqlc version (devel) DO NOT EDIT.
package client

const (
	Api_keyScopes     = "api_key.Scopes"
	Bearer_authScopes = "bearer_auth.Scopes"
)

// Defines values for Role.
const (
	RoleAdmin        Role = "admin"
	RoleAdvancedUser Role = "advanced user"
	RoleGuest        Role = "guest"
	RoleManager      Role = "manager"
	RoleSuperadmin   Role = "superadmin"
	RoleUser         Role = "user"
)

// Defines values for Scope.
const (
	ScopeSettingsWrite  Scope = "settings:write"
	ScopeUsersRead      Scope = "users:read"
	ScopeUsersWrite     Scope = "users:write"
	ScopeWorkItemReview Scope = "work-item:review"
)

// GetCurrentUserRes represents a user
type GetCurrentUserRes struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`

	// Orgs are organizations a user belongs to
	Orgs     *interface{} `json:"orgs,omitempty"`
	Password *string      `json:"password,omitempty"`
	Phone    *string      `json:"phone,omitempty"`

	// Role User role.
	Role     *Role   `json:"role,omitempty"`
	UserId   *int64  `json:"user_id,omitempty"`
	Username *string `json:"username,omitempty"`
}

// HTTPValidationError defines model for HTTPValidationError.
type HTTPValidationError struct {
	Detail *[]ValidationError `json:"detail,omitempty"`
}

// Organization Organization a user belongs to.
type Organization = string

// Role User role.
type Role string

// Scope defines model for Scope.
type Scope string

// UpdateUserRequest represents User data to update
type UpdateUserRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`

	// Role User role.
	Role *Role `json:"role,omitempty"`
}

// ValidationError defines model for ValidationError.
type ValidationError struct {
	Loc  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
}

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = UpdateUserRequest
