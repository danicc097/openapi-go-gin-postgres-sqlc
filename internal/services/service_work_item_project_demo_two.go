package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"go.uber.org/zap"
)

type DemoTwoWorkItem struct {
	logger *zap.SugaredLogger
	repos  *repos.Repos
	wiSvc  *WorkItem
}

type DemoTwoWorkItemCreateParams struct {
	repos.DemoTwoWorkItemCreateParams
	WorkItemCreateParams
}

// NewDemoTwoWorkItem returns a new DemoTwoWorkItem service.
func NewDemoTwoWorkItem(logger *zap.SugaredLogger, repos *repos.Repos) *DemoTwoWorkItem {
	wiSvc := NewWorkItem(logger, repos)

	return &DemoTwoWorkItem{
		logger: logger,
		repos:  repos,
		wiSvc:  wiSvc,
	}
}

// ByID gets a work item by ID.
func (w *DemoTwoWorkItem) ByID(ctx context.Context, d models.DBTX, id models.WorkItemID) (*models.WorkItem, error) {
	defer newOTelSpan().Build(ctx).End()

	wi, err := w.repos.DemoTwoWorkItem.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.DemoTwoWorkItem.ByID: %w", err)
	}

	return wi, nil
}

// Create creates a new work item.
func (w *DemoTwoWorkItem) Create(ctx context.Context, d models.DBTX, caller CtxUser, params DemoTwoWorkItemCreateParams) (*models.WorkItem, error) {
	defer newOTelSpan().Build(ctx).End()

	if err := w.wiSvc.validateCreateParams(d, caller, &params.Base); err != nil {
		return nil, err
	}

	demoTwoWi, err := w.repos.DemoTwoWorkItem.Create(ctx, d, params.DemoTwoWorkItemCreateParams)
	if err != nil {
		return nil, fmt.Errorf("repos.DemoTwoWorkItem.Create: %w", err)
	}

	if err := w.wiSvc.postCreate(ctx, d, demoTwoWi.WorkItemID, params.WorkItemCreateParams); err != nil {
		return nil, err
	}

	opts := append(w.wiSvc.getSharedDBOpts(), models.WithWorkItemJoin(models.WorkItemJoins{DemoTwoWorkItem: true}))
	wi, err := w.repos.DemoTwoWorkItem.ByID(ctx, d, demoTwoWi.WorkItemID, opts...)
	if err != nil {
		return nil, fmt.Errorf("repos.DemoTwoWorkItem.ByID: %w", err)
	}

	return wi, nil
}

// Update updates an existing work item.
func (w *DemoTwoWorkItem) Update(ctx context.Context, d models.DBTX, caller CtxUser, id models.WorkItemID, params repos.DemoTwoWorkItemUpdateParams) (*models.WorkItem, error) {
	defer newOTelSpan().Build(ctx).End()

	if err := w.wiSvc.validateUpdateParams(d, caller, params.Base); err != nil {
		return nil, err
	}

	wi, err := w.repos.DemoTwoWorkItem.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("repos.DemoTwoWorkItem.Update: %w", err)
	}

	return wi, nil
}

// repo has Update only, then service has Close() (Update with closed=True), Move() (Update with kanban step change), ...)
// params for dedicated workItem require workItemID (FK-as-PK)
// TBD if useful: ByTag, ByType (for closed workitem searches. open ones simply return everything and filter in client)

func (w *DemoTwoWorkItem) ListDeleted(ctx context.Context, d models.DBTX, teamID models.TeamID) ([]models.WorkItem, error) {
	// WorkItemsByTeamID with deleted opt, orderby createdAt
	return []models.WorkItem{}, errors.New("not implemented")
}

func (w *DemoTwoWorkItem) List(ctx context.Context, d models.DBTX, teamID models.TeamID) ([]models.WorkItem, error) {
	// WorkItemsByTeamID with orderby createdAt
	return []models.WorkItem{}, errors.New("not implemented")
}
