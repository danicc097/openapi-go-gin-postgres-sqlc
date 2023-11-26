package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
)

// InitializeProject.
func (h *Handlers) InitializeProject(c *gin.Context, project models.Project) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// GetProjectBoard.
func (h *Handlers) GetProjectBoard(c *gin.Context, project models.Project) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// GetProjectWorkitems.
func (h *Handlers) GetProjectWorkitems(c *gin.Context, project models.Project, params models.GetProjectWorkitemsParams) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// GetProjectConfig.
func (h *Handlers) GetProjectConfig(c *gin.Context, project models.Project) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// UpdateProjectConfig.
func (h *Handlers) UpdateProjectConfig(c *gin.Context, project models.Project) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// GetProject.
func (h *Handlers) GetProject(c *gin.Context, project models.Project) {
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

func (h *Handlers) CreateWorkitemTag(c *gin.Context, project models.Project) {
	defer newOTelSpanWithUser(c).End()

	caller := getUserFromCtx(c)

	body := &WorkItemTagCreateRequest{}
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
