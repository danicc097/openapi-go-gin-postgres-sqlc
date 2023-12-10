package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) CreateWorkItemTag(c *gin.Context, projectName models.ProjectName) {
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
