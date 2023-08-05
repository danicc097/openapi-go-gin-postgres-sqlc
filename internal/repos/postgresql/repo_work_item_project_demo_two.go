package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// DemoTwoWorkItem represents the repository used for interacting with DemoTwoWorkItem records.
type DemoTwoWorkItem struct {
	q *db.Queries
}

// NewDemoTwoWorkItem instantiates the DemoTwoWorkItem repository.
func NewDemoTwoWorkItem() *DemoTwoWorkItem {
	return &DemoTwoWorkItem{
		q: db.New(),
	}
}

var _ repos.DemoTwoWorkItem = (*DemoTwoWorkItem)(nil)

func (u *DemoTwoWorkItem) ByID(ctx context.Context, d db.DBTX, id int, opts ...db.WorkItemSelectConfigOption) (*db.WorkItem, error) {
	extraOpts := []db.WorkItemSelectConfigOption{db.WithWorkItemJoin(db.WorkItemJoins{DemoTwoWorkItem: true})}
	return db.WorkItemByWorkItemID(ctx, d, id, (append(extraOpts, opts...))...)
}

func (u *DemoTwoWorkItem) Create(ctx context.Context, d db.DBTX, params repos.DemoTwoWorkItemCreateParams) (*db.WorkItem, error) {
	workItem, err := db.CreateWorkItem(ctx, d, &params.Base)
	if err != nil {
		return nil, fmt.Errorf("could not create workItem: %w", parseErrorDetail(err))
	}

	params.DemoTwoProject.WorkItemID = workItem.WorkItemID
	demoWorkItem, err := db.CreateDemoTwoWorkItem(ctx, d, &params.DemoTwoProject)
	if err != nil {
		return nil, fmt.Errorf("could not create workItem: %w", parseErrorDetail(err))
	}

	workItem.DemoTwoWorkItemJoin = demoWorkItem

	return workItem, nil
}

func (u *DemoTwoWorkItem) Update(ctx context.Context, d db.DBTX, id int, params repos.DemoTwoWorkItemUpdateParams) (*db.WorkItem, error) {
	workItem, err := u.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItem by id: %w", parseErrorDetail(err))
	}
	demoWorkItem := workItem.DemoTwoWorkItemJoin

	if params.Base != nil {
		workItem.SetUpdateParams(params.Base)
	}

	if params.DemoTwoProject != nil {
		demoWorkItem.SetUpdateParams(params.DemoTwoProject)
	}

	workItem, err = workItem.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update workItem: %w", parseErrorDetail(err))
	}
	demoWorkItem, err = demoWorkItem.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update demoWorkItem: %w", parseErrorDetail(err))
	}

	workItem.DemoTwoWorkItemJoin = demoWorkItem

	return workItem, err
}