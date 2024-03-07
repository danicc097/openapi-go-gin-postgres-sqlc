package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type DemoWorkItem struct {
	logger *zap.SugaredLogger
	repos  *repos.Repos
	wiSvc  *WorkItem
}

type Member struct {
	Role   models.WorkItemRole `json:"role"   ref:"#/components/schemas/WorkItemRole" required:"true"`
	UserID db.UserID           `json:"userID" required:"true"`
}

type DemoWorkItemCreateParams struct {
	repos.DemoWorkItemCreateParams
	WorkItemCreateParams
}

// NewDemoWorkItem returns a new DemoWorkItem service.
func NewDemoWorkItem(logger *zap.SugaredLogger, repos *repos.Repos) *DemoWorkItem {
	wiSvc := NewWorkItem(logger, repos)

	return &DemoWorkItem{
		logger: logger,
		repos:  repos,
		wiSvc:  wiSvc,
	}
}

// ByID gets a work item by ID.
func (w *DemoWorkItem) ByID(ctx context.Context, d db.DBTX, id db.WorkItemID) (*db.WorkItem, error) {
	defer newOTelSpan().Build(ctx).End()

	wi, err := w.repos.DemoWorkItem.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.DemoWorkItem.ByID: %w", err)
	}

	return wi, nil
}

// Create creates a new work item.
func (w *DemoWorkItem) Create(ctx context.Context, d db.DBTX, caller CtxUser, params DemoWorkItemCreateParams) (*db.WorkItem, error) {
	defer newOTelSpan().Build(ctx).End()

	if err := w.wiSvc.validateCreateParams(d, caller, &params.Base); err != nil {
		return nil, err
	}

	switch internal.DemoKanbanStepsNameByID[params.Base.KanbanStepID] {
	case models.DemoKanbanStepsDisabled:
		// something
	}

	switch internal.DemoWorkItemTypesNameByID[params.Base.WorkItemTypeID] {
	case models.DemoWorkItemTypesType1:
		// something
	}

	demoWi, err := w.repos.DemoWorkItem.Create(ctx, d, params.DemoWorkItemCreateParams)
	if err != nil {
		return nil, fmt.Errorf("repos.DemoWorkItem.Create: %w", err)
	}

	err = w.wiSvc.AssignTags(ctx, d, demoWi.WorkItemID, params.TagIDs)
	if err != nil {
		return nil, internal.WrapErrorWithLocf(err, "", []string{"tagIDs"}, "could not assign tags")
	}

	err = w.wiSvc.AssignUsers(ctx, d, demoWi.WorkItemID, params.Members)
	if err != nil {
		return nil, internal.WrapErrorWithLocf(err, "", []string{"members"}, "could not assign members")
	}

	opts := append(w.wiSvc.getSharedDBOpts(), db.WithWorkItemJoin(db.WorkItemJoins{DemoWorkItem: true}))
	wi, err := w.repos.DemoWorkItem.ByID(ctx, d, demoWi.WorkItemID, opts...)
	if err != nil {
		return nil, fmt.Errorf("repos.DemoWorkItem.ByID: %w", err)
	}

	return wi, nil
}

// Update updates an existing work item.
func (w *DemoWorkItem) Update(ctx context.Context, d db.DBTX, caller CtxUser, id db.WorkItemID, params repos.DemoWorkItemUpdateParams) (*db.WorkItem, error) {
	defer newOTelSpan().Build(ctx).End()

	if err := w.wiSvc.validateUpdateParams(d, caller, params.Base); err != nil {
		return nil, err
	}

	wi, err := w.repos.DemoWorkItem.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("repos.DemoWorkItem.Update: %w", err)
	}

	return wi, nil
}

// repo has Update only, then service has Close() (Update with closed=True), Move() (Update with kanban step change), ...)
// params for dedicated workItem require workItemID (FK-as-PK)
// TBD if useful: ByTag, ByType (for closed workitem searches. open ones simply return everything and filter in client)

func (w *DemoWorkItem) ListDeleted(ctx context.Context, d db.DBTX, teamID db.TeamID) ([]db.WorkItem, error) {
	// WorkItemsByTeamID with deleted opt, orderby createdAt
	return []db.WorkItem{}, errors.New("not implemented")
}

func (w *DemoWorkItem) List(ctx context.Context, d db.DBTX, teamID int) ([]db.WorkItem, error) {
	// WorkItemsByTeamID with orderby createdAt
	return []db.WorkItem{}, errors.New("not implemented")
}
