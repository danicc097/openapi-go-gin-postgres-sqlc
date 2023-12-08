// Code generated by 'yaegi extract github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest'. DO NOT EDIT.

package symbols

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/rest"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AdminPing":                   reflect.ValueOf(rest.AdminPing),
		"Alphanumspace":               reflect.ValueOf(&rest.Alphanumspace).Elem(),
		"CreateTeam":                  reflect.ValueOf(rest.CreateTeam),
		"CreateWorkItemTag":           reflect.ValueOf(rest.CreateWorkItemTag),
		"CreateWorkItemType":          reflect.ValueOf(rest.CreateWorkItemType),
		"CreateWorkitem":              reflect.ValueOf(rest.CreateWorkitem),
		"CreateWorkitemComment":       reflect.ValueOf(rest.CreateWorkitemComment),
		"CustomSchemaErrorFunc":       reflect.ValueOf(rest.CustomSchemaErrorFunc),
		"DeleteTeam":                  reflect.ValueOf(rest.DeleteTeam),
		"DeleteUser":                  reflect.ValueOf(rest.DeleteUser),
		"DeleteWorkItemTag":           reflect.ValueOf(rest.DeleteWorkItemTag),
		"DeleteWorkItemType":          reflect.ValueOf(rest.DeleteWorkItemType),
		"DeleteWorkitem":              reflect.ValueOf(rest.DeleteWorkitem),
		"Events":                      reflect.ValueOf(rest.Events),
		"GetCurrentUser":              reflect.ValueOf(rest.GetCurrentUser),
		"GetPaginatedNotifications":   reflect.ValueOf(rest.GetPaginatedNotifications),
		"GetProject":                  reflect.ValueOf(rest.GetProject),
		"GetProjectBoard":             reflect.ValueOf(rest.GetProjectBoard),
		"GetProjectConfig":            reflect.ValueOf(rest.GetProjectConfig),
		"GetProjectWorkitems":         reflect.ValueOf(rest.GetProjectWorkitems),
		"GetTeam":                     reflect.ValueOf(rest.GetTeam),
		"GetWorkItem":                 reflect.ValueOf(rest.GetWorkItem),
		"GetWorkItemTag":              reflect.ValueOf(rest.GetWorkItemTag),
		"GetWorkItemType":             reflect.ValueOf(rest.GetWorkItemType),
		"InitializeProject":           reflect.ValueOf(rest.InitializeProject),
		"MyProviderCallback":          reflect.ValueOf(rest.MyProviderCallback),
		"MyProviderLogin":             reflect.ValueOf(rest.MyProviderLogin),
		"NewHandlers":                 reflect.ValueOf(rest.NewHandlers),
		"NewPubSub":                   reflect.ValueOf(rest.NewPubSub),
		"NewServer":                   reflect.ValueOf(rest.NewServer),
		"OpenapiYamlGet":              reflect.ValueOf(rest.OpenapiYamlGet),
		"OtelName":                    reflect.ValueOf(constant.MakeFromLiteral("\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest\"", token.STRING, 0)),
		"Ping":                        reflect.ValueOf(rest.Ping),
		"ReadOpenAPI":                 reflect.ValueOf(rest.ReadOpenAPI),
		"RegisterHandlers":            reflect.ValueOf(rest.RegisterHandlers),
		"RegisterHandlersWithOptions": reflect.ValueOf(rest.RegisterHandlersWithOptions),
		"Run":                         reflect.ValueOf(rest.Run),
		"SSEHeadersMiddleware":        reflect.ValueOf(rest.SSEHeadersMiddleware),
		"UpdateProjectConfig":         reflect.ValueOf(rest.UpdateProjectConfig),
		"UpdateTeam":                  reflect.ValueOf(rest.UpdateTeam),
		"UpdateUser":                  reflect.ValueOf(rest.UpdateUser),
		"UpdateUserAuthorization":     reflect.ValueOf(rest.UpdateUserAuthorization),
		"UpdateWorkItemTag":           reflect.ValueOf(rest.UpdateWorkItemTag),
		"UpdateWorkItemType":          reflect.ValueOf(rest.UpdateWorkItemType),
		"UpdateWorkitem":              reflect.ValueOf(rest.UpdateWorkitem),
		"ValidateRequestFromContext":  reflect.ValueOf(rest.ValidateRequestFromContext),
		"ValidationErrorSeparator":    reflect.ValueOf(constant.MakeFromLiteral("\"$$$$\"", token.STRING, 0)),
		"WithMiddlewares":             reflect.ValueOf(rest.WithMiddlewares),

		// type definitions
		"AuthRestriction":              reflect.ValueOf((*rest.AuthRestriction)(nil)),
		"ClientChan":                   reflect.ValueOf((*rest.ClientChan)(nil)),
		"Config":                       reflect.ValueOf((*rest.Config)(nil)),
		"CreateDemoTwoWorkItemRequest": reflect.ValueOf((*rest.CreateDemoTwoWorkItemRequest)(nil)),
		"CreateDemoWorkItemRequest":    reflect.ValueOf((*rest.CreateDemoWorkItemRequest)(nil)),
		"CreateProjectBoardRequest":    reflect.ValueOf((*rest.CreateProjectBoardRequest)(nil)),
		"CreateTeamRequest":            reflect.ValueOf((*rest.CreateTeamRequest)(nil)),
		"CreateWorkItemCommentRequest": reflect.ValueOf((*rest.CreateWorkItemCommentRequest)(nil)),
		"CreateWorkItemTagRequest":     reflect.ValueOf((*rest.CreateWorkItemTagRequest)(nil)),
		"CreateWorkItemTypeRequest":    reflect.ValueOf((*rest.CreateWorkItemTypeRequest)(nil)),
		"DemoTwoWorkItems":             reflect.ValueOf((*rest.DemoTwoWorkItems)(nil)),
		"DemoWorkItems":                reflect.ValueOf((*rest.DemoWorkItems)(nil)),
		"ErrorHandler":                 reflect.ValueOf((*rest.ErrorHandler)(nil)),
		"Event":                        reflect.ValueOf((*rest.Event)(nil)),
		"GinServerOptions":             reflect.ValueOf((*rest.GinServerOptions)(nil)),
		"Handlers":                     reflect.ValueOf((*rest.Handlers)(nil)),
		"MiddlewareFunc":               reflect.ValueOf((*rest.MiddlewareFunc)(nil)),
		"MultiErrorHandler":            reflect.ValueOf((*rest.MultiErrorHandler)(nil)),
		"Notification":                 reflect.ValueOf((*rest.Notification)(nil)),
		"OAValidatorOptions":           reflect.ValueOf((*rest.OAValidatorOptions)(nil)),
		"OperationID":                  reflect.ValueOf((*rest.OperationID)(nil)),
		"PaginationPage":               reflect.ValueOf((*rest.PaginationPage)(nil)),
		"ProjectBoard":                 reflect.ValueOf((*rest.ProjectBoard)(nil)),
		"ProjectName":                  reflect.ValueOf((*rest.ProjectName)(nil)),
		"PubSub":                       reflect.ValueOf((*rest.PubSub)(nil)),
		"Server":                       reflect.ValueOf((*rest.Server)(nil)),
		"ServerInterface":              reflect.ValueOf((*rest.ServerInterface)(nil)),
		"ServerInterfaceWrapper":       reflect.ValueOf((*rest.ServerInterfaceWrapper)(nil)),
		"ServerOption":                 reflect.ValueOf((*rest.ServerOption)(nil)),
		"SharedWorkItemFields":         reflect.ValueOf((*rest.SharedWorkItemFields)(nil)),
		"Team":                         reflect.ValueOf((*rest.Team)(nil)),
		"UpdateTeamRequest":            reflect.ValueOf((*rest.UpdateTeamRequest)(nil)),
		"UpdateWorkItemTagRequest":     reflect.ValueOf((*rest.UpdateWorkItemTagRequest)(nil)),
		"UpdateWorkItemTypeRequest":    reflect.ValueOf((*rest.UpdateWorkItemTypeRequest)(nil)),
		"User":                         reflect.ValueOf((*rest.User)(nil)),
		"UserNotificationsChan":        reflect.ValueOf((*rest.UserNotificationsChan)(nil)),
		"WorkItemTag":                  reflect.ValueOf((*rest.WorkItemTag)(nil)),
		"WorkItemType":                 reflect.ValueOf((*rest.WorkItemType)(nil)),

		// interface wrapper definitions
		"_ServerInterface": reflect.ValueOf((*_github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface)(nil)),
	}
}

