package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

// InitializeProject
func (h *Handlers) InitializeProject(c *gin.Context, project models.Project) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// GetProjectBoard
func (h *Handlers) GetProjectBoard(c *gin.Context, project models.Project) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// GetProjectWorkitems
func (h *Handlers) GetProjectWorkitems(c *gin.Context, project models.Project, params models.GetProjectWorkitemsParams) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// GetProjectConfig
func (h *Handlers) GetProjectConfig(c *gin.Context, project models.Project) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// UpdateProjectConfig
func (h *Handlers) UpdateProjectConfig(c *gin.Context, project models.Project) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// GetProject
func (h *Handlers) GetProject(c *gin.Context, project models.Project) {
	ctx := c.Request.Context()

	defer newOTELSpan(ctx, "GetProject", trace.WithAttributes(userIDAttribute(c))).End()

	// TODO project service (includes project, team, board...)
	// role, ok := h.authzsvc.RoleByRank(user.RoleRank)
	// if !ok {
	// 	msg := fmt.Sprintf("role with rank %d not found", user.RoleRank)
	// 	renderErrorResponse(c, msg, errors.New(msg))

	// 	return
	// }

	// res := User{UserPublic: user.ToPublic(), Role: role.Name, Scopes: user.Scopes}

	// c.JSON(http.StatusOK, res)
}

// create workitem tag
func (h *Handlers) CreateWorkitemTag(c *gin.Context, project models.Project) {
	ctx := c.Request.Context()

	defer newOTELSpan(ctx, "CreateWorkitemTag", trace.WithAttributes(userIDAttribute(c))).End()

	user := getUserFromCtx(c)

	h.workitemtagsvc.Create(c, h.pool, user, &db.WorkItemTagCreateParams{
		ProjectID: internal.ProjectIDByName[project],
		// TODO params + oapi path parameters override name
	})
	c.JSON(http.StatusNotImplemented, "not implemented")
}
