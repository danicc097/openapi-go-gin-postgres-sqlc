package rest

import (
	"encoding/json"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) UpdateWorkitem(c *gin.Context, request UpdateWorkitemRequestObject) (UpdateWorkitemResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) DeleteWorkitem(c *gin.Context, request DeleteWorkitemRequestObject) (DeleteWorkitemResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) CreateWorkitem(c *gin.Context, request CreateWorkitemRequestObject) (CreateWorkitemResponseObject, error) {
	ctx := c.Request.Context()

	caller, _ := GetUserCallerFromCtx(c)
	tx := GetTxFromCtx(c)

	addRequestBodyToSpan(c)

	var res any // depends on project

	project, b := projectAndBodyByDiscriminator(c, request.Body)

	//exhaustive:enforce
	switch project {
	case models.ProjectNameDemo:
		body, _ := b.(models.CreateDemoWorkItemRequest)
		workItem, err := h.svc.DemoWorkItem.Create(ctx, tx, caller, services.DemoWorkItemCreateParams{
			DemoWorkItemCreateParams: repos.DemoWorkItemCreateParams{
				DemoProject: body.DemoProject,
				Base:        body.Base,
			},
			WorkItemCreateParams: services.WorkItemCreateParams{
				TagIDs:  body.TagIDs,
				Members: restMembersToServices(body.Members),
			},
		})
		if err != nil {
			renderErrorResponse(c, "Could not create work item", err)
		}

		res = DemoWorkItemResponse{
			WorkItemBase: fillBaseWorkItemResponse(workItem),
			DemoWorkItem: *workItem.DemoWorkItemJoin,
		}
	case models.ProjectNameDemoTwo:
		body, _ := b.(models.CreateDemoTwoWorkItemRequest)
		workItem, err := h.svc.DemoTwoWorkItem.Create(ctx, tx, caller, services.DemoTwoWorkItemCreateParams{
			DemoTwoWorkItemCreateParams: repos.DemoTwoWorkItemCreateParams{
				DemoTwoProject: body.DemoTwoProject,
				Base:           body.Base,
			},
			WorkItemCreateParams: services.WorkItemCreateParams{
				TagIDs:  body.TagIDs,
				Members: restMembersToServices(body.Members),
			},
		})
		if err != nil {
			renderErrorResponse(c, "Could not create work item", err)
		}

		res = DemoTwoWorkItemResponse{
			WorkItemBase:    fillBaseWorkItemResponse(workItem),
			DemoTwoWorkItem: *workItem.DemoTwoWorkItemJoin,
		}
	default:
		renderErrorResponse(c, "Unknown discriminator", internal.NewErrorf(models.ErrorCodeUnknown, "%+v", b))
	}

	var resJson *CreateWorkitem201JSONResponse
	json.Unmarshal(rawMessage(res), &resJson)
	return resJson, nil
}

func (h *StrictHandlers) GetWorkItem(c *gin.Context, request GetWorkItemRequestObject) (GetWorkItemResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func fillBaseWorkItemResponse(workItem *models.WorkItem) WorkItemBase {
	return WorkItemBase{
		WorkItem: *workItem,
		SharedWorkItemJoins: SharedWorkItemJoins{
			Members:          workItem.AssigneesJoin,
			WorkItemTags:     workItem.WorkItemTagsJoin,
			TimeEntries:      workItem.TimeEntriesJoin,
			WorkItemComments: workItem.WorkItemCommentsJoin,
			WorkItemType:     workItem.WorkItemTypeJoin,
		},
	}
}

func (h *StrictHandlers) GetPaginatedWorkItem(c *gin.Context, request GetPaginatedWorkItemRequestObject) (GetPaginatedWorkItemResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}