// _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface is an interface wrapper for ServerInterface type
type _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface struct {
	IValue                     interface{}
	WAdminPing                 func(c *gin.Context)
	WCreateTeam                func(c *gin.Context, projectName models.Project)
	WCreateWorkItemTag         func(c *gin.Context, projectName models.Project)
	WCreateWorkItemType        func(c *gin.Context, projectName models.Project)
	WCreateWorkitem            func(c *gin.Context)
	WCreateWorkitemComment     func(c *gin.Context, id int)
	WDeleteTeam                func(c *gin.Context, projectName models.Project, id int)
	WDeleteUser                func(c *gin.Context, id uuid.UUID)
	WDeleteWorkItemTag         func(c *gin.Context, projectName models.Project, id int)
	WDeleteWorkItemType        func(c *gin.Context, projectName models.Project, id int)
	WDeleteWorkitem            func(c *gin.Context, id int)
	WEvents                    func(c *gin.Context, params models.EventsParams)
	WGetCurrentUser            func(c *gin.Context)
	WGetPaginatedNotifications func(c *gin.Context, params models.GetPaginatedNotificationsParams)
	WGetProject                func(c *gin.Context, projectName models.Project)
	WGetProjectBoard           func(c *gin.Context, projectName models.Project)
	WGetProjectConfig          func(c *gin.Context, projectName models.Project)
	WGetProjectWorkitems       func(c *gin.Context, projectName models.Project, params models.GetProjectWorkitemsParams)
	WGetTeam                   func(c *gin.Context, projectName models.Project, id int)
	WGetWorkItem               func(c *gin.Context, id int)
	WGetWorkItemTag            func(c *gin.Context, projectName models.Project, id int)
	WGetWorkItemType           func(c *gin.Context, projectName models.Project, id int)
	WInitializeProject         func(c *gin.Context, projectName models.Project)
	WMyProviderCallback        func(c *gin.Context)
	WMyProviderLogin           func(c *gin.Context)
	WOpenapiYamlGet            func(c *gin.Context)
	WPing                      func(c *gin.Context)
	WUpdateProjectConfig       func(c *gin.Context, projectName models.Project)
	WUpdateTeam                func(c *gin.Context, projectName models.Project, id int)
	WUpdateUser                func(c *gin.Context, id uuid.UUID)
	WUpdateUserAuthorization   func(c *gin.Context, id uuid.UUID)
	WUpdateWorkItemTag         func(c *gin.Context, projectName models.Project, id int)
	WUpdateWorkItemType        func(c *gin.Context, projectName models.Project, id int)
	WUpdateWorkitem            func(c *gin.Context, id int)
}

