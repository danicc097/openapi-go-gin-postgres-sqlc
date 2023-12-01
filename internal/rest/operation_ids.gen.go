// Code generated by pregen. DO NOT EDIT.

package rest

type OperationID string

const (
	// Operation IDs for the 'admin' tag.
	AdminPing OperationID = "AdminPing"

	// Operation IDs for the 'default' tag.
	OpenapiYamlGet OperationID = "OpenapiYamlGet"
	Ping           OperationID = "Ping"

	// Operation IDs for the 'events' tag.
	Events OperationID = "Events"

	// Operation IDs for the 'notifications' tag.
	GetPaginatedNotifications OperationID = "GetPaginatedNotifications"

	// Operation IDs for the 'oidc' tag.
	MyProviderCallback OperationID = "MyProviderCallback"
	MyProviderLogin    OperationID = "MyProviderLogin"

	// Operation IDs for the 'project' tag.
	CreateWorkitemTag   OperationID = "CreateWorkitemTag"
	GetProject          OperationID = "GetProject"
	GetProjectBoard     OperationID = "GetProjectBoard"
	GetProjectConfig    OperationID = "GetProjectConfig"
	GetProjectWorkitems OperationID = "GetProjectWorkitems"
	InitializeProject   OperationID = "InitializeProject"
	UpdateProjectConfig OperationID = "UpdateProjectConfig"

	// Operation IDs for the 'user' tag.
	DeleteUser              OperationID = "DeleteUser"
	GetCurrentUser          OperationID = "GetCurrentUser"
	UpdateUser              OperationID = "UpdateUser"
	UpdateUserAuthorization OperationID = "UpdateUserAuthorization"

	// Operation IDs for the 'workitem' tag.
	CreateWorkitem        OperationID = "CreateWorkitem"
	CreateWorkitemComment OperationID = "CreateWorkitemComment"
	DeleteWorkitem        OperationID = "DeleteWorkitem"
	GetWorkitem           OperationID = "GetWorkitem"
	UpdateWorkitem        OperationID = "UpdateWorkitem"
)
