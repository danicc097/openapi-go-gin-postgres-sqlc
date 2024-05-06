package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// DemoTwoWorkItem represents the repository used for interacting with DemoTwoWorkItem records.
type DemoTwoWorkItem struct {
	q models.Querier
}

// NewDemoTwoWorkItem instantiates the DemoTwoWorkItem repository.
func NewDemoTwoWorkItem() *DemoTwoWorkItem {
	return &DemoTwoWorkItem{
		q: NewQuerierWrapper(models.New()),
	}
}

var _ repos.DemoTwoWorkItem = (*DemoTwoWorkItem)(nil)

func (u *DemoTwoWorkItem) ByID(ctx context.Context, d models.DBTX, id models.WorkItemID, opts ...models.WorkItemSelectConfigOption) (*models.WorkItem, error) {
	extraOpts := []models.WorkItemSelectConfigOption{models.WithWorkItemJoin(models.WorkItemJoins{DemoTwoWorkItem: true})}

	return models.WorkItemByWorkItemID(ctx, d, id, (append(extraOpts, opts...))...)
}

func (u *DemoTwoWorkItem) Create(ctx context.Context, d models.DBTX, params repos.DemoTwoWorkItemCreateParams) (*models.WorkItem, error) {
	workItem, err := models.CreateWorkItem(ctx, d, &params.Base)
	if err != nil {
		return nil, fmt.Errorf("could not create workItem: %w", ParseDBErrorDetail(err))
	}

	params.DemoTwoProject.WorkItemID = workItem.WorkItemID
	demoTwoWorkItem, err := models.CreateDemoTwoWorkItem(ctx, d, &params.DemoTwoProject)
	if err != nil {
		return nil, fmt.Errorf("could not create demoTwoWorkItem: %w", ParseDBErrorDetail(err))
	}

	workItem.DemoTwoWorkItemJoin = demoTwoWorkItem

	return workItem, nil
}

func (u *DemoTwoWorkItem) Update(ctx context.Context, d models.DBTX, id models.WorkItemID, params repos.DemoTwoWorkItemUpdateParams) (*models.WorkItem, error) {
	workItem, err := u.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItem by id: %w", ParseDBErrorDetail(err))
	}
	demoTwoWorkItem := workItem.DemoTwoWorkItemJoin

	if params.Base != nil {
		workItem.SetUpdateParams(params.Base)
	}

	if params.DemoTwoProject != nil {
		demoTwoWorkItem.SetUpdateParams(params.DemoTwoProject)
	}

	workItem, err = workItem.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update workItem: %w", ParseDBErrorDetail(err))
	}
	demoTwoWorkItem, err = demoTwoWorkItem.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update demoTwoWorkItem: %w", ParseDBErrorDetail(err))
	}

	workItem.DemoTwoWorkItemJoin = demoTwoWorkItem

	return workItem, err
}
