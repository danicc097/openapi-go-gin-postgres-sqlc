package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
)

// InitializeProject.
func (h *dummyStrictHandlers) InitializeProject(c *gin.Context, project models.Project) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// GetProjectBoard.
func (h *dummyStrictHandlers) GetProjectBoard(c *gin.Context, project models.Project) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// GetProjectWorkitems.
func (h *dummyStrictHandlers) GetProjectWorkitems(c *gin.Context, project models.Project, params models.GetProjectWorkitemsParams) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// GetProjectConfig.
func (h *dummyStrictHandlers) GetProjectConfig(c *gin.Context, project models.Project) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// UpdateProjectConfig.
func (h *dummyStrictHandlers) UpdateProjectConfig(c *gin.Context, project models.Project) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// GetProject.
func (h *dummyStrictHandlers) GetProject(c *gin.Context, project models.Project) {
	defer newOTelSpanWithUser(c).End()

	// TODO project service (includes project, team, board...)
	// role, ok := h.svc.authz.RoleByRank(user.RoleRank)
	// if !ok {
	// 	msg := fmt.Sprintf("role with rank %d not found", user.RoleRank)
	// 	renderErrorResponse(c, msg, errors.New(msg))

	// 	return
	// }

	// res := User{UserPublic: user.ToPublic(), Role: role.Name, Scopes: user.Scopes}

	// c.JSON(http.StatusOK, res)
}

func (h *dummyStrictHandlers) CreateWorkitemTag(c *gin.Context, project models.Project) {
	defer newOTelSpanWithUser(c).End()

	caller := getUserFromCtx(c)

	body := &CreateWorkItemTagRequest{}
	if shouldReturn := parseBody(c, body); shouldReturn {
		return
	}

	wit, err := h.svc.WorkItemTag.Create(c, h.pool, caller, &db.WorkItemTagCreateParams{
		ProjectID:   internal.ProjectIDByName[project],
		Name:        body.Name,
		Description: body.Description,
		Color:       body.Color,
	})
	if err != nil {
		renderErrorResponse(c, "Could not create workitem tag", err)

		return
	}

	c.JSON(http.StatusCreated, wit)
}
