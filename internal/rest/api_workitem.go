package rest

import (
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/slices"
	"github.com/gin-gonic/gin"
)

// create workitem.
func (h *Handlers) CreateWorkitem(c *gin.Context) {
	ctx := c.Request.Context()

	defer newOTELSpanWithUser(c, "CreateWorkitem").End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	body := &models.WorkItemCreateRequest{}
	if shouldReturn := parseBody(c, body); shouldReturn {
		return
	}
	var err error
	var workItem *db.WorkItem

	switch disc, _ := body.Discriminator(); models.Project(disc) {
	case models.ProjectDemo:
		params, _ := body.AsDemoWorkItemCreateRequest()

		workItem, err = h.demoworkitemsvc.Create(ctx, tx, services.DemoWorkItemCreateParams{
			DemoWorkItemCreateParams: repos.DemoWorkItemCreateParams{
				DemoProject: db.DemoWorkItemCreateParams(params.DemoProject),
				Base:        db.WorkItemCreateParams(params.Base),
			},
			TagIDs: params.TagIDs,
			Members: slices.Map(params.Members, func(item models.ServicesMember, _ int) services.Member {
				return services.Member(item)
			}),
		})
		if err != nil {
			renderErrorResponse(c, "Could not create work item", err)

			return
		}
	case models.ProjectDemoTwo:
		params, _ := body.AsDemoTwoWorkItemCreateRequest()

		workItem, err = h.demotwoworkitemsvc.Create(ctx, tx, services.DemoTwoWorkItemCreateParams{
			DemoTwoWorkItemCreateParams: repos.DemoTwoWorkItemCreateParams{
				DemoTwoProject: db.DemoTwoWorkItemCreateParams(params.DemoTwoProject),
				Base:           db.WorkItemCreateParams(params.Base),
			},
			TagIDs: params.TagIDs,
			Members: slices.Map(params.Members, func(item models.ServicesMember, _ int) services.Member {
				return services.Member(item)
			}),
		})
		if err != nil {
			renderErrorResponse(c, "Could not create work item", err)

			return
		}
	default:
		renderErrorResponse(c, fmt.Sprintf("Unknown project %q", disc), nil)

		return
	}

	c.JSON(http.StatusCreated, workItem)
}

// delete workitem.
func (h *Handlers) DeleteWorkitem(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTELSpanWithUser(c, "DeleteWorkitem").End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	c.JSON(http.StatusNotImplemented, "not implemented")
}

// get workitem.
func (h *Handlers) GetWorkitem(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTELSpanWithUser(c, "GetWorkitem").End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	c.JSON(http.StatusNotImplemented, "not implemented")
}

// update workitem.
func (h *Handlers) UpdateWorkitem(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTELSpanWithUser(c, "UpdateWorkitem").End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	c.JSON(http.StatusNotImplemented, "not implemented")
}

// create workitem comment.
func (h *Handlers) CreateWorkitemComment(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTELSpanWithUser(c, "CreateWorkitemComment").End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	c.JSON(http.StatusNotImplemented, "not implemented")
}
