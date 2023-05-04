package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type DemoWorkItem struct {
	logger *zap.Logger
	wiRepo repos.DemoWorkItem
}

type DemoWorkItemCreateParams struct {
	repos.DemoWorkItemCreateParams
	TagIDs []int `json:"tagIDs" required:"true"`
}

// NewDemoWorkItem returns a new DemoWorkItem service.
func NewDemoWorkItem(logger *zap.Logger, wiRepo repos.DemoWorkItem) *DemoWorkItem {
	return &DemoWorkItem{
		logger: logger,
		wiRepo: wiRepo,
	}
}

// ByID gets a work item by ID.
func (a *DemoWorkItem) ByID(ctx context.Context, d db.DBTX, id int64) (*db.DemoWorkItem, error) {
	defer newOTELSpan(ctx, "DemoWorkItem.ByID").End()

	wi, err := a.wiRepo.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("wiRepo.ByID: %w", err)
	}

	return wi, nil
}

// Create creates a new work item.
func (a *DemoWorkItem) Create(ctx context.Context, d db.DBTX, params DemoWorkItemCreateParams) (*db.DemoWorkItem, error) {
	defer newOTELSpan(ctx, "DemoWorkItem.Create").End()

	wi, err := a.wiRepo.Create(ctx, d, params.DemoWorkItemCreateParams)
	if err != nil {
		return nil, fmt.Errorf("wiRepo.Create: %w", err)
	}

	for _, id := range params.TagIDs {
		_, err := db.CreateWorkItemWorkItemTag(ctx, d, &db.WorkItemWorkItemTagCreateParams{
			WorkItemTagID: id,
			WorkItemID:    wi.WorkItemID,
		})
		// TODO should see what kind of pgx extracted error that is and ignore if it's of type internal.ErrorCodeAlreadyExists (wrapped unique violation error)
		if err != nil {
			return nil, fmt.Errorf("db.CreateWorkItemWorkItemTag: %w", err)
		}
	}

	// TODO accept []Members with their role as well

	// TODO now query workitem.xo.go joining with tags, members, demoWorkItem, etc... and return that instead.
	// all work_item_***_project.go services must return db.WorkItem which has specific project joins, which
	// allows us to share the same generic logic and extend it based on current project, etc.

	return wi, nil
}

// Update updates an existing work item.
func (a *DemoWorkItem) Update(ctx context.Context, d db.DBTX, id int64, params repos.DemoWorkItemUpdateParams) (*db.DemoWorkItem, error) {
	defer newOTELSpan(ctx, "DemoWorkItem.Update").End()

	wi, err := a.wiRepo.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("wiRepo.Update: %w", err)
	}

	return wi, nil
}

// Delete deletes a work item by ID.
func (a *DemoWorkItem) Delete(ctx context.Context, d db.DBTX, id int64) (*db.DemoWorkItem, error) {
	defer newOTELSpan(ctx, "DemoWorkItem.Delete").End()

	wi, err := a.wiRepo.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("wiRepo.Delete: %w", err)
	}

	return wi, nil
}

// repo has Update only, then service has Close() (Update with closed=True), Move() (Update with kanban step change), ...)
// params for dedicated workItem require workItemID (FK-as-PK)
// TBD if useful: ByTag, ByType (for closed workitem searches. open ones simply return everything and filter in client)
