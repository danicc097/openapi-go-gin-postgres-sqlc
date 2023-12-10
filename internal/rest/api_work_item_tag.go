package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
)

func (h *dummyStrictHandlers) CreateWorkItemTag(c *gin.Context, projectName models.ProjectName) {
	tx := getTxFromCtx(c)
	u := getUserFromCtx(c)

	body := &models.CreateWorkItemTagJSONRequestBody{}
	if shouldReturn := parseBody(c, body); shouldReturn {
		return
	}

	wit, err := h.svc.WorkItemTag.Create(c, tx, u, &db.WorkItemTagCreateParams{
		Color:       body.Color,
		Description: body.Description,
		Name:        body.Name,
		ProjectID:   internal.ProjectIDByName[projectName],
	})
	if err != nil {
		renderErrorResponse(c, "Could not create work item tag", err)

		return
	}

	renderResponse(c, wit, http.StatusCreated)
}

func (h *dummyStrictHandlers) UpdateWorkItemTag(c *gin.Context, projectName models.ProjectName, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *dummyStrictHandlers) GetWorkItemTag(c *gin.Context, projectName models.ProjectName, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *dummyStrictHandlers) DeleteWorkItemTag(c *gin.Context, projectName models.ProjectName, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *StrictHandlers) UpdateWorkItemTag(c *gin.Context, request UpdateWorkItemTagRequestObject) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *StrictHandlers) DeleteWorkItemTag(c *gin.Context, request DeleteWorkItemTagRequestObject) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *StrictHandlers) CreateWorkItemTag(c *gin.Context, request CreateWorkItemTagRequestObject) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *StrictHandlers) GetWorkItemTag(c *gin.Context, request GetWorkItemTagRequestObject) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}
