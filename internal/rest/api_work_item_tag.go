package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) CreateWorkItemTag(c *gin.Context, request CreateWorkItemTagRequestObject) (CreateWorkItemTagResponseObject, error) {
	tx := GetTxFromCtx(c)
	u := getUserFromCtx(c)

	body := &models.CreateWorkItemTagJSONRequestBody{}
	if shouldReturn := parseBody(c, body); shouldReturn {
		return nil, nil
	}

	wit, err := h.svc.WorkItemTag.Create(c, tx, u, &db.WorkItemTagCreateParams{
		Color:       body.Color,
		Description: body.Description,
		Name:        body.Name,
		ProjectID:   internal.ProjectIDByName[request.ProjectName],
	})
	if err != nil {
		renderErrorResponse(c, "Could not create work item tag", err)

		return nil, nil

	}

	renderResponse(c, wit, http.StatusCreated)
	return nil, nil
}

func (h *StrictHandlers) GetWorkItemTag(c *gin.Context, request GetWorkItemTagRequestObject) (GetWorkItemTagResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) DeleteWorkItemTag(c *gin.Context, request DeleteWorkItemTagRequestObject) (DeleteWorkItemTagResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) UpdateWorkItemTag(c *gin.Context, request UpdateWorkItemTagRequestObject) (UpdateWorkItemTagResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}