func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) AdminPing(c *gin.Context) {
	W.WAdminPing(c)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) CreateTeam(c *gin.Context, projectName models.Project) {
	W.WCreateTeam(c, projectName)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) CreateWorkItemTag(c *gin.Context, projectName models.Project) {
	W.WCreateWorkItemTag(c, projectName)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) CreateWorkItemType(c *gin.Context, projectName models.Project) {
	W.WCreateWorkItemType(c, projectName)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) CreateWorkitem(c *gin.Context) {
	W.WCreateWorkitem(c)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) CreateWorkitemComment(c *gin.Context, id int) {
	W.WCreateWorkitemComment(c, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) DeleteTeam(c *gin.Context, projectName models.Project, id int) {
	W.WDeleteTeam(c, projectName, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) DeleteUser(c *gin.Context, id uuid.UUID) {
	W.WDeleteUser(c, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) DeleteWorkItemTag(c *gin.Context, projectName models.Project, id int) {
	W.WDeleteWorkItemTag(c, projectName, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) DeleteWorkItemType(c *gin.Context, projectName models.Project, id int) {
	W.WDeleteWorkItemType(c, projectName, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) DeleteWorkitem(c *gin.Context, id int) {
	W.WDeleteWorkitem(c, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) Events(c *gin.Context, params models.EventsParams) {
	W.WEvents(c, params)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) GetCurrentUser(c *gin.Context) {
	W.WGetCurrentUser(c)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) GetPaginatedNotifications(c *gin.Context, params models.GetPaginatedNotificationsParams) {
	W.WGetPaginatedNotifications(c, params)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) GetProject(c *gin.Context, projectName models.Project) {
	W.WGetProject(c, projectName)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) GetProjectBoard(c *gin.Context, projectName models.Project) {
	W.WGetProjectBoard(c, projectName)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) GetProjectConfig(c *gin.Context, projectName models.Project) {
	W.WGetProjectConfig(c, projectName)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) GetProjectWorkitems(c *gin.Context, projectName models.Project, params models.GetProjectWorkitemsParams) {
	W.WGetProjectWorkitems(c, projectName, params)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) GetTeam(c *gin.Context, projectName models.Project, id int) {
	W.WGetTeam(c, projectName, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) GetWorkItem(c *gin.Context, id int) {
	W.WGetWorkItem(c, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) GetWorkItemTag(c *gin.Context, projectName models.Project, id int) {
	W.WGetWorkItemTag(c, projectName, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) GetWorkItemType(c *gin.Context, projectName models.Project, id int) {
	W.WGetWorkItemType(c, projectName, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) InitializeProject(c *gin.Context, projectName models.Project) {
	W.WInitializeProject(c, projectName)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) MyProviderCallback(c *gin.Context) {
	W.WMyProviderCallback(c)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) MyProviderLogin(c *gin.Context) {
	W.WMyProviderLogin(c)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) OpenapiYamlGet(c *gin.Context) {
	W.WOpenapiYamlGet(c)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) Ping(c *gin.Context) {
	W.WPing(c)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) UpdateProjectConfig(c *gin.Context, projectName models.Project) {
	W.WUpdateProjectConfig(c, projectName)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) UpdateTeam(c *gin.Context, projectName models.Project, id int) {
	W.WUpdateTeam(c, projectName, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) UpdateUser(c *gin.Context, id uuid.UUID) {
	W.WUpdateUser(c, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) UpdateUserAuthorization(c *gin.Context, id uuid.UUID) {
	W.WUpdateUserAuthorization(c, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) UpdateWorkItemTag(c *gin.Context, projectName models.Project, id int) {
	W.WUpdateWorkItemTag(c, projectName, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) UpdateWorkItemType(c *gin.Context, projectName models.Project, id int) {
	W.WUpdateWorkItemType(c, projectName, id)
}
func (W _github_com_danicc097_openapi_go_gin_postgres_sqlc_internal_rest_ServerInterface) UpdateWorkitem(c *gin.Context, id int) {
	W.WUpdateWorkitem(c, id)
}
